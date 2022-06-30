package inmemory_test

import (
	"testing"

	. "github.com/boke0/hippocampus/engine/inmemory"
)

func SeedTrie() (Trie, map[string][]byte) {
    trie := NewTrie()
    dict := make(map[string][]byte)
    for i := 0; i<1000; i++ {
        key := randomByte()
        trie.Insert(key, key)
        dict[string(key)] = key
    }
    return trie, dict
}

func TestTrieGet(t *testing.T) {
    trie, dict := SeedTrie()
    for key, value := range dict {
        t.Run(key, func(t *testing.T) {
            result := trie.Get([]byte(key))
            if result != nil {
                resultBytes := result.([]byte)
                if string(resultBytes) != string(value) {
                    t.Errorf("not matched")
                }
            }else{
                t.Errorf("not found")
            }
        })
    }
}

func TestTrieExists(t *testing.T) {
    trie, dict := SeedTrie()
    for key := range dict {
        t.Run(key, func(t *testing.T) {
            if !trie.Exists([]byte(key)) {
                t.Errorf("not found")
            }
        })
    }
}

func TestTrieDelete(t *testing.T) {
    trie, dict := SeedTrie()
    for key := range dict {
        t.Run(key, func(t *testing.T) {
            trie.Delete([]byte(key))
            result := trie.Exists([]byte(key))
            if result {
                t.Errorf("found")
            }
        })
    }
}

func TestTrieKeys(t *testing.T) {
    trie, _ := SeedTrie()
    keys := trie.Keys()
    for _, key := range keys {
        if !trie.Exists(key) {
            t.Errorf("invalid key")
        }
    }
}
