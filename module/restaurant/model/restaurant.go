package restaurantmodel

import (
	"errors"
	"food-delivery/common"
	"strings"
)

const EntityName = "restaurant"

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr" gorm:"column:addr;"`
	UserId          int                `json:"-" gorm:"column:user_id"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"` // preload:false -> user bị insert theo khi create restaurant
	LikedCount      int                `json:"liked_count" gorm:"column:liked_count"`
	//Cover           *common.Images     `json:"cover" gorm:"column:cover"`
	//Logo            *common.Image      `json:"logo" gorm:"column:logo"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
	UserId          int            `json:"-" gorm:"column:user_id"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (m *RestaurantCreate) GenObjID() {
	//m.ObjectID = primitive.NewObjectID().Hex()
}

func (m *RestaurantCreate) Validate() error {
	m.Name = strings.TrimSpace(m.Name)

	if m.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo"`
	Cover *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
