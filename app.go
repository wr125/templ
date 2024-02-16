package main

import (
	"context"
	"net/http"

	"github.com/wr125/templ/templates"

	"log"

	"github.com/labstack/echo/v4"

	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	e := echo.New()

	component := templates.Index()

	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)

	})
	e.POST("/home", func(c echo.Context) error {
		// Handle the button click on the server side
		// Perform any necessary actions or return a response
		return c.HTML(http.StatusOK, "<h1>Happy that you did!</h1>")
	})
	e.POST("/clicked", func(c echo.Context) error {
		// Handle the button click on the server side
		// Perform any necessary actions or return a response
		return c.HTML(http.StatusOK, "<h1>Button Clicked!</h1>")
	})
	e.GET("/404", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "The Page You Were Looking For Was Not Found")
	})

	//log.Fatal(http.ListenAndServe(":"+port, nil))

	e.Static("/static", "static")
	e.Static("/css", "css")
	e.Static("/fonts", "fonts")
	e.Static("/img", "img")

	log.Println("listening on", port)
	e.Logger.Fatal(e.Start(":"+port), nil)
	/*e.IPExtractor = echo.ExtractIPFromXFFHeader(
			echo.TrustLoopback(false), // e.g. ipv4 start with 127.
			echo.TrustLinkLocal(false), // e.g. ipv4 start with 169.254
			echo.TrustPrivateNet(false), // e.g. ipv4 start with 10. or 192.168
			echo.TrustIPRange(lbIPRange),
	   )*/
}
