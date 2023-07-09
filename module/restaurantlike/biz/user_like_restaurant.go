package biz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type InLikedCountResStore interface {
	IncreaseLikesCount(tx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore InLikedCountResStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore InLikedCountResStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:    store,
		incStore: incStore,
	}
}

// LikeRestaurant ...
func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		if err := biz.incStore.IncreaseLikesCount(ctx, data.RestaurantID); err != nil {
			log.Println(" Err IncreaseLikesCount: ", err)
		}
	}()

	return nil
}
