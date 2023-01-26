package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {


    port := flag.Int("port", 3000, "the port to start the CYOA web application")

    filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")

    flag.Parse()

    fmt.Printf("Using the story in %s.\n", *filename)

    f, err := os.Open(*filename)
    if err != nil {
        panic(err)
    }
    story, err := cyoa.JsonStory(f)


    if err != nil {
        panic(err)
    }

    h := cyoa.NewHandler(story)

    fmt.Printf("starting server on the port %d\n", *port)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}