package main

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
)

type Hacker struct {
	Name string `redis:"name"`
	Score int `redis:"score"`
}

var pool *redis.Pool

var ErrNoAlbum = errors.New("no album found")

func init ()  {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	
}


func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/json/:hackers", GetHacker)


	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Start server
	log.Fatal(app.Listen(":8010"))
}

// Handler
func GetHacker(key string) (error, *Hacker) {
	conn := pool.Get()
	defer conn.Close()
	values, err := redis.Values(conn.Do("ZRANGE", "key"))
	if err != nil {
		return err, nil

	}
	var hacker Hacker
	err = redis.ScanStruct(values, &hacker)
	if err != nil {
		return err, nil
	}
	return nil, &hacker
}