package inmemory

type Trie struct {
	dict map[byte]*Trie
	data    interface{}
}

func NewTrie() Trie {
	return Trie {
		dict: make(map[byte]*Trie),
		data: nil,
	}
}

func (t *Trie) Insert(key []byte, data interface{}) {
	c := key[0]
	next, ok := t.dict[c]
	if !ok || next == nil {
		newTrie := NewTrie()
		t.dict[c] = &newTrie
		next = &newTrie
	}
	if len(key) > 1 {
		next.Insert(key[1:], data)
	} else {
		next.data = data
	}
}

func (t Trie) Find(key []byte) *Trie {
	c := key[0]
	nextTrie, ok := t.dict[c]
	if !ok {
		return nil 
	}
	if len(key) > 1 {
		return nextTrie.Find(key[1:])
	} else {
		return nextTrie
	}
}

func (t Trie) Get(key []byte) interface{} {
	result := t.Find(key)
	if result != nil {
		return result.data
	}else{
		return nil
	}
}

func (t *Trie) Delete(key []byte) {
	c := key[0]
	next := t.dict[c]
	if next == nil {
		return
	}
	if len(key) > 2 {
		next.Delete(key[1:])
	} else {
		delete(next.dict, key[1])
	}
}

func (t Trie) Exists(key []byte) bool {
	result := t.Find(key)
	if result == nil {
		return false
	}
	return result.data != nil
}
