package main

import (
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"errors"

	// jwtmiddleware
	"github.com/auth0/go-jwt-middleware"
	// jwt
	"github.com/dgrijalva/jwt-go"
	// abadonded url "github.com/gin-gonic/contrib/static" new url:
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

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

var jwtMiddleWare *jwtmiddleware.JWTMiddleware

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N string `json:"n"`
	E string `json:"e"`
	X5c string `json:"x5c"`
}

func main() {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options {
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			aud := os.Getenv("AUTH0_API_AUDIENCE")
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAudience {
				return token, errors.New("Invalid audience.")
			}
			// verify iss claim
			iss := os.Getenv("AUTH0_DOMAIN")
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss,false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := getPermCert(token)
			if err != nil {
				log.Fatalf("count not get cert: %+v", err)
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	//register jwtMiddleware
	jwtMiddleWare = jwtMiddleware

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

	api.GET("/jokes", authMiddleware(), JokeHandler)
	api.POST("/jokes/like/:jokeID", authMiddleware(), LikeJoke)

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

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get client secret key
		err := jwtMiddleWare.CheckJWT(c.Writer, c.Request)
		if err != nil {
			//token not found
			fmt.Println(err)
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Unauthorized"))
			return
		}
	}
}

func getPermCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(os.Getenv("AUTH0_DOMAIN") + ".well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	// x5c := jwks.Keys[0].X5c
	// for k, v := range x5c {
	// 	if token.Header["kid"] == jwks.Keys[k].Kid {
	// 		cert = "-----BEGIN CERTIFICATE-----\n" + v + "\n-----END CERTIFICATE-----"
	// 	}
	// }

	if cert == "" {
		return cert, errors.New("unable to find appropriate key.")
	}

	return cert, nil
}
