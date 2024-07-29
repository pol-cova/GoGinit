package templates

const GoFrTemplate = `package main
import "gofr.dev/pkg/gofr"

func main() {
    // initialise gofr object
    app := gofr.New()

    // register route greet
    app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

        return "Hello GoFr!", nil
    })

    // Runs the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
   app.Run()
}`
