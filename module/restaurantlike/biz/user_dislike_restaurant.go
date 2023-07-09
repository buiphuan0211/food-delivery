package biz

import (
	"context"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"log"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecLikedCountResStore interface {
	DecreaseLikesCount(tx context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store    UserDislikeRestaurantStore
	desStore DecLikedCountResStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, desStore DecLikedCountResStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, desStore: desStore}
}

func (biz *userDislikeRestaurantBiz) UserDislikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDisLikeRestaurant(err)
	}

	if err := biz.desStore.DecreaseLikesCount(ctx, restaurantId); err != nil {
		log.Println("Error DecreaseLikesCount: ", err)
	}

	return nil
}
