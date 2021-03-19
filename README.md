# bucket
Bundle embedded databases with fixed api

```
go run cmd/main.go
```

used by 
```
import "github.com/ucwong/bucket"

...

badger := bucket.Badger()
bolt := bucket.Bolt()
```
