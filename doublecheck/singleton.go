package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton
var mutex sync.Mutex

func GetInstance() *Singleton {
	mutex.Lock()
	if singleton == nil {
		singleton = new(Singleton)
	}
	mutex.Unlock()
	return singleton
}
