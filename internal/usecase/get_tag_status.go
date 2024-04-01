package usecase

import (
	"github.com/mist-gopher/feature-tag/internal/tag"
)

type GetTagStatusInput struct {
	AppId   string
	ApiKey  string
	TagName string
}

type GetTagStatusOutput struct {
	Active bool
}

type GetTagByNameRepository interface {
	GetByName(name string) (*tag.Tag, error)
}

func GetTagStatus(input GetTagStatusInput, repo GetTagByNameRepository) (GetTagStatusOutput, error) {
	output := GetTagStatusOutput{Active: false}
	tag, err := repo.GetByName(tag.New(input.AppId, input.TagName, false).Id)

	if err != nil {
		return output, err
	}

	if tag != nil {
		output.Active = tag.Value
	}

	return output, nil
}
