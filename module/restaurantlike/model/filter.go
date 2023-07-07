package restaurantlikemodel

type Filter struct {
	RestaurantID int `json:"-" form:"restaurant_id"`
	UserID       int `json:"-" form:"user_id"`
}
