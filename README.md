# wifi-name

Get current wifi name

### Install

```
$ go get github.com/yelinaung/wifi-name
```

### Example

```go
package main

import (
	"fmt"
	"github.com/yelinaung/wifi-name"
)

func main() {
	fmt.Println("wifi name ", wifiname.WifiName())
}
```

### License

MIT

