package main

import (
	"net/http"
	"strconv"

	// abadonded url "github.com/gin-gonic/contrib/static" new url:
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// struct for jokes
type Joke struct {
	ID 		int 	`json:"id" binding:"required"`
	Likes 	int 	`json:"likes"`
	Joke	string	`json:"joke" binding:"required"`
}

// slice of jokes
var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
  }

func main() {

	router := gin.Default()

	//serve static files (html / css / js)
	router.Use(static.Serve("/", static.LocalFile("../reactFront/views", true)))

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
	c.JSON(http.StatusOK, jokes)
}

// increment likes fora given joke
func LikeJoke(c *gin.Context) {
	// check if JokeID exists
	if jokeid, err := strconv.Atoi(c.Param("jokeID"));
		err == nil {
			// find joke and increment its Likes
			for i:= 0; i < len(jokes); i++ {
				if jokes[i].ID == jokeid {
					jokes[i].Likes += 1
				}
			}
			// return pointer to updated jokes list
			c.JSON(http.StatusOK, &jokes)
	} else {
		// if JokeID invalid
		c.AbortWithStatus(http.StatusNotFound)
	}

}


