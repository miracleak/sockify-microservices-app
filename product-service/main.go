// product-service/main.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/product/:id", func(w http.ResponseWriter, r *http.Request) {
        // Implement product retrieval logic here
        fmt.Fprintln(w, "Product details retrieved")
    })

    // Other product-related routes

    http.ListenAndServe(":3001", nil)
}
