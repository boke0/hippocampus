package inmemory_test

import (
	"encoding/hex"
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

type TestData struct {
	D string `json:"d"`
}

func SeedEngine() (InmemoryEngine[TestData], map[string]TestData) {
	engine := NewInmemoryEngine[TestData]()
	dict := make(map[string]TestData)
	for i := 0; i < 1000; i++ {
		key := randomByte()
		t := TestData {
			D: hex.EncodeToString(key),
		}
		engine.Set(hex.EncodeToString(key), t)
		dict[hex.EncodeToString(key)] = t
	}
	return engine, dict
}

func TestEngineGet(t *testing.T) {
	engine, dict := SeedEngine()
	for key, value := range dict {
		t.Run(key, func(t *testing.T) {
			result := engine.Get(key)
			if result != nil {
				if result.D != value.D {
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

func TestEngineExportAndImport(t *testing.T) {
	engine, dict := SeedEngine()
	engine.Export("./test/exported.json")
	engine.Clear()
	engine.Import("./test/exported.json")
	for key := range dict {
		if !t.Run(key, func(t *testing.T) {
			result := engine.Exists(key)
			if !result {
				t.Errorf("not found")
			}
		}) {
			break
		}
	}
}
