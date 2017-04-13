Example:

```
import (
    "github.com/pleamon/gopool"
    "log"
)

func callback(counter string) {
    log.Println(counter)
}

func main() {
    worker := gopool.Worker{}

    // 设置callback与协程池大小
    worker.Init(callback, 10)

    for i := 0; i < 10000; i++ {
        worker.Push("127.0.0.1")
    }

    worker.Start()
    worker.Wait()
}
```
