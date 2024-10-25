package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wanchanok6698/web-blogs/api/middleware"
	"github.com/wanchanok6698/web-blogs/api/v1/controllers"
	"github.com/wanchanok6698/web-blogs/api/v1/models"
	"github.com/wanchanok6698/web-blogs/api/v1/services"
)

func BlogRoutes(prefix string, app *fiber.App) {
	blogService, err := services.NewBlogService()
	if err != nil {
		log.Fatalf("Failed to initialize blog service: %v", err)
	}

	blogController := controllers.NewBlogsController(*blogService)

	routeGrop := app.Group(prefix)
	routeGrop.Get("/blogs", blogController.GetAllBlogs)
	routeGrop.Get("/blogs/:id", blogController.GetBlogByID)
	routeGrop.Post("/blogs", middleware.ValidateData(&models.BlogPost{}), blogController.CreateBlog)
	routeGrop.Patch("/blogs/:id", middleware.ValidateData(&models.BlogUpdate{}), blogController.UpdateBlog)
	routeGrop.Delete("/blogs/:id", blogController.DeleteBlog)
}
