package api

import "github.com/gin-gonic/gin"

type Greeting struct {
	Message string `json:"message"`
}

func init() {
	r := gin.Default()

	// Define route for the /hello endpoint
	r.GET("/hello", func(c *gin.Context) {
		// Create a Greeting object
		greeting := Greeting{
			Message: "Hello, World!",
		}

		// Return JSON response
		c.JSON(200, greeting)
	})

	// Start the server
	r.Run(":8080")

}
