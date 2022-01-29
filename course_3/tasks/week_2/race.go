package main

import "fmt"

func printY() {
	for y:=10; y<20; y++ {
		fmt.Println(y)
	}
}

func main() {
	go printY()
	for x:=0; x<10; x++ {
		fmt.Println(x)
	}
}

/*
Race Condition
A design error in a multi-threaded system or application in which 
the operations of the system or application depend on the order of 
uncontrollable events.
It becomes a bug when some (or all) sequncies of operations are 
undesirable.

Why Race condition occurs?
It occurs when several processes can access and change the shared data 
at the same time. If one of the threads depends on this shared data, than 
its behaviour becomes uncotrollable and can lead to errors. Also, the error 
(resources conflict) might occur, if two threads are trying to write in 
the same file simulteneously.
*/
