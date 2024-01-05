package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton

var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Singleton Instance"}
	})
	return instance
}

func Singleton() {	
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	fmt.Println("Are singleton1 and singleton2 the same instance?", singleton1 == singleton2)

	fmt.Println("Data from singleton1:", singleton1.data)
	fmt.Println("Data from singleton2:", singleton2.data)
}
