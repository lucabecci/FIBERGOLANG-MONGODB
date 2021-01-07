package main

import (
	"context"
	"encoding/json"
	"fiber-golang-mongo/database"
	"fiber-golang-mongo/models"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	port := 4000

	app.Post("/person", createPerson)
	// app.Post("/person", createPerson)
	// app.Put("/person/:id", updatePerson)
	// app.Delete("/person/:id", deletePerson)
	app.Listen(port)
}

func createPerson(c *fiber.Ctx) {
	collection, err := database.GetMongoDBCollection("fiber_go", "person")

	if err != nil {
		c.Status(500).Send(err)
	}

	var person models.Person

	json.Unmarshal([]byte(c.Body()), &person)

	res, err := collection.InsertOne(context.Background(), person)

	if err != nil {
		c.Status(500).Send(err)
	}

	response, _ := json.Marshal(res)

	c.Send(response)
}
