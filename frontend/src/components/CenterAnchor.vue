<template>
  <el-row style="height: 20px"></el-row>
  <el-row>
    <el-col>
      <span class="task-list">task-list</span>
    </el-col>
  </el-row>
  <el-row style="height: 55px"></el-row>
  <el-row justify="center" align="middle">
    <el-col :span="3">
      <div style="text-align: center;vertical-align: middle">
        <el-image
            fit="fill"
            src="/src/public/main/start.png"
            alt="start"
            class="otherImg"
        >
        </el-image>
      </div>
    </el-col>
    <el-col :span="10">
      <span class="desc">DownLoad</span>
    </el-col>
    <el-col :span="6"></el-col>
  </el-row>
  <el-row style="height: 20px"></el-row>
  <el-row justify="center" align="middle">
    <el-col :span="3">
      <div style="text-align: center;vertical-align: middle">
        <el-image
            fit="fill"
            src="/src/public/main/waiting.png"
            alt="waiting"
            class="otherImg"
        >
        </el-image>
      </div>
    </el-col>
    <el-col :span="10">
      <span class="desc">Waiting</span>
    </el-col>
    <el-col :span="6"></el-col>
  </el-row>
  <el-row style="height: 20px"></el-row>
  <el-row justify="center" align="middle">
    <el-col :span="3">
      <div style="text-align: center;vertical-align: middle">
        <el-image
            fit="fill"
            src="/src/public/main/stop.png"
            alt="stop"
            class="otherImg"
        >
        </el-image>
      </div>
    </el-col>
    <el-col :span="10">
      <span class="desc">Stop</span>
    </el-col>
    <el-col :span="6"></el-col>
  </el-row>
</template>

<script setup>

</script>

<style scoped>
.task-list{
  font-size: 22px;
  color: black;
  -moz-user-select: none;
  -ms-user-select: none;
  -webkit-user-select: none;
  user-select: none;
}
.desc{
  font-size: 18px;
  color: black;
  -moz-user-select: none;
  -ms-user-select: none;
  -webkit-user-select: none;
  user-select: none;
}
.otherImg {
  padding-top: 3px;
  width: 18px;
  height: 18px;
  margin: 0 auto;
  display: inline-block;
}
</style>