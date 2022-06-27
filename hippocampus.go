package hippocampus

type Engine interface {
    Set(key string, data interface{})
    Get(key string) interface{}
    Delete(key string)
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

func (h Hippocampus) Get(key string, result interface{}) {
    result = h.engine.Get(key)
} 

func (h *Hippocampus) Delete(key string) {
    h.engine.Delete(key)
}

func (h Hippocampus) Exists(key string) bool {
    return h.engine.Exists(key)
}

func (h *Hippocampus) Fetch(key string, result interface{}, callback func() interface{}) {
    if h.Exists(key) {
        h.Get(key, result)
    }else{
        data := callback()
        h.Set(key, data)
        result = data
    }
}
