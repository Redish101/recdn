package recdn

import "encoding/json"

// 文件
type File struct {
	// 文件类型
	Type string

	// 文件内容
	Content interface{}
}

// 将文件转换为字符串
func (f *File) String() (string, error) {
	res, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// 从字符串加载文件
func (f *File) FromString(data string) error {
	err := json.Unmarshal([]byte(data), f)
	if err != nil {
		return err
	}
	return nil
}
