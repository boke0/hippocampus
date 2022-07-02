package inmemory

import (
	"encoding/json"
	"os"
	"path"
)

type InmemoryEngine[T any] struct {
	trie Trie[T]
}

func NewInmemoryEngine[T any]() InmemoryEngine[T] {
	return InmemoryEngine[T]{
		trie: NewTrie[T](),
	}
}

func (engine InmemoryEngine[T]) Get(key string) *T {
	return engine.trie.Get([]byte(key))
}

func (engine InmemoryEngine[T]) Set(key string, data T) {
	engine.trie.Insert([]byte(key), data)
}

func (engine InmemoryEngine[T]) Delete(key string) {
	engine.trie.Delete([]byte(key))
}

func (engine InmemoryEngine[T]) Exists(key string) bool {
	return engine.trie.Exists([]byte(key))
}

func (engine InmemoryEngine[T]) Keys() []string {
	keys := []string{}
	for _, byteKeys := range engine.trie.Keys() {
		keys = append(keys, string(byteKeys))
	}
	return keys
}

func (engine InmemoryEngine[T]) Export(filename string) {
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

func (engine *InmemoryEngine[T]) Import(filename string) {
	cwd, _ := os.Getwd()
	jsonBytes, err := os.ReadFile(path.Join(cwd, filename))
	if err != nil {
		panic(err)
	}
	dict := make(map[string]T)
	if err := json.Unmarshal(jsonBytes, &dict); err != nil {
		panic(err)
	}
	for key, data := range dict {
		engine.Set(key, data)
	}
}

func (engine *InmemoryEngine[T]) Clear() {
	engine.trie = NewTrie[T]()
}
