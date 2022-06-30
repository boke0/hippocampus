package hippocampus

type Engine interface {
	Set(key string, data interface{})
	Get(key string) interface{}
	Delete(key string)
	Keys() []string
	Exists(key string) bool
}

type Hippocampus struct {
	engine Engine
}

func NewHippocampus(engine Engine) Hippocampus {
	return Hippocampus{
		engine,
	}
}

func (h *Hippocampus) Set(key string, data interface{}) {
	h.engine.Set(key, data)
}

func (h Hippocampus) Get(key string) (interface{}, bool) {
	if h.engine.Exists(key) {
		result := h.engine.Get(key)
		return result, true
	} else {
		return nil, false
	}
}

func (h *Hippocampus) Delete(key string) {
	h.engine.Delete(key)
}

func (h Hippocampus) Exists(key string) bool {
	return h.engine.Exists(key)
}

func (h Hippocampus) Keys() []string {
	return h.engine.Keys()
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
