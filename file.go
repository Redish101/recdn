package recdn

import "encoding/json"

type File struct {
	// 文件类型
	Type string

	// 文件内容
	Content interface{}
}

func (f *File) String() (string, error) {
	res, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (f *File) FromString(data string) error {
	err := json.Unmarshal([]byte(data), f)
	if err != nil {
		return err
	}
	return nil
}
