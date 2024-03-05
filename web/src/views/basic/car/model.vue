<!-- eslint-disable vue/no-parsing-error -->
<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="车型管理" />
    </div>
    <n-card :bordered="false" class="proCard">
      <n-tabs
        type="card"
        class="card-tabs"
        :value="defaultTab"
        animated
        @before-leave="handleBeforeLeave"
      >
        <n-tab-pane
          v-for="item in [
            { key: '-1', value: '-1', label: '全部' },
            ...(options.service_car_type || []),
          ]"
          :name="item.key.toString()"
          :tab="item.label"
          :key="item.key"
          default-
        >
          <List :type="defaultTab" />
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import List from './list.vue';
  import { useRouter } from 'vue-router';
  import { options, loadOptions } from './model';
  const defaultTab = ref('-1');
  const router = useRouter();

  onMounted(() => {
    if (router.currentRoute.value.query?.type) {
      defaultTab.value = router.currentRoute.value.query.type as string;
    }
    loadOptions();
  });
  function handleBeforeLeave(tabName: string): boolean | Promise<boolean> {
    defaultTab.value = tabName;
    return true;
  }
</script>
