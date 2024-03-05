package hook

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var CarModelInfo = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		if err != nil {
			return
		}

		var list = make([]entity.SysDictData, 0)
		err = g.Model("sys_dict_data").Ctx(ctx).
			WhereIn("type", []string{
				consts.CAR_MODEL_FUEL_TYPE, consts.CAR_MODEL_GEAR_TYPE,
				consts.CAR_MODEL_HUB_TYPE, consts.CAR_MODEL_TIRE_TYPE, consts.CAR_MODEL_TYPE}).Scan(&list)

		if err != nil {
			return
		}
		var m = make(map[string]map[string]entity.SysDictData)
		for _, item := range list {
			if m[item.Type] == nil {
				m[item.Type] = make(map[string]entity.SysDictData)
			}
			m[item.Type][item.Value] = item

		}

		for i, record := range result {
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

			if rd := record["type"]; !rd.IsEmpty() {
				items, ok := m[consts.CAR_MODEL_TYPE]
				if ok && items != nil {
					v, ok := items[rd.String()]
					if ok {
						record["typeDict"] = gvar.New(v)
					}
				}
			}

			if rd := record["fuel"]; !rd.IsEmpty() {
				items, ok := m[consts.CAR_MODEL_FUEL_TYPE]
				if ok && items != nil {
					v, ok := items[rd.String()]
					if ok {
						record["fuelDict"] = gvar.New(v)
					}
				}
			}

			// 变速箱
			if rd := record["gearbox"]; !rd.IsEmpty() {
				items, ok := m[consts.CAR_MODEL_GEAR_TYPE]
				if ok && items != nil {
					v, ok := items[rd.String()]
					if ok {
						record["gearBoxDict"] = gvar.New(v)
					}
				}
			}
			//
			if rd := record["front_tire_spec"]; !rd.IsEmpty() {
				items, ok := m[consts.CAR_MODEL_TIRE_TYPE]
				if ok && items != nil {
					v, ok := items[rd.String()]
					if ok {
						record["frontTireDict"] = gvar.New(v)
					}
				}
			}

			if rd := record["back_tire_spec"]; !rd.IsEmpty() {
				items, ok := m[consts.CAR_MODEL_TIRE_TYPE]
				if ok && items != nil {
					v, ok := items[rd.String()]
					if ok {
						record["backTireDict"] = gvar.New(v)
					}
				}
			}
			if rd := record["front_hub_spec"]; !rd.IsEmpty() {
				items, ok := m[consts.CAR_MODEL_HUB_TYPE]
				if ok && items != nil {
					v, ok := items[rd.String()]
					if ok {
						record["frontHubDict"] = gvar.New(v)
					}
				}
			}
			if rd := record["back_hub_spec"]; !rd.IsEmpty() {
				items, ok := m[consts.CAR_MODEL_HUB_TYPE]
				if ok && items != nil {
					v, ok := items[rd.String()]
					if ok {
						record["backHubDict"] = gvar.New(v)
					}
				}
			}
			result[i] = record
		}
		return
	},
}
