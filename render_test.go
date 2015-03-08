package pongo2

import (
	"github.com/rocwong/neko"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

	Convey("Initial Whith All Options", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		m := neko.New()
		m.Use(Renderer(Options{BaseDir: "fixtures", Extension: ".html"}))
		m.GET("/", func(ctx *neko.Context) {
			ctx.Render("home", map[string]interface{}{"user": "pongo2.v3"}, 200)
		})
		m.ServeHTTP(w, req)
		So(w.Body.String(), ShouldEqual, "layout hello pongo2.v3")
	})

	Convey("Initial Whith 'BaseDir' Option", t, func() {
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

	Convey("Initial Whith 'Extension' Option", t, func() {
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

}
