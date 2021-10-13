package main

import (
	"fmt"
	ants "github.com/panjf2000/ants/v2"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main(){
	//err := cmd.Execute()
	//if err != nil{
	//	log.Fatalf("cmd Execute err: %v",err)
	//}
	defer ants.Release()
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	for i := 0;i < 1000; i++{
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	log.Printf("running goroutines: %d\n", ants.Running())
	log.Printf("finish all tasks.\n")
	log.Printf("hello world")
	ch1 := make(chan string,10000)
	p,_ := ants.NewPool(10)
	defer p.Release()
	for i := 0;i < 1000; i++{
		wg.Add(1)
		_ = p.Submit(func(){
			ch1 <- "f"
			log.Printf("hello submit %d",i)
			defer wg.Done()
		})
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)

}

var sum int32

func myFunc(i interface{}){
	n := i.(int32)
	atomic.AddInt32(&sum,n)
	log.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

