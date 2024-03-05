package car

import (
	"context"
	"hotgo/api/business/car"
	"hotgo/internal/model/input/businessIn"
	"hotgo/internal/service"
)

var (
	CarModel = &cCarModel{}
)

type cCarModel struct{}

func (model *cCarModel) List(ctx context.Context, req *car.ListCarModelReq) (res *car.ListCarModelRes, err error) {
	list, err := service.CarModel().List(ctx, &req.CarModelSearchInp)
	if err != nil {
		return
	}
	res = new(car.ListCarModelRes)
	res.List = list
	res.PageRes.Pack(req, len(list))
	return
}

func (model *cCarModel) Edit(ctx context.Context, req *car.EditCarModelReq) (res *car.EditCarModelRes, err error) {
	err = service.CarModel().Edit(ctx, &businessIn.CarEditInp{BasicCarModel: req.BasicCarModel})
	return
}

func (model *cCarModel) Delete(ctx context.Context, req *car.DeleteCarModelReq) (res *car.DeleteCarModelRes, err error) {
	err = service.CarModel().DeleteList(ctx, req.Ids)
	return
}
