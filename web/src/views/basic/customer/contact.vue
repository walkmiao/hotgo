<template>
  <div>
    <BasicTable
      :pagination="false"
      :openChecked="true"
      :columns="contactColumns"
      :request="loadDataTable"
      ref="actionRef"
      @edit-end="editRowEnd"
      @update:checked-row-keys="onCheckedRow"
      :scroll-x="1800"
    >
      <template #tableTitle>
        <n-button
          type="primary"
          @click="addRow"
          class="min-left-space"
          v-if="hasPermission(['/member/edit'])"
        >
          <template #icon>
            <n-icon>
              <PlusOutlined />
            </n-icon>
          </template>
          添加联系人
        </n-button>
        <n-button
          type="error"
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
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, onBeforeUnmount, onBeforeMount } from 'vue';
  import { useMessage } from 'naive-ui';
  import { BasicTable } from '@/components/Table';
  import { GetCustomerContactList } from '@/api/customer/customer';
  import { contactColumns } from './columns';
  import { PlusOutlined, DeleteOutlined } from '@vicons/antd';

  import { adaModalWidth } from '@/utils/hotgo';

  import { EditCustomer } from '@/api/customer/customer';
  import { defaultInfo } from './model';
  import bus from '@/libs/eventBus';
  import { loadOptions } from './model';
  import { usePermission } from '@/hooks/web/usePermission';

  interface Props {
    type?: string;
    customerId?: number;
  }

  const props = withDefaults(defineProps<Props>(), {
    customerId: 1000,
  });

  const { hasPermission } = usePermission();

  const message = useMessage();
  const actionRef = ref();
  const showModal = ref(false);
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const dialogWidth = ref('100%');
  defineExpose({
    actionRef,
  });
  //全局bus
  function editCustomer(formParams: any) {
    let contacts = actionRef.value?.getDataSource() as any[];
    contacts.forEach((item, _) => {
      delete item['cancelCbs'];
      delete item['editValueRefs'];
      delete item['onCancelEdit'];
      delete item['onEdit'];
      delete item['onSubmitEdit'];
      delete item['submitCbs'];
      delete item['validCbs'];
    });

    const contactJson = JSON.stringify(contacts);
    formParams.contacts = contactJson;
    EditCustomer({ customer: formParams })
      .then((res) => {
        if (!res) {
          message.error('返回res为空');
          return;
        }
        console.log('res:', res.data);
        if (res.code != 0) {
          message.error('操作失败:', res.message);
          return;
        }
        message.success('操作成功');
        setTimeout(() => {
          showModal.value = false;
          reloadTable();
        });
      })
      .catch((e) => {
        message.error('操作失败:', e.message);
      })
      .finally(() => console.log('finally'));
  }

  onBeforeMount(() => {
    // 监听全局事件
    bus.on('editCustomer', editCustomer);
    console.log('监听全局事件');
  });
  // 在组件卸载之前移除侦听;
  onBeforeUnmount(() => {
    console.log('移除event editCustomer!');
    bus.off('editCustomer', editCustomer);
  });

  function addRow() {
    const tableRef = actionRef.value?.tableElRef;
    tableRef.data.push(defaultInfo);
  }

  const loadDataTable = async (_) => {
    adaModalWidth(dialogWidth, 1200);
    return await GetCustomerContactList(props.customerId);
  };

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function editRowEnd(row: Recordable) {
    let contacts = actionRef.value?.getDataSource() as any[];

    if (row.record?.firstChoose) {
      contacts.forEach((item) => {
        if (item.name != row.record.name) {
          item.firstChoose = false;
        }
      });
    }
  }

  onMounted(async () => {
    await loadOptions();
  });
</script>

<style lang="less" scoped></style>
