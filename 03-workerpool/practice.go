package workerpool

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gammazero/workerpool"
)

func Run() {
	start := time.Now()
	wp := workerpool.New(3)
	requests := []string{"P1", "P2", "P3", "P4", "P5", "P6", "P7", "P8"}

	for _, r := range requests {
		r := r
		wp.Submit(func() {
			second := rand.Intn(3)
			doTicketSync(second, r)
		})
	}
	wp.StopWait()
	fmt.Println("concurrentProcessWorkerPool ELAPSED UNTIL ALL DONE", time.Since(start))
}

func doTicketSync(secondNo int, data string) {
	fmt.Printf("Job Started PartnerID :%v,Started at:%v, Paused Seconds:%v \n", data, time.Now(), secondNo)
	time.Sleep(time.Duration(secondNo) * time.Second)
	fmt.Printf("Job Ended PartnerID :%v,Ended at:%v, Paused Seconds:%v \n", data, time.Now(), secondNo)
}
