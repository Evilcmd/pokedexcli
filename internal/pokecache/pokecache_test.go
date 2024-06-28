package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(time.Microsecond * 10)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	cache := NewCache(time.Microsecond * 10)

	cases := []struct {
		inputKey     string
		ExcpectedVal []byte
	}{
		{"key1", []byte("val1")},
		{"key2", []byte("val2")},
		{"key3", []byte("val3")},
		{"key4", []byte("val4")},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.ExcpectedVal)
		val, ok := cache.Get(cas.inputKey)

		if !ok {
			t.Error("key1 not found")
		}

		if string(val) != string(cas.ExcpectedVal) {
			t.Error("values dont match")
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	k, v := "key1", []byte("val1")

	cache.Add(k, v)

	time.Sleep(interval * 2)

	_, ok := cache.Get(k)

	if ok {
		t.Error("key not reaped")
	}
}

func TestNotReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	k, v := "key1", []byte("val1")

	cache.Add(k, v)

	time.Sleep(interval / 2)

	_, ok := cache.Get(k)

	if !ok {
		t.Error("key should not reaped")
	}
}
