package instagram

type Client struct {
	sessionid string
}

func NewClient(sessionid string) *Client {
	return &Client{
		sessionid: sessionid,
	}
}
