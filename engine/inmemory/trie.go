package inmemory

type Trie[T any] struct {
	dict map[byte]*Trie[T]
	data *T
}

func NewTrie[T any]() Trie[T] {
	return Trie[T]{
		dict: make(map[byte]*Trie[T]),
	}
}

func (t *Trie[T]) Insert(key []byte, data T) {
	c := key[0]
	next, ok := t.dict[c]
	if !ok || next == nil {
		newTrie := NewTrie[T]()
		t.dict[c] = &newTrie
		next = &newTrie
	}
	if len(key) > 1 {
		next.Insert(key[1:], data)
	} else {
		next.data = &data
	}
}

func (t Trie[T]) Find(key []byte) *Trie[T] {
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

func (t Trie[T]) Get(key []byte) *T {
	result := t.Find(key)
	if result != nil {
		return result.data
	} else {
		return nil
	}
}

func (t *Trie[T]) Delete(key []byte) {
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

func (t Trie[T]) Exists(key []byte) bool {
	result := t.Find(key)
	if result == nil {
		return false
	}
	return result.data != nil
}

func (t Trie[T]) Keys() [][]byte {
	keys := [][]byte{}
	for c, p := range t.dict {
		if p != nil {
			for _, k := range p.Keys() {
				k_ := []byte{c}
				k_ = append(k_, k...)
				keys = append(keys, k_)
			}
		}
	}
	if t.data != nil {
		keys = append(keys, []byte{})
	}
	return keys
}
