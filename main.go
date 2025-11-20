package main

import (
	"fmt"

	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/gogo/protobuf/proto"
	_ "github.com/hashicorp/golang-lru"
	_ "github.com/owncast/owncast/logging"
)

func main() {
	fmt.Println("Hello world")
}
