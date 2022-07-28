<template>
  <el-dialog
      v-model="visible"
      width="60%"
  >
    <el-tabs type="border-card">
      <!-- 链接下载 -->
      <el-tab-pane label="Link Task">
        <el-form :model="linkTask_form">
          <el-row>
            <el-col>
              <el-form-item>
                <el-input
                    type="textarea"
                    v-model="linkTask_form.downLink"
                    placeholder="One task url per line (supports magnet)"
                    rows="3"
                    resize="none"
                    clearable
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row style="height: 10px"></el-row>
          <el-row align="middle" style="height: 32px">
            <el-col :span="3" style="font-size: 16px">
              rename:
            </el-col>
            <el-col :span="1"></el-col>
            <el-col :span="11">
              <el-form-item>
                <el-input
                    v-model="linkTask_form.rename"
                    placeholder="rename"
                ></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="3" style="font-size: 16px">
              slice:
            </el-col>
            <el-col :span="4">
              <el-form-item>
                <el-input-number
                    v-model="linkTask_form.sliceNum"
                    class="mx-4"
                    :min="1"
                    :max="64"
                    controls-position="right"
                />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </el-tab-pane>
      <!-- 种子下载 -->
      <el-tab-pane label="BT Task">
        种子任务
      </el-tab-pane>
    </el-tabs>
  </el-dialog>
</template>

<script setup>
import {reactive, ref} from "vue";

// 弹窗控件
const visible = ref(false);
const linkTask_form = reactive({
  // 下载链接
  downLink: '',
  // 重命名
  rename: '',
  // 分片数量
  sliceNum: 64,
});
</script>

<style scoped>
.el-input-number .is-without-controls {
  width: 130px;
  line-height: 30px;
  height: 28px;
}
</style>