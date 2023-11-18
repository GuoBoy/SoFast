# 过程梳理

全局用户列表
    - 用户ws连接
    - 用户地址
    - 用户名
    - 头像

消息类型
Offline = iota	// 下线
Online // 上线
SuccessOnline // 上线成功
Text // 文本消息
MsgOk // 消息发送成功
Other // 未知类型，作为字符串解析

开启广播，每3秒发送一次
开启路由
    展示地址二维码
    ws消息

