// package main

// import (
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin",
// 			"https://www.google.com, https://novalagung.com")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "*")
// 		w.Header().Set("Access-Control-Allow-Headers", "*")

// 		if r.Method == "OPTIONS" {
// 			w.Write([]byte("allowed"))
// 			return
// 		}

// 		w.Write([]byte("hello"))
// 	})

// 	log.Println("Starting app at :9000")
// 	http.ListenAndServe(":9000", nil)
// }

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/cors"
)

func main() {
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://novalagung.com", "https://www.google.com"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/index", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
