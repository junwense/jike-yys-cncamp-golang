package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

/**
使用 cond 进行多生产者消费者
*/
func main() {
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	for i := 0; i < 2; i++ {
		go func(i int) {
			for {
				q.produce(i)
			}
		}(i)
	}

	for j := 0; j < 4; j++ {
		go func(j int) {
			for {
				q.consume(j)
			}
		}(j)
	}

	time.Sleep(time.Second * 30)
}

func (q *Queue) consume(i int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	s := fmt.Sprintf("Thread-%d", i)
	for len(q.queue) == 0 {
		fmt.Printf("%s consumer %s no data process, wait \n", time.Now().String(), s)
		//cond阻塞
		q.cond.Wait()
	}

	result := q.queue[0]
	q.queue = q.queue[1:]
	time.Sleep(time.Second)
	fmt.Printf("%s consumer %s has consume %s one data\n", time.Now().String(), s, result)
}

func (q *Queue) produce(i int) {
	s := fmt.Sprintf("Thread-%d", i)
	fmt.Printf("%s producer %s has produce one data\n", time.Now().String(), s)
	q.addData(s)
	q.cond.Broadcast()
	time.Sleep(time.Second * 2)
}

func (q *Queue) addData(s string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, s)
}
