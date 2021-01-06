package router
import (
	"github.com/gorilla/mux"
	"goWeb/middleware"

	"goWeb/handler"
)
func RegisterRoutes(r *mux.Router){
	r.Use(middleware.Logging())
	indexRouter := r.PathPrefix("/index").Subrouter()
	indexRouter.Handle("/",&handler.HelloHandler{})

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/names/{name}/countries/{country}",handler.ShowVisitorInfo)
	userRouter.Use(middleware.Method("GET"))
}
