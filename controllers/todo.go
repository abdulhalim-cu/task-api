package controllers

import (
	"github.com/abdulhalim-cu/task-api/database"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	"github.com/abdulhalim-cu/task-api/models"
)

// GET /api/todos
func GetAllTodos(ctx *fiber.Ctx) {
	db := database.DbCon
	var todos []models.Todo
	//errors := db.Find(&todos).GetErrors()
	err := db.Find(&todos).Error
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err,
		})
		return
	}

	//if errors != nil {
	//	ctx.Status(500).JSON(fiber.Map{
	//		"ok": false,
	//		"errors": errors,
	//	})
	//	return
	//}
	ctx.JSON(fiber.Map{
		"ok": true,
		"todos": todos,
	})
}

// GET /api/todos/:id
func GetTodoById(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DbCon
	var todo models.Todo
	if err := db.Find(&todo, id).Error; gorm.IsRecordNotFoundError(err) {
		ctx.Status(404).JSON(fiber.Map{
			"ok": false,
			"error": "Todo not found",
		})
		return
	}
	ctx.JSON(fiber.Map{
		"ok": true,
		"todo": todo,
	})
}

// POST /api/todos
func CreateTodo(ctx *fiber.Ctx) {
	params := new(struct{
		Title		string
		Description	string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok": false,
			"error": "Title or description not specified",
		})
		return
	}

	db := database.DbCon
	todo := models.CreateTodo(params.Title, params.Description)

	if err := db.Create(&todo).Error; err !=nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err,
		})
		return
	}
	ctx.JSON(fiber.Map{
		"ok": true,
		"todo": todo,
	})
}

// PATCH /api/todos/:id
func ToggleTodoStatus(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DbCon
	var todo models.Todo
	if err := db.Find(&todo, id).Error; gorm.IsRecordNotFoundError(err) {
		ctx.Status(404).JSON(fiber.Map{
			"ok": false,
			"error": "todo not found",
		})
		return
	}

	err := db.Model(&todo).Update("Done", !todo.Done).Error
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err,
		})
		return
	}
	ctx.JSON(fiber.Map{
		"ok": true,
		"todo": todo,
	})
}

// DELETE /api/todos/:id
func DeleteTodo(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	db := database.DbCon
	var todo models.Todo
	if err := db.Find(&todo, id).Error; gorm.IsRecordNotFoundError(err) {
		ctx.Status(404).JSON(fiber.Map{
			"ok": false,
			"error": "Todo not found",
		})
		return
	}
	err := db.Delete(&todo).Error
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err,
		})
		return
	}
	ctx.JSON(fiber.Map{
		"ok": true,
		"todo": todo,
	})
}

