package test

import "fmt"

func test() {

	fmt.Println("checj")

}

type UserInterface interface {
	GetUser() string
	GetEmail() string
	GetName() string
}

type userImpl struct {
}

func NewUser() UserInterface {
	return userImpl{}
}

func (u userImpl) GetUser() string {
	return ""
}

func (u userImpl) GetEmail() string { return "" }
func (u userImpl) GetName() string  { return "" }

type user1Impl struct {
}

func NewUser1() UserInterface {
	return userImpl{}
}
