package pongo2

import (
	"github.com/rocwong/neko"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

type user struct {
	Name string
	age  int
}

func Test_Render(t *testing.T) {
	Convey("Normal Render", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		m := neko.New()
		m.Use(Renderer())
		m.GET("/", func(ctx *neko.Context) {
			ctx.Render("home", map[string]interface{}{"user": "pongo2.v3"}, 200)
		})
		m.ServeHTTP(w, req)
		So(w.Body.String(), ShouldEqual, "")
	})

	Convey("Initial with 'BaseDir' Option", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		m := neko.New()
		m.Use(Renderer(Options{BaseDir: "fixtures"}))
		m.GET("/", func(ctx *neko.Context) {
			ctx.Render("home", map[string]interface{}{"user": "pongo2.v3"}, 200)
		})
		m.ServeHTTP(w, req)
		So(w.Body.String(), ShouldEqual, "layout hello pongo2.v3")
	})

	Convey("Initial with 'MultiDir' Option ", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		m := neko.New()
		m.Use(Renderer(
			Options{
				MultiDir: map[string]string {
					"dir2": "fixtures",
				},
			}),
		)
		m.GET("/", func(ctx *neko.Context) {
			ctx.Render("#dir2/user", neko.JSON{"user": &user{Name: "pongo2", age: 3}}, 200)
		})
		m.ServeHTTP(w, req)
		So(w.Body.String(), ShouldEqual, "hello pongo2, i am 3")
	})

	Convey("Initial with 'Extension' Option", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		m := neko.New()
		m.Use(Renderer(Options{Extension: ".html"}))
		m.GET("/", func(ctx *neko.Context) {
			ctx.Render("home", map[string]interface{}{"user": "pongo2.v3"}, 200)
		})
		m.ServeHTTP(w, req)
		So(w.Body.String(), ShouldEqual, "layout hello pongo2.v3")
	})

	Convey("Unsupported data type", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		m := neko.New()
		m.Use(Renderer(Options{Extension: ".html"}))
		m.GET("/", func(ctx *neko.Context) {
			ctx.Render("home", "pongo2.v3", 200)
		})
		So(func() { m.ServeHTTP(w, req) }, ShouldPanic)
	})

}
