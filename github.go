package recdn

import "fmt"

type GitHubConfig struct {
	Raw string
}

func (a *App) GitHub(repo string, branch string, file string) (*File, error) {
	url := fmt.Sprintf(a.Config.GitHub.Raw, repo, branch, file)
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}
