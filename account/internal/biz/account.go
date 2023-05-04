package biz

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mall-account-ser/account/internal/data"
)

type AccountControl struct {
	ID       uint
	Name     string
	Password string
}

func (a *AccountControl) Add() error {
	encode, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败！")
	}
	a.Password = string(encode)

	model := data.Account{
		Name:     a.Name,
		Password: a.Password,
	}
	err = model.CreateAccount()
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountControl) UpdateName() error {
	if len(a.Name) == 0 || len(a.Name) > 12 {
		return errors.New("账号长度不符合规则，请重新输入！")
	}

	model := data.Account{
		Base: data.Base{ID: a.ID},
		Name: a.Name,
	}
	err := model.UpdateAccount()
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountControl) UpdatePassword() error {
	if len(a.Password) == 0 {
		return errors.New("请填写密码！")
	}

	ecode, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败！")
	}
	a.Password = string(ecode)

	model := data.Account{
		Base:     data.Base{ID: a.ID},
		Password: a.Password,
	}
	err = model.UpdateAccount()
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountControl) Delete() error {
	model := data.Account{
		Base: data.Base{ID: a.ID},
	}
	err := model.DeleteAccount()
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountControl) List(offset, limit int) ([]data.Account, error) {
	model := data.Account{}
	list, err := model.AccountList(a.Name, offset, limit)
	if err != nil {
		return nil, err
	}
	return list, nil
}
