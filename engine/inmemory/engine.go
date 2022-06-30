package inmemory

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

