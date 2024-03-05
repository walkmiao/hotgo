package hook

import (
	"context"
	"encoding/json"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/businessIn"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberInfo 后台用户信息
var CustomerInfo = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		if err != nil {
			return
		}
		for i, record := range result {
			// 客户类型
			if !record["type"].IsEmpty() {
				typeName, err := g.Model("sys_dict_data").Ctx(ctx).
					Fields("label").
					Where(g.Map{
						"type":  consts.CUSTOMER_TYPE,
						"value": record["type"].Interface(),
					}).
					Value()
				if err != nil {
					break
				}
				record["typeName"] = typeName
			}

			// 客户来源
			if !record["source"].IsEmpty() {
				sourceName, err := g.Model("sys_dict_data").Ctx(ctx).
					Fields("label").
					Where(g.Map{
						"type":  consts.CUSTOMER_RESOURCE,
						"value": record["source"].Interface(),
					}).
					Value()
				if err != nil {
					break
				}
				record["sourceName"] = sourceName
			}
			// 创建者
			if !record["createdById"].IsEmpty() {
				createdName, err := g.Model("admin_member").Ctx(ctx).
					Fields("real_name").
					Where(g.Map{
						"id": record["createdById"],
					}).
					Value()
				if err != nil {
					break
				}
				record["createdName"] = createdName
			}

			// 创建者
			if !record["updatedById"].IsEmpty() {
				updatedName, err := g.Model("admin_member").Ctx(ctx).
					Fields("real_name").
					Where(g.Map{
						"id": record["updatedById"],
					}).
					Value()
				if err != nil {
					break
				}
				record["updatedName"] = updatedName
			}
			// 联系人
			if !record["contacts"].IsEmpty() {
				contacts := record["contacts"].Bytes()
				var contactsList = make([]businessIn.Contact, 0)
				if err = json.Unmarshal(contacts, &contactsList); err != nil {
					return
				}
				var selected bool
				if len(contactsList) == 1 {
					person := contactsList[0].Name
					phone := contactsList[0].Phone
					record["person"] = gvar.New(person)
					record["phone"] = gvar.New(phone)
					selected = true
				}
				if !selected {
					for _, contact := range contactsList {
						if contact.FirstChoose {
							person := contact.Name
							phone := contact.Phone
							record["person"] = gvar.New(person)
							record["phone"] = gvar.New(phone)
							selected = true
						}
					}
				}
				if !selected && len(contactsList) > 0 {
					person := contactsList[0].Name
					phone := contactsList[0].Phone
					record["person"] = gvar.New(person)
					record["phone"] = gvar.New(phone)
				}

			}

			result[i] = record
		}
		return
	},
}
