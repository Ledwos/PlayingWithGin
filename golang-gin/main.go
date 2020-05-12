package main

import (
	"net/http"

	// abadonded url "github.com/gin-gonic/contrib/static" new url:
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//serve static files (html / css / js)
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// route group for API
	api := router.Group("/api") 
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}

	api.GET("/jokes", JokeHandler)
	api.POST("/jokes/like/:jokeID", LikeJoke)

	// start / run server
	router.Run(":3000")
}

// gets list of avialable jokes
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message":"Jokes handler not implemented yet",
	})
}

// increment likes fora given joke
func LikeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message":"Can't give likes just yet!",
	})
}


