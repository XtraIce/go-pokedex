package lcache_test

import (
	"testing"
	"time"

	"github.com/xtraice/pokedexcli/internal/lcache"
)

// TestCache_Add tests the Add method of the Cache struct.
func TestCache_Add(t *testing.T) {
	cache := lcache.NewCache(time.Second)

	key := "key"
	val := []byte("value")

	cache.Add(key, val)

	// Verify that the value was added to the cache
	if _, ok := cache.Get(key); !ok {
		t.Errorf("expected value to be added to the cache")
	}
}

// TestCache_Get tests the Get method of the Cache struct.
func TestCache_Get(t *testing.T) {
	cache := lcache.NewCache(time.Second)

	key := "key"
	val := []byte("value")

	// Add a value to the cache
	cache.Add(key, val)

	// Retrieve the value from the cache
	result, ok := cache.Get(key)

	// Verify that the value was retrieved successfully
	if !ok {
		t.Errorf("expected value to be retrieved from the cache")
	}

	// Verify that the retrieved value matches the original value
	if string(result) != string(val) {
		t.Errorf("expected retrieved value to match original value")
	}
}

// TestCache_Reap tests the Reap method of the Cache struct.
func TestCache_Reap(t *testing.T) {
	cache := lcache.NewCache(time.Second)

	key := "key"
	val := []byte("value")

	// Add a value to the cache
	cache.Add(key, val)

	// Wait for the cache entry to expire
	time.Sleep(2 * time.Second)

	// Retrieve the expired value from the cache
	_, ok := cache.Get(key)

	// Verify that the expired value was reaped from the cache
	if ok {
		t.Errorf("expected expired value to be reaped from the cache")
	}
}
