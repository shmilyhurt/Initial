package dao

import (
	"Initial/conf"
	"gorm.io/gorm"
	"time"
)

type DashProStatus struct {
	Id        int       `gorm:"primary_key"`
	OneStatus int       `gorm:"column:one_status;default:0"`
	TwoStatus int       `gorm:"column:two_status;default:0"`
	ThrStatus int       `gorm:"column:thr_status;default:0"`
	FouStatus int       `gorm:"column:fou_status;default:0"`
	FivStatus int       `gorm:"column:fiv_status;default:0"`
	Created   time.Time `gorm:"autoCreateTime;column:created;type:datetime"`
}

type DashProType struct {
	Id      int       `gorm:"primary_key"`
	SCount  int       `gorm:"column:s_count;default:0"`
	MCount  int       `gorm:"column:m_count;default:0"`
	LCount  int       `gorm:"column:l_count;default:0"`
	DCount  int       `gorm:"column:d_count;default:0"`
	Created time.Time `gorm:"autoCreateTime;column:created;type:datetime"`
}

func (d *DashProStatus) TableName() string {
	return "dash_pro_status"
}

func (d *DashProType) TableName() string {
	return "dash_pro_type"
}

func CreateDashProStatus() {
	var projects []Project
	var dashProStatus DashProStatus
	var (
		oneStatusCount  int64
		twoStatusCount  int64
		secStatusCount  int64
		fourStatusCount int64
		fiveStatusCount int64
	)
	conf.Db.Find(&projects).Where("status", "N").Count(&oneStatusCount)
	conf.Db.Find(&projects).Where("status", "P").Count(&twoStatusCount)
	conf.Db.Find(&projects).Where("status", "E").Count(&secStatusCount)
	conf.Db.Find(&projects).Where("status", "S").Count(&fourStatusCount)
	conf.Db.Find(&projects).Where("status", "C").Count(&fiveStatusCount)
	dashProStatus.OneStatus = int(oneStatusCount)
	dashProStatus.TwoStatus = int(twoStatusCount)
	dashProStatus.ThrStatus = int(secStatusCount)
	dashProStatus.FouStatus = int(fourStatusCount)
	dashProStatus.FivStatus = int(fiveStatusCount)
	conf.Db.Create(&dashProStatus)
}

func CreateDashProType() {
	var projects []Project
	var dashProType DashProType
	var (
		SCount int64
		MCount int64
		LCount int64
		DCount int64
	)
	conf.Db.Find(&projects).Where("type", "S").Count(&SCount)
	conf.Db.Find(&projects).Where("type", "M").Count(&MCount)
	conf.Db.Find(&projects).Where("type", "L").Count(&LCount)
	conf.Db.Find(&projects).Where("type", "D").Count(&DCount)
	dashProType.SCount = int(SCount)
	dashProType.MCount = int(MCount)
	dashProType.LCount = int(LCount)
	dashProType.DCount = int(DCount)
	conf.Db.Create(&dashProType)
}

func GetDashProStatus(db *gorm.DB, dashProStatus *DashProStatus) (err error) {
	err = db.Last(dashProStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDashProType(db *gorm.DB, dashProType *[]DashProType) (err error) {
	err = db.Order("id desc").Limit(7).Find(dashProType).Error
	if err != nil {
		return err
	}
	return nil
}
