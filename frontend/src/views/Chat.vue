<template>
  <div class="chat-center">
    <!-- 消息区域 -->
    <el-scrollbar class="scroll" ref="scroll">
      <div v-for="item in store.messages" :key="item.data" class="scroll-item">
        <msg-item :msg="item"></msg-item>
      </div>
    </el-scrollbar>
    <!-- 输入区 -->
    <div class="input-box">
      <div class="tools">
        <div class="tool-item" @click="onPaste">Paste</div>
        <div class="tool-item">File</div>
      </div>
      <div class="input-msg-area">
        <textarea class="msg-box" placeholder="输入要发送的消息" v-model="msg" rows="5"></textarea>
        <div class="send-btns">
          <el-button type="primary" size="small" @click="onSend">发送</el-button>
          <el-button type="danger" size="small" @click="onClear">重置</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStore } from '@/store';
import MsgItem from '@/components/MsgItem.vue'
import { watch, ref, nextTick, onActivated } from 'vue';
import { ElMessage, ElScrollbar } from 'element-plus';
import { sendMsg, newMessage, MsgType } from '@/common/websocket'
const store = useStore()
const scroll = ref<InstanceType<typeof ElScrollbar>>()
watch(() => store.messages, () => {
  nextTick(() => {
    scroll.value?.setScrollTop(scroll.value.wrapRef?.scrollHeight as number)
  })
});
onActivated(() => {
  nextTick(() => {
    scroll.value?.setScrollTop(scroll.value.wrapRef?.scrollHeight as number)
  })
})
// 发送消息
const msg = ref('')
const onSend = () => {
  if (msg.value === '') {
    ElMessage.warning('请输入要发送消息')
    return
  }
  sendMsg(newMessage(MsgType.Text, msg.value))
  onClear()
}
// 清空输入
const onClear = () => msg.value = ''
// 粘贴信息
const onPaste = () => {
  navigator.clipboard.readText().then(v => msg.value = v)
}
</script>

<style lang="scss" scoped>
textarea {
  resize: none;
  padding: 0.5rem;
  font-family: Inter, Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  -webkit-text-size-adjust: 100%;
}

.chat-center {
  height: 100%;
  display: flex;
  flex-direction: column;

  .scroll {
    flex: 1;
    background-color: rgb(212, 243, 216);
  }

  .input-box {
    display: flex;
    flex-direction: column;

    // 工具栏
    .tools {
      display: flex;

      .tool-item {
        cursor: pointer;
        border-left: 1px solid #000;
        padding: 0 5px;
      }
    }

    .tools:first-child {
      border-left: 0;
    }

    .input-msg-area {
      display: grid;
      grid-template-columns: 1fr auto;

      .send-btns {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
      }
    }
  }
}
</style>
