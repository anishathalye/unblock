package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"sync"
)

const chunkSize = 4096

var chunks list.List
var done bool
var mu sync.Mutex
var cond = sync.NewCond(&mu)

func main() {
	go reader()
	writer()
}

func reader() {
	for {
		buf := make([]byte, chunkSize)
		n, err := os.Stdin.Read(buf)
		if err == io.EOF {
			mu.Lock()
			done = true
			cond.Signal()
			mu.Unlock()
			return
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		mu.Lock()
		chunks.PushBack(buf[:n])
		cond.Signal()
		mu.Unlock()
	}
}

func writer() {
	for {
		mu.Lock()
		for !done && chunks.Len() == 0 {
			cond.Wait()
		}
		if chunks.Len() == 0 && done {
			mu.Unlock()
			return
		}
		front := chunks.Front()
		buf := front.Value.([]byte)
		chunks.Remove(front)
		mu.Unlock()
		_, err := os.Stdout.Write(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
	}
}
