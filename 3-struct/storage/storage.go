package storage

import (
	"demo/CLI/bins"
	"demo/CLI/file"
	"encoding/json"
	"fmt"
	"time"
)

type Storage struct {
	List bins.BinList `json:"list"`
}

func (storage *Storage) ToBytes() ([]byte, error) {
	file, err := json.Marshal(storage)
	if err != nil {
		return nil, err
	}
	return file, nil
}

var LocalBinsStorageFileName = "bins.json"

func ReadStorage() (*Storage, error) {
	file, err := file.ReadFiles(LocalBinsStorageFileName)
	if err != nil {
		return nil, err
	}
	var storage Storage
	err = json.Unmarshal(file, &storage)
	if err != nil {
		fmt.Println(err)
	}
	return &storage, nil
}

func WriteStorage() {
	bin1 := bins.NewBin("1", true, time.Now(), "Первый")
	bin2 := bins.NewBin("2", false, time.Now(), "Второй")
	binList := bins.BinList{*bin1, *bin2}
	storage := Storage{binList}
	bytes, err := storage.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	file.WriteFiles(bytes, LocalBinsStorageFileName)
}
