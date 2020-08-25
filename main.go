package main

import (
	"fmt"
	"github.com/abdulhalim-cu/task-api/controllers"
	"github.com/abdulhalim-cu/task-api/database"
	"github.com/abdulhalim-cu/task-api/models"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func Routers(a *fiber.App) {
	a.Get("/api/todos", controllers.GetAllTodos)
	a.Get("/api/todos/:id", controllers.GetTodoById)
	a.Post("/api/todos", controllers.CreateTodo)
	a.Patch("/api/todos/:id", controllers.ToggleTodoStatus)
	a.Delete("/api/todos/:id", controllers.DeleteTodo)
}

func initDB() {
	var err error
	database.DbCon, err = gorm.Open("sqlite3", "todos.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database successfully created")
	database.DbCon.AutoMigrate(&models.Todo{})
	fmt.Println("Database migration done")
}

func main() {
	app := fiber.New()
	// Initialize database
	initDB()
	defer database.DbCon.Close()
	// Setup router
	Routers(app)
	app.Listen(3000)
}
