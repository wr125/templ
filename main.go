package main

import (
	"context"
	"net/http"

	"github.com/wr125/templ/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	component := templates.Index("Costfood.com")

	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)

	})
	e.POST("/clicked", func(c echo.Context) error {
		// Handle the button click on the server side
		// Perform any necessary actions or return a response
		return c.String(http.StatusOK, "Button Clicked!")
	})
	e.GET("/404", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "The Page You Were Looking For Was Not Found")
	})
	/*e.IPExtractor = echo.ExtractIPFromXFFHeader(
			echo.TrustLoopback(false), // e.g. ipv4 start with 127.
			echo.TrustLinkLocal(false), // e.g. ipv4 start with 169.254
			echo.TrustPrivateNet(false), // e.g. ipv4 start with 10. or 192.168
			echo.TrustIPRange(lbIPRange),
	   )*/
	e.Static("/static", "static")
	e.Static("/css", "css")
	e.Static("/fonts", "fonts")

	e.Logger.Fatal(e.Start(":3000"))
}
