package service

import (
	"context"
	"gim/internal/user/cache"
	"gim/internal/user/dao"
	"gim/internal/user/model"
)

type userService struct{}

var UserService = new(userService)

// Get 获取用户信息
func (*userService) Get(ctx context.Context, userId int64) (*model.User, error) {
	user, err := cache.UserCache.Get(userId)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = dao.UserDao.Get(userId)
	if err != nil {
		return nil, err
	}

	if user != nil {
		err = cache.UserCache.Set(*user)
		if err != nil {
			return nil, err
		}
	}
	return user, err
}

// Get 获取用户信息
func (*userService) GetByIds(ctx context.Context, userIds []int64) ([]model.User, error) {
	return dao.UserDao.GetByIds(userIds)
}

// Get 获取用户信息
func (*userService) Update(ctx context.Context, user model.User) error {
	err := dao.UserDao.Update(user)
	if err != nil {
		return err
	}

	err = cache.UserCache.Del(user.Id)
	if err != nil {
		return err
	}

	return nil
}
