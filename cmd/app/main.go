package main

import (
	"fmt"
	"net/http"

	"cortico/internal"

	// Load environment variables
	_ "github.com/leapkit/leapkit/core/tools/envload"

	// postgres driver
	_ "github.com/lib/pq"
)

func main() {
	s := internal.New()
	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println("[error] starting app:", err)
	}
}
