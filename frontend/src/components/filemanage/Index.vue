<template>
  <div class="filemanage">
    <div class="search-box">
      <el-input v-model="searchKey" placeholder="输入文件名" clearable size="small" @input="filteFile"></el-input>
      <el-button size="small" @click="refreshFiles">刷新</el-button>
      <el-button size="small" type="warning" @click="showDialog = true; uploadFiles = []">上传文件</el-button>
    </div>
    <el-table :data="files" stripe size="small" style="width: 100%;height: 100%;">
      <el-table-column type="index" />
      <el-table-column prop="filename" label="name"></el-table-column>
      <el-table-column prop="filesize" label="size"></el-table-column>
      <el-table-column>
        <template #default="scope">
          <el-button @click="onDownload(scope.row)" size="small">下载</el-button>
          <el-button type="danger" size="small" @click="onDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
  <el-dialog title="上传文件" v-model="showDialog">
    <el-upload class="upload-demo" :action="uploadUrl" :on-error="onUploadErr" :file-list="uploadFiles">
      <el-button size="small" type="primary">点击上传</el-button>
    </el-upload>
    <template #footer>
      <el-button type="primary" @click="showDialog = false; refreshFiles()">完成</el-button>
    </template>
  </el-dialog>
</template>
<script lang="ts" setup>
import { ElLoading, ElMessage, ElMessageBox, ElUpload } from 'element-plus'
import { onMounted, ref } from 'vue'
import { Request } from '@/service'
import axios from 'axios'
// 格式化文件大小
const formatFileSize = (fileSize: number) => {
  if (fileSize < 1024) {
    return fileSize + 'B';
  } else if (fileSize < (1024 * 1024)) {
    var temp = fileSize / 1024;
    return `${temp.toFixed(2)}KB`
  } else if (fileSize < (1024 * 1024 * 1024)) {
    var temp = fileSize / (1024 * 1024);
    return `${temp.toFixed(2)}MB`
  } else {
    var temp = fileSize / (1024 * 1024 * 1024);
    return `${temp.toFixed(2)}GB`
  }
}
// 文件列表
const fileList = ref<FileItem[]>([])
const files = ref<FileItem[]>([])
const baseUrl = import.meta.env.DEV ? `http://${document.location.hostname}:23456` : window.location.origin
const refreshFiles = () => {
  const loading = ElLoading.service()
  fileList.value = []
  Request.get(`${import.meta.env.DEV ? `http://${document.location.hostname}:23456` : window.location.origin}/files`).then(res => {
    if (res.code !== 200) {
      ElMessage.error('获取文件列表失败')
      return
    }
    fileList.value = res.data.map((e: FileItem) => {
      e.filesize = formatFileSize(e.filesize as number)
      return e
    })
    files.value = fileList.value
    loading.close()
  }).catch(e => {
    ElMessage.error(e)
  })
}

onMounted(() => {
  refreshFiles()
})

const searchKey = ref('')
const filteFile = () => {
  files.value = fileList.value.filter(({ filename }) => !searchKey.value || filename.toLowerCase().includes(searchKey.value.toLowerCase()))
}

const onDownload = (f: FileItem) => {
  const ele = document.createElement('a')
  ele.setAttribute('href', `http://${document.location.hostname}:23456/download/${f.filename}`)
  ele.target = '_blank'
  ele.style.display = 'none'
  document.body.appendChild(ele)
  ele.click()
  document.body.removeChild(ele)
}

const showDialog = ref(false)
const uploadUrl = import.meta.env ? `http://${document.location.hostname}:23456/upload` : '/upload'
const onUploadErr = (err: any, file: { name: any; }, _: any) => {
  ElMessage.error(`${file.name}上传失败,${err}`)
}
const uploadFiles = ref([])

const onDelete = (f: FileItem) => {
  ElMessageBox({
    title: '真的要删除吗？',
    showCancelButton: true
  }).then(() => {
    axios.delete(`${baseUrl}/delete/${f.filename}`, {
      timeout: 3000
    }).then(res => {
      if (res.data.code === 200) {
        ElMessage.success('删除成功')
        refreshFiles()
      } else {
        ElMessage.error('删除失败' + res.data.msg)
      }
    })
  }).catch(() => {
    ElMessage.warning('已取消删除')
  })
}
</script>
<style lang="scss" scoped>
.filemanage {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-rows: auto 1fr;
}

.search-box {
  display: flex;
  padding: 0.1rem;
}
</style>
