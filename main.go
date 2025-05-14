package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4" //Create Token
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"` //Metadata
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Globla Varrible
var books []Book

// Check MiddleWare
func checkMiddleWare(c *fiber.Ctx) error {
	start := time.Now()

	fmt.Printf("URL: %s ,Method: %s, Time: %s\n", c.OriginalURL(), c.Method(), start)

	return c.Next()
}

// Main Func
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Load Env Error")
	}
	engine := html.New("./views", ".html") //use engie html and CreateFolder View for use html

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "Devil Mecry", Author: "Teh"})
	books = append(books, Book{ID: 2, Title: "Devil Mecry2", Author: "Teh"})

	//Login
	app.Post("/login", getLogin)

	//CheckMiddle
	app.Use(checkMiddleWare)

	//JWT MiddleWare
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	//Router
	app.Get("/books", getBookS)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHTML)
	app.Get("/config", getENV)

	//Listen เชื่อม ServerPort
	app.Listen(":5000")
}

// Func UploadFile
func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("กรุณาเลือกไฟล์ให้มันถูกดิ้!!")
	}
	//Savefile in Foloder uploads
	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Server is Suck")
	}
	return c.SendString("File Upload Complete!!")
}

// Test Html
func testHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{ //Map key = Title
		"Title": "Hello Bro !!",
	})
}

// ENV
func getENV(c *fiber.Ctx) error {
	//Use Env in os = getenv
	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "defalutsecret"
	}

	return c.JSON(fiber.Map{
		"SECRET": secret,
	})

}

//Login With Struct
type User = struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}
var memberUser = User{
	Email: "kunach@gmail.com",
	Password: "taehuhu555",
}

// Login
func getLogin(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("Login ไม่ถูกหว่ะน้อง")
	}

	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("นาย Login ไม่ถูกหว่ะแม่ง")
	}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = user.Email
		claims["role"] = "admin"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

	return c.JSON(fiber.Map{
		"Message":"Login Success",
		"token":t,
		
	})
}


