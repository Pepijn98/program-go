[![Go Report Card](https://goreportcard.com/badge/github.com/KurozeroPB/program-go)](https://goreportcard.com/report/github.com/KurozeroPB/program-go)

# program-go
__Small and simple program-o wrapper in Golang__

## Install
`go get github.com/KurozeroPB/program-go`

## Usage
Small example:
```go
package main

import (
  "fmt"

  "github.com/KurozeroPB/program-go"
)

func main() {
  resp, err := pgo.Say(6, "Testing 123 hello", "test_id_123456")
  if err != nil {
    fmt.Printf("Error: %s\n", err)
    return
  }
  fmt.Printf("Response: %s\n", resp.BotSay)
}
```

## Docs

#### Say(botID, query, convoID)
| Parameter | Type          | Description |
|-----------|:-------------:|-------------|
| botID     | int           | The program-o bot id you want to use
| query     | string        | The query you want to send
| convoID   | string        | The conversation id to recognize you
