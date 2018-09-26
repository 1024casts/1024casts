package model

import (
	"1024casts/backend/pkg/auth"
	"gopkg.in/go-playground/validator.v9"
	"sync"
	"time"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username          string    `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password          string    `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Email             string    `json:"email" gorm:"column:email;not null"`
	Avatar            string    `json:"avatar" gorm:"column:avatar" binding:"omitempty"`
	RealName          string    `json:"real_name" gorm:"column:real_name" binding:"omitempty"`
	City              string    `json:"city" gorm:"column:city" binding:"omitempty"`
	Company           string    `json:"company" gorm:"column:company" binding:"omitempty"`
	WeiboUrl          string    `json:"weibo_url" gorm:"column:weibo_url" binding:"omitempty"`
	WechatId          string    `json:"wechat_id" gorm:"column:wechat_id" binding:"omitempty"`
	PersonalWebsite   string    `json:"personal_website" gorm:"column:personal_website" binding:"omitempty"`
	Introduction      string    `json:"introduction" gorm:"column:introduction" binding:"omitempty"`
	TopicCount        int       `json:"topic_count" gorm:"column:topic_count" binding:"omitempty"`
	ReplyCount        int       `json:"reply_count" gorm:"column:reply_count" binding:"omitempty"`
	FollowerCount     int       `json:"follower_count" gorm:"column:follower_count" binding:"omitempty"`
	NotificationCount int       `json:"notification_count" gorm:"column:notification_count" binding:"omitempty"`
	Status            int       `json:"status" gorm:"column:status" binding:"omitempty"`
	LastLoginTime     time.Time `json:"last_login_time" gorm:"column:last_login_time" binding:"omitempty"`
	LastLoginIp       string    `json:"last_login_ip" gorm:"column:last_login_ip" binding:"omitempty"`
	GithubId          string    `json:"github_id" gorm:"column:github_id" binding:"omitempty"`
	RememberToken     string    `json:"remember_token" gorm:"column:remember_token" binding:"omitempty"`
	IsActivated       int       `json:"is_activated" gorm:"column:is_activated" binding:"omitempty"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserModel
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}

func (c *UserModel) TableName() string {
	return "users"
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
