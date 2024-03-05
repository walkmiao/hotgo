package business

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/input/businessIn"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

type sCarModel struct{}

func init() {
	service.RegisterCarModel(New())
}
func New() *sCarModel {
	return &sCarModel{}
}

var columns = dao.BasicCarModel.Columns()

func (s *sCarModel) List(ctx context.Context, in *businessIn.CarModelSearchInp) (list []*businessIn.CarListModel, err error) {
	list = make([]*businessIn.CarListModel, 0)
	cols := dao.BasicCarModel.Columns()
	mod := dao.BasicCarModel.Ctx(ctx).Where(columns.Deleted, 0)
	if in.Type > 0 {
		mod = mod.Where(columns.Type, in.Type)
	}
	if in.Key != "" {
		key := fmt.Sprintf("%%%s%%", in.Key)
		mod = mod.WhereLike(columns.Code, key).
			WhereOrLike(columns.SaleName, key)
	}
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}
	err = mod.Hook(hook.CarModelInfo).Page(in.Page, in.PerPage).OrderDesc(columns.UpdatedAt).OrderDesc(columns.CreatedAt).Scan(&list)
	return
}

func (s *sCarModel) Edit(ctx context.Context, in *businessIn.CarEditInp) (err error) {
	mod := dao.BasicCarModel.Ctx(ctx).Where(columns.Deleted, 0)
	userId := contexts.GetUserId(ctx)
	//编辑
	if in.Id > 0 {
		in.UpdatedById = userId
		_, err = mod.WherePri(in.Id).Data(in).Update()
		return
	}
	//新增
	in.CreatedById = userId
	_, err = mod.Data(in).Insert()
	return
}

func (s *sCarModel) DeleteList(ctx context.Context, ids []int) (err error) {
	var r sql.Result
	if len(ids) == 1 {
		r, err = dao.BasicCarModel.Ctx(ctx).Where(dao.BasicCarModel.Columns().Id, ids[0]).Data(dao.BasicCarModel.Columns().Deleted, 1).Update()
	} else {
		r, err = dao.BasicCarModel.Ctx(ctx).WhereIn(dao.BasicCarModel.Columns().Id, ids).Data(dao.BasicCarModel.Columns().Deleted, 1).Update()

	}
	if err != nil {
		return errors.New("删除车辆类型发生错误")
	}
	count, err := r.RowsAffected()
	if err != nil {
		return
	}
	if int(count) != len(ids) {
		return fmt.Errorf("需要删除%d个记录,实际删除%d个记录", len(ids), count)
	}
	return
}
