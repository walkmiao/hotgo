import { defineStore } from 'pinia';
import { createStorage, storage } from '@/utils/Storage';
import { store } from '@/store';
import { GetCustomerOption } from '@/api/customer/customer'


const Storage = createStorage({ storage: localStorage });




export interface ICustomerState {
  typeTabs: LabelValueOptions[];
  sourceTabs: LabelValueOptions[];
}

export const useCustomerStore = defineStore({
  id: 'app-customer',
  state: (): ICustomerState => ({
    typeTabs: [] as LabelValueOptions[],
    sourceTabs: [] as LabelValueOptions[]

  }),
  getters: {
    getTypeTabs(): LabelValueOptions[] {
      return this.typeTabs;
    },
    getSourceTabs(): LabelValueOptions[] {
      return this.sourceTabs;
    }
  },
  actions: {

    //加载tab选项
    async loadOptions() {
      const customerOpts = await GetCustomerOption();
      if (customerOpts.mix) {
        if (customerOpts.mix.options)
          for (const i in customerOpts.mix.options) {
            this.typeTabs.push(customerOpts.mix.options[i]);
          }
        if (customerOpts.mix.sources)
          for (const i in customerOpts.mix.sources) {
            this.sourceTabs.push(customerOpts.mix.sources[i]);
          }
      }

    }
  },
});

// Need to be used outside the setup
export function useCustomerStoreWidthOut() {
  return useCustomerStore(store);
}
