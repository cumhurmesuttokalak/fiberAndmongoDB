package main

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/cumhurmesuttokalak/go-fiberAndmongoDB/configs"
	"www.github.com/cumhurmesuttokalak/go-fiberAndmongoDB/routes"
)

func main() {
	app := fiber.New()
	//veritabanı bağlantısı başlatma işlemi
	configs.ConnectDB()
	//rotaların tanımı
	routes.Routes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})
	app.Listen(":8088")
}
