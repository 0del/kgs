# kgs - Key generate service

Generation key support for shorten url, unique key.

### Usage 
#### 1. Go code
```go
package main

import (
	"fmt"
	"github.com/9bany/kgs/pkg/key0"
)

func main() {
	gen := key0.NewGen(key0.WithLen(int8(len)))
	k, _ := gen.New()
    fmt.Println(k)
}

```
#### 2. Docker

1. Build
```sh
docker build . -t kgs
```
2. Run
```sh
docker run -p 8080:8080 kgs:latest
```

Thanks !