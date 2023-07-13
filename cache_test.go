package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   interface{}
		ttl     time.Duration
		wantErr bool
	}{
		{"Test 1", "key1", "value1", 1 * time.Second, false},
		{"Test 2", "key2", 12345, 1 * time.Second, false},
		{"Test 3", "key3", []string{"value", "test"}, 1 * time.Second, false},
	}

	cache := NewCache()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.Set(tt.key, tt.value, tt.ttl)

			got, found := cache.Get(tt.key)

			if found != !tt.wantErr {
				t.Errorf("Get() found = %v, wantFound %v", found, !tt.wantErr)
				return
			}

			if found && !reflect.DeepEqual(got, tt.value) {
				t.Errorf("Get() = %v, want %v", got, tt.value)
			}

			time.Sleep(2 * time.Second)

			_, found = cache.Get(tt.key)
			if found {
				t.Errorf("Get() expected cache miss, found = %v, want %v", found, false)
			}
		})
	}
}

func BenchmarkCache(b *testing.B) {
	cache := NewCache()
	cache.Set("key", "value", 1*time.Second)
	for i := 0; i < b.N; i++ {
		cache.Get("key")
	}
}
