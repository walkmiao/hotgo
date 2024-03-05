import { ref } from 'vue';
import { Dicts } from '@/api/dict/dict';

import { FormSchema, useForm } from '@/components/Form';
import { statusOptions } from '@/enums/optionsiEnum';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { SelectOption } from 'naive-ui';
import { SelectMixedOption } from 'naive-ui/es/select/src/interface';

export interface SelectData {
  key: number | string;
  label: string;
  value: number | string;
  listClass: string;
}

// 用户列表.

export const defaultState: CarModel = {
  id: 0,
};

export interface CarModel {
  id?: number;
  code?: string;
  brand?: string;
  cc?: string;
  year?: string;
  model?: string;
  saleName?: string;
  country?: string;
  type?: number;
  vendor?: string;
  fuel?: number;
  gearBox?: number;
  gearDesc?: string;
  frontTireSpec?: number;
  backTireSpec?: number;
  frontHubSpec?: number;
  backHubSpec?: number;
  series?: string;
}

const schemas: FormSchema[] = [
  {
    field: 'key',
    component: 'NInput',
    label: '关键字',
    labelMessage: '关键字查询',
    componentProps: {
      placeholder: '请输入关键字查询',
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

export const options = ref<any>({});

export async function loadOptions() {
  const prefix = 'service_car_';
  const types = ['model', 'type', 'fuel', 'gearbox', 'hub', 'tire'];
  const params = types.map((item) => prefix + item);
  const carOpts = (await Dicts({ types: params })) as Map<string, SelectData[]>;
  Object.keys(carOpts).forEach((key) => {
    options.value[key] = carOpts[key];
  });
}

export const countryOptions: SelectMixedOption[] = [
  { value: '阿尔及利亚', label: '阿尔及利亚' },
  { value: '安哥拉', label: '安哥拉' },
  { value: '阿根廷', label: '阿根廷' },
  { value: '亚美尼亚', label: '亚美尼亚' },
  { value: '阿鲁巴岛', label: '阿鲁巴岛' },
  { value: '澳大利亚', label: '澳大利亚' },
  { value: '奥地利', label: '奥地利' },
  { value: '阿塞拜疆', label: '阿塞拜疆' },
  { value: '巴哈马群岛', label: '巴哈马群岛' },
  { value: '巴林', label: '巴林' },
  { value: '孟加拉', label: '孟加拉' },
  { value: '比利时', label: '比利时' },
  { value: '伯利兹', label: '伯利兹' },
  { value: '百慕大', label: '百慕大' },
  { value: '不丹', label: '不丹' },
  { value: '玻利维亚', label: '玻利维亚' },
  { value: '博兹瓦纳', label: '博兹瓦纳' },
  { value: '巴西', label: '巴西' },
  { value: '保加利亚', label: '保加利亚' },
  { value: '柬埔寨', label: '柬埔寨' },
  { value: '加拿大', label: '加拿大' },
  { value: '智利', label: '智利' },
  { value: '中国', label: '中国' },
  { value: '哥伦比亚', label: '哥伦比亚' },
  { value: '哥斯达黎加', label: '哥斯达黎加' },
  { value: '克罗地亚', label: '克罗地亚' },
  { value: '塞浦路斯', label: '塞浦路斯' },
  { value: '捷克共和国', label: '捷克共和国' },
  { value: '丹麦', label: '丹麦' },
  { value: '多米尼加共和国', label: '多米尼加共和国' },
  { value: '厄瓜多尔', label: '厄瓜多尔' },
  { value: '埃及', label: '埃及' },
  { value: '萨尔瓦多', label: '萨尔瓦多' },
  { value: '爱沙尼亚', label: '爱沙尼亚' },
  { value: '埃塞俄比亚', label: '埃塞俄比亚' },
  { value: '芬兰', label: '芬兰' },
  { value: '法国', label: '法国' },
  { value: '格鲁吉亚', label: '格鲁吉亚' },
  { value: '德国', label: '德国' },
  { value: '加纳', label: '加纳' },
  { value: '希腊', label: '希腊' },
  { value: '格林纳达', label: '格林纳达' },
  { value: '危地马拉', label: '危地马拉' },
  { value: '洪都拉斯', label: '洪都拉斯' },
  { value: '匈牙利', label: '匈牙利' },
  { value: '冰岛', label: '冰岛' },
  { value: '印度', label: '印度' },
  { value: '印度尼西亚', label: '印度尼西亚' },
  { value: '爱尔兰', label: '爱尔兰' },
  { value: '以色列', label: '以色列' },
  { value: '意大利', label: '意大利' },
  { value: '日本', label: '日本' },
  { value: '约旦', label: '约旦' },
  { value: '哈萨克斯坦', label: '哈萨克斯坦' },
  { value: '肯尼亚', label: '肯尼亚' },
  { value: '科威特', label: '科威特' },
  { value: '拉脱维亚', label: '拉脱维亚' },
  { value: '黎巴嫩', label: '黎巴嫩' },
  { value: '立陶宛', label: '立陶宛' },
  { value: '马来西亚', label: '马来西亚' },
  { value: '墨西哥', label: '墨西哥' },
  { value: '摩洛哥', label: '摩洛哥' },
  { value: '莫桑比克', label: '莫桑比克' },
  { value: '荷兰', label: '荷兰' },
  { value: '新西兰', label: '新西兰' },
  { value: '尼加拉瓜', label: '尼加拉瓜' },
  { value: '尼日利亚', label: '尼日利亚' },
  { value: '挪威', label: '挪威' },
  { value: '阿曼', label: '阿曼' },
  { value: '巴基斯坦', label: '巴基斯坦' },
  { value: '巴拿马', label: '巴拿马' },
  { value: '巴拉圭', label: '巴拉圭' },
  { value: '秘鲁', label: '秘鲁' },
  { value: '菲律宾', label: '菲律宾' },
  { value: '波兰', label: '波兰' },
  { value: '葡萄牙', label: '葡萄牙' },
  { value: '波多黎各', label: '波多黎各' },
  { value: '卡塔尔', label: '卡塔尔' },
  { value: '罗马尼亚', label: '罗马尼亚' },
  { value: '俄罗斯', label: '俄罗斯' },
  { value: '新加坡', label: '新加坡' },
  { value: '南非', label: '南非' },
  { value: '韩国', label: '韩国' },
  { value: '西班牙', label: '西班牙' },
  { value: '斯里兰卡', label: '斯里兰卡' },
  { value: '瑞典', label: '瑞典' },
  { value: '瑞士', label: '瑞士' },
  { value: '台湾', label: '台湾' },
  { value: '坦桑尼亚', label: '坦桑尼亚' },
  { value: '泰国', label: '泰国' },
  { value: '突尼斯', label: '突尼斯' },
  { value: '土耳其', label: '土耳其' },
  { value: '乌干达', label: '乌干达' },
  { value: '乌克兰', label: '乌克兰' },
  { value: '阿联酋', label: '阿联酋' },
  { value: '英国', label: '英国' },
  { value: '乌拉圭', label: '乌拉圭' },
  { value: '乌兹别克斯坦', label: '乌兹别克斯坦' },
  { value: '委内瑞拉', label: '委内瑞拉' },
  { value: '越南', label: '越南' },
  { value: '赞比亚', label: '赞比亚' },
];
