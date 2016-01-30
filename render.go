package pongo2

import (
	"github.com/rocwong/neko"
	"gopkg.in/flosch/pongo2.v3"
	"sync"
	"regexp"
	"strings"
)

type (
	Options struct {
		// BaseDir represents a base directory of the pongo2 templates.
		BaseDir string
		// MultiDir multiple directories of the pongo2 templates.
		MultiDir map[string]string
		// Extension represents an extension of files.
		Extension string
	}
	pongoRenderer struct {
		context *neko.Context
	}
)

var reg *regexp.Regexp
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
	if !strings.HasSuffix(opt.BaseDir, "/") {
		opt.BaseDir += "/"
	}
	for k, v := range opt.MultiDir {
		if !strings.HasSuffix(v, "/") {
			opt.MultiDir[k] += "/"
		}
	}
	reg, _ = regexp.Compile(`^#([\w-]+)/(.+)$`)

	return func(ctx *neko.Context) {
		ctx.HtmlEngine = &pongoRenderer{context: ctx}
	}
}

func (c *pongoRenderer) Render(view string, context interface{}, status ...int) (err error) {
	mutex.RLock()
	template, ok := pongoCache[view]
	mutex.RUnlock()
	if !ok {
		match := reg.FindStringSubmatch(view)

		if len(match) == 3 && len(opt.MultiDir[match[1]]) > 0 {
			template, err = pongo2.FromFile(opt.MultiDir[match[1]] + match[2] + opt.Extension)
		} else {
			template, err = pongo2.FromFile(opt.BaseDir + view + opt.Extension)
		}
		if err != nil {
			c.context.Writer.WriteHeader(500)
			c.context.Writer.Write([]byte(err.Error()))
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
	var d map[string]interface{} = make(map[string]interface{})
	switch data.(type) {
	case neko.JSON:
		d = map[string]interface{}(data.(neko.JSON))
	case map[string]interface{}:
		d = data.(map[string]interface{})
	case nil:
		d = nil
	default:
		panic("Error: Data type does not supported")
	}
	return pongo2.Context(d)
}

