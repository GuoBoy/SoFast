<template>
  <div class="chat-history">
    <div class="tools">
      <!-- <el-button type="primary" @click="" size="small">刷新</el-button> -->
    </div>
    <el-scrollbar class="scroll" ref="scroll">
      <div v-for="item in store.messages" :key="item.data" class="scroll-item">
        <msg-item :msg="item"></msg-item>
      </div>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { useStore } from '@/store';
import MsgItem from '@/components/MsgItem.vue'
import { watch, ref, nextTick, onActivated } from 'vue';
import { ElScrollbar } from 'element-plus';
const store = useStore()
const scroll = ref<InstanceType<typeof ElScrollbar>>()
watch(store.messages, () => {
  nextTick(() => {
    scroll.value?.setScrollTop(scroll.value.wrapRef?.scrollHeight as number)
  })
});
onActivated(() => {
  nextTick(() => {
    scroll.value?.setScrollTop(scroll.value.wrapRef?.scrollHeight as number)
  })
})
</script>

<style lang="scss" scoped>
.chat-history {
  height: 100%;
  display: flex;
  flex-direction: column;

  .scroll {
    flex: 1;
    background-color: rgb(212, 243, 216);
  }
}
</style>
