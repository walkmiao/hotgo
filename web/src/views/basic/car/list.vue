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
      :row-key="(row) => row.id"
      :request="loadDataTable"
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
          v-if="hasPermission(['/car/model/edit'])"
        >
          <template #icon>
            <n-icon>
              <PlusOutlined />
            </n-icon>
          </template>
          添加车型
        </n-button>
        <n-button
          type="error"
          @click="batchDelete"
          :disabled="batchDeleteDisabled"
          class="min-left-space"
          v-if="hasPermission(['/car/model/delete'])"
        >
          <template #icon>
            <n-icon>
              <DeleteOutlined />
            </n-icon>
          </template>
          批量删除
        </n-button>
        <n-button
          type="primary"
          @click="batchDelete"
          :disabled="batchDeleteDisabled"
          class="min-left-space"
          v-if="hasPermission(['/car/model/delete'])"
        >
          <template #icon>
            <n-icon>
              <DownCircleOutlined />
            </n-icon>
          </template>
          导出
        </n-button>
      </template>
    </BasicTable>

    <n-modal
      v-model:show="showModal"
      :show-icon="false"
      preset="dialog"
      :title="formParams.id || 0 > 0 ? '编辑车型 #' + formParams?.id : '添加车型'"
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
          <n-form-item-gi :span="6" label="车型编码" path="code">
            <n-input placeholder="输入编码" v-model:value="formParams.code" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="品牌" path="brand">
            <n-input v-model:value="formParams.brand" placeholder="请输入品牌" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="排量" path="cc">
            <n-input v-model:value="formParams.cc" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="年款" path="year">
            <n-input v-model:value="formParams.year" />
          </n-form-item-gi>
        </n-grid>
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="6" label="车型" path="model">
            <n-input placeholder="输入车型" v-model:value="formParams.model" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="车型全称">
            <n-input v-model:value="fullName" disabled />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="销售名称">
            <n-input placeholder="" v-model:value="formParams.saleName" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="国别">
            <n-select
              placeholder=""
              v-model:value="formParams.country"
              :options="countryOptions"
              filterable
            />
          </n-form-item-gi>
        </n-grid>
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="6" label="国产合资进口">
            <n-select
              v-model:value="formParams.type"
              :options="options.service_car_type"
              @update-value="handleUpdateType"
            />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="燃油类型">
            <n-select
              v-model:value="formParams.fuel"
              :options="options.service_car_fuel"
              @update-value="handleUpdateFuel"
            />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="厂家">
            <n-input v-model:value="formParams.vendor" />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="系列">
            <n-input v-model:value="formParams.series" />
          </n-form-item-gi>
        </n-grid>
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="6" label="变速箱类型">
            <n-select
              v-model:value="formParams.gearBox"
              :options="options.service_car_gearbox"
              @update-value="handleUpdateGearBox"
            />
          </n-form-item-gi>
          <n-form-item-gi :span="18" label="变速箱描述">
            <n-input v-model:value="formParams.gearDesc" />
          </n-form-item-gi>
        </n-grid>
        <n-grid :x-gap="2" :cols="24">
          <n-form-item-gi :span="6" label="前轮胎规格">
            <n-select
              v-model:value="formParams.frontTireSpec"
              :options="options.service_car_tire"
              @update-value="handleUpdateTire.bind(null, true)"
            />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="后轮胎规格">
            <n-select
              v-model:value="formParams.backTireSpec"
              :options="options.service_car_tire"
              @update-value="handleUpdateTire.bind(null, false)"
            />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="前轮毂规格">
            <n-select
              v-model:value="formParams.frontHubSpec"
              :options="options.service_car_hub"
              @update-value="handleUpdateHub.bind(null, true)"
            />
          </n-form-item-gi>
          <n-form-item-gi :span="6" label="后轮毂规格">
            <n-select
              v-model:value="formParams.backHubSpec"
              :options="options.service_car_hub"
              @update-value="handleUpdateHub.bind(null, false)"
            />
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
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onMounted, computed } from 'vue';
  import { useDialog, useMessage, FormRules } from 'naive-ui';
  import { ActionItem, BasicTable, TableAction } from '@/components/Table';
  import { BasicForm } from '@/components/Form/index';
  import { columns } from './columns';
  import { PlusOutlined, DeleteOutlined, DownOutlined, DownCircleOutlined } from '@vicons/antd';
  import { adaModalWidth } from '@/utils/hotgo';
  import { getRandomString } from '@/utils/charset';
  import { cloneDeep } from 'lodash-es';
  import { register, defaultState, CarModel, countryOptions, options, loadOptions } from './model';
  import { usePermission } from '@/hooks/web/usePermission';
  import { useUserStore } from '@/store/modules/user';
  import { LoginRoute } from '@/router';
  import { getNowUrl } from '@/utils/urlUtils';
  import { EditCar, GetCarModelList, DeleteCarModel } from '@/api/car/model';
  interface Props {
    type?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    type: '-1',
  });

  const rules: FormRules = {
    code: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入编码',
    },
    brand: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入品牌',
    },
    year: {
      required: true,
      trigger: ['blur', 'input'],
      validator: (_: any, value: any) => {
        if (value === '') {
          return Promise.reject('请输入年份');
        }
        if (!/^\d{4}$/.test(value)) {
          return Promise.reject('请输入正确的年份');
        }
        return Promise.resolve();
      },
    },
    model: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入车型',
    },
    cc: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入排量',
    },
  };

  const { hasPermission } = usePermission();
  const userStore = useUserStore();
  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const searchFormRef = ref<any>({});
  const formRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const dialogWidth = ref('100%');
  const formParams = ref<CarModel>(cloneDeep(defaultState));
  const showQrModal = ref(false);
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
              return true;
            },
            auth: ['/car/model/status'],
          },
          {
            label: '禁用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return true;
            },
            auth: ['/car/model/status'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            ifShow: () => {
              return true;
            },
            auth: ['/car/model/edit'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            ifShow: () => {
              return true;
            },
            auth: ['/car/model/delete'],
          },
        ],
        //dropDownActions: downActions,
        select: (key) => {
          if (key === 0) {
            return handleResetPwd(record);
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

  function handleUpdateType(value) {
    formParams.value.type = Number(value);
  }
  function handleUpdateBrand(value) {
    formParams.value.brand = Number(value);
  }

  function handleUpdateSeries(value) {
    formParams.value.series = Number(value);
  }

  function handleUpdateFuel(value) {
    formParams.value.fuel = Number(value);
  }
  function handleUpdateGearBox(value) {
    formParams.value.gearBox = Number(value);
  }

  function handleUpdateTire(fronted = false, value) {
    if (fronted) {
      formParams.value.frontTireSpec = Number(value);
    } else {
      formParams.value.backTireSpec = Number(value);
    }
  }

  function handleUpdateHub(fronted = false, value) {
    if (fronted) {
      formParams.value.frontHubSpec = Number(value);
    } else {
      formParams.value.backHubSpec = Number(value);
    }
  }

  function addTable() {
    showModal.value = true;
    formParams.value = cloneDeep(defaultState);
    if (!options.value.service_car_type) {
      loadOptions();
    }
  }

  const loadDataTable = async (res) => {
    adaModalWidth(dialogWidth, 1200);
    return await GetCarModelList({
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
        EditCar(formParams.value)
          .then((_res) => {
            if (_res.code != 0) {
              message.error('操作失败');
              return;
            }
            message.success('操作成功');
            setTimeout(() => {
              showModal.value = false;
              reloadTable();
            });
          })
          .catch((err) => {
            message.error('操作失败:' + err.message || '');
          });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function handleEdit(record: Recordable) {
    formParams.value = cloneDeep(record) as CarModel;
    loadOptions();
    showModal.value = true;
  }

  function handleResetPwd(record: Recordable) {
    record.password = getRandomString(12);
    dialog.warning({
      title: '警告',
      content: '你确定要重置密码？\r\n重置成功后密码为：' + record.password + '\r\n 请先保存',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {},
    });
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: `你确定要删除id为${record.id}的记录吗？`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        DeleteCarModel([record.id]).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            reloadTable();
          });
        });
      },
    });
  }

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        DeleteCarModel(checkedIds.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            reloadTable();
          });
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
    // Status(record.id, status).then((_res) => {
    //   message.success(status === 1 ? '禁用成功' : '启用成功');
    //   setTimeout(() => {
    //     reloadTable();
    //   });
    // });
  }

  function handleInviteQR(code: any) {
    const domain = getNowUrl() + '#';
    qrParams.value.qrUrl = domain + LoginRoute.path + '?scope=register&inviteCode=' + code;
    showQrModal.value = true;
  }
  onMounted(() => {});
  const fullName = computed(() => {
    const brand = formParams.value.brand;
    const series = formParams.value.series;
    const year = formParams.value.year;
    const cc = formParams.value.cc;
    const saleName = formParams.value.saleName;

    return brand + series + year + cc + saleName;
  });
</script>

<style lang="less" scoped></style>
@/api/car/model
