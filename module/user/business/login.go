package userbusiness

import (
	"context"
	"food-delivery/common"
	"food-delivery/component/appcontext"
	"food-delivery/component/tokenprovider"
	usermodel "food-delivery/module/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, data map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type LoginBusiness struct {
	appCtx        appcontext.AppContext
	StoreUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *LoginBusiness {
	return &LoginBusiness{
		StoreUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user
// 2. Hash password from input and compare with password in DB
// 3.Provider: issueJWT Token for client
// 3.1. Access token and refresh token
// 4. Return token

func (biz *LoginBusiness) Login(ctx context.Context, data usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, _ := biz.StoreUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passHasher := biz.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHasher {
		return nil, usermodel.ErrEmailOrPasswordInvalid

	}

	payload := tokenprovider.TokenPayload{
		UserId: user.ID,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
