package tag

type Tag struct {
	Id    string
	Name  string
	Value bool
}

func New(clientId, name string, value bool) Tag {
	return Tag{
		Id: MakeId(name, clientId),
    Name:  name,
		Value: value,
	}
}

func MakeId(name, clientId string) string {
  return "tag:" + clientId + ":" + name
}

