package api

import (
	config2 "demo/CLI/config"
	"fmt"
)

func init() {
	config := config2.NewConfig()
	envKeyValue := config.Key
	fmt.Println(envKeyValue)
}
