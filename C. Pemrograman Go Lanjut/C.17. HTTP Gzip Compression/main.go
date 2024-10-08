package main

import (
	"io"
	// "net/http"
	"os"

	// "github.com/NYTimes/gziphandler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// mux := new(http.ServeMux)

	// mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
	// 	f, err := os.Open("sample.png")
	// 	if f != nil {
	// 		defer f.Close()
	// 	}
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	_, err = io.Copy(w, f)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	// server := new(http.Server)
	// server.Addr = ":9000"
	// server.Handler = gziphandler.GzipHandler(mux)

	// server.ListenAndServe()

	e := echo.New()

	e.Use(middleware.Gzip())

	e.GET("/image", func(c echo.Context) error {
		f, err := os.Open("sample.png")
		if err != nil {
			return err
		}

		_, err = io.Copy(c.Response(), f)
		if err != nil {
			return err
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))
}
