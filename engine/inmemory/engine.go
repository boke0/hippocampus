package inmemory

import (
	"encoding/json"
	"os"
	"path"
)

type InmemoryEngine struct {
    trie Trie
}

func NewInmemoryEngine() InmemoryEngine {
    return InmemoryEngine{
        trie: NewTrie(),
    }
}

func (engine InmemoryEngine) Get(key string) interface{} {
    return engine.trie.Get([]byte(key))
}

func (engine InmemoryEngine) Set(key string, data interface{}) {
    engine.trie.Insert([]byte(key), data)
}

func (engine InmemoryEngine) Delete(key string) {
    engine.trie.Delete([]byte(key))
}

func (engine InmemoryEngine) Exists(key string) bool {
    return engine.trie.Exists([]byte(key))
}

func (engine InmemoryEngine) Keys() []string {
    keys := []string{}
    for _, byteKeys := range engine.trie.Keys() {
        keys = append(keys, string(byteKeys))
    }
    return keys
}

func (engine InmemoryEngine) Export(filename string) {
    dict := make(map[string]interface{})
    for _, key := range engine.Keys() {
        dict[key] = engine.Get(key)
    }
    jsonBytes, err := json.Marshal(dict)
    if err != nil {
        panic(err)
    }
    cwd, _ := os.Getwd()
    if err := os.WriteFile(path.Join(cwd, filename), jsonBytes, 0o644); err != nil {
        panic(err)
    }
}

func (engine *InmemoryEngine) Import(filename string, parse func(bytes []byte) (map[string]interface{}, error)) {
    cwd, _ := os.Getwd()
    jsonBytes, err := os.ReadFile(path.Join(cwd, filename))
    if err != nil {
        panic(err)
    }
    dict, err := parse(jsonBytes)
    if err != nil {
        panic(err)
    }
    for key, data := range dict {
        engine.Set(key, data)
    }
}

func (engine *InmemoryEngine) Clear() {
    engine.trie = NewTrie()
}
