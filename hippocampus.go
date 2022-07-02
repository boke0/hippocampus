package hippocampus

type Engine[T any] interface {
	Set(key string, data T)
	Get(key string) *T
	Delete(key string)
	Keys() []string
	Exists(key string) bool
}

type Hippocampus[T any] struct {
	Engine Engine[T]
}

func NewHippocampus[T any](Engine Engine[T]) Hippocampus[T] {
	return Hippocampus[T]{
		Engine,
	}
}

func (h *Hippocampus[T]) Set(key string, data T) {
	h.Engine.Set(key, data)
}

func (h Hippocampus[T]) Get(key string) (*T, bool) {
	if h.Engine.Exists(key) {
		result := h.Engine.Get(key)
		return result, true
	} else {
		return nil, false
	}
}

func (h *Hippocampus[T]) Delete(key string) {
	h.Engine.Delete(key)
}

func (h Hippocampus[T]) Exists(key string) bool {
	return h.Engine.Exists(key)
}

func (h Hippocampus[T]) Keys() []string {
	return h.Engine.Keys()
}

func (h *Hippocampus[T]) Fetch(key string, callback func() T) *T {
	if h.Exists(key) {
		result, _ := h.Get(key)
		return result
	} else {
		data := callback()
		h.Set(key, data)
		result, _ := h.Get(key)
		return result
	}
}
