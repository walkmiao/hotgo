import { BasicColumn } from '@/components/Table/src/types/table';
import { h } from 'vue';
import { NSwitch, NTag, NSelect } from 'naive-ui';

export interface CustomerType {
  key: any;
  label: string;
  value: any;
  valueType: string;
  type: string;
  listClass: string;
}
export const contactColumns: BasicColumn[] = [
  {
    title: '联系人',
    key: 'name',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '性别',
    key: 'sex',
    width: 100,
    editComponent: 'NInput',
    edit: true,
  },
  {
    title: '电话',
    key: 'phone',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '手机',
    key: 'telephone',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: 'QQ',
    key: 'qq',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '微信',
    key: 'wx',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '生日',
    key: 'birthday',
    width: 100,
    editComponent: 'NInput',
    edit: true,
  },
  {
    title: '邮件',
    key: 'email',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '地址',
    key: 'addr',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '微信OPENID',
    key: 'openId',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '是否首选',
    key: 'firstChoose',
    width: 100,
    edit: true,
    editComponent: 'NSwitch',
    render(row) {
      return h(NSwitch, {
        style: {
          marginRight: '6px',
        },
        value: row.firstChoose as boolean,
        disabled: true,
        bordered: false,
      });
    },
  },
];
export const columns = [
  {
    title: '客户ID',
    key: 'id',
    width: 100,
    sorter: (row1, row2) => {
      return row1.id - row2.id;
    },
  },
  {
    title: '客户类型',
    key: 'typeName',
    width: 100,
  },
  {
    title: '客户名称',
    key: 'name',
    width: 100,
  },
  {
    title: '助记符',
    key: 'zjf',
    width: 100,
  },
  {
    title: '地址',
    key: 'addr',
    width: 100,
  },
  {
    title: '电话',
    key: 'phone',
    width: 100,
  },
  {
    title: '首选联系人',
    key: 'person',
    width: 100,
  },
  {
    title: '客户来源',
    key: 'sourceName',
    width: 100,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'info',
          bordered: false,
        },
        { default: () => (row.sourceName != '' ? row.sourceName : '未添加字典') }
      );
    },
  },
  {
    title: '既是客户又是供应商',
    key: 'supply',
    width: 100,
    render(row) {
      return h(NSwitch, {
        style: {
          marginRight: '6px',
        },
        value: row.supply == 1 ? true : false,
        defaultValue: false,
        disabled: true,
        bordered: false,
      });
    },
  },
  {
    title: '客户账户',
    key: 'account',
    width: 100,
  },
  {
    title: '客户身份号码',
    key: 'certificate',
    width: 100,
  },
  {
    title: '客户编码',
    key: 'code',
    width: 100,
  },
  {
    title: '邀请码',
    key: 'vite_code',
    width: 100,
  },
  {
    title: '描述',
    key: 'description',
    width: 100,
  },
  {
    title: '对账账户',
    key: 'username',
    width: 100,
  },
  {
    title: '对账密码',
    key: 'password',
    width: 100,
  },
  {
    title: '创建者',
    key: 'createdName',
    width: 100,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 100,
  },
  {
    title: '更新者',
    key: 'updatedName',
    width: 100,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    width: 100,
  },
  {
    title: '状态',
    key: 'status',
    width: 80,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.status == 0 ? 'info' : 'error',
          bordered: false,
        },
        {
          default: () => (row.status == 0 ? '正常' : '已禁用'),
        }
      );
    },
  },
];
