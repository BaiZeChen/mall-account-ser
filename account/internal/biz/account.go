package biz

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mall-account-ser/account/internal/data"
	"mall-account-ser/account/internal/pkg"
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

	encode, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败！")
	}
	a.Password = string(encode)

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

func (a *AccountControl) List(offset, limit int) ([]data.Account, int64, error) {
	model := data.Account{}
	list, err := model.AccountList(a.Name, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	count, err := model.AccountCount(a.Name)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (a *AccountControl) checkAccount() (data.Account, error) {
	model := &data.Account{
		Name: a.Name,
	}
	account, err := model.FindAccountByName()
	if err != nil {
		return account, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(a.Password))
	if err != nil {
		return account, errors.New("账号校验失败")
	}
	return account, nil
}

func (a *AccountControl) Login() (string, error) {
	account, err := a.checkAccount()
	if err != nil {
		return "", err
	}
	jwt := &pkg.JWTClaims{
		AccountId:   uint32(account.ID),
		AccountName: account.Name,
	}
	token, err := jwt.Generate()
	if err != nil {
		return "", err
	}
	return token, nil
}
