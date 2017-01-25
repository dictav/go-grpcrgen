package grpcrgen

import (
	"net/http"

	Common "github.com/dictav/go-grpcrgen/example/myservice"
	Blog "github.com/dictav/go-grpcrgen/example/myservice/blog"
	User "github.com/dictav/go-grpcrgen/example/myservice/user"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// NewCommonRouter is httprouter for Common
func NewCommonRouter(conn *grpc.ClientConn) httprouter.Handle {
	c := Common.NewCommonClient(conn)

	return func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		body, err := readBody(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.Background()

		b := newBuilderForFlatbuffersRawCodec(body)
		var tab tabler
		action := params.ByName("action")
		switch action {
		case "Hello":
			tab, err = c.Hello(ctx, b)
		default:
			http.Error(rw, "NotFound", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		var res []byte
		if tab != nil {
			res = tab.Table().Bytes
		}

		_, err = rw.Write(res)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// NewBlogRouter is httprouter for Blog
func NewBlogRouter(conn *grpc.ClientConn) httprouter.Handle {
	c := Blog.NewBlogClient(conn)

	return func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		body, err := readBody(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.Background()

		b := newBuilderForFlatbuffersRawCodec(body)
		var tab tabler
		action := params.ByName("action")
		switch action {
		case "Search":
			tab, err = c.Search(ctx, b)
		default:
			http.Error(rw, "NotFound", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		var res []byte
		if tab != nil {
			res = tab.Table().Bytes
		}

		_, err = rw.Write(res)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// NewUserRouter is httprouter for User
func NewUserRouter(conn *grpc.ClientConn) httprouter.Handle {
	c := User.NewUserClient(conn)

	return func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		body, err := readBody(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.Background()

		b := newBuilderForFlatbuffersRawCodec(body)
		var tab tabler
		action := params.ByName("action")
		switch action {
		case "Get":
			tab, err = c.Get(ctx, b)
		default:
			http.Error(rw, "NotFound", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		var res []byte
		if tab != nil {
			res = tab.Table().Bytes
		}

		_, err = rw.Write(res)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
