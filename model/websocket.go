package model

type WebSocketTicket struct {
	Ticket string `json:"ticket"`
}

type WebSocketType struct {
	Type           string   `json:"type"`
	Codes          []string `json:"codes"`
	IsOnlySnapshot *bool    `json:"isOnlySnapshot,omitempty"`
	IsOnlyRealtime *bool    `json:"isOnlyRealtime,omitempty"`
}

type WebSocketFormat struct {
	Format string `json:"format"`
}
