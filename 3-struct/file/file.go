package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFiles(name string) ([]byte, error) {
	isJson := checkIsJson(name)
	if !isJson {
		return nil, errors.New("формат файла не соответствует json")
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFiles(content []byte, name string) {

	file, err := os.Create(name)

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

func checkIsJson(name string) bool {
	return strings.HasSuffix(name, ".json")
}
