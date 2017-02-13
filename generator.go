package grpcrgen

//go:generate go-assets-builder --package=grpcrgen --strip-prefix="/template" --output=bindata.go template

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	clientSuffix      = "Client"
	serviceFileSuffix = "_grpc.go"
	routerTemplate    = "/router.tmpl"
	helperTemplate    = "/helper.tmpl"
)

type api struct {
	Name       string
	ImportPath string
	Functions  []string
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "", 0)
}

func executeTemplate(name string, params interface{}) (out []byte, err error) {
	w := bytes.NewBuffer(out)

	t := template.New("t")
	tmpl := string(Assets.Files[name].Data)
	template.Must(t.Parse(tmpl))
	t.Execute(w, params)

	return w.Bytes(), nil
}

// Generate http HandleFunc
func Generate(inputDir, outputDir string) error {
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}
	}

	packageName := path.Base(outputDir)
	fn := filepath.Join(outputDir, "helper.go")
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := generateHelper(f, packageName); err != nil {
		return err
	}

	fn = filepath.Join(outputDir, "router.go")
	f, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return generateRouter(f, inputDir, packageName)
}

func generateHelper(w io.Writer, packageName string) error {
	params := struct {
		PackageName string
	}{
		PackageName: packageName,
	}

	out, err := executeTemplate(helperTemplate, params)
	if err != nil {
		return err
	}

	_, err = w.Write(out)
	return err
}

func extractAPI(dir string) ([]api, error) {
	logger.Println("Searching ", dir, " ...")

	ctx := build.Default
	var pkg *build.Package
	var err error
	if build.IsLocalImport(dir) {
		pkg, err = ctx.ImportDir(dir, 0)
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		wd = filepath.Join(wd, dir)
		src := filepath.Join(ctx.GOPATH, "src")
		path, err := filepath.Rel(src, wd)
		if err != nil {
			return nil, err
		}
		pkg, err = ctx.Import(path, dir, 0)
	}
	if err != nil {
		return nil, fmt.Errorf("Cannot import directory %s: %s", dir, err)
	}

	apis, err := extractAPIImpl(dir, pkg)
	if err != nil {
		return nil, err
	}

	if info, err := os.Stat(dir); err != nil {
		return nil, err
	} else if !info.IsDir() {
		return nil, fmt.Errorf(dir + " is not directory")
	}

	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range infos {
		if f.IsDir() {
			if ret, err := extractAPI(filepath.Join(dir, f.Name())); err == nil {
				apis = append(apis, ret...)
			}
		}
	}

	return apis, nil
}

func extractAPIImpl(dir string, pkg *build.Package) ([]api, error) {
	apis := make([]api, 0, 20)
	fs := token.NewFileSet()
	for _, v := range pkg.GoFiles {
		if !strings.HasSuffix(v, serviceFileSuffix) {
			continue
		}

		importPath := pkg.ImportPath
		if build.IsLocalImport(dir) {
			importPath = "../" + importPath
		}
		logger.Println("Extract a service ", v, " in ", importPath)
		fn := filepath.Join(dir, v)
		parsedFile, err := parser.ParseFile(fs, fn, nil, 0)
		if err != nil {
			return nil, err
		}

		for _, decl := range parsedFile.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			if genDecl.Tok != token.TYPE {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				tn := typeSpec.Name.Name
				if !strings.HasSuffix(tn, clientSuffix) {
					continue
				}
				serviceName := tn[0 : len(tn)-len(clientSuffix)] // remove clientSuffix

				interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
				if !ok {
					continue
				}

				funcNames := make([]string, 0, 20)
				for _, m := range interfaceType.Methods.List {
					if len(m.Names) == 0 {
						continue
					}

					name := m.Names[0].Name

					if _, ok := m.Type.(*ast.FuncType); !ok {
						continue
					}

					funcNames = append(funcNames, name)
				}

				apis = append(apis, api{
					serviceName,
					importPath,
					funcNames,
				})
			}
		}
	}

	return apis, nil
}

func generateRouter(w io.Writer, dir, packageName string) error {

	logger.Println("Generating for " + packageName)

	ret, err := extractAPI(dir)
	if err != nil {
		return err
	}

	params := struct {
		PackageName string
		API         []api
	}{
		PackageName: packageName,
		API:         ret,
	}

	out, err := executeTemplate(routerTemplate, params)
	if err != nil {
		logger.Println("HERE2")
		return err
	}

	_, err = w.Write(out)
	return err
}
