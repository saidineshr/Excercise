package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var avg int
func studentRating(wg *sync.WaitGroup,mt *sync.Mutex){
	mt.Lock()
	var timetaken = int((rand.Intn(5)))
	time.Sleep(time.Duration(timetaken* int(time.Second)))

	m:=rand.Intn(10)
	avg+=m
	mt.Unlock()
	wg.Done()
}
func main() {
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(200)
	for i:=1;i<=200;i++{
		go studentRating(&wg,&m)
	}
	wg.Done()
	fmt.Println(float32(avg)/200.0)
}
