package restaurantbusiness

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	Delete(ctx context.Context, id int) error
	FindDataWithCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (data *restaurantmodel.Restaurant, err error)
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{
		store: store,
	}
}

func (b *deleteRestaurantBusiness) Delete(ctx context.Context, id int) error {
	var cond = map[string]interface{}{"id": id}
	restaurant, err := b.store.FindDataWithCondition(ctx, cond)
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if restaurant.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, err)
	}

	if err = b.store.Delete(ctx, id); err != nil {
		return common.ErrCanNotDeleteEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
