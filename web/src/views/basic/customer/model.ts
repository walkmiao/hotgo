import { cloneDeep } from 'lodash-es';
import { ref, unref } from 'vue';
import { Dicts } from '@/api/dict/dict';
import { GetCustomerOption } from '@/api/customer/customer';
import { getRoleOption } from '@/api/system/role';
import { getPostOption } from '@/api/org/post';
import { FormSchema, useForm } from '@/components/Form';
import { statusOptions } from '@/enums/optionsiEnum';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { SelectOption } from 'naive-ui';
export interface Customer {
  type: any;
  name: string;
  zjf: string;
  addr?: string;
  source?: number;
  supply?: number;
  account?: string;
  certificate?: string;
}
// 增加余额/积分.

export interface addState {
  id: number;
  username: string;
  realName: string;
  integral: number;
  balance: number;
  operateMode: number;
  num: number | null;
}

export const addDefaultState = {
  id: 0,
  realName: '',
  username: '',
  integral: 0,
  balance: 0,
  operateMode: 1,
  num: null,
};

export function addNewState(state: addState | null): addState {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(addDefaultState);
}

export const operateModes = [
  {
    value: 1,
    label: '加款',
  },
  {
    value: 2,
    label: '扣款',
  },
];

export const addRules = {};

// 用户列表.

export const defaultState: State = {
  id: 0,
  types: [],
  curTypeId: 1,
  zjf: '',
  addr: '',
  city: '',
  name: '',
  sources: [],
  curSourceId: 1,
};

export interface State {
  id: number;
  types: SelectOption[];
  name: string;
  type: number;
  zjf?: string;
  city: number;
  addr?: string;
  fullAddr?: string;
  sources: SelectOption[];
  source: number;
  code?: string;
  account?: string;
  supply?: boolean;
  certificate?: string;
  viteCode?: string;
  username?: string;
  password?: string;
  description?: string;
  contacts?: Contact[];
}

export interface Contact {
  name: string;
  sex: string;
  phone: string;
  telephone: string;
  qq?: string;
  wx?: string;
  birthday?: string;
  email?: string;
  addr?: string;
  openId?: string;
  firstChoose?: boolean;
}
const info = 'xxx';
export const defaultInfo: Contact = {
  name: info,
  phone: info,
  telephone: info,
  qq: info,
  wx: info,
  sex: '男',
  addr: info,
  openId: info,
  firstChoose: false,
};

const schemas: FormSchema[] = [
  {
    field: 'key',
    component: 'NInput',
    label: '关键字',
    labelMessage: '是客户不是用户',
    componentProps: {
      placeholder: '请输入客户名/编码/联系人/电话',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: [{ message: '请输入用户名', trigger: ['blur'] }],
  },
  {
    field: 'created_at',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
];

export const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

export const options = ref<any>({
  sourceTabs: [{ id: -1, name: '全部' }],
  typeTabs: [{ id: -1, name: '全部' }],
});

export async function loadOptions() {
  const customerOpts = await GetCustomerOption();
  if (customerOpts.mix) {
    if (customerOpts.mix.options) options.value.typeTabs = customerOpts.mix.options;
    if (customerOpts.mix.sources) options.value.sourceTabs = customerOpts.mix.sources;
  }
}
