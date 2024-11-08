package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Element struct {
	FirstName string
	Surname   string
}

func (e Element) Display(formatName func(element Element) string) {
	//fmt.Printf("Your full name is %v %v\n", e.FirstName, e.Surname)
	fmt.Println(formatName(e))
}

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	server.Run(":8080") // localhost:8080

	element := Element{"Eduard", "Garcia"}
	element.Display(func(element Element) string { return element.FirstName + " " + element.Surname })
	element.Display(func(element Element) string { return element.FirstName + " and " + element.Surname })

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
