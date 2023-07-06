package restaurantbiz

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) (result []restaurantmodel.Restaurant, err error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) // map[int]int -> restaurantId, number user liked
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBusiness(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{
		repo: repo,
	}
}

// ListRestaurant ...
func (b *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) (result []restaurantmodel.Restaurant, err error) {

	result, err = b.repo.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCanNotListEntity(restaurantmodel.EntityName, err)
	}

	return
}

// Interface của Golang sẽ khai báo ở nơi ta sử dụng nó
