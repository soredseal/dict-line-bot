package dictionary

type Request struct {
	Word string
	Token string
}

type Client interface {
	Fetch(request Request)
}
