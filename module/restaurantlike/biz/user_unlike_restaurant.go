package biz

import (
	"context"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store}
}

func (biz *userDislikeRestaurantBiz) UserDislikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDisLikeRestaurant(err)
	}

	return nil
}
