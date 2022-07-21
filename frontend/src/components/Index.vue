<template>
  <el-row class="main-view">
    <el-col :span="2" class="left-view">
      <LeftAnchor/>
    </el-col>
    <el-col :span="6"></el-col>
    <el-col :span="16"></el-col>
  </el-row>
</template>

<script setup>
import LeftAnchor from "./LeftAnchor.vue";</script>

<style scoped>
.main-view {
  width: 100%;
  height: calc(100% - 33px);
}

.left-view {
  background-color: rgba(0, 0, 0, 0.7);
}
</style>