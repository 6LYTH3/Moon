package route

import (
	"../controller"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func Load() *httprouter.Router {
	fmt.Println("Load Route")
	r := httprouter.New()

	r.GET("/home", controller.Home)
	r.GET("/hello", controller.Hello)
	r.GET("/hello/:name", controller.Hello)
	r.POST("/blog/new", controller.Post)

	return r
}
