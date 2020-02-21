# asuleymanov/golos-go

[![GoDoc](https://godoc.org/github.com/asuleymanov/golos-go?status.svg)](https://godoc.org/github.com/asuleymanov/golos-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/asuleymanov/golos-go)](https://goreportcard.com/report/github.com/asuleymanov/golos-go)

Golang RPC client library for [Golos](https://golos.io).

## Usage

```go
import "github.com/asuleymanov/golos-go"
```


## Example

This is just a code snippet. Please check the `examples` directory
for more complete and ready to use examples.

```go
package main

import (
	"fmt"

	"github.com/asuleymanov/golos-go"
)

func main() {
	cls, _ := golos.NewClient("wss://golos.lexa.host/ws")
	defer cls.Close()

  fmt.Println("Config --> ")
  fmt.Printf("%#v \n",cls.Config)
  fmt.Println("")
  fmt.Println("")
  
  fmt.Println("DatabaseInfo --> ")
  ans,err:= cls.API.GetDatabaseInfo()
  fmt.Printf("%#v \n Error : %s\n",ans,err)
  fmt.Println("")
  fmt.Println("")
  
  fmt.Println("ChainProperties --> ")
  ans1,err1:= cls.API.GetChainProperties()
  fmt.Printf("%#v \n Error : %s\n",ans1,err1)
  fmt.Println("")
  fmt.Println("")
}
```

## Package Organisation


You need to create a `Client` object to be able to do anything.
Then you just need to call `NewClient()`.


## License

MIT, see the `LICENSE` file.
