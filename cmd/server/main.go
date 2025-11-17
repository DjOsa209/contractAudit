package main

import (
    "contractAudit/internal/server"
)

func main() {
    s := server.New()
    s.Start()
}