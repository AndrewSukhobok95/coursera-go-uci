package main

import (
	"fmt"
	"sync"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct{
	num int
	leftCS, rightCS *ChopStick
}

func (p Philosopher) eat(host chan int, wg *sync.WaitGroup) {
	for i:=0; i<3; i++ {

		host <- p.num

		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat %d\n", p.num)
		fmt.Printf("finishing eating %d\n", p.num)
		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<- host
	}
	wg.Done()
}

func InitPhilosophers() []*Philosopher {
	sticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, sticks[i], sticks[(i+1)%5]}
	}

	return philosophers
}

func main() {
	var wg sync.WaitGroup
	host := make(chan int, 2)
	philosophers := InitPhilosophers()
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat(host, &wg)
	}

	wg.Wait()
}

