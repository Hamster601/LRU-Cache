package LRU_Cache

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_CLRUCache(t *testing.T) {
	obj := NewCache(2)
	time.Sleep(150 * time.Millisecond)
	obj.Put("1", 1)
	time.Sleep(150 * time.Millisecond)
	obj.Put("2", 2)
	time.Sleep(150 * time.Millisecond)
	param1 := obj.Get("1")
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, 1, param1)
	obj.Put("3", 3)
	time.Sleep(150 * time.Millisecond)
	param1 = obj.Get("2")
	assert.Equal(t, nil, param1)
	obj.Put("4", 4)
	time.Sleep(150 * time.Millisecond)
	param1 = obj.Get("1")
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, nil, param1)
	param1 = obj.Get("3")
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, 3, param1)
	param1 = obj.Get("4")
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, 4, param1)
}

func BenchmarkGetAndPut1(b *testing.B) {
	b.ResetTimer()
	obj := NewCache(128)
	wg := sync.WaitGroup{}
	wg.Add(b.N * 2)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			obj.Get(strconv.Itoa(rand.Intn(200)))
		}()
		go func() {
			defer wg.Done()
			obj.Put(strconv.Itoa(rand.Intn(200)), strconv.Itoa(rand.Intn(200)))
		}()
	}
	wg.Wait()
}
