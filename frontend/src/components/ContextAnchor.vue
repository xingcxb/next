<template>
  <link href="./src/public/css/hr.css">
  <el-row style="height: 20px"></el-row>
  <el-row>
    <el-col :span="5">
      <span class="task-list">DownLoad</span>
    </el-col>
    <el-col :span="14"></el-col>
    <el-col :span="1">
      <div style="text-align: center;vertical-align: middle">
        <el-image
            fit="fill"
            src="/src/public/main/delete.png"
            alt="delete"
            class="smallImg"
        >
        </el-image>
      </div>
    </el-col>
    <el-col :span="1">
      <div style="text-align: center;vertical-align: middle">
        <el-image
            fit="fill"
            src="/src/public/main/refresh.png"
            alt="refresh"
            class="smallImg"
        >
        </el-image>
      </div>
    </el-col>
    <el-col :span="1">
      <div style="text-align: center;vertical-align: middle">
        <el-image
            fit="fill"
            src="/src/public/main/restore.png"
            alt="restore"
            class="smallImg"
        >
        </el-image>
      </div>
    </el-col>
    <el-col :span="1">
      <div style="text-align: center;vertical-align: middle">
        <el-image
            fit="fill"
            src="/src/public/main/waiting_line.png"
            alt="waiting"
            class="smallImg"
        >
        </el-image>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="22" :offset="1">
      <hr class="hr-fade-content" data-content="分隔线">
    </el-col>
  </el-row>
  <!-- 此处是循环 -->
  <el-row>
    <el-col :offset="1" :span="22">
      <div class="things">
        <el-row style="height: 10px"></el-row>
        <el-row style="height: 25px" align="middle">
          <el-col :offset="1" :span="9" style="cursor: default;color: #464d49;text-align: left">
              <!-- 下载文件名 -->
              {{ fileName }}
          </el-col>
          <el-col :span="6"></el-col>
          <el-col :span="7" class="handle">
              <el-row style="height: 25px;" align="middle">
                <el-col :span="1" :offset="3">
                  <icon-play-one theme="outline" size="24" class="smallImg" :strokeWidth="1"/>
                </el-col>
                <el-col :span="1" :offset="3">
                  <icon-close-small theme="outline" size="24" class="smallImg" :strokeWidth="1"/>
                </el-col>
                <el-col :span="1" :offset="3">
                  <icon-folder-open theme="outline" size="17" class="smallImg" :strokeWidth="1"/>
                </el-col>
                <el-col :span="1" :offset="3">
                  <icon-link-two theme="outline" size="17" class="smallImg" :strokeWidth="1"/>
                </el-col>
                <el-col :span="1" :offset="3">
                  <icon-view-list theme="outline" size="17" class="smallImg" :strokeWidth="1"/>
                </el-col>
              </el-row>
          </el-col>
        </el-row>
        <el-row style="height: 30px"></el-row>
        <el-row>
          <el-col :offset="1" :span="22">
            <el-progress
                :percentage="100"
                :text-inside="true"
                :show-text="false"
            />
          </el-col>
        </el-row>
        <el-row style="height: 10px"></el-row>
        <el-row>
          <el-col :offset="1" :span="2" style="color: #909090;font-size: 14px">
            {{downloadedSize}}MB/{{totalSize}}MB
          </el-col>
        </el-row>
      </div>
    </el-col>
  </el-row>
</template>

<script setup>
import {ref} from 'vue'

// 文件名
let fileName = ref('测试.txt')
// 已下载的大小
let downloadedSize = ref(0)
// 文件大小
let totalSize = ref(0)


</script>

<style scoped>
.task-list {
  font-size: 22px;
  color: black;
  -moz-user-select: none;
  -ms-user-select: none;
  -webkit-user-select: none;
  user-select: none;
}

.smallImg {
  width: 15px;
  height: 15px;
  color: #000000;
}

.smallImg:hover {
  cursor: pointer;
  color: #ffffff;
}

.things {
  height: 112px;
  width: 100%;
  border-radius: 5px;
  border: 1px solid var(--el-border-color);
  margin-top: 20px;
}

.things:hover {
  border: 1px solid #6dd5fa;
}

.handle{
  height: 24px;
  width: 100%;
  border-radius: 10px;
  border: 1px solid var(--el-border-color);
  margin: 0;
}
.handle:hover{
  border: 1px solid #66CCFF;
  background: #66CCFF;
}
</style>