<template>
  <!-- logo -->
  <el-row>
    <el-col>
      <div style="text-align: center; vertical-align: middle" @click="toSite">
        <el-image fit="fill" class="logoImg" src="/src/public/logo.png" alt="go website"></el-image>
      </div>
    </el-col>
  </el-row>
  <el-row style="height: 30px" data-wails-drag></el-row>
  <!-- 任务列表 -->
  <el-row>
    <el-col>
      <div style="text-align: center;vertical-align: middle">
        <el-image fit="fill" class="otherImg" src="/src/public/main/download.png" alt="download"></el-image>
      </div>
    </el-col>
  </el-row>
  <el-row style="height: 30px" data-wails-drag></el-row>
  <!-- 添加任务 -->
  <el-row>
    <el-col>
      <div style="text-align: center;vertical-align: middle" >
        <el-image
            @click="createDownLoad()"
            fit="fill"
            class="otherImg"
            src="/src/public/main/add.png"
            alt="add"></el-image>
      </div>
    </el-col>
  </el-row>
  <el-row style="height: 365px" data-wails-drag></el-row>
  <!-- 设置 -->
  <el-row>
    <el-col>
      <div style="text-align: center;vertical-align: middle">
        <el-image fit="fill" class="otherImg" src="/src/public/main/setting.png" alt="setting"></el-image>
      </div>
    </el-col>
  </el-row>
  <el-row style="height: 30px" data-wails-drag></el-row>
  <!-- 关于 -->
  <el-row>
    <el-col>
      <div style="text-align: center;vertical-align: middle">
        <el-image fit="fill" class="otherImg" src="/src/public/main/about.png" alt="about"></el-image>
      </div>
    </el-col>
  </el-row>
  <AddTask
      v-model = "visible"
  ></AddTask>
</template>

<script setup>
import {ref} from "vue";
import {BrowserOpenURL} from "../../wailsjs/runtime";
import AddTask from "./AddTask.vue";
import {GetAppInfo} from "../../wailsjs/go/main/App";

let visible = ref(false);

// 打开我的主页
function toSite() {
  GetAppInfo().then(function (result) {
    const appInfo = JSON.parse(result);
    BrowserOpenURL(appInfo.webSite);
  });
}

// 新建任务
function createDownLoad() {
  visible.value = true;
}
</script>

<style scoped>
.logoImg {
  width: 70px;
  height: 70px;
  margin: 0 auto;
  display: inline-block;
}

.otherImg {
  width: 40px;
  height: 40px;
  margin: 0 auto;
  display: inline-block;
}
</style>