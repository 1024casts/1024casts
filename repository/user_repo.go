package repository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
)

type UserRepo struct {
	db *model.Database
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		db: model.DB,
	}
}

func (repo *UserRepo) CreateUser(user model.UserModel) (id uint64, err error) {
	err = repo.db.Self.Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (repo *UserRepo) GetUserById(id uint64) (*model.UserModel, error) {
	user := &model.UserModel{}
	result := repo.db.Self.Where("id = ?", id).First(user)

	return user, result.Error
}

// GetUser gets an user by the user identifier.
func (repo *UserRepo) GetUserByUsername(username string) (*model.UserModel, error) {
	user := &model.UserModel{}
	result := repo.db.Self.Where("username = ?", username).First(user)

	return user, result.Error
}

func (repo *UserRepo) GetUserByGithubId(githubId string) (*model.UserModel, error) {
	user := &model.UserModel{}
	result := repo.db.Self.Where("github_id = ?", githubId).First(user)

	return user, result.Error
}

func (repo *UserRepo) GetUserByUserNames(username []string) ([]*model.UserModel, error) {
	user := make([]*model.UserModel, 0)
	result := repo.db.Self.Where("username in (?)", username).Find(&user)

	return user, result.Error
}

func (repo *UserRepo) GetUserList(username string, offset, limit int) ([]*model.UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*model.UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := repo.db.Self.Model(&model.UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := repo.db.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (repo *UserRepo) GetUserByEmail(email string) (*model.UserModel, error) {
	user := model.UserModel{}
	result := repo.db.Self.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func (repo *UserRepo) FindByChangePasswordHash(hash string) (*model.UserModel, error) {
	users := model.UserModel{}
	return &users, nil
}

func (repo *UserRepo) Update(userMap map[string]interface{}, id uint64) error {
	user, err := repo.GetUserById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(user).Updates(userMap).Error
}

func (repo *UserRepo) IncrReplyCount(userId uint64) error {
	user := model.UserModel{}
	return repo.db.Self.Model(&user).Where("id=?", userId).
		Update("reply_count", gorm.Expr("reply_count + ?", 1)).Error
}

func (repo *UserRepo) DeleteUser(id uint64) error {
	user, err := repo.GetUserById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&user).Error
}

func (repo *UserRepo) GetResetPasswordInfoByEmail(email string) (*model.PasswordResetModel, error) {
	user := model.PasswordResetModel{}
	result := repo.db.Self.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func (repo *UserRepo) DeleteResetPasswordByEmail(email string) error {
	user := model.PasswordResetModel{}
	return repo.db.Self.Where("email = ?", email).Delete(&user).Error
}

func (repo *UserRepo) GetUserMember(userId uint64, status int) (*model.UserMemberModel, error) {
	userMember := model.UserMemberModel{}
	result := repo.db.Self.Where("user_id = ? and status=?", userId, status).Order("id desc").First(&userMember)
	if result.Error == gorm.ErrRecordNotFound {
		return &userMember, nil
	}

	return &userMember, result.Error
}

func (repo *UserRepo) UpdateUserMemberStatus(userId uint64) error {
	userMember := model.UserMemberModel{}
	return repo.db.Self.Model(&userMember).Where("user_id=?", userId).
		Update("status", gorm.Expr("reply_count + ?", 1)).Error
}

func (repo *UserRepo) UpdateUserMemberEndTime(db *gorm.DB, userId uint64, endTime time.Time) error {
	userMember := model.UserMemberModel{}
	return db.Model(&userMember).Where("user_id=?", userId).
		Update("end_time", endTime).Error
}

func (repo *UserRepo) UpdateUserMember(db *gorm.DB, userId uint64, startTime, endTime time.Time) error {
	userMember := model.UserMemberModel{}
	return db.Model(&userMember).Where("user_id=?", userId).
		Updates(model.UserMemberModel{StartTime: startTime, EndTime: endTime}).Error
}
