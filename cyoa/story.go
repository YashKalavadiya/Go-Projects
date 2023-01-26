package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
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

func init() { 
    tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

func NewHandler(s Story) http.Handler {
    return handler{s}
}
type handler struct {
    s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := tpl.Execute(w, h.s["intro"])
    if err != nil {
        panic(err)
    }

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



