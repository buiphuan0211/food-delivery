package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
	"gorm.io/gorm"
)

// IncreaseLikesCount ...
func (s *sqlStore) IncreaseLikesCount(tx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).
		Error; err != nil {

		return common.ErrDB(err)
	}

	return nil
}

// DecreaseLikesCount ...
func (s *sqlStore) DecreaseLikesCount(tx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).
		Error; err != nil {

		return common.ErrDB(err)
	}

	return nil
}
