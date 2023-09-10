package recdn

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Redish101/recdn/cache"
	"github.com/evanw/esbuild/pkg/api"
)

func RemoveMin(input string) string {
	start := strings.Index(input, ".min")
	end := strings.Index(input[start+1:], ".min") + start + 1
	if start != -1 && end != -1 {
		result := input[:start] + input[end:]
		return result
	}
	return input
}

func Fetch(path string) (*File, error) {
	cachedData, err := cache.Get(path)
	if err != nil {
		return nil, err
	}
	if cachedData != "" {
		var cookedData *File
		err := json.Unmarshal([]byte(cachedData), cookedData)
		if err != nil {
			return nil, err
		}
		return cookedData, nil
	}
	resp, err := http.Get(path)
	if strings.HasSuffix(path, ".min.js") {
		path = strings.ReplaceAll(path, ".min", "")
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		content := Min(string(body), api.LoaderJS)
		var cookedData = &File{
			Type:    resp.Header.Get("Content-Type"),
			Content: string(content),
		}
		return cookedData, nil
	}
	if strings.HasSuffix(path, ".min.css") {
		path = strings.ReplaceAll(path, ".min", "")
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		content := Min(string(body), api.LoaderCSS)
		var cookedData = &File{
			Type:    resp.Header.Get("Content-Type"),
			Content: string(content),
		}
		return cookedData, nil
	}
	if strings.HasSuffix(path, ".min.json") {
		path = strings.ReplaceAll(path, ".min", "")
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		content := Min(string(body), api.LoaderJSON)
		var cookedData = &File{
			Type:    resp.Header.Get("Content-Type"),
			Content: string(content),
		}
		return cookedData, nil
	}
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	var cookedData = &File{
		Type:    resp.Header.Get("Content-Type"),
		Content: string(body),
	}
	return cookedData, nil
}
