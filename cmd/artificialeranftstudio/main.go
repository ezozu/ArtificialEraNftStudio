// cmd/artificialeranftstudio/main.go
package main

import (
"flag"
"log"
"os"

"artificialeranftstudio/internal/artificialeranftstudio"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeranftstudio.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
