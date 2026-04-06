package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type File struct {
	name string
}

func (f *File) ReadFiles() ([]byte, error) {
	data, err := os.ReadFile(f.name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *File) WriteFiles(content []byte) {

	file, err := os.Create(f.name)

	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	fmt.Println("Запись успешна")

}

func NewFile(name string) (*File, error) {
	isValidName := checkIsJson(name)
	if !isValidName {
		return nil, errors.New("формат файла должен быть .json")
	}
	return &File{name}, nil
}

func checkIsJson(name string) bool {
	return strings.HasSuffix(name, ".json")
}
