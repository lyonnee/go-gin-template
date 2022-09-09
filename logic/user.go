package logic

import (
	"context"
	"time"

	"github.com/LyonNee/app-layout/dal"
	"github.com/LyonNee/app-layout/database/mysql"
	"github.com/LyonNee/app-layout/entity"
	"github.com/LyonNee/app-layout/pkg/errors"
)

type UserLogic struct{}

func (logic *UserLogic) Register(ctx context.Context, name string, age uint8, phoneNum string, password string) (uint64, error) {
	conn, err := mysql.GetConn(ctx)
	if err != nil {
		return 0, err
	}

	now := time.Now().Unix()
	userEntity := entity.User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNum,
		Password:    password,
		CreateAt:    now,
		UpdateAt:    now,
	}

	uid, err := dal.InsertOneUser(ctx, conn, userEntity)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (logic *UserLogic) Login(ctx context.Context, phoneNum string, password string) (uint64, string, string, error) {
	conn, err := mysql.GetConn(ctx)
	if err != nil {
		return 0, "", "", err
	}

	userEntity, err := dal.QueryUserByPhoneNum(ctx, conn, phoneNum)
	if err != nil {
		return 0, "", "", err
	}

	if userEntity.Password != password {
		return 0, "", "", errors.ERR_INVALID_PASSWORD
	}

	return userEntity.ID, userEntity.Name, userEntity.PhoneNumber, nil
}
