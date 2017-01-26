# go-googleanalytics
Google Analytics custom report extractor

    package main

    import (
      "log"
      "os"

      ga "github.com/santacruz123/go-googleanalytics"
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