# Golang-kv
Bundle embedded databases with fixed api

```
go run cmd/main.go
```

used by 
```
import "github.com/ri0day/golang-kv"


...

bolt := kv.Bolt("")
defer bolt.Close()

bolt.setTTL([]byte("k"), []byte("v"), time.Second)

...

