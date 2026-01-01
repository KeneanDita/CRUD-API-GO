package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Task struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func create(c *fiber.Ctx, db *sql.DB) error {
	var task Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	id := uuid.NewString()

	query := `
		INSERT INTO tasks (id, title, description, status)
		VALUES ($1, $2, $3, $4)
	`

	if _, err := db.Exec(query, id, task.Title, task.Description, task.Status); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create task"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Task created successfully",
		"id":      id,
	})
}

func list(c *fiber.Ctx, db *sql.DB) error {
	rows, err := db.Query("SELECT id, title, description, status FROM tasks")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Scan error"})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func getTask(c *fiber.Ctx, db *sql.DB) error {
	var task Task

	err := db.QueryRow(
		"SELECT id, title, description, status FROM tasks WHERE id=$1",
		c.Params("id"),
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status)

	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(task)
}

func update(c *fiber.Ctx, db *sql.DB) error {
	var task Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	query := `
		UPDATE tasks
		SET title=$2, description=$3, status=$4
		WHERE id=$1
	`

	if _, err := db.Exec(query, c.Params("id"), task.Title, task.Description, task.Status); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update task"})
	}

	return c.JSON(fiber.Map{"message": "Task updated successfully"})
}

func remove(c *fiber.Ctx, db *sql.DB) error {
	if _, err := db.Exec("DELETE FROM tasks WHERE id=$1", c.Params("id")); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete task"})
	}

	return c.JSON(fiber.Map{"message": "Task deleted successfully"})
}
