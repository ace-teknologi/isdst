# isdst

A golang helper to determine if a time in a location is in daylight savings time or not.

## Why?

Unlike other languages, go doesn't expose the ability to determine if
the time is in DST or not.

See [proposal: time: add Time.IsDST() bool method #42102](https://github.com/golang/go/issues/42102)

## How to use?

```go
package main

import "github.com/ace-teknologi/isdst"

func main() {
  loc, _ := time.LoadLocation("Australia/Broken_Hill")

  // DST
  t := time.Date(2020, time.January, 1, 0, 0, 0, 0, loc)
  fmt.Printf("Is daylight savings? %t", isdst.IsDST(t))

  // Non-DST
  t := time.Date(2020, time.June, 1, 0, 0, 0, 0, loc)
  fmt.Printf("Is daylight savings? %t", isdst.IsDST(t))
}
```
