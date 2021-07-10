<template>
  <div id="all">
    <el-row>
      <el-col :span="24">
        <div class="grid-content bg-purple-dark"><h1 id="ecsList">Get ECS list</h1></div>
      </el-col>
    </el-row>
    <Auth v-on:success="ShowSuccess($event)"></Auth>
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Project ID">
        <el-input v-model="form.projID"></el-input>
      </el-form-item>
      <el-form-item label="Offset">
        <el-input-number v-model="form.offset" :min="1"></el-input-number>
      </el-form-item>
      <el-form-item label="Limit">
        <el-input-number v-model="form.limit" :min="1"></el-input-number>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" v-on:click="showInfo">Query</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
//import TodoListItem from './TodoListItem.vue'
import axios from "axios"
import Auth from "@/components/auth";

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_IP,
  //baseURL: process.env.BACKEND,
  // baseURL: 'http://localhost:9999/v1',
  // baseURL: 'http://37.230.196.108/v1/'
});

export default {
  name: "ListECS",
  components: {Auth},
  data() {
    return {
      form: {
        projID: '',
        offset: 0,
        limit: 0,
        info: '',
        jsonInfo: {},
      },
      servers: [],
    }
  },
  methods: {
    ShowError(msg) {
      this.$notify({
        title: 'Error',
        type: 'error',
        message: msg,
        duration: 10000 // 10 sec
      });
    },
    ShowSuccess(msg) {
      this.$notify({
        type: 'success',
        title: 'Success',
        message: msg,
        duration: 5000 // 5 sec
      });
    },
    showInfo() {
      console.log(this.form)
      axios_instance.get(this.form.projID + "/cloudservers/detail?",
          {
            params: {
              offset: this.form.offset,
              limit: this.form.limit,
              aKey: this.form.accessKey,
              sKey: this.form.secretKey,
            }
          }).then(result => {
        this.form.info = result.data
        let jsonInfo = JSON.parse(result.data)
        this.form.jsonInfo = jsonInfo
        if (jsonInfo.error_msg) {
          this.ShowError(jsonInfo.error_msg)
        } else if (jsonInfo.error) {
          this.ShowError(jsonInfo.error.message)
        }
        console.log(result)
      }, error => {
        console.error(error);
      });
    },
  }
}
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
}

#all {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  color: #303133;
}

#ecsList {
  font-weight: 30;
  font-size: 25pt;
  text-align: center;
}

h2 {
  font-weight: 15;
  font-size: 15pt;
  margin: 30px 30px;
}

</style>