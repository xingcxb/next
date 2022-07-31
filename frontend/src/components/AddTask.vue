<template>
  <el-dialog
      v-model="visible"
      width="60%"
      center
      destroy-on-close
  >
    <el-tabs type="border-card" v-model="selectTab" @tab-click="tabClick">
      <!-- 链接下载 -->
      <el-tab-pane label="Link Task" name="uri">
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
        </el-form>
      </el-tab-pane>
      <!-- 种子下载 -->
      <el-tab-pane label="BT Task" name="torrent">
        <el-form-item>
          <el-upload
              class="upload-torrent"
              drag
          >
            <icon-inbox-in theme="outline" size="30" fill="#000000" :strokeWidth="1"/>
            <div class="el-upload__text">
              Drop file here or <em>click to upload</em>
            </div>
          </el-upload>
        </el-form-item>
      </el-tab-pane>
      <el-row style="height: 5px"></el-row>
      <!-- 重命名、切片 -->
      <el-row align="middle" style="height: 32px">
        <el-col :span="15">
          <el-form-item label="rename:">
            <el-input
                v-model="linkTask_form.rename"
                size="small"
                placeholder="rename"
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2"></el-col>
        <el-col :span="7" style="font-size: 16px">
          <el-form-item label="slice:">
            <el-input-number
                v-model="linkTask_form.sliceNum"
                size="small"
                :min="1"
                :max="64"
                controls-position="right"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row style="height: 10px"></el-row>
      <!-- 存储路径 -->
      <el-row>
        <el-col>
          <el-form-item label="filePath:">
            <el-input
                v-model="linkTask_form.filePath"
                size="small"
                placeholder="filePath"
            >
              <template #append>
                <icon-folder-close
                    theme="outline"
                    size="17"
                    fill="#000000"
                    :strokeWidth="1"
                    class="smallImg"
                />
              </template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <!-- 高级选项 -->
      <template v-if="linkTask_form.senior">
        <el-row>
          <el-col>
            <el-form-item label="User-Agent:">
              <el-input
                  v-model="linkTask_form.userAgent"
                  size="small"
                  placeholder="User-Agent"
              ></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col>
            <el-form-item label="Referer:">
              <el-input
                  v-model="linkTask_form.referer"
                  size="small"
                  placeholder="Referer"
              ></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col>
            <el-form-item label="Cookie:">
              <el-input
                  v-model="linkTask_form.cookie"
                  size="small"
                  placeholder="Cookie"
              ></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="12">
          <el-col :span="15" :xs="24">
            <el-form-item label="Proxy:">
              <el-input
                  v-model="linkTask_form.proxy"
                  size="small"
                  placeholder="[http://][USER:PASSWORD@]HOST[:PORT]"
              ></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="9" :xs="24">
            <div class="help-link">
              <a target="_blank" href="https://github.com/agalwood/Motrix/wiki/Proxy" rel="noopener noreferrer">
                查看代理配置说明
                <icon-link-three theme="outline" size="17" fill="#000000" :strokeWidth="1"/>
              </a>
            </div>
          </el-col>
        </el-row>
      </template>
      <!-- 提交 -->
      <el-row>
        <el-col :span="3" :xs="9">
          <el-checkbox v-model="linkTask_form.senior" label="senior"/>
        </el-col>
        <el-col :span="13"></el-col>
        <el-col :span="8" :xs="13">
          <el-button @click="addTask('link')">cancel</el-button>
          <el-button type="primary" @click="addTask('link')">submit</el-button>
        </el-col>
      </el-row>
    </el-tabs>
  </el-dialog>
</template>

<script setup>
import {reactive, ref} from "vue";

// 弹窗控件
const visible = ref(false);
// tab标签
const selectTab = ref('uri');
const linkTask_form = reactive({
  // 下载链接
  downLink: '',
  // 重命名
  rename: '',
  // 分片数量
  sliceNum: 64,
  // 高级
  senior: false,
  // userAgent
  userAgent: '',
  // referer
  referer: '',
  // cookie
  cookie: '',
  //proxy
  proxy: '',
});

function tabClick(tab, event) {
  this.getActiveName = tab.name
}

</script>

<style scoped>
.el-input-number .is-without-controls {
  width: 130px;
  line-height: 30px;
  height: 28px;
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

.help-link {
  font-size: 12px;
  line-height: 14px;
  padding-top: 7px;

> a {
  color: #909399;
}

}

.upload-torrent {
  width: 100%;

.el-upload, .el-upload-dragger {
  width: 100%;
}

.el-upload-dragger {
  border-radius: 4px;
  padding: 24px;
  height: auto;
}

.upload-inbox-icon {
  display: inline-block;
  margin-bottom: 12px;
}

.torrent-name {
  margin-top: 4px;
  font-size: $ --font-size-small;
  color: $ --color-text-secondary;
  line-height: 16px;
}

}
</style>