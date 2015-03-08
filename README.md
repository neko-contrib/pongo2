#Pongo2 Render
[![wercker status](https://app.wercker.com/status/d4def3154c15de48715ae974744df9f5/s "wercker status")](https://app.wercker.com/project/bykey/d4def3154c15de48715ae974744df9f5)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/neko-contrib/pongo2)
[![GoCover](http://gocover.io/_badge/github.com/neko-contrib/pongo2)](http://gocover.io/github.com/neko-contrib/pongo2)

[Pongo2.v3](https://github.com/flosch/pongo2) template engine for [Neko](https://github.com/rocwong/neko).

## Usage
~~~go
package main

import (
  "github.com/rocwong/neko"
  "github.com/neko-contrib/pongo2"
)

func main() {
  app := neko.Classic()
  //default: Options{BaseDir: "views", Extension: ".html"}
  //app.Use(pongo2.Renderer())
  app.Use(pongo2.Renderer(pongo2.Options{BaseDir: "template/", Extension: ".html"}))
  app.Run(":3000")
}
~~~

####type Options
~~~go
type Options struct {
  // BaseDir represents a base directory of the pongo2 templates.
  BaseDir string
  // Extension represents an extension of files.
  Extension string
}
~~~


