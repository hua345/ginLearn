package main

import (
	"ginLearn/pkg/config"
	"ginLearn/pkg/redigo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	config.Init()
	redigo.Init()
}
func main() {
	router := gin.Default()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	redigo.SetWithExpire("name", "fang", 60*60)
	value, err := redigo.Get("name")
	if err != nil {
		log.Println("Read redis key name failed")
		panic("Read redis key name failed")
	}
	log.Printf("redis Name: %s", value)
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8090")
}
