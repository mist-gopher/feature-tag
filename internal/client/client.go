package client

type Client struct {
	Id   string
	Keys []ClientKey
}

type ClientKey struct {
	Name  string
	Value string
}
