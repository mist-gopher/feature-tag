package usecase_test

import (
	"testing"

	"github.com/mist-gopher/feature-tag/internal/repository/memory_repo"
	"github.com/mist-gopher/feature-tag/internal/tag"
	"github.com/mist-gopher/feature-tag/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func MockTagData() map[string]tag.Tag {
	activeTag := tag.New("valid key", "tagname", true)
	return map[string]tag.Tag{
		activeTag.Id: activeTag,
	}
}

func TestGetTagStatus(t *testing.T) {
	type TestCase struct {
		input       usecase.GetTagStatusInput
		resultValue bool
		resultError error
	}

	cases := map[string]TestCase{
		"with valid input for active tag": {
			input: usecase.GetTagStatusInput{
				AppId:   "valid key",
				ApiKey:  "valid id",
				TagName: "tagname",
			},
			resultValue: true,
			resultError: nil,
		},
		"unknown tag": {
			input: usecase.GetTagStatusInput{
				AppId:   "valid key",
				ApiKey:  "valid id",
				TagName: "not a created tag",
			},
			resultValue: false,
			resultError: nil,
		},
	}

	repo := memory_repo.NewTagRepositoryInMemory()
	repo.Data = MockTagData()

	for testName, testCase := range cases {
		t.Run(testName, func(t *testing.T) {
			result, err := usecase.GetTagStatus(testCase.input, &repo)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.resultValue, result.Active)
		})
	}
}
