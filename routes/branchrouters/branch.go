package branchrouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/need/go-backend/controllers/essential-functions/branchcontroller"
)

func BranchRouter(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-branch", branchcontroller.CreateBranch)
	api.Get("/get-branches", branchcontroller.GetAllBranches)
	api.Get("/get-branch/:branch", branchcontroller.GetBranch)
	api.Delete("/del-branch/:branch", branchcontroller.DeleteBranch)
	api.Put("/update-branch/:branch", branchcontroller.UpdateBranch)
}
