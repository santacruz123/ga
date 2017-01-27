# ga
Google Analytics custom report extractor

```go
package main

import (
  "log"
  "os"

  "gopkg.in/santacruz123/ga.v1"
)

func main() {
  req := ga.New(os.Getenv("ACCESS"))

  req.ViewID("112236938")
  req.DateRange("2017-01-01", "2017-01-02")

  req.Dimension("ga:deviceCategory")
  req.Dimension("ga:campaign")
  req.Dimension("ga:adGroup")

  req.Metric("ga:users", "", "")
  req.Metric("ga:sessions", "", "")

  res, err := req.Do()

  if err != nil {
    log.Fatal(err.Error())
  }

  if err := res.CSV(os.Stdout); err != nil {
    log.Fatal(err.Error())
  }
}
```
