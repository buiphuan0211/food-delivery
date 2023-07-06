package restaurantlikemodel

type Filter struct {
	RestaurantID string `json:"-" form:"restaurant_id"`
	UserID       string `json:"-" form:"user_id"`
}
