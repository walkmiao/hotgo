package businessIn

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

type CarListModel struct {
	entity.BasicCarModel
	UpdatedName   string             `json:"updatedName"` // 更新人名称
	CreatedName   string             `json:"createdName"` // 创建人名称
	BrandDict     entity.SysDictData `json:"brandDict"`
	TypeDict      entity.SysDictData `json:"typeDict"`
	FuelDict      entity.SysDictData `json:"fuelDict"`
	GearboxDict   entity.SysDictData `json:"gearBoxDict"`
	FrontTireDict entity.SysDictData `json:"frontTireDict"`
	BackTireDict  entity.SysDictData `json:"backTireDict"`
	FrontHubDict  entity.SysDictData `json:"frontHubDict"`
	BackHubDict   entity.SysDictData `json:"backHubDict"`
	SeriesDict    entity.SysDictData `json:"seriesDict"`
}

type CarEditInp struct {
	entity.BasicCarModel
}

type CarModelSearchInp struct {
	form.PageReq
	Key       string  `json:"key"        dc:"搜索关键字"`
	CreatedAt []int64 `json:"createdAt"  dc:"创建时间"`
	Type      int     `json:"type" dc:"车辆类型进口国产合资"`
}
