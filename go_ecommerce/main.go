package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Cloud-Hacks/go_dev_prac/go_commerce/api"
	"github.com/Cloud-Hacks/go_dev_prac/go_commerce/store"

	"github.com/anthdm/weavebox"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleAPIError(ctx *weavebox.Context, err error) {
	fmt.Println("API error:", err)
	ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
}

func main() {
	app := weavebox.New()
	app.ErrorHandler = handleAPIError

	adminMW := &api.AdminAuthMiddleware{}
	adminW := &api.AuthenticationHandler{}
	adminRoute := app.Box("/admin")
	adminRoute.Post("/", adminW.AuthenticateUser)
	adminRoute.Use(adminMW.Authenticate)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	productStore := store.NewMongoProductStore(client.Database("ggcommerce"))
	productHandler := api.NewProductHandler(productStore)

	// admin/product
	adminProductRoute := adminRoute.Box("/product")
	adminProductRoute.Get("/:id", productHandler.HandleGetProductByID)
	adminProductRoute.Get("/", productHandler.HandleGetProducts)
	adminProductRoute.Post("/", productHandler.HandlePostProduct)

	app.Serve(3001)
}
