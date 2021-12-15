package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var port string

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello "+port+"\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// використання флагів

	flag.StringVar(&port, "config", "", "config path")
	flag.Parse()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	e.POST("/containers/start", start)
	e.DELETE("/containers/:port/stop", stop)
	e.GET("/containers", getList)
	e.GET("/user/:id", getUser)

	/*http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8091", nil)*/
}

func start(c echo.Context) error {

	return c.String(http.StatusOK, "")
}
func stop(c echo.Context) error {

	return c.String(http.StatusOK, "")
}

func getList(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
	//return c.String(http.StatusOK, "")
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

/*func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}*/
