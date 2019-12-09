package APIGO

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"user":    "localhost:8080/v1/user",
		"article": "localhost:8080/v1/article",
		"review":  "localhost:8080/v1/review",
		"column":  "localhost:8080/v1/column"}
	j, _ := json.Marshal(m)
	fmt.Fprintf(w, string(j))
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"GetUserArticleByName",
		"GET",
		"/v1/article/{username}",
		GetUserArticleByName,
	},

	Route{
		"GetUserArticleByNameAndID",
		"GET",
		"/v1/article/{username}/{id}",
		GetUserArticleByNameAndID,
	},

	Route{
		"GetColumnByName",
		"GET",
		"/v1/column/{username}",
		GetColumnByName,
	},

	Route{
		"GetReviewByNameAndID",
		"GET",
		"/v1/review/{username}/{id}",
		GetReviewByNameAndID,
	},

	Route{
		"CreateUser",
		"POST",
		"/v1/user",
		CreateUser,
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/v1/user/{username}",
		DeleteUser,
	},

	Route{
		"LoginUser",
		"GET",
		"/v1/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		"GET",
		"/v1/user/logout",
		LogoutUser,
	},

	Route{
		"GetUserByName",
		"GET",
		"/v1/user/{username}",
		GetUserByName,
	},

	Route{
		"UpdateUser",
		"PUT",
		"/v1/user/{username}",
		UpdateUser,
	},
}
