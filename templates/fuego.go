package templates

const FuegoTemplate = `

package main

import "github.com/go-fuego/fuego"

func main() {
	s := fuego.NewServer()

	fuego.Get(s, "/", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, from Fuego!", nil
	})

	s.Run()
}`
