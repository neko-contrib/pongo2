package pongo2

import (
	"github.com/rocwong/neko"
	"gopkg.in/flosch/pongo2.v3"
	"sync"
)

type (
	Options struct {
		// BaseDir represents a base directory of the pongo2 templates.
		BaseDir string
		// Extension represents an extension of files.
		Extension string
	}
	pongoRenderer struct {
		context *neko.Context
	}
)

var pongoCache = map[string]*pongo2.Template{}
var mutex = &sync.RWMutex{}
var opt Options

func Renderer(options ...Options) neko.HandlerFunc {
	if options != nil {
		opt = options[0]
	} else {
		opt = Options{BaseDir: "views/", Extension: ".html"}
	}
	if opt.BaseDir == "" {
		opt.BaseDir = "views/"
	}
	if opt.Extension == "" {
		opt.Extension = ".html"
	}
	if lastChar(opt.BaseDir) != '/' {
		opt.BaseDir += "/"
	}
	return func(ctx *neko.Context) {
		ctx.HtmlEngine = &pongoRenderer{context: ctx}
	}
}

func (c *pongoRenderer) Render(view string, context interface{}, status ...int) (err error) {
	mutex.RLock()
	template, ok := pongoCache[view]
	mutex.RUnlock()

	if !ok {
		template, err = pongo2.FromFile(opt.BaseDir + view + opt.Extension)
		if err != nil {
			return err
		}
		mutex.Lock()
		pongoCache[view] = template
		mutex.Unlock()
	}
	code := 200
	if status != nil {
		code = status[0]
	}
	c.context.Writer.WriteHeader(code)
	err = template.ExecuteWriter(getContext(context), c.context.Writer)
	return
}

func getContext(data interface{}) pongo2.Context {
	return pongo2.Context(data.(map[string]interface{}))
}

func lastChar(str string) uint8 {
	size := len(str)
	return str[size-1]
}
