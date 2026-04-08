package storage

import (
	"demo/CLI/bins"
	"encoding/json"
	"fmt"
	"time"
)

type Db interface {
	ReadFiles() ([]byte, error)
	WriteFiles([]byte)
}

type Storage struct {
	List bins.BinList `json:"list"`
}

type StorageWithDb struct {
	Storage
	Db Db
}

func NewStorage(db Db) *StorageWithDb {
	file, err := db.ReadFiles()
	if err != nil {
		return &StorageWithDb{
			Storage: Storage{
				List: bins.BinList{},
			},
			Db: db,
		}
	}
	var vault Storage
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println(err)
	}
	return &StorageWithDb{
		Storage: vault,
		Db:      db,
	}

}

func (storage *Storage) ToBytes() ([]byte, error) {
	file, err := json.Marshal(storage)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (storageWithDb *StorageWithDb) ReadStorage() (*Storage, error) {
	file, err := storageWithDb.Db.ReadFiles()
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

func (storageWithDb *StorageWithDb) WriteStorage() {
	bin1 := bins.NewBin("1", true, time.Now(), "Первый")
	bin2 := bins.NewBin("2", false, time.Now(), "Второй")
	binList := bins.BinList{*bin1, *bin2}
	storage := Storage{binList}
	bytes, err := storage.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	storageWithDb.Db.WriteFiles(bytes)
}
