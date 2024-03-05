// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/businessIn"
	"hotgo/internal/model/output/businessOut"
)

type (
	ICarModel interface {
		List(ctx context.Context, in *businessIn.CarModelSearchInp) (list []*businessIn.CarListModel, err error)
		Edit(ctx context.Context, in *businessIn.CarEditInp) (err error)
		DeleteList(ctx context.Context, ids []int) (err error)
	}
	ICustomer interface {
		VerifyUnique(ctx context.Context, in *businessIn.VerifyUniqueInp) (err error)
		Edit(ctx context.Context, in *businessIn.CustomerEditInp) (err error)
		List(ctx context.Context, in *businessIn.CustomerSearchInp) (list []*businessIn.CustomerListModel, totalCount int, err error)
		DeleteList(ctx context.Context, ids []int) (err error)
		Options(ctx context.Context) (mix *businessOut.CustomerOptionsMix, err error)
		UpdateStatus(ctx context.Context, in *businessIn.CustomerStatusIn) (err error)
		GetContactList(ctx context.Context, id int) (contacts []businessIn.Contact, err error)
		EditCustomerContact(ctx context.Context, id int, contacts []businessIn.Contact) (err error)
	}
)

var (
	localCarModel ICarModel
	localCustomer ICustomer
)

func CarModel() ICarModel {
	if localCarModel == nil {
		panic("implement not found for interface ICarModel, forgot register?")
	}
	return localCarModel
}

func RegisterCarModel(i ICarModel) {
	localCarModel = i
}

func Customer() ICustomer {
	if localCustomer == nil {
		panic("implement not found for interface ICustomer, forgot register?")
	}
	return localCustomer
}

func RegisterCustomer(i ICustomer) {
	localCustomer = i
}
