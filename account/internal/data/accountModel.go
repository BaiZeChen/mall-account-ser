package data

import (
	"errors"
	"fmt"
	"mall-account-ser/account/pkg"
)

type Account struct {
	Base
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

func (a *Account) CreateAccount() error {
	res := pkg.GlobalGorm.Create(a)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("添加账户失败，原因：%s", res.Error.Error()))
	}
	return nil
}

func (a *Account) UpdateAccountName(id uint, name string) error {
	a.ID = id
	res := pkg.GlobalGorm.Model(a).Update("name", name)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("修改账户名称失败，原因：%s", res.Error.Error()))
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
		Limit(limit).Find(list)
	if res.Error != nil {
		return nil, errors.New(fmt.Sprintf("查询账户列表失败，原因：%s", res.Error.Error()))
	}
	return list, nil
}

func (a *Account) AccountCount(name string) (int64, error) {
	var length int64
	res := pkg.GlobalGorm.Where("name LIKE ?", "%"+name+"%").Count(&length)
	if res.Error != nil {
		return 0, errors.New(fmt.Sprintf("统计账户列表失败，原因：%s", res.Error.Error()))
	}
	return length, nil
}
