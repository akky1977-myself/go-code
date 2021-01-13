package entity

// Hash is struct
type Hash struct {
	ID      int64  `json:"id"`
	Hash    string `json:"hash"`
	Country string `json:"country"`
	Name    string `json:"name"`
}
