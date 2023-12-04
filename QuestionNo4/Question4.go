package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(6)
	var muex sync.Mutex
	mreaders := 13
	nwriters := 5
	sharedbuffer := []int{10, 10, 10, 10, 10}

	reader := func(maxCounter int) {

		for i := 0; i < maxCounter; i++ {
			go func(Counter int) {
				muex.Lock()
				defer muex.Unlock()

				time.Sleep(time.Millisecond * 5)
				fmt.Printf("max counter %d ", Counter)
				fmt.Println(sharedbuffer)
			}(i)
		}
	}
	writer := func(maxCounter int) {
		for i := 0; i < maxCounter; i++ {
			go func(i int) {
				muex.Lock()
				defer muex.Unlock()
				for j := 0; j < len(sharedbuffer); j++ {

					time.Sleep(time.Millisecond * 5)
					sharedbuffer[j] = i

				}
			}(i)
		}
	}

	writer(nwriters)
	reader(mreaders)
	time.Sleep(time.Second * 10)
}
