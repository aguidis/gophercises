package main

import (
	"flag"
	"fmt"
	"github.com/aguidis/cyoa/internal/layout"
	"log"
	"net/http"
	"os"

	"github.com/aguidis/cyoa/internal/handler"
	"github.com/aguidis/cyoa/internal/model"
)

func main() {
	// Create flags for our optional variables
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// Open the JSON file and parse the story in it.
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := model.JsonStory(f)
	if err != nil {
		panic(err)
	}

	// Create our custom CYOA story handler
	h := handler.NewHandler(story, handler.WithTemplate(layout.StoryTmpl), handler.WithPathFunc(handler.PathFn))

	// Create a ServeMux to route our requests
	mux := http.NewServeMux()
	// This story handler is using a custom function and template
	// Because we use /story/ (trailing slash) all web requests
	// whose path has the /story/ prefix will be routed here.
	mux.Handle("/story/", h)
	// This story handler is using the default functions and templates
	// Because we use / (base path) all incoming requests not
	// mapped elsewhere will be sent here.
	mux.Handle("/", handler.NewHandler(story))
	// Start the server using our ServeMux
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
