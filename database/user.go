package database

import (
	"hako-server/model"
)

func (d *GormDatabase) GetUserInfo(address string) (model.UserInfo, error) {
	var userInfo model.UserInfo
	// SELECT * FROM `user_infos` WHERE address = '0xE31c9fF6a8A1b952098CfeaF60c521cf68435503'
	err := d.DB.Where("address = ?", address).Find(&userInfo).Error
	return userInfo, err
}
