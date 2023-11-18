<template>
  <div class="center">
    <div class="self">
      <div class="item">
        头像:
        <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
      </div>
      <div class="item">
        用户名:
        <el-tag>{{ store.username }}</el-tag>
      </div>
    </div>
    <div class="join">
      <div class="title">加入当前服务</div>
      <img :src="addr" class="img" />
      <div>
        <el-tag>{{ addr }}</el-tag>
        <el-button type="primary" size="small" @click="copyAddr">点击复制</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStore } from '@/store'
import { ElMessage } from 'element-plus';
const store = useStore()
const addr = `${import.meta.env.DEV ? "http://" + document.location.hostname + ":23456" : window.location.origin}/addr`
const copyAddr = () => {
  navigator.clipboard.writeText(addr.replace("/addr", ''))
  ElMessage.success("已复制到剪切板")
}
</script>

<style lang="scss" scoped>
.center {
  display: flex;
  align-items: center;
  flex-direction: column;
}

.self {
  display: flex;
  flex-direction: column;
  margin-top: 2vh;

  .item {
    margin-top: 1vh;
    text-align: left;
  }
}

.join {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 10vh;

  .title {
    font-size: large;
  }

  .img {
    height: 15vw;
  }
}
</style>
