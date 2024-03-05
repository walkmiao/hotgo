package car

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/businessIn"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

type ListCarModelReq struct {
	g.Meta `path:"/car/model/list" method:"get" tags:"车辆类型" summary:"获取车辆类型列表"`
	businessIn.CarModelSearchInp
}

type ListCarModelRes struct {
	List []*businessIn.CarListModel `json:"list"`
	form.PageRes
}

type EditCarModelReq struct {
	g.Meta `path:"/car/model/edit" method:"post" tags:"车辆类型" summary:"编辑车辆类型"`
	entity.BasicCarModel
}

type EditCarModelRes struct {
}

type DeleteCarModelReq struct {
	g.Meta `path:"/car/model/delete" method:"delete" tags:"车辆类型" summary:"删除车辆类型"`
	Ids    []int `json:"ids" dc:"需要删除的ID列表"`
}
type DeleteCarModelRes struct {
}
