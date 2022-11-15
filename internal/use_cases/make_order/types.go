package makeorder

type request struct {
	Room      string `json:"room"`
	UserEmail string `json:"email"`
	From      string `json:"from"`
	To        string `json:"to"`
}
