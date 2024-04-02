package memory_repo

import (
	"github.com/mist-gopher/feature-tag/internal/tag"
)

type Tag struct {
	Data map[string]tag.Tag
	Err  error
}

func NewTagRepositoryInMemory() Tag {
	return Tag{
		Data: map[string]tag.Tag{},
		Err:  nil,
	}
}

func (repo *Tag) GetByName(name string) (*tag.Tag, error) {
	if repo.Err != nil {
		return nil, repo.Err
	}

	value, exist := repo.Data[name]

	if !exist {
		return nil, nil
	}

	return &value, nil
}
