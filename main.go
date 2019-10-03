package main

import (
	"github.com/labstack/echo"
	"github.com/pokpitch/goclean/api/v1/promotion/di"
	"github.com/pokpitch/goclean/api/v1/promotion/gateway/route"
)

func main() {
	e := echo.New()

	// Routes
	route.NewPromotionRoute(di.ProvidePromotionHandler()).Initial(e)

	// Listener
	e.Logger.Fatal(e.Start(":1323"))
}