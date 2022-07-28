<template>
  <el-row class="header-title" data-wails-drag>
    <el-col :span="2" style="background-color: rgba(228, 228, 226, 0.7);"></el-col>
    <el-col :span="5" style="background: #f2f3f6;"></el-col>
    <el-col :span="17" style="background: white;"></el-col>
  </el-row>
  <el-row class="main-view">
    <el-col :span="2" class="left-view">
      <LeftAnchor/>
    </el-col>
    <el-col :span="5" class="center-view">
      <CenterAnchor/>
    </el-col>
    <el-col :span="17" class="context-view">
      <ContextAnchor/>
    </el-col>
  </el-row>
</template>

<script setup>
import LeftAnchor from "./LeftAnchor.vue";
import CenterAnchor from "./CenterAnchor.vue";
import ContextAnchor from "./ContextAnchor.vue";
</script>

<style scoped>
.header-title{
  width: 100%;
  height: 33px;
}
.main-view {
  width: 100%;
  height: calc(100% - 33px);
}

.left-view {
  height: 100%;
  min-height: 735px;
  background-color: rgba(228, 228, 226, 0.7);
}
.center-view{
  height: 100%;
  min-height: 735px;
  background: #f2f3f6;
}
.context-view{
  height: 100%;
  min-height: 735px;
  background: white;
}
</style>