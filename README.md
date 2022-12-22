## Yeelight Golang

A simple Golang library to control Xiaomi Yeelights device over LAN with TCP.

This solution offers a 1:1 implementation of the [official docs from Xiaomi](http://www.yeelight.com/download/Yeelight_Inter-Operation_Spec.pdf).

### Installation
```bash
go get github.com/LordAur/yeelight
```

### Usage
```go
import "github.com/LordAur/yeelight"

func main() {
    y := yeelight.New(&yeelight.Config{
        IpAddress: "192.168.0.0",
        Port:      55443,
    })

    defer y.Close()

    r, err := y.GetProps("bright", "power", "ct")
    if err != nil {
        // ...
    }
}
```