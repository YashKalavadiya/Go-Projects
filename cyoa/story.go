package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var tpl *template.Template


var defaultHandlerTemplate = `
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
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </body>
</html>`


//function option design pattern
//super complex to remember, need more practice with this one

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
    return func(h *handler) {
        if t != nil {
            h.t = t
        }
    }
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
    return func(h *handler) {
        h.pathFn = fn
    }
}

func init() { 
    tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {

    h := handler{s, tpl, defaultPathFn}
    for _, opt := range opts {
        opt(&h)
    }
    return h
}
type handler struct {
    s Story
    t *template.Template
    pathFn func(r *http.Request) string
}

func defaultPathFn(r *http.Request) string {
    path := strings.TrimSpace(r.URL.Path)
    if path == "" || path == "/" {
        path = "/intro"
    }
    return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    path := h.pathFn(r)

    if chapter, ok := h.s[path]; ok {
        err := h.t.Execute(w, chapter)

        if err != nil {

            log.Printf("%v", err)
            http.Error(w, "Something went wrong...", http.StatusInternalServerError)
        }
        return
    }
    http.Error(w, "Chapter Not Found", http.StatusNotFound)

}

func JsonStory(r io.Reader) (Story, error) {
    d := json.NewDecoder(r)
    var story Story
    if err := d.Decode(&story); err != nil {
        return nil, err
    }

    return story, nil
}

type Story map[string]Chapter

type Chapter struct {
    Title string `json:"title"`
    Paragraphs []string `json:"story"`
    Options []Option `json:"options"`
}

type Option struct {
    Text string `json:"text"`
    Chapter string `json:"arc"`
}



