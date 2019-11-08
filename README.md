# gopool

## 结构

```txt
                         ┌ Worker
                         ├ ...
Task --> Job Channel --> ├ Worker
                         ├ ...
                         └ Worker
```

## 示例

```golang
package main
import (
    "fmt"
    "time"

    "github.com/boxungo/gopool"
)

func main() {
    poolOpt := gopool.Option{
        MaxWorkers: 2,
        MaxJobs: 0,
    }
    pool := gopool.NewPool(option)

    for i := 0; i < 10; i++ {
        out := fmt.Sprintf("Job: %d", i)
        job := gopool.NewJob(func() {
            fmt.Println(out)
        })
        pool.Submit(job)
    }

    time.Sleep(time.Second * 10)
}
```
