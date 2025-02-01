package main

import (
	"example.com/controllers"
	"example.com/handlers"
	"example.com/middlewares"
	"example.com/store"
	"example.com/utils"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || dbName == "" {
		log.Fatal("Missing necessary environment variables")
		return
	}

	db, err := utils.Connection(user, password, host, dbName)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	// Initialize a new Fiber app
	app := fiber.New()

	// Dependency Injection
	userstore := store.NewUserStore(db)
	userController := controllers.NewUserController(userstore)
	userHandler := handlers.NewAuthhandler(userController)

	app.Post("/register", userHandler.Register)

	app.Post("/login", userHandler.Login)

	app.Get("/admin", middlewares.VerifyUser, middlewares.CheckAdmin, userHandler.GetAdminDetails)

	// Start the server on port 3000
	log.Fatal(app.Listen(":4000"))
}
