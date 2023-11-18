import { useStore } from '@/store'
import { ElMessage } from 'element-plus'

export enum MsgType {
  Offline = 0,	// 下线
  Online = 1, // 上线
  SuccessOnline = 2, // 上线成功
  Text = 3, // 文本消息
  MsgOk = 4, // 消息发送成功
  Other = 5, // 未知类型，作为字符串解析
  UserList = 6 // 获取用户列表
}

export const newMessage = (tp: MsgType, data: any = undefined) => {
  const store = useStore()

  return <Message>{
    username: store.username,
    msg_type: tp,
    data: data,
    created_at: (new Date()).toLocaleString()
  }
}

let socket: WebSocket

export function init() {
  const store = useStore()
  if (import.meta.env.DEV) {
    socket = new WebSocket(`ws://${document.location.hostname}:23456/ws`)
  } else {
    socket = new WebSocket(`ws://${window.location.host}/ws`)
  }
  socket.onopen = function () {
    store.connected = true
    // 发送上线提醒
    sendMsg(newMessage(MsgType.Online))
  }

  socket.onmessage = function (e) {
    try {
      console.log(e.data, "消息来了")
      const data = JSON.parse(e.data)
      switch (data.msg_type) {
        case MsgType.Online:
          store.addMsg(newMessage(data.msg_type, `${data.username} 上线`))
          break
        case MsgType.Offline:
          store.addMsg(newMessage(data.msg_type, `${data.username} 离线`))
          break
        case MsgType.SuccessOnline:
          store.username = data.username
          break
        case MsgType.Text:
          store.addMsg(data)
          break
        case MsgType.UserList:
          store.userlist = []
          data.data.forEach((u: string) => {
            const uu = u.split('^')
            store.addUser(uu[0], uu[1])
          })
          break
        default:
          store.addMsg(newMessage(data.msg_type, `未知消息：${data.data}`))
          break
      }
    } catch (err) {
      console.log(err, "错误");

      ElMessage.error(`解析消息：${e.data}, 错误原因：${err}`)
    }
  }
  socket.onerror = function (e) {
    console.log('错误', e)
    socket.close()
    ElMessage.error('连接服务器失败，正在重连。。。')
    if (socket.CONNECTING) return
    reConnect()
  }
  socket.onclose = function (e) {
    console.log('关闭连接', e)
    if (socket.CONNECTING) return
    reConnect()
  }
  const reConnect = () => setTimeout(() => init(), 9000)
}

export function sendMsg(msg: Message) {
  socket.send(JSON.stringify(msg))
}
