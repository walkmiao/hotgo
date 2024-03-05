package customer

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/businessIn"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/output/businessOut"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取客户列表
type ListCustomerReq struct {
	g.Meta `path:"/customer" method:"get" tags:"客户" summary:"获取客户列表"`
	businessIn.CustomerSearchInp
}

type ListCustomerRes struct {
	List []*businessIn.CustomerListModel `json:"list"   dc:"客户列表"`
	form.PageRes
}

// 删除一个或者多个客户
type DeleteCustomerReq struct {
	g.Meta `path:"/customer" method:"delete" tags:"客户" summary:"删除客户"`
	Ids    []int `json:"ids" dc:"需要删除的客户ID"`
}
type DeleteCustomerRes struct {
}

// 获取客户类型字典
type GetCustomerOptionReq struct {
	g.Meta `path:"/customer/option" method:"get" tags:"客户" summary:"获取客户类型"`
}

type GetCustomerOptionRes struct {
	Mix *businessOut.CustomerOptionsMix `json:"mix" dc:"客户字典清单"`
}

// 更新客户状态
type UpdateStatusReq struct {
	g.Meta `path:"/customer/status" method:"put" tags:"客户" summary:"更新客户状态"`
	Id     int `json:"id"`
	Status int `json:"status"`
}

type UpdateStatusRes struct {
}

// 编辑单个客户，更新或者新增
type EditCustomerReq struct {
	g.Meta   `path:"/customer" method:"post" tags:"客户" summary:"编辑单个客户"`
	Customer entity.BasicCustomer `json:"customer"`
}

type EditCustomerRes struct {
}

// 获取客户联系人
type GetCustomerContactReq struct {
	g.Meta `path:"/customer/contact" method:"get" tags:"客户" summary:"获取客户联系人"`
	Id     int `json:"id" v:"required#客户ID不能为空" dc:"客户ID"`
}

type GetCustomerContactRes struct {
	ContactList []businessIn.Contact `json:"list" dc:"获取客户联系人列表"`
}

// 编辑客户联系人
type EditCustomerContactReq struct {
	g.Meta   `path:"/customer/contact" method:"post" tags:"客户" summary:"编辑客户联系人"`
	Contacts []businessIn.Contact `json:"contacts" v:"required#客户联系人不能为空" dc:"客户联系人列表"`
	Id       int                  `json:"id" dc:"客户ID"`
}

type EditCustomerContactRes struct {
}
