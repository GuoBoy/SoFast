import { defineStore } from 'pinia'
import { ref } from 'vue'

// 全局存储
export const useStore = defineStore('main', () => {

  const messages = ref<Message[]>([])

  const addMsg = (msg: Message) => {
    console.log("收到消息", msg.data);

    messages.value.push(msg)
  }
  const username = ref("")
  const connected = ref(false)
  const userlist = ref<UserItem[]>([])
  const addUser = (ua: string, ra: string) => {
    console.log("添加用户", ua, ra);

    userlist.value.push({ username: ua, remoteaddr: ra })
  }

  return {
    messages, addMsg,
    username,
    connected,
    userlist, addUser,
  }
})
