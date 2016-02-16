package router

import (
	"net/http"

	"github.com/tksasha/balance/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/tksasha/balance/Godeps/_workspace/src/github.com/justinas/alice"

	. "github.com/tksasha/balance/handlers"
	. "github.com/tksasha/balance/handlers/middleware"
)

func NewRouter() http.Handler {
	router := httprouter.New()

	//
	// /categories
	//
	router.GET("/categories", Categories.Index)

	router.POST("/categories", Categories.Create)

	router.POST("/categories/:category_id/recovery", Recoveries.Create)

	router.PATCH("/categories/:id", Categories.Update)

	router.PUT("/categories/:id", Categories.Update)

	router.DELETE("/categories/:id", Categories.Destroy)

	//
	// /items
	//
	router.GET("/items", Items.Index)

	router.POST("/items", Items.Create)

	router.DELETE("/items/:id", Items.Destroy)

	router.PATCH("/items/:id", Items.Update)

	router.PUT("/items/:id", Items.Update)

	//
	// /consolidates
	//
	router.GET("/consolidates", Consolidates.Index)

	//
	// /cashes
	//
	router.POST("/cashes", Cashes.Create)

	router.GET("/cashes", Cashes.Index)

	router.DELETE("/cashes/:id", Cashes.Destroy)

	router.PATCH("/cashes/:id", Cashes.Update)

	return alice.New(ContentTypeMiddleware).Then(router)
}
