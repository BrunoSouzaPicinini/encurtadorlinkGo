package service

import (
	"encurtadorLink/src/model"
	"encurtadorLink/src/repository"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Create(short model.Shortener) (model.Shortener, error) {
	if short.Slug == "" {
		id, err := gonanoid.New(10)

		if err != nil {
			return model.Shortener{}, err
		}

		short.Slug = id
	}

	createdShort, err := repository.Insert(short)

	if err != nil {
		return model.Shortener{}, err
	}

	return createdShort, nil
}

func FindBySlug(slug string) (model.Shortener, error) {
	short, err := repository.FindBySlug(slug)

	if err != nil {
		return model.Shortener{}, err
	}

	return short, nil
}
