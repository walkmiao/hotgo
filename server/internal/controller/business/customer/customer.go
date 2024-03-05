package customer

import (
	"context"
	"hotgo/api/business/customer"
	"hotgo/internal/library/response"
	"hotgo/internal/model/input/businessIn"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/container/gmap"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Customer = cCustomer{}
)

type cCustomer struct{}

func (c *cCustomer) AddCustomer(ctx context.Context, req *customer.EditCustomerReq) (res *customer.EditCustomerRes, err error) {
	if err = service.Customer().Edit(ctx, &businessIn.CustomerEditInp{BasicCustomer: req.Customer}); err != nil {
		return
	}
	response.RText(g.RequestFromCtx(ctx), "更新用户成功")
	return
}

func (c *cCustomer) List(ctx context.Context, req *customer.ListCustomerReq) (res *customer.ListCustomerRes, err error) {
	gmap.New()
	list, totalCount, err := service.Customer().List(ctx, &req.CustomerSearchInp)
	if err != nil {
		return
	}
	res = new(customer.ListCustomerRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *cCustomer) Delete(ctx context.Context, req *customer.DeleteCustomerReq) (res *customer.DeleteCustomerRes, err error) {
	err = service.Customer().DeleteList(ctx, req.Ids)
	return
}

func (c *cCustomer) Options(ctx context.Context, req *customer.GetCustomerOptionReq) (res *customer.GetCustomerOptionRes, err error) {
	list, err := service.Customer().Options(ctx)
	if err != nil {
		return
	}
	res = new(customer.GetCustomerOptionRes)
	res.Mix = list
	return
}

func (c *cCustomer) GetCustomerContacts(ctx context.Context, req *customer.GetCustomerContactReq) (res *customer.GetCustomerContactRes, err error) {
	list, err := service.Customer().GetContactList(ctx, req.Id)
	if err != nil {
		return
	}
	res = new(customer.GetCustomerContactRes)
	res.ContactList = list
	return
}

func (c *cCustomer) EditCustomerContact(ctx context.Context, req *customer.EditCustomerContactReq) (res *customer.EditCustomerContactRes, err error) {
	err = service.Customer().EditCustomerContact(ctx, req.Id, req.Contacts)
	return
}

func (c *cCustomer) Status(ctx context.Context, req *customer.UpdateStatusReq) (res *customer.UpdateStatusRes, err error) {
	err = service.Customer().UpdateStatus(ctx, &businessIn.CustomerStatusIn{Id: req.Id, Status: req.Status})
	return
}
