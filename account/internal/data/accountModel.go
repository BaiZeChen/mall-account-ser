package data

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mall-ser/account/pkg"
)

type Account struct {
	Base
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

func (a *Account) TableName() string {
	return "account"
}

func (a *Account) CreateAccount() error {
	res := pkg.GlobalGorm.Create(a)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("添加账户失败，原因：%s", res.Error.Error()))
	}
	return nil
}

func (a *Account) UpdateAccount() error {
	if len(a.Name) > 0 {
		res := pkg.GlobalGorm.Model(a).Update("name", a.Name)
		if res.Error != nil {
			return errors.New(fmt.Sprintf("修改账户名称失败，原因：%s", res.Error.Error()))
		}
		if res.RowsAffected == 0 {
			return errors.New("没有找到对应的用户")
		}
	}
	if len(a.Password) > 0 {
		res := pkg.GlobalGorm.Model(a).Update("password", a.Password)
		if res.Error != nil {
			return errors.New(fmt.Sprintf("修改账户名称失败，原因：%s", res.Error.Error()))
		}
		if res.RowsAffected == 0 {
			return errors.New("没有找到对应的用户")
		}
	}
	return nil
}

func (a *Account) DeleteAccount() error {
	res := pkg.GlobalGorm.Delete(a)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("删除账户失败，原因：%s", res.Error.Error()))
	}
	return nil
}

func (a *Account) AccountList(name string, offset, limit int) ([]Account, error) {
	var list []Account
	res := pkg.GlobalGorm.
		Select("id", "name", "create_time", "update_time").
		Where("name LIKE ?", "%"+name+"%").Order("id desc").Offset(offset).
		Limit(limit).Find(&list)
	if res.Error != nil {
		return nil, errors.New(fmt.Sprintf("查询账户列表失败，原因：%s", res.Error.Error()))
	}
	return list, nil
}

func (a *Account) AccountCount(name string) (int64, error) {
	var length int64
	res := pkg.GlobalGorm.Model(a).Where("name LIKE ?", "%"+name+"%").Count(&length)
	if res.Error != nil {
		return 0, errors.New(fmt.Sprintf("统计账户列表失败，原因：%s", res.Error.Error()))
	}
	return length, nil
}

func (a *Account) FindAccountByName() (Account, error) {
	var account Account
	res := pkg.GlobalGorm.
		Where("name = ?", a.Name).
		First(&account)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// 没有找到账号
			return account, errors.New("没有找到对应的账号信息")
		}
		return account, res.Error
	}
	return account, nil
}
