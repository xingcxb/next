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
        <el-row style="height: 25px" justify="start" align="middle">
          <el-col :offset="1" :span="8">
            <span
                style="font-size: 20px;
                color: black;
                text-align: left;
                align-content: center;
                width: 100%"
            >
              <!-- 下载文件名 -->
              {{ fileName }}
            </span>
          </el-col>
          <el-col :span="7"></el-col>
          <el-col :span="7" class="handle">
<!--            <div class="handle">-->
              <el-row style="height: 25px;" align="middle">
                <el-col :span="1" :offset="3">
<!--                  <div style="text-align: center;vertical-align: middle">-->
                    <el-image
                        fit="fill"
                        src="/src/public/main/start_line.png"
                        alt="start"
                        class="smallImg"
                    >
                    </el-image>
<!--                  </div>-->
                </el-col>
                <el-col :span="1" :offset="3">
                  <!--<div style="text-align: center;vertical-align: middle">-->
                    <el-image
                        fit="fill"
                        src="/src/public/main/delete.png"
                        alt="delete"
                        class="smallImg"
                    >
                    </el-image>
<!--                  </div>-->
                </el-col>
                <el-col :span="1" :offset="3">
<!--                  <div style="text-align: center;vertical-align: middle">-->
                    <el-image
                        fit="fill"
                        src="/src/public/main/file.png"
                        alt="delete"
                        class="smallImg"
                    >
                    </el-image>
<!--                  </div>-->
                </el-col>
                <el-col :span="1" :offset="3">
<!--                  <div style="text-align: center;vertical-align: middle">-->
                    <el-image
                        fit="fill"
                        src="/src/public/main/link.png"
                        alt="delete"
                        class="smallImg"
                    >
                    </el-image>
<!--                  </div>-->
                </el-col>
                <el-col :span="1" :offset="3">
<!--                  <div style="text-align: center;vertical-align: middle">-->
                    <el-image
                        fit="fill"
                        src="/src/public/main/about.png"
                        alt="delete"
                        class="smallImg"
                    >
                    </el-image>
<!--                  </div>-->
                </el-col>
              </el-row>
<!--            </div>-->
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
}

.smallImg:hover {
  cursor: pointer;
}

.things {
  height: 112px;
  width: 100%;
  border-radius: 5px;
  border: 1px solid var(--el-border-color);
  margin-top: 20px;
}

.things:hover {
  border: 1px solid #504ff3;
}

.handle{
  height: 24px;
  width: 100%;
  border-radius: 10px;
  border: 1px solid var(--el-border-color);
  margin: 0;
}
.handle:hover{
  border: 1px solid #504ff3;
  background: #504ff3;
}
</style>