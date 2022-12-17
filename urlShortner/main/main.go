package main

import (
	"github.com/yashkalavadiya/urlshortner"
	"fmt"
	"net/http"
)



func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/my-test": "https://www.google.com",
		"/my-2-test": "https://www.github.com",
	}

	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://www.google.com
- path: /anothershort
  url: https://www.github.com/yashkalavadiya
`
	yamlHandler, err := urlshortner.YAMLHandler([]byte(yaml), mapHandler)

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", yamlHandler)
}

func defaultMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola amigo!! como estas")
}
