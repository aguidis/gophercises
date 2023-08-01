package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aguidis/cyoa/internal/handler"
	"github.com/aguidis/cyoa/internal/model"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web app on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parsed()
	fmt.Printf("Using the story %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := model.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
