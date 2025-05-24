package repository

import "encurtadorLink/src/model"

func Insert(short model.Shortener) (model.Shortener, error) {
	result := dbConnection.Create(&short)

	if result.Error != nil {
		return model.Shortener{}, result.Error
	}
	return short, nil
}

func FindBySlug(slug string) (model.Shortener, error) {
	var short model.Shortener
	result := dbConnection.Where("slug = ?", slug).First(&short)

	if result.Error != nil {
		return model.Shortener{}, result.Error
	}

	return short, nil
}
