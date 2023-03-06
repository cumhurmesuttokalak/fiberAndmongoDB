package routes

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/cumhurmesuttokalak/go-fiberAndmongoDB/controllers"
)

func Routes(app *fiber.App) {
	//TODO: routes
	app.Post("/createuser", controllers.CreateUser)
	app.Get("/getsingleuser/:userId", controllers.GetSingleUser)
	app.Get("/getallusers", controllers.GetAllUsers)
	app.Put("updateuser", controllers.UpdateUser)
	app.Delete("deleteuser:/userId", controllers.DeleteUser)
}
