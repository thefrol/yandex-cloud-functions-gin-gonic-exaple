package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	//router inits here, so newer call
	//to cloud function would use this global value
	//and wont create router again
	router = gin.Default()
	router.GET("/", MyHandler)
	router.POST("/stuff", PostHandler)
}

func main() {
	//used for runnig localy
	//can test your api
	//go run .
	//then open http://localhost:8080 in your browser
	router.Run()
}

func MyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": "", "value": "test"})
}

// this function handles post requests to http://your.site/stuff
// it reads body as json,
// and returns its field "name"
func PostHandler(c *gin.Context) {
	type Request struct {
		Name string `json:"name"`
	}
	var r Request
	if err := c.ShouldBindJSON(&r); err != nil {
		fmt.Println("Received bad data")
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"user": r.Name, "value": "its working"})
}

func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(*r)
	router.ServeHTTP(w, r) // ServeHTTP conforms to the http.Handler interface (https://godoc.org/github.com/gin-gonic/gin#Engine.ServeHTTP)
}
