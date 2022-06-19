# hippocampus
Low-level cache for Golang.

```go
import (
    . "github.com/boke0/hippocampus"
    . "github.com/boke0/hippocampus/engine"
)

type Foo struct {
    aaa string
}

func main() {
    hippocampus := NewHippocampus(NewInmemoryEngine())

    // Set value
    hippocampus.Set("hogehoge", Foo {
        aaa string
    })

    // Get value
    var foo Foo
    hippocampus.Get("hogehoge", &foo)
}
```

## Usage

### Initialize

```go
engine := NewInmemoryEngine()
hippocampus := NewHippocampus(engine)
```

NewHippocampus expects an argument that implements following interface.

```go
type Engine interface {
    Set(key string, data interface{})
    Get(key string) interface{}
    Delete(key string)
    Exists(key string) bool
}
```

### Set value

```go
hippocampus.Set("hogehoge", Foo {
    aaa: "hogehoge"
})
```

### Get value

```go
var foo Foo
hippocampus.Get("hogehoge", &Foo)
```

### Check if value exists

```go
if hippocampus.Exists("hogehoge") {
    fmt.Println("Exists!")
}
```

### Get value and set if not exists

```go
hippocampus.Fetch("hogehoge", func() {
    return Foo {
        aaa: "new hogehoge"
    }
})
```

### Delete value

```go
hippocampus.Delete("hogehoge")
```

