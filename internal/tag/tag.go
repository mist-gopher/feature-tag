package tag

type Tag struct {
	Id    string
	Name  string
	Value bool
}

func New(clientId, name string, value bool) Tag {
	return Tag{
		Id:    "tag:" + clientId + ":" + name,
		Name:  name,
		Value: value,
	}
}
