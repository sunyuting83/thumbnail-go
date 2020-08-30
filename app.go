package main

import (
	"flag"
	"zipimg/zipimg"

	"github.com/gofiber/fiber"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8567, "端口号，默认为8567")
	flag.Parse()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		filepath := c.Query("path")
		d := zipimg.MakeThumbnail(filepath, filepath)
		if d == true {
			c.JSON(fiber.Map{
				"status": 0,
			})
		} else {
			c.JSON(fiber.Map{
				"status": 1,
			})
		}
	})

	app.Listen(port)
}
