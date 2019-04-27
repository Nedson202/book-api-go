package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nedson202/book-api-go/config"
	"github.com/nedson202/book-api-go/controllers"
)

var controller = controllers.Controller{}

func getRoutes() Routes {
	var routes = Routes{
		Route{
			"GetBooks",
			"GET",
			"/books",
			controller.GetBooks(),
		},
		Route{
			"GetBook",
			"GET",
			"/books/{id}",
			controller.GetBook(),
		},
		Route{
			"AddBook",
			"POST",
			"/books",
			controller.AddBook(),
		},
		Route{
			"UpdateBook",
			"PUT",
			"/books/{id}",
			controller.UpdateBook(),
		},
		Route{
			"DeleteBook",
			"DELETE",
			"/books{id}",
			controller.RemoveBook(),
		},
	}

	return routes
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := getRoutes()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = config.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
