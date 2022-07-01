package hippocampus

type Engine interface {
	Set(key string, data interface{})
	Get(key string) interface{}
	Delete(key string)
	Keys() []string
	Exists(key string) bool
}

type Hippocampus struct {
	Engine Engine
}

func NewHippocampus(Engine Engine) Hippocampus {
	return Hippocampus{
		Engine,
	}
}

func (h *Hippocampus) Set(key string, data interface{}) {
	h.Engine.Set(key, data)
}

func (h Hippocampus) Get(key string) (interface{}, bool) {
	if h.Engine.Exists(key) {
		result := h.Engine.Get(key)
		return result, true
	} else {
		return nil, false
	}
}

func (h *Hippocampus) Delete(key string) {
	h.Engine.Delete(key)
}

func (h Hippocampus) Exists(key string) bool {
	return h.Engine.Exists(key)
}

func (h Hippocampus) Keys() []string {
	return h.Engine.Keys()
}

func (h *Hippocampus) Fetch(key string, callback func() interface{}) interface{} {
	if h.Exists(key) {
		result, _ := h.Get(key)
		return result
	} else {
		data := callback()
		h.Set(key, data)
		return data
	}
}
