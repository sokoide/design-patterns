package adapter

import (
	"fmt"
	"singleton-example/domain"
	"sync"
)

type singletonDatabase struct{}

func (db *singletonDatabase) Connect() {
	fmt.Println("Connected to the database instance.")
}

var instance *singletonDatabase
var once sync.Once

func GetDatabaseInstance() domain.Database {
	once.Do(func() {
		fmt.Println("Creating single database instance now.")
		instance = &singletonDatabase{}
	})
	return instance
}
