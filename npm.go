package recdn

import "fmt"

// 从Npm获取文件
func (a *App) Npm(pack string, version string, file string) (*File, error) {
	url := fmt.Sprintf("https://registry.npmmirror.com/%s/%s/files/%s", pack, version, file)
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}
