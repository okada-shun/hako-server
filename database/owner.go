package database

import (
	"hako-server/model"
)

func (d *GormDatabase) GetOwnerInfo() (model.OwnerInfo, error) {
	var ownerInfo model.OwnerInfo
	// SELECT * FROM `owner_infos`
	err := d.DB.Find(&ownerInfo).Error
	return ownerInfo, err
}
