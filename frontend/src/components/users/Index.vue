<template>
  <div class="title">在线用户</div>
  <ul class="user-container">
    <li v-for="(u, index) in store.userlist" :key="u.remoteaddr" :class="['user', index % 2 === 0 ? 'userou' : 'userji']">
      <img src="../../assets/logo.png" class="avatar" />
      <span class="username">{{ u.username }}</span>
    </li>
  </ul>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useStore } from '@/store'
import { MsgType, newMessage, sendMsg } from '@/common/websocket'
const store = useStore()

// 发送获取用户列表请求
const getUsers = () => sendMsg(newMessage(MsgType.UserList))
onMounted(() => {
  getUsers()
})
</script>

<style lang="scss" scoped>
.title {
  font-size: 1.5rem;
  text-align: center;
}

.user-container {
  height: 100%;
  overflow-y: scroll;
  padding-left: 0;
}

.userji {
  background-color: aqua;
}

.userou {
  background-color: #a2e7e7;
}

.user {
  padding: 0.5rem 1.5rem;
  display: flex;
  align-items: center;

  .avatar {
    height: 2rem;
  }

  .username {
    margin-left: 1rem;
  }
}
</style>
