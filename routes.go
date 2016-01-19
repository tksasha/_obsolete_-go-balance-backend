package main

import (
  "github.com/julienschmidt/httprouter"
)

type Route struct {
  Method      string
  Path        string
  HandlerFunc httprouter.Handle
}

var routes = []Route {
  //Route { "GET",  "/items",     ItemsIndex  },
  //Route { "POST", "/items",     ItemsCreate },
  //Route { "GET",  "/items/:id", ItemsShow   },

  Route { "GET",  "/categories", CategoriesIndex  },
  Route { "POST", "/categories", CategoriesCreate },

  //Route { "GET", "/consolidates", ConsolidatesIndex },
}
