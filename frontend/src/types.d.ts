interface Message {
  username: string,
  msg_type: number,
  data: any,
  created_at: string
}

interface UserItem {
  username: string,
  remoteaddr: string
}

interface FileItem {
  filename: string,
  filesize: string | number
}
