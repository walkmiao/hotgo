package business

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/businessIn"
	"hotgo/internal/model/output/businessOut"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sCustomer struct{}

func NewCustomer() *sCustomer {
	return &sCustomer{}
}

func init() {
	customer := NewCustomer()
	service.RegisterCustomer(customer)
}
func (s *sCustomer) VerifyUnique(ctx context.Context, in *businessIn.VerifyUniqueInp) (err error) {
	if in.Where == nil {
		return
	}

	cols := dao.BasicCustomer.Columns()
	msgMap := g.MapStrStr{
		cols.Code: "客户编码已存在，请换一个",
	}

	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = hgorm.IsUnique(ctx, &dao.BasicCustomer, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}

// Edit 修改/新增
func (s *sCustomer) Edit(ctx context.Context, in *businessIn.CustomerEditInp) (err error) {
	if in.Name == "" {
		return gerror.New("用户名称不能为空")
	}
	if in.Type == 0 {
		return gerror.New("用户类型不能为空")
	}

	//获取当前用户ID
	curId := contexts.GetUserId(ctx)
	//修改
	if in.Id > 0 {
		in.UpdatedById = curId
		_, err = dao.BasicCustomer.Ctx(ctx).WherePri(in.Id).Data(in).Update()
		return
	}
	in.CreatedById = curId
	if err = s.VerifyUnique(ctx, &businessIn.VerifyUniqueInp{
		Id: in.Id,
		Where: g.Map{
			dao.AdminDept.Columns().Code: in.Code,
		},
	}); err != nil {
		return
	}
	// 新增
	_, err = dao.BasicCustomer.Ctx(ctx).Data(in).Insert()
	if err != nil {
		return
	}
	return
}

// List 获取列表
func (s *sCustomer) List(ctx context.Context, in *businessIn.CustomerSearchInp) (list []*businessIn.CustomerListModel, totalCount int, err error) {
	mod := dao.BasicCustomer.Ctx(ctx)
	cols := dao.BasicCustomer.Columns()

	if in.Type > 0 {
		mod = mod.Where(cols.Type, in.Type)
	}
	// 部门名称
	if in.Key != "" {
		key := fmt.Sprintf("%%%s%%", in.Key)
		mod = mod.WhereOrLike(cols.Name, key).
			WhereOrLike(cols.Addr, key).
			WhereOrLike(cols.Code, key).
			WhereOrLike(cols.Zjf, key).
			WhereOrLike(cols.Contacts, key)
	}

	if in.Supply > 0 {
		mod = mod.Where(cols.Supply, 1)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	mod = mod.Where(dao.BasicCustomer.Columns().Deleted, 0)
	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取客户数据行失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Hook(hook.CustomerInfo).Page(in.Page, in.PerPage).OrderDesc(cols.Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取客户列表失败！")
		return
	}

	return
}

func (s *sCustomer) DeleteList(ctx context.Context, ids []int) (err error) {
	var r sql.Result
	if len(ids) == 1 {
		r, err = dao.BasicCustomer.Ctx(ctx).Where(dao.BasicCustomer.Columns().Id, ids[0]).Data(dao.BasicCustomer.Columns().Deleted, 1).Update()
	} else {
		r, err = dao.BasicCustomer.Ctx(ctx).WhereIn(dao.BasicCustomer.Columns().Id, ids).Data(dao.BasicCustomer.Columns().Deleted, 1).Update()

	}
	if err != nil {
		return errors.New("删除客户发生错误")
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

// Options 获取客户类型
func (s *sCustomer) Options(ctx context.Context) (mix *businessOut.CustomerOptionsMix, err error) {
	listMapping := make([]struct {
		Label string `json:"label"`
		Value int    `json:"value"`
		Type  string `json:"type"`
	}, 0)

	err = dao.SysDictData.Ctx(ctx).
		Fields("label", "value", "type").Where("type", consts.CUSTOMER_TYPE).
		WhereOr("type", consts.CUSTOMER_RESOURCE).Scan(&listMapping)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return
	}

	opts := []*businessIn.CustomerOptionModel{{
		Type:     -1,
		TypeName: "全部",
	}}
	sources := []*businessIn.CustomerSourceModel{{
		Source:     -1,
		SourceName: "全部",
	}}
	for _, item := range listMapping {
		switch item.Type {
		case consts.CUSTOMER_TYPE:
			opts = append(opts, &businessIn.CustomerOptionModel{
				Type:     item.Value,
				TypeName: item.Label,
			})
		case consts.CUSTOMER_RESOURCE:
			sources = append(sources, &businessIn.CustomerSourceModel{
				Source:     item.Value,
				SourceName: item.Label,
			})
		}

	}
	mix = &businessOut.CustomerOptionsMix{
		Options: opts,
		Sources: sources,
	}
	return
}

func (s *sCustomer) UpdateStatus(ctx context.Context, in *businessIn.CustomerStatusIn) (err error) {
	_, err = dao.BasicCustomer.Ctx(ctx).WherePri(in.Id).Data("status", in.Status).Update()
	return
}

func (s *sCustomer) GetContactList(ctx context.Context, id int) (contacts []businessIn.Contact, err error) {

	var customer entity.BasicCustomer
	contacts = make([]businessIn.Contact, 0)
	err = dao.BasicCustomer.Ctx(ctx).Fields(dao.BasicCustomer.Columns().Contacts).WherePri(id).Scan(&customer)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return contacts, nil
		}
		return
	}
	contacts = make([]businessIn.Contact, 0)
	var str = customer.Contacts.String()
	if str != "" {
		if err = json.Unmarshal([]byte(str), &contacts); err != nil {
			return
		}
	}
	return
}

func (s *sCustomer) EditCustomerContact(ctx context.Context, id int, contacts []businessIn.Contact) (err error) {
	data, err := json.Marshal(contacts)
	if err != nil {
		return
	}
	col := dao.BasicCustomer.Columns().Contacts
	dao.BasicCustomer.Ctx(ctx).Fields(col).WherePri(id).Data(g.Map{col: data}).Update()
	return

}
