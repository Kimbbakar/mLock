package mlock_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/Kimbbakar/mlock"
)

var wg sync.WaitGroup
var count [10000]int

func TestYourFunction(t *testing.T) {
	testCases := []struct {
		input1 int
		input2 int
		output int
	}{
		{1, 10, 10},
		{10, 100, 100},
		{100, 1000, 1000},
		{1000, 10000, 10000},
		{1000, 100000, 100000},
	}

	d := time.Second * 2
	mlock.KeepClean(&d)

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input_%d_%d", tc.input1, tc.input2), func(t *testing.T) {
			testFunc(tc.input1, tc.input2)
			for i := 0; i < tc.input1; i++ {
				if count[i] != tc.output {
					t.Errorf("Expected %d, got %d", tc.output, count[i])
				}
			}
		})
	}
}

func testFunc(lockCount, concurrentReq int) {
	count = [10000]int{}
	for j := 0; j < concurrentReq; j++ {
		for i := 0; i < lockCount; i++ {
			wg.Add(1)
			go jobFunc(i, j)
		}
	}

	wg.Wait()
}

func jobFunc(i, j int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
	mlock.Lock(i)
	defer mlock.UnLock(i)
	count[i] += 1
	wg.Done()
}
