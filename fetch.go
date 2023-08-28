package recdn

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Redish101/recdn/cache"
)

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
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)  
	var cookedData = &File{
		Type: resp.Header.Get("Content-Type"),
		Content: string(body),
	}
	return cookedData, nil
}