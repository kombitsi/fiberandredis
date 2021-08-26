package main

import (
	"log"
	"time"
	"github.com/gomodule/redigo/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/kombitsi/fiberandredis/handler"

)

func setRoutes(app *fiber.App)  {
	app.Get("/json/hackers", GetHacker)

	app.Get("/json/hackers:key", DelHacker)

	app.Get("/json/albums:id", GetAlbums)

	app.Get("/json/hackers:id", DelAlbums)



}

func main() {
	app := fiber.New()


	app.Listen(8010)
}

