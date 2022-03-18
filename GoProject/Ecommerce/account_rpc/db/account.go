package db

import (
	"context"
	"time"

	"github.com/XiaoXinTao/goproject/account_rpc/model"
)

func CreateUser(ctx context.Context, passportUid int64, passwordCode string) error {
	User := &model.Users{
		PassportUid: passportUid,
		PasswordCode: passwordCode,
		CreateTime: time.Now(),
	}
	return getDb(ctx).Debug().Model(&model.Users{}).Create(&User).Error
}

func CheckUserIsExist(ctx context.Context, passportUid int64, passwordCode string) (bool, error) {
	var count int64 = 0
	err := getDb(ctx).Debug().Model(&model.Users{}).Where("passport_uid = ?", passportUid).Where("password_code = ?", passwordCode).
		Count(&count).Error
	return count > 0, err
}