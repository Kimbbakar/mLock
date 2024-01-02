package mlock

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	cleanInterval = 30 * time.Minute
	once          sync.Once
)

var mLock multiLock

type multiLock struct {
	cLck sync.RWMutex
	l    sync.Map
}

type lockDetails struct {
	lck *sync.Mutex
	c   int64
}

func KeepClean(intervalInMinute *time.Duration) {
	once.Do(func() {
		if intervalInMinute != nil {
			cleanInterval = *intervalInMinute
		}

		go func() {
			defer log.Println("I'm closing")
			ticker := time.NewTicker(cleanInterval)
			for {
				<-ticker.C
				func() {
					mLock.cLck.Lock()
					defer mLock.cLck.Unlock()

					mLock.l.Range(func(key, value any) bool {
						ld := value.(*lockDetails)
						if ld.c == 0 {
							mLock.l.Delete(key)
						}
						return true
					})
				}()
			}
		}()
	})
}

func Lock(keys ...interface{}) {
	mLock.cLck.RLock()
	defer mLock.cLck.RUnlock()

	uKey := getKey(keys)
	lDetails := getOrStoreLock(uKey)
	atomic.AddInt64(&lDetails.c, 1)
	lDetails.lck.Lock()
}

func UnLock(keys ...interface{}) {
	uKey := getKey(keys)
	lDetails := getLock(uKey)
	if lDetails == nil {
		return
	}
	atomic.AddInt64(&lDetails.c, -1)
	lDetails.lck.Unlock()
}

func getOrStoreLock(key string) *lockDetails {
	ld, _ := mLock.l.LoadOrStore(key, &lockDetails{
		lck: &sync.Mutex{},
	})
	return ld.(*lockDetails)
}

func getLock(key string) *lockDetails {
	ld, ok := mLock.l.Load(key)
	if !ok {
		return nil
	}
	return ld.(*lockDetails)
}

func getKey(keys ...interface{}) string {
	if len(keys) == 0 {
		panic("mLock key necessary")
	}

	k := []string{}
	for _, key := range keys {
		k = append(k, fmt.Sprint(key))
	}

	return strings.Join(k, "_")
}
