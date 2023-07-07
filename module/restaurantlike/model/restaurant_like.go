package restaurantlikemodel

import (
	"fmt"
	"food-delivery/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantID int                `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserID       int                `json:"user_id" gorm:"column:user_id"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string { return "restaurant_likes" }

func (l *Like) GetRestautantId() int {
	return l.RestaurantID
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotDisLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot dislike this restaurant"),
		fmt.Sprintf("ErrCannotDisLikeRestaurant"),
	)
}
