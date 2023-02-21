package models

import (
	"errors"
	"log"
	"sync"
	"time"
)

// Video `gorm:"-"`  读写操作均会忽略该字段
type Video struct {
	Id         int64       `json:"id,omitempty"`
	UserInfoId int64       `json:"-"`
	Author     UserInfo    `json:"author,omitempty" gorm:"-"`
	PlayUrl    string      `json:"play_url,omitempty"`
	CoverUrl   string      `json:"cover_url,omitempty"`
	Title      string      `json:"title,omitempty"`
	Users      []*UserInfo `json:"-" gorm:"many2many:user_favor_videos;"`
	CreatedAt  time.Time   `json:"-"`
	UpdatedAt  time.Time   `json:"-"`
}

type T struct {
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Location  struct {
		Street   string `json:"street"`
		City     string `json:"city"`
		State    string `json:"state"`
		Postcode int    `json:"postcode"`
	} `json:"location"`
	Username string `json:"username"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
}

type VideoDAO struct {
}

var (
	videoDAO  *VideoDAO
	videoOnce sync.Once
)

func NewVideoDAO() *VideoDAO {
	videoOnce.Do(func() {
		videoDAO = new(VideoDAO)
	})
	return videoDAO
}

// AddVideo 添加视频
// 注意：由于视频和userinfo有多对一的关系，所以传入的Video参数一定要进行id的映射处理！
func (v *VideoDAO) AddVideo(video *Video) error {
	if video == nil {
		return errors.New("AddVideo video 空指针")
	}
	return DB.Create(video).Error
}

// QueryVideoByVideoId 通过id查询视频
func (v *VideoDAO) QueryVideoByVideoId(videoId int64, video *Video) error {
	if video == nil {
		return errors.New("QueryVideoByVideoId 空指针")
	}
	return DB.Where("id=?", videoId).
		Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title"}).
		First(video).Error
}

// QueryVideoCountByUserId 通过用户id查询视频数量
func (v *VideoDAO) QueryVideoCountByUserId(userId int64, count *int64) error {
	if count == nil {
		return errors.New("QueryVideoCountByUserId count 空指针")
	}
	return DB.Model(&Video{}).Where("user_info_id=?", userId).Count(count).Error
}

// QueryVideoListByUserId 通过用户id查询视频列表
func (v *VideoDAO) QueryVideoListByUserId(userId int64, videoList *[]*Video) error {
	if videoList == nil {
		return errors.New("QueryVideoListByUserId videoList 空指针")
	}
	return DB.Where("user_info_id=?", userId).
		Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title"}).
		Find(videoList).Error
}

// QueryVideoListByLimitAndTime  返回按投稿时间倒序的视频列表，并限制为最多limit个
func (v *VideoDAO) QueryVideoListByLimitAndTime(limit int, latestTime time.Time, videoList *[]*Video) error {
	if videoList == nil {
		return errors.New("QueryVideoListByLimit videoList 空指针")
	}
	return DB.Model(&Video{}).Where("created_at<?", latestTime).
		Order("created_at ASC").Limit(limit).
		Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title", "created_at", "updated_at"}).
		Find(videoList).Error
}

// IsVideoExistById 视频是否存在
func (v *VideoDAO) IsVideoExistById(id int64) bool {
	var video Video
	if err := DB.Where("id=?", id).Select("id").First(&video).Error; err != nil {
		log.Println(err)
	}
	if video.Id == 0 {
		return false
	}
	return true
}
