package biz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	GetUsersLikeRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{
		store: store,
	}
}

// ListUsers ...
func (biz *listUserLikeRestaurantBiz) ListUsers(ctx context.Context, filter *restaurantlikemodel.Filter, paging *common.Paging) ([]common.SimpleUser, error) {

	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCanNotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
