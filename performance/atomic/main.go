/*
* package main starts a http server that counts nums of the visitors.
*
* It gives an example of how atomic operation
*/
package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

var visitCount atomic.Uint64
var visitCountWithMutex uint64
var mu sync.Mutex

func IncVisitCount() uint64 {
	return visitCount.Add(1)
}

func IncVisitCountWithMutex() uint64 {
	mu.Lock()
	defer mu.Unlock()
	visitCountWithMutex++
	return visitCountWithMutex
}

func VisitCountHandler(w http.ResponseWriter, r *http.Request) {
	num := IncVisitCount()
	//num := IncVisitCountWithMutex()
	fmt.Fprintf(w, "There have been %d visitors now\n", num)
}

func main() {
	http.HandleFunc("/visitcount", VisitCountHandler)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
