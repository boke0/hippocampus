package hippocampus_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/boke0/hippocampus"
	. "github.com/boke0/hippocampus/engine/inmemory"
)

func randomByte() []byte {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(25) + 6
	b := make([]byte, num)
	rand.Read(b)
	return b
}

func SeedHippocampus() (Hippocampus[[]byte], map[string][]byte) {
	hippocampus := NewHippocampus[[]byte](NewInmemoryEngine[[]byte]())
	dict := make(map[string][]byte)
	for i := 0; i < 1000; i++ {
		key := randomByte()
		hippocampus.Set(string(key), key)
		dict[string(key)] = key
	}
	return hippocampus, dict
}

func TestHippocampusGet(t *testing.T) {
	hippocampus, dict := SeedHippocampus()
	for key, value := range dict {
		t.Run(key, func(t *testing.T) {
			result, _ := hippocampus.Get(key)
			if string(*result) != string(value) {
				t.Errorf("not matched")
			}
		})
	}
}

func TestHippocampusExists(t *testing.T) {
	hippocampus, dict := SeedHippocampus()
	for key := range dict {
		t.Run(key, func(t *testing.T) {
			if !hippocampus.Exists(key) {
				t.Errorf("not found")
			}
		})
	}
}

func TestHippocampusDelete(t *testing.T) {
	hippocampus, dict := SeedHippocampus()
	for key := range dict {
		t.Run(key, func(t *testing.T) {
			hippocampus.Delete(key)
			result := hippocampus.Exists(key)
			if result {
				t.Errorf("found")
			}
		})
	}
}
