package businessIn

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

type CustomerEditInp struct {
	entity.BasicCustomer
}

type CustomerSearchInp struct {
	form.PageReq
	Key       string  `json:"key"        dc:"搜索关键字"`
	Supply    int     `json:"supply"      dc:"即是客户又是供应商"`
	CreatedAt []int64 `json:"createdAt"  dc:"创建时间"`
	Type      int     `json:"type" dc:"客户类型"`
}

type CustomerListModel struct {
	entity.BasicCustomer
	TypeName    string `json:"typeName"`
	SourceName  string `json:"sourceName"`
	CreatedName string `json:"createdName"`
	UpdatedName string `json:"updatedName"`
	Person      string `json:"person"`
	Phone       string `json:"phone"`
}

type CustomerOptionModel struct {
	Type     int    `json:"id"`
	TypeName string `json:"name"`
}

type CustomerSourceModel struct {
	Source     int    `json:"id"`
	SourceName string `json:"name"`
}

// VerifyUniqueInp 验证客户唯一属性
type VerifyUniqueInp struct {
	Id    int64
	Where g.Map
}

type CustomerUpdateFields struct {
	Id          int64  `json:"id"          description:"客户ID"`
	Type        int64  `json:"type"        description:"客户类别-字典"`
	Name        string `json:"name"        description:"客户名称"`
	Zjf         string `json:"zjf"         description:"助记符"`
	Addr        string `json:"addr"        description:"客户地址"`
	Source      int64  `json:"source"      description:"客户来源"`
	Supply      int    `json:"supply"      description:"即是客户又是供应商"`
	Account     string `json:"account"     description:"客户账户"`
	Certificate string `json:"certificate" description:"客户身份号码"`
	Code        string `json:"code"        description:"客户编码"`
	ViteCode    string `json:"viteCode"    description:"邀请码"`
	Description string `json:"description" description:"描述"`
	Username    string `json:"username"    description:"对账账号"`
	Password    string `json:"password"    description:"对账密码"`
	UpdatedById int64  `json:"updatedById" description:"更新用户的id"`
}
type CustomerInsertFields struct {
	Type        int64  `json:"type"        description:"客户类别-字典"`
	Name        string `json:"name"        description:"客户名称"`
	Zjf         string `json:"zjf"         description:"助记符"`
	Addr        string `json:"addr"        description:"客户地址"`
	Source      int64  `json:"source"      description:"客户来源"`
	Supply      int    `json:"supply"      description:"即是客户又是供应商"`
	Account     string `json:"account"     description:"客户账户"`
	Certificate string `json:"certificate" description:"客户身份号码"`
	Code        string `json:"code"        description:"客户编码"`
	ViteCode    string `json:"viteCode"    description:"邀请码"`
	Description string `json:"description" description:"描述"`
	Username    string `json:"username"    description:"对账账号"`
	Password    string `json:"password"    description:"对账密码"`
	CreatedById int64  `json:"createdById" description:"创建用户的id"`
}

type CustomerStatusIn struct {
	Id     int `json:"id,omitempty"`
	Status int `json:"status,omitempty"`
}

type Contact struct {
	Name        string `json:"name" bson:"name"`
	Sex         string `json:"sex" bson:"sex"`
	Phone       string `json:"phone" bson:"phone"`
	TelePhone   string `json:"telePhone" bson:"tele_phone"`
	QQ          string `json:"qq"`
	WX          string `json:"wx" bson:"wx"`
	Birthday    string `json:"birthday" bson:"birthday"`
	Email       string `json:"email" bson:"email"`
	Addr        string `json:"addr" bson:"addr"`
	OpenId      string `json:"openId" bson:"open_id"`
	FirstChoose bool   `json:"firstChoose" bson:"first_choose"` //是否首选
}
