package api

import (
    "encoding/json"
    "net/http"
)

// Coin balance params
type CoinBalanceParams struct {
    Username string
}

type CoinBalanceResponse struct {
    // Response code
    Code int
    // Account balance
    Balance int64
}

type Error struct {
    // Error code
    Code int
    // Error message
    Message string
}
