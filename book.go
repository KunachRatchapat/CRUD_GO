package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	
)



//GetBook
func getBookS(c *fiber.Ctx) error {
	return c.JSON(books)
}

//Find Get book Id
func getBook(c *fiber.Ctx) error {
	//Req Id
	bookId,err := strconv.Atoi(c.Params("id"))
	//Check Errors
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())	
	}
	//Find Id in Slice
	for _, book := range books{ //range วนลูปตัวที่อยู่ใน slice, _  = index
		if book.ID == bookId {
			//Send Res the Book
			return c.JSON(book)
		}
		 
	}
	return c.SendStatus(fiber.StatusNotFound)
}

//Create Book
func createBook(c *fiber.Ctx) error {
	//Instance book for Struct;l
	book := new(Book)
	if err := c.BodyParser(book); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

//Update Book
func updateBook(c *fiber.Ctx) error {
	bookId,err := strconv.Atoi(c.Params("id"))

	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//Instance Bookupdate for struct Book
	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); 
	err !=nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//Check Book in Slice
	for i,book := range books{
		if book.ID == bookId{
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

//Delete Books
func deleteBook(c *fiber.Ctx) error {
	bookId , err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//Find Book For Delete
	for i,book := range books{
		if book.ID == bookId{
			books = append(books[:i] ,books[i+1:]...) //จนถึงตัวสุดท้าย
			return c.SendStatus(fiber.StatusNoContent)
	}

	} 
	return c.SendStatus(fiber.StatusNotFound)
}
