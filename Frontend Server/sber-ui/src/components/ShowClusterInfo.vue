<template>
  <div id="all">
    <el-row>
      <el-col :span="24">
        <div class="grid-content bg-purple-dark"><h1 id="ecsList">Brief info about project</h1></div>
      </el-col>
    </el-row>
    <Auth v-on:success="ShowSuccess($event)"></Auth>
    <ProjectID v-on:success="ShowSuccess($event)"></ProjectID>
    <el-form ref="form" :model="form" label-width="120px">
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
    <Table v-bind:content="info.ecs.servers"></Table>
    <Table v-bind:content="info.vpcs.vpcs"></Table>
    <el-collapse accordion>
      <el-collapse-item title="Raw Data" name="1">
        <div><span style="white-space: pre-wrap;">{{ this.info }}</span></div>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import axios from "axios"
import Auth from "@/components/auth";
import ProjectID from "@/components/ProjectID";
import Table from "@/components/Table";

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_IP,
});

export default {
  name: "Info",
  components: {
    Auth,
    Table,
    ProjectID
  },
  data() {
    return {
      form: {
        projID: '',
        offset: 0,
        limit: 0,
      },
      info: {
        ecs: {
          error_msg: "",
          error: {
            message: ""
          },
          count: 4,
          servers: [
            {
              name: "",
              id: "",
              status: ""
            }
          ]
        },
        vpcs: {
          error_msg: "",
          error: {
            message: ""
          },
          count: 0,
          vpcs: [
            {
              name: "",
              id: "",
              status: ""
            }
          ]
        }
      },
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
      axios_instance.get(this.form.projID + "/getinfo?",
          {
            params: {
              offset: this.form.offset,
              limit: this.form.limit,
            }
          }).then(result => {
        this.info = result.data
        if (this.info.error_msg.length !== 0) {
          this.ShowError(this.info.error_msg)
          this.info = {}
        } else if (this.info.error.message.length !== 0) {
          if (this.info.error.message) {
            this.ShowError(this.info.error.message)
          }
          this.info = {}
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