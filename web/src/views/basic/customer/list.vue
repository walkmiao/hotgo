<template>
  <div>
    <BasicForm
      @register="register"
      @submit="handleSubmit"
      @reset="handleReset"
      @keyup.enter="handleSubmit"
      ref="searchFormRef"
    >
      <template #statusSlot="{ model, field }">
        <n-input v-model:value="model[field]" />
      </template>
    </BasicForm>
    <BasicTable
      :openChecked="true"
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      :actionColumn="actionColumn"
      @update:checked-row-keys="onCheckedRow"
      :scroll-x="1800"
    >
      <template #tableTitle>
        <n-button
          type="primary"
          @click="addTable"
          class="min-left-space"
          v-if="hasPermission(['/member/edit'])"
        >
          <template #icon>
            <n-icon>
              <PlusOutlined />
            </n-icon>
          </template>
          添加客户
        </n-button>
        <n-button
          type="error"
          @click="batchDelete"
          :disabled="batchDeleteDisabled"
          class="min-left-space"
          v-if="hasPermission(['/member/delete'])"
        >
          <template #icon>
            <n-icon>
              <DeleteOutlined />
            </n-icon>
          </template>
          批量删除
        </n-button>
      </template>
    </BasicTable>

    <n-modal
      v-model:show="showModal"
      :show-icon="false"
      preset="dialog"
      :title="formParams?.id > 0 ? '编辑客户 #' + formParams?.id : '添加客户'"
      :style="{ width: dialogWidth }"
    >
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
      >
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="6" label="客户类型" size="large" path="type">
            <n-select
              placeholder="选择类型"
              :options="formParams.types"
              v-model:value="formParams.type"
            />
            <n-button @click="showDictCustomer">...</n-button>
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="客户名称" path="name">
            <n-input placeholder="请输入客户名称" v-model:value="formParams.name" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="客户编码">
            <n-input placeholder="" v-model:value="formParams.code" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="助记符">
            <n-input placeholder="输入一个好记的名字" v-model:value="formParams.zjf" />
          </n-form-item-gi>
        </n-grid>
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="6" label="客户来源">
            <n-select
              placeholder="选择类型"
              :options="formParams.sources"
              v-model:value="formParams.source"
            />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="城市选择">
            <CitySelector v-model:value="formParams.city" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="详细地址">
            <n-input placeholder="输入详细地址" v-model:value="formParams.addr" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="既是客户又是供应商" :label-width="140">
            <n-switch placeholder="" v-model:value="formParams.supply" />
          </n-form-item-gi>
        </n-grid>
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="6" label="邀请码">
            <n-input v-model:value="formParams.viteCode" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="用户名">
            <n-input v-model:value="formParams.username" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="密码">
            <n-input
              type="password"
              show-password-on="mousedown"
              v-model:value="formParams.password"
            />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="生成链接">
            <n-button @click="null" type="info">生成链接</n-button>
          </n-form-item-gi>
        </n-grid>
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="24" label="描述(备注)">
            <n-input v-model:value="formParams.description" />
          </n-form-item-gi>
        </n-grid>
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="24" label="">
            <n-tabs
              type="line"
              class="card-tabs"
              default-value="联系人"
              animated
              @before-leave="handleBeforeLeave"
            >
              <n-tab-pane
                :name="item.name"
                :tab="item.name"
                v-for="item in customerTabs"
                :key="item.name"
              >
                <Contact
                  v-if="item.name == '联系人'"
                  :customerId="formParams.id"
                  ref="contactRef"
                />
              </n-tab-pane>
            </n-tabs>
          </n-form-item-gi>
        </n-grid>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="() => (showModal = false)">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
        </n-space>
      </template>
    </n-modal>
    <n-modal v-model:show="showDictModal">
      <List :checkedId="40" @update-modal="updateModal" />
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onMounted } from 'vue';
  import { useDialog, useMessage, FormRules } from 'naive-ui';
  import { ActionItem, BasicTable, TableAction } from '@/components/Table';
  import { BasicForm } from '@/components/Form/index';
  import { GetCustomerList, Status, DeleteCustomer } from '@/api/customer/customer';
  import { columns } from './columns';
  import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
  import { customerTabs } from '@/enums/optionsiEnum';
  import { adaModalWidth } from '@/utils/hotgo';
  import { getRandomString } from '@/utils/charset';
  import { cloneDeep } from 'lodash-es';
  import Contact from './contact.vue';
  import { toRaw } from '@vue/reactivity';
  import bus from '@/libs/eventBus';
  import List from '@/views/system/dict/list.vue';
  import {
    addNewState,
    addState,
    register,
    defaultState,
    State,
    options,
    loadOptions,
  } from './model';
  import { usePermission } from '@/hooks/web/usePermission';
  import { useUserStore } from '@/store/modules/user';
  import { LoginRoute } from '@/router';
  import { getNowUrl } from '@/utils/urlUtils';
  import CitySelector from '@/components/CitySelector/citySelector.vue';
  interface Props {
    type?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    type: '-1',
  });

  const rules: FormRules = {
    name: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入用户名',
    },
  };

  const { hasPermission } = usePermission();
  const userStore = useUserStore();
  const showIntegralModal = ref(false);
  const showBalanceModal = ref(false);
  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const searchFormRef = ref<any>({});
  const formRef = ref<any>({});
  const contactRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const dialogWidth = ref('100%');
  const formParams = ref<State>({
    id: 0,
    name: '',
    city: 0,
    type: -1,
    source: -1,
    types: [],
    sources: [],
  });
  const showQrModal = ref(false);
  const showDictModal = ref(false);
  const qrParams = ref({
    name: '',
    qrUrl: '',
  });
  const actionColumn = reactive({
    width: 250,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '启用',
            onClick: handleStatus.bind(null, record, 0),
            ifShow: () => {
              return record.status === 1;
            },
            auth: ['/business/customer/status'],
          },
          {
            label: '禁用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 0;
            },
            auth: ['/business/customer/status'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            auth: ['/member/edit'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            auth: ['/member/delete'],
          },
          {
            label: '联系人',
            onClick: showContact.bind(null, record),
            ifShow: true,
          },
        ],
        //dropDownActions: downActions,
        select: (key) => {
          if (key === 0) {
            return handleResetPwd(record);
          }
          if (key === 100) {
            return handleAddBalance(record);
          }
          if (key === 101) {
            return handleAddIntegral(record);
          }
          if (key === 102) {
            if (userStore.loginConfig?.loginRegisterSwitch !== 1) {
              message.error('管理员暂未开启此功能');
              return;
            }
            return handleInviteQR(record.inviteCode);
          }
        },
      });
    },
  });

  function getDropDownActions(record: Recordable): ActionItem[] {
    if (record.id === 1) {
      return [];
    }

    let list = [
      {
        label: '重置密码',
        key: 0,
      },
      {
        label: '变更余额',
        key: 100,
      },
      {
        label: '变更积分',
        key: 101,
      },
    ];

    if (userStore.loginConfig?.loginRegisterSwitch === 1) {
      list.push({
        label: 'TA的邀请码',
        key: 102,
      });
    }

    return list;
  }
  function handleBeforeLeave(name: string | number, oldname: string | number | null): boolean {
    //todo 添加离开tab事件
    console.log('name:', name, 'oldname:', oldname);
    return true;
  }
  function addTable() {
    showModal.value = true;
    formParams.value = cloneDeep(defaultState);
    addOptions();
  }

  function addOptions() {
    const opts = options.value.typeTabs.filter((item) => item.id !== -1);
    const sources = options.value.sourceTabs.filter((item) => item.id !== -1);
    formParams.value.types = [];
    opts.forEach((item) => {
      formParams.value.types.push({
        label: item.name,
        value: item.id,
      });
    });
    formParams.value.sources = [];
    sources.forEach((item) => {
      formParams.value.sources.push({
        label: item.name,
        value: item.id,
      });
    });
  }
  const loadDataTable = async (res) => {
    adaModalWidth(dialogWidth, 1200);
    return await GetCustomerList({
      ...res,
      ...searchFormRef.value?.formModel,
      ...{ type: props.type },
    });
  };

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        bus.emit('editCustomer', toRaw(formParams.value));
        showModal.value = false;
        loadDataTable(null)
          .then(() => {
            reloadTable();
          })
          .catch((e) => {
            message.error(e.message);
          });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function handleEdit(record: Recordable) {
    formParams.value = cloneDeep(record) as State;
    addOptions();
    showModal.value = true;
  }

  function handleResetPwd(record: Recordable) {
    record.password = getRandomString(12);
    dialog.warning({
      title: '警告',
      content: '你确定要重置密码？\r\n重置成功后密码为：' + record.password + '\r\n 请先保存',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        ResetPwd(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
    });
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: `你确定要删除id为${record.id}的记录吗？`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        DeleteCustomer([record.id]).then((_res) => {
          message.success('删除记录成功');
          reloadTable();
        });
      },
    });
  }

  function showContact(record: Recordable) {
    formParams.value = cloneDeep(record);
    addOptions();
    showModal.value = true;
  }

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        DeleteCustomer(checkedIds.value).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
    });
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    reloadTable();
  }

  function handleStatus(record: Recordable, status) {
    Status(record.id, status).then((_res) => {
      message.success(status === 1 ? '禁用成功' : '启用成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  function handleUpdateDeptValue(value) {
    formParams.value.deptId = Number(value);
  }

  function handleUpdateRoleValue(value) {
    formParams.value.roleId = Number(value);
  }

  function handleUpdatePostValue(value) {
    formParams.value.postIds = value;
  }

  function updateBalanceShowModal(value) {
    showBalanceModal.value = value;
  }

  function handleAddBalance(record: Recordable) {
    showBalanceModal.value = true;
    formParams.value = addNewState(record as addState);
  }

  function updateIntegralShowModal(value) {
    showIntegralModal.value = value;
  }

  function handleAddIntegral(record: Recordable) {
    showIntegralModal.value = true;
    formParams.value = addNewState(record as addState);
  }

  function handleInviteQR(code: any) {
    const domain = getNowUrl() + '#';
    qrParams.value.qrUrl = domain + LoginRoute.path + '?scope=register&inviteCode=' + code;
    showQrModal.value = true;
  }

  function showDictCustomer() {
    showDictModal.value = true;
  }
  function updateModal() {
    console.log('update modal');
    showDictModal.value = false;
    loadOptions().then(() => {
      addOptions();
    });
  }
  onMounted(async () => {
    await loadOptions();
  });
</script>

<style lang="less" scoped></style>
