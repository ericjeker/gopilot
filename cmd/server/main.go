package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopilot/pkg/markdown"
	"log"
	"net/http"
)

func main() {
	files, err := markdown.FindMarkdownFiles([]string{"./data/markdown"})
	if err != nil {
		log.Fatalf("Error finding markdown files: %v", err)
	}

	fmt.Println("Found mardown files:")
	for _, file := range files {
		fmt.Println(file)
	}

	fmt.Println("--------------------------------")
	for _, file := range files {
		content, err := markdown.ParseMarkdownFile(file)
		if err != nil {
			log.Fatalf("Error parsing markdown file: %v", err)
		}

		fmt.Printf("Parsed: \n%s\n%s\n", file, content)
		fmt.Println("--------------------------------")
	}

	//r := setupRouter()
	//err := r.Run(`:8080`)
	//if err != nil {
	//	return
	//}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
