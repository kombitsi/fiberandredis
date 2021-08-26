package handler

import (
	"github.com/gofiber/fiber/v2\"
)

type hackers struct {
	name string `json:"name"`
	score int `json:"score"`


}

	func GetHacker(c *fiber.Ctx)  {
	c.Send("All Books")
}

	func DelHacker(c *fiber.Ctx) {
	c.Send("del Books")
}

	func GetAlbums(c *fiber.Ctx) {
	c.Send("get albums")
}

	func DelAlbums(c *fiber.Ctx) {
		c.Send("del album")
	}
