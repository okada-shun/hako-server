package database

import (
	"hako-server/model"
)

func (d *GormDatabase) GetHakoInfo() (model.HakoInfo, error) {
	var hakoInfo model.HakoInfo
	// SELECT * FROM `hako_infos`
	err := d.DB.Find(&hakoInfo).Error
	return hakoInfo, err
}
