package biz

import (
	"context"
	"food-delivery/component/asyncjob"
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

// NewUserDislikeRestaurantBiz ...
func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, desStore DecLikedCountResStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, desStore: desStore}
}

// UserDislikeRestaurant ...
func (biz *userDislikeRestaurantBiz) UserDislikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDisLikeRestaurant(err)
	}

	// Side effect
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.desStore.DecreaseLikesCount(ctx, restaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	//go func() {
	//	defer common.AppRecover()
	//
	//	if err := biz.desStore.DecreaseLikesCount(ctx, restaurantId); err != nil {
	//		log.Println("Error DecreaseLikesCount: ", err)
	//	}
	//
	//	for i := 1; i <= 3; i++ {
	//		if err := biz.desStore.DecreaseLikesCount(ctx, restaurantId); err != nil {
	//			break
	//		}
	//
	//		time.Sleep(time.Second * 3)
	//	}
	//}()

	return nil
}
