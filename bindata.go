package grpcrgen

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9fb84cfea443acbb7e9448be50708b2897c65aa8 = "package {{.PackageName}}\n\nimport (\n\t\"errors\"\n\t\"fmt\"\n\t\"io\"\n\n\t\"github.com/google/flatbuffers/go\"\n)\n\nvar (\n\t// DefaultMaxChannels is number of buffer chennels. It is used to receive HTTP Reuest's body. Default value is 100.\n\tDefaultMaxChannels = 100\n\n\t// DefaultMaxBufferSize is limit of buffer size. It is used to receive HTTP Reuest's body. Default value is 4KB.\n\tDefaultMaxBufferSize = 4 * 1024\n)\n\nvar bufCh chan []byte\n\ntype tabler interface {\n\tTable() flatbuffers.Table\n}\n\n// InitBuffer creates buffers. This function is called by readBody if no buffers are created. If you care about the performance at runtime, call this method when initializing the application.\nfunc InitBuffer() {\n\tbufCh = make(chan []byte, DefaultMaxBufferSize)\n\tfor i := 0; i < DefaultMaxBufferSize; i++ {\n\t\tbufCh <- make([]byte, DefaultMaxBufferSize)\n\t}\n}\n\nfunc newBuilderForFlatbuffersRawCodec(buf []byte) *flatbuffers.Builder {\n\treturn &flatbuffers.Builder{Bytes: buf}\n}\n\nfunc readBody(r io.Reader) ([]byte, error) {\n\tif cap(bufCh) == 0 {\n\t\tInitBuffer()\n\t}\n\n\tbuf := <-bufCh\n\tdefer func() {\n\t\tbufCh <- buf\n\t}()\n\n\tn, err := r.Read(buf)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tif n < 5 {\n\t\tstr := fmt.Sprintf(\"Invalid request body: %v\", buf)\n\t\treturn nil, errors.New(str)\n\t}\n\n\treturn buf[:n], nil\n}\n"
var _Assets7a93af52afbb617e967b241958bb67b6e4fd9034 = "package {{.PackageName}}\n\nimport (\n\t\"net/http\"\n{{range .API}}\n  {{.Name}} \"{{.ImportPath}}\"\n{{- end}}\n\t\"github.com/julienschmidt/httprouter\"\n\t\"golang.org/x/net/context\"\n\t\"google.golang.org/grpc\"\n)\n\n{{range .API}}\n// New{{.Name}}Router is httprouter for {{.Name}}\nfunc New{{.Name}}Router(conn *grpc.ClientConn) httprouter.Handle {\n\tc := {{.Name}}.New{{.Name}}Client(conn)\n\n\treturn func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {\n\t\tbody, err := readBody(r.Body)\n\t\tif err != nil {\n\t\t\thttp.Error(rw, err.Error(), http.StatusBadRequest)\n\t\t\treturn\n\t\t}\n\n\t\tctx := context.Background()\n\n\t\tb := newBuilderForFlatbuffersRawCodec(body)\n\t\tvar tab tabler\n\t\taction := params.ByName(\"action\")\n\t\tswitch action {\n\t\t{{- range .Functions}}\n\t\tcase \"{{.}}\":\n\t\t\ttab, err = c.{{.}}(ctx, b)\n\t\t{{- end}}\n\t\tdefault:\n\t\t\thttp.Error(rw, \"NotFound\", http.StatusNotFound)\n\t\t\treturn\n\t\t}\n\n\t\tif err != nil {\n\t\t\thttp.Error(rw, err.Error(), http.StatusBadRequest)\n\t\t\treturn\n\t\t}\n\n\t\tvar res []byte\n\t\tif tab != nil {\n\t\t\tres = tab.Table().Bytes\n\t\t}\n\n\t\t_, err = rw.Write(res)\n\t\tif err != nil {\n\t\t\thttp.Error(rw, err.Error(), http.StatusInternalServerError)\n\t\t\treturn\n\t\t}\n\t}\n}\n{{end}}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"helper.tmpl", "router.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1486366354, 1486366354000000000),
		Data:     nil,
	}, "/helper.tmpl": &assets.File{
		Path:     "/helper.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1486366354, 1486366354000000000),
		Data:     []byte(_Assets9fb84cfea443acbb7e9448be50708b2897c65aa8),
	}, "/router.tmpl": &assets.File{
		Path:     "/router.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1486366354, 1486366354000000000),
		Data:     []byte(_Assets7a93af52afbb617e967b241958bb67b6e4fd9034),
	}}, "")
