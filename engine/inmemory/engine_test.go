package inmemory_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/boke0/hippocampus/engine/inmemory"
)

func randomByte() []byte {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(25) + 6
	b := make([]byte, num)
	rand.Read(b)
	return b
}

func SeedEngine() (InmemoryEngine, map[string][]byte) {
	engine := NewInmemoryEngine()
	dict := make(map[string][]byte)
	for i := 0; i < 1000; i++ {
		key := randomByte()
		engine.Set(string(key), key)
		dict[string(key)] = key
	}
	return engine, dict
}

func TestEngineGet(t *testing.T) {
	engine, dict := SeedEngine()
	for key, value := range dict {
		t.Run(key, func(t *testing.T) {
			result := engine.Get(key)
			if result != nil {
				resultBytes := result.([]byte)
				if string(resultBytes) != string(value) {
					t.Errorf("not matched")
				}
			} else {
				t.Errorf("not found")
			}
		})
	}
}

func TestEngineExists(t *testing.T) {
	engine, dict := SeedEngine()
	for key := range dict {
		t.Run(key, func(t *testing.T) {
			if !engine.Exists(key) {
				t.Errorf("not found")
			}
		})
	}
}

func TestEngineDelete(t *testing.T) {
	engine, dict := SeedEngine()
	for key := range dict {
		t.Run(key, func(t *testing.T) {
			engine.Delete(key)
			result := engine.Exists(key)
			if result {
				t.Errorf("found")
			}
		})
	}
}
