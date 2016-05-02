package singleton

import "sync"

// singleton struct
type singleton struct {
}

// global instance
var instance *singleton
var once sync.Once

// Once guarentee the instance is created only once
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
