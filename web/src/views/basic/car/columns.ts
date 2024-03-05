import { BasicColumn } from '@/components/Table/src/types/table';
import { h } from 'vue';
import { NSwitch, NTag, NSelect } from 'naive-ui';
import { options, SelectData } from './model';
import { Dict } from '@/api/dict/dict';

export const columns: BasicColumn[] = [
  {
    title: '编码',
    key: 'code',
    width: 100,
  },
  {
    title: '品牌',
    key: 'brand',
    width: 100,
  },
  {
    title: '排量',
    key: 'cc',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '年款',
    key: 'year',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '车型',
    key: 'model',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '销售名称',
    key: 'name',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '国家',
    key: 'country',
    width: 100,
    editComponent: 'NInput',
    edit: true,
  },
  {
    title: '类型',
    key: 'type',
    width: 100,

    render(row: any) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.typeDict?.listClass,
          bordered: false,
        },
        { default: () => row.typeDict?.label || '' }
      );
    },
  },
  {
    title: '厂家',
    key: 'vender',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '燃油类型',
    key: 'fuel',
    width: 100,

    render(row: any) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.fuelDict?.listClass,
          bordered: false,
        },
        { default: () => row.fuelDict?.label || '' }
      );
    },
  },
  {
    title: '变速箱类型',
    key: 'gearBox',
    width: 100,
    render(row: any) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.gearBoxDict?.listClass,
          bordered: false,
        },
        { default: () => row.gearBoxDict?.label || '' }
      );
    },
  },
  {
    title: '变速箱描述',
    key: 'gearDesc',
    width: 100,
    edit: true,
    editComponent: 'NInput',
  },
  {
    title: '前轮胎规格',
    key: 'frontTire',
    width: 100,
    render(row: any) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.frontTireDict?.listClass,
          bordered: false,
        },
        { default: () => row.frontTireDict?.label || '' }
      );
    },
  },
  {
    title: '后轮胎规格',
    key: 'backTire',
    width: 100,
    render(row: any) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.backTireDict?.listClass,
          bordered: false,
        },
        { default: () => row.backTireDict?.label || '' }
      );
    },
  },
  {
    title: '前轮毂规格',
    key: 'frontHub',
    width: 100,
    render(row: any) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.frontHubDict?.listClass,
          bordered: false,
        },
        { default: () => row.frontHubDict?.label || '' }
      );
    },
  },
  {
    title: '后轮毂规格',
    key: 'backHub',
    width: 100,
    render(row: any) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.backHubDict?.listClass,
          bordered: false,
        },
        { default: () => row.backHubDict?.label || '' }
      );
    },
  },
  {
    title: '车系',
    key: 'series',
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
];
