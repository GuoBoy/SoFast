<template>
  <!-- 全局消息 -->
  <div class="global-msg" v-if="msg.username === undefined || ''">{{ msg.data }}</div>
  <!-- 本人消息 -->
  <div class="msgitem" style="justify-content: flex-end;" v-else-if="store.username === msg.username">
    <div class="msg-container">
      <div class="location" style="text-align: right;">{{ msg.username }}</div>
      <div class="msg">{{ msg.data }}</div>
    </div>
    <img :src="loadImg('logo.png')" class="avatar" />
  </div>
  <!-- 其它人消息 -->
  <div class="msgitem" v-else>
    <img :src="loadImg('logo.png')" class="avatar" />
    <div class="msg-container">
      <div class="location">{{ msg.username }}</div>
      <div class="msg">{{ msg.data }}</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useStore } from '@/store'
const props = defineProps<{
  msg: Message
}>()
const store = useStore()
const loadImg = (filename: string) => new URL(`../assets/${filename}`, import.meta.url).href
</script>

<style lang="scss" scoped>
.global-msg {
  text-align: center;
  font-size: 0.5rem;
}

.msgitem {
  // border: 1px solid #000;
  display: flex;
  margin-bottom: 2vh;

  .msg-container {
    max-width: 70%;

    .location {
      font-size: 0.5rem;
    }

    .msg {
      width: fit-content;
      word-wrap: break-word;
      background-color: rgb(57, 195, 137);
      padding: 3px 5px;
    }
  }

  .avatar {
    height: 3vh;
    border: 1px #c3c2c2 solid;
    border-radius: 2px;
    margin: 0.5rem;
  }
}
</style>
