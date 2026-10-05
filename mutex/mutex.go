package main

import (
	"fmt"
	"sync"
)

type post struct{
	views int
	mu sync.Mutex
}

func increment(myPost *post, wg *sync.WaitGroup){
	

	for i:=1; i<=100 ;i++{
		myPost.views+=1
	}

	wg.Wait()
}

func main() {
	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}
	wg.Add(1)
	go increment(&myPost,&wg)
	wg.Add(-1)
	fmt.Println(myPost.views)
}