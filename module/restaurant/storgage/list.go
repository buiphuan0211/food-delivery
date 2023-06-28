package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
	"strconv"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var (
		result []restaurantmodel.Restaurant
	)

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("user_id = ?", f.OwnerId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		db = db.Where("id < ?", v)
		paging.FakeCursor = v
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		lastItem := result[len(result)-1]
		id := strconv.Itoa(lastItem.ID)
		paging.NextCursor = id
	}

	println("result: ", result)

	return result, nil
}
