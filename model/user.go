package model

import (
	"fmt"

	"1024casts/backend/pkg/auth"
	"1024casts/backend/pkg/constvar"

	"gopkg.in/go-playground/validator.v9"
	"time"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

type UserProfile struct {
	UserModel
	Email string `json:"email" gorm:"column:email;not null" binding:"required" validate:"min=6,max=32"`
	Avatar string `json:"avatar" gorm:"column:avatar" binding:"omitempty"`
	RealName string `json:"real_name" gorm:"column:real_name" binding:"omitempty" validate:"min=6,max=32"`
	City string `json:"city" gorm:"column:city" binding:"omitempty"`
	Company string `json:"city" gorm:"column:city" binding:"omitempty"`
	WeiboUrl string `json:"weibo_url" gorm:"column:weibo_url" binding:"omitempty"`
	WechatId string `json:"wechat_id" gorm:"column:wechat_id" binding:"omitempty"`
	PersonalWebsite string `json:"personal_website" gorm:"column:personal_website" binding:"omitempty"`
	Introduction string `json:"introduction" gorm:"column:introduction" binding:"omitempty"`
	TopicCount int `json:"topic_count" gorm:"column:topic_count" binding:"omitempty"`
	ReplyCount int `json:"reply_count" gorm:"column:reply_count" binding:"omitempty"`
	FollowerCount int `json:"follower_count" gorm:"column:follower_count" binding:"omitempty"`
	NotificationCount int `json:"notification_count" gorm:"column:notification_count" binding:"omitempty"`
	Status int `json:"status" gorm:"column:status" binding:"omitempty"`
	LastLoginTime time.Time `json:"last_login_time" gorm:"column:last_login_time" binding:"omitempty"`
	LastLoginIp string `json:"last_login_ip" gorm:"column:last_login_ip" binding:"omitempty"`
	GithubId string `json:"github_id" gorm:"column:github_id" binding:"omitempty"`
	RememberToken string `json:"remember_token" gorm:"column:remember_token" binding:"omitempty"`
	IsActivated int `json:"is_activated" gorm:"column:is_activated" binding:"omitempty"`
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

// Update updates an user account information.
func (u *UserModel) Update(username string) error {
	return DB.Self.Model(u).Updates(UserModel{Username: username}).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// GetUser gets an user by the user identifier.
func GetUserById(id int) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ?", id).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
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
