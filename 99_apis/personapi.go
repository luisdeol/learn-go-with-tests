package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people = []Person{{Id: 1, Name: "LuisDev", Age: 30}, {Id: 2, Name: "DevLuis", Age: 31}}

func main() {
	router := gin.Default()

	router.GET("/people", getPeople)
	router.GET("/people/:id", getPerson)
	router.Run("localhost:8080")
}

func getPeople(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, people)
}

func getPerson(context *gin.Context) {
	id := context.Param("id")
	idInt, error := strconv.Atoi(id)

	if error != nil {
		fmt.Println("Error: ", error)

		return
	}

	person, err := getPersonById(idInt)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, person)
}

func getPersonById(id int) (*Person, error) {
	for i, p := range people {
		if p.Id == id {
			return &people[i], nil
		}
	}

	return nil, errors.New("Person not found")
}
