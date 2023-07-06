package restaurantrepo

import (
	"context"
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
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) // map[int]int -> restaurantId, number user liked
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		likeStore: likeStore,
	}
}

func (b *listRestaurantRepo) ListRestaurant(
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

	likeMap, err := b.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println(err)
		return
	}

	for i, item := range result {
		result[i].LikeCount = likeMap[item.ID]
	}

	return
}
