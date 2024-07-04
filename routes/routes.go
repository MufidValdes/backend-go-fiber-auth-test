package routes

import (
	"example.com/go-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/users/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}

/*
Ketentuan API :
Pada bagian User Endpoint :
POST : /users/register, dan gunakan atribut berikut ini
ID (primary key, required)
Username (required)
Email (unique & required)
Password (required & minlength 6)
Relasi dengan model Photo (Gunakan constraint cascade)
Created At (timestamp)
Updated At (timestamp)
GET: /users/login
Using email & password (required)
PUT : /users/:userId (Update User)
DELETE : /users/:userId (Delete User)
Photos Endpoint

*/
