package models

import (
	"errors"
	"log"
	"sync"
)

// UserInfo 用户信息结构体
type UserInfo struct {
	Id     int64      `json:"id" gorm:"id,omitempty"`
	Name   string     `json:"name" gorm:"name,omitempty"`
	User   *UserLogin `json:"-"` //用户与密码之间的多对多
	Videos []*Video   `json:"-"` //用户与投稿视频的一对多
}

type UserInfoDAO struct {
}

var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once
)

func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDAO = new(UserInfoDAO)
	})
	return userInfoDAO
}

// QueryUserInfoById 查询用户信息通过id
func (u *UserInfoDAO) QueryUserInfoById(userId int64, userinfo *UserInfo) error {
	if userinfo == nil {
		return errors.New("空指针错误")
	}
	//DB.Where("id=?",userId).First(userinfo)
	DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	//id为零值，说明sql执行失败
	if userinfo.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

// AddUserInfo 添加用户并返回用户信息
func (u *UserInfoDAO) AddUserInfo(userinfo *UserInfo) error {
	if userinfo == nil {
		return errors.New("空指针错误")
	}
	return DB.Create(userinfo).Error
}

// IsUserExistById 判断用户是否存在
func (u *UserInfoDAO) IsUserExistById(id int64) bool {
	var userinfo UserInfo
	if err := DB.Where("id=?", id).Select("id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.Id == 0 {
		return false
	}
	return true
}
