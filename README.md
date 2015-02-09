#Pongo2 Render
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

app := neko.Classic()
//default: Options{BaseDir: "views", Extension: ".html"}
//app.Use(pongo2.Renderer())
app.Use(pongo2.Renderer(&Options{BaseDir: "template/", Extension: ".html"}))
app.Run(":3000")
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


