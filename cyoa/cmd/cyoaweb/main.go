package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

    //This line on uncommenting will handle CYOA on path prefixed with /story
    // tmpl := template.Must(template.New("").Parse(withPrefixPathTemplate))
    //h := cyoa.NewHandler(story, cyoa.WithTemplate(tmpl), cyoa.WithPathFunc(pathFn))
    
    h := cyoa.NewHandler(story, cyoa.WithTemplate(nil))

    fmt.Printf("starting server on the port %d\n", *port)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

var withPrefixPathTemplate = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Chooose your Own Adventure</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link href="css/style.css" rel="stylesheet">
    </head>
    <body>

        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}

        <ul>
            {{range .Options}}
            <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </body>
</html>`

//Just to try out Functional Options Design Pattern
func pathFn(r *http.Request) string {
    path := strings.TrimSpace(r.URL.Path)

    if path == "/story" || path == "/story/" {
        path = "/story/intro"
    }
    return path[len("/story/"):] 
}
