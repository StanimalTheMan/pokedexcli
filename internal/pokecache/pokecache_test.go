package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "https://example.com",
			inputVal: []byte("testdata"),
		},
		{
			inputKey: "https://example.com/path",
			inputVal: []byte("moretestdata"),
		},
	}

	for i, cas := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(cas.inputKey, cas.inputVal)
			val, ok := cache.Get(cas.inputKey)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(cas.inputVal) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
