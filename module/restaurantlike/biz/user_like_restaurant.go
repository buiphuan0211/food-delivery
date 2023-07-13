package biz

import (
	"context"
	"food-delivery/component/asyncjob"
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

// NewUserLikeRestaurantBiz ...
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

	// Side effect
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.incStore.IncreaseLikesCount(ctx, data.RestaurantID)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	//go func() {
	//	defer common.AppRecover()
	//	if err := biz.incStore.IncreaseLikesCount(ctx, data.RestaurantID); err != nil {
	//		log.Println(" Err IncreaseLikesCount: ", err)
	//	}
	//}()

	return nil
}
