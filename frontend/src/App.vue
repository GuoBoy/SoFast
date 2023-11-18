<template>
  <el-menu class="c-top" :default-active="activeIndex" router>
    <el-menu-item v-for="m in menus" :index="m.path" :key="m.path">{{ m.title }}</el-menu-item>
  </el-menu>
  <el-card class="c-body">
    <router-view v-slot="{ Component }">
      <keep-alive include="chat">
        <component :is="Component"></component>
      </keep-alive>
    </router-view>
  </el-card>
</template>

<script  lang="ts" setup>
import { init } from '@/common/websocket'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const menus = ref([
  { path: '/chat', title: '聊天室' },
  { path: "/history", title: "聊天记录" },
  { path: '/users', title: '在线用户' },
  { path: '/user', title: '个人中心' },
  { path: '/filemanage', title: '文件共享' }])
const activeIndex = ref(menus.value[0].path)
const router = useRouter()
router.push(menus.value[0].path)
onMounted(() => {
  init()
})
</script>

<style lang="scss">
html,
body,
#app {
  margin: 0;
  height: 100%;
}

body {
  background: 100% 100% url("./assets/bg.jpg") no-repeat;
}

// 核心内容
#app {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  width: 80%;
  min-width: 375px;
  display: grid;
  grid-template-columns: auto 1fr;

  // 导航栏
  .c-top {
    justify-content: center;
  }

  // 消息区域
  .el-card__body {
    padding: 0 0.1rem;
    height: 100%;
  }
}
</style>
