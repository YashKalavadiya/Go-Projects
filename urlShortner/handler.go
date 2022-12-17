package urlshortner

import (
	"net/http"
	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {

	
	pathUrls, err := parseYAML(yamlBytes)

	if err != nil {
		return nil, err
	}


	pathsToUrls := buildMap(pathUrls)
	
	return MapHandler(pathsToUrls, fallback), nil

}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for _, p := range pathUrls {
		pathToUrls[p.Path] = p.URL
	}
	return pathToUrls
}

func parseYAML(data []byte) ([]pathUrl, error) {

	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

type pathUrl struct {
	Path	string `yaml:"path"`
	URL	string `yasm:"url"`
}
