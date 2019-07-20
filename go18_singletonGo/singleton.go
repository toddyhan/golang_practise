package single

import (
	"fmt"
	"sync"
)

/**
* leetcode no: singleton pattern implemented by golang
*
* @description:
*
* @author: toddyhan
*
* @create: 2019/05/24 17:46
**/

type Student struct {

	name string
	number int
}
var lock *sync.Mutex
var t *Student

func GetInstance() *Student {
	if t == nil {
		lock.Lock()
		defer lock.Unlock()
		if t == nil {
			t = &Student{}
		}
	}
	return t
}

var once sync.Once
func GetInstanceOnce() *Student {
	once.Do(func() {
		t = &Student{}
	})
	return t
}

func (t Student) Learn() {
	fmt.Println("I am learning!")
}


