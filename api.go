package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	lib := Library{
		Name:  "My libary",
		Books: []BookList{},
		Users: make(map[strings]*Users),
	}

	//Books add Get method
	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(lib.Books)

	})
	app.Post("/books", func(c *fiber.Ctx) error {
		var bookdetail struct {
			Name string
			ID string
			Subject string 
			ISBN int 

		}
		if err:=c.BodyParser($bookbookdetail); err!=nill{
			return c.Status(400).JSON((fiber.Map{err.Error()}))
		}
		return c.JSON(lib.Books)
	})


// user 
	app.Get("/users",func (c*fiber.ctx)error  {
		return c.JSON(lib.Users)
		
	})


	app.Post("/users",func (C*fiber.Ctx) error  {
		var {
			var user struct {
			Name     string `json:"name"`
			ID       string `json:"id"`
			Email    string `json:"email"`
			userName string `json:"username"`
		}
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		msg, err := AddUser(lib.Users, user.Name, user.ID, user.Email, user.UserName)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"message": msg})
		}
	})
	
	app.Listen(":3000")

}
