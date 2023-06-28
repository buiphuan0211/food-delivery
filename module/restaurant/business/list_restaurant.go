package restaurantbusiness

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) (result []restaurantmodel.Restaurant, err error)
}

type listRestaurantBusiness struct {
	store ListRestaurantStore
}

func NewListRestaurantBusiness(store ListRestaurantStore) *listRestaurantBusiness {
	return &listRestaurantBusiness{
		store: store,
	}
}

func (b *listRestaurantBusiness) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) (result []restaurantmodel.Restaurant, err error) {
	result, err = b.store.ListDataWithCondition(ctx, filter, paging, "User")

	if err != nil {
		return
	}
	return
}

// Interface của Golang sẽ khai báo ở nơi ta sử dụng nó
