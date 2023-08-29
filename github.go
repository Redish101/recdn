package recdn

import "fmt"

// GitHub回源配置
type GitHubConfig struct {
	Raw string
}

// 从GitHub获取文件
func (a *App) GitHub(repo string, branch string, file string) (*File, error) {
	url := fmt.Sprintf(a.Config.GitHub.Raw, repo, branch, file)
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}
