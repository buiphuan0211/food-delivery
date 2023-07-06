package restaurantbiz

import (
	"context"
	"fmt"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
	"log"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) (result []restaurantmodel.Restaurant, err error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

func NewListRestaurantBusiness(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{
		store:     store,
		likeStore: likeStore,
	}
}

func (b *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) (result []restaurantmodel.Restaurant, err error) {
	result, err = b.store.ListDataWithCondition(ctx, filter, paging, "User")

	if err != nil {
		return
	}

	ids := make([]int, len(result))
	for i := range ids {
		ids[i] = result[i].ID
	}

	fmt.Println("b: ", b)
	fmt.Println("b.likeStore: ", b.likeStore)

	// FIXME: Giá trị b.like bị nil
	likeMap, err := b.likeStore.GetRestaurantLikes(ctx, ids)
	fmt.Println("check 2")

	if err != nil {
		fmt.Println("check 3")
		log.Println(err)
		return
	}

	for i, item := range result {
		result[i].LikeCount = likeMap[item.ID]
	}

	return
}

// Interface của Golang sẽ khai báo ở nơi ta sử dụng nó
