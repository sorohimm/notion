package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"notion/internal/infrastructure"
	"os"
)

var (
	log *zap.SugaredLogger
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error loading logger: %s", err)
		os.Exit(1)
		return
	}

	log = logger.Sugar()
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	injector, err := infrastructure.Injector(log, client)
	if err != nil {
		log.Fatal("main :: inject failing")
	}

	hackersController := injector.InjectBalanceController()

	app := fiber.New()

	app.Get("json/hackers", hackersController.GetHackers)

	log.Fatal(app.Listen(":8010"))
}
