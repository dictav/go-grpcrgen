package example

import (
	"testing"

	"github.com/dictav/go-grpcrgen/example/myrouter"
	"github.com/dictav/go-grpcrgen/example/myservice"
	"github.com/dictav/go-grpcrgen/example/myservice/blog"
	"github.com/dictav/go-grpcrgen/example/myservice/user"
)

func TestImport(t *testing.T) {
	_ = &myservice.Geo{}
	_ = &blog.Blog{}
	_ = &user.User{}
	_ = myrouter.NewBlogRouter(nil)
}
