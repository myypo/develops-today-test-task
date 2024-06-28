package cat

import (
	"encoding/json"
	"fmt"
	"io"
	"sca/internal/context"
	herr "sca/internal/error/http"

	"github.com/google/uuid"
)

type Cat struct {
	Id uuid.UUID

	Name     string
	YearsExp uint
	Breed
	SalaryCents uint
}

type Breed string

func NewBreed(ctx *context.Context, maybeBreed string) (Breed, herr.ErrorHttp) {
	type breedRecord struct {
		Name string `json:"name" binding:"required"`
	}

	resp, err := ctx.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		return "", herr.NewErrInternal(err)
	}
	defer resp.Body.Close()

	var records []breedRecord
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", herr.NewErrInternal(err)
	}
	if err := json.Unmarshal(body, &records); err != nil {
		return "", herr.NewErrInternal(err)
	}

	for _, v := range records {
		if v.Name == maybeBreed {
			return Breed(maybeBreed), nil
		}
	}

	return "", herr.NewErrBadRequest(
		fmt.Errorf("the provided breed does not exist: %s", maybeBreed),
	)
}
