<template>
  <div id="all">
    <el-row>
      <el-col :span="24">
        <div class="grid-content bg-purple-dark"><h1 id="ecsList">Brief info about project entities</h1></div>
      </el-col>
    </el-row>
    <Auth v-on:success="ShowSuccess($event)" v-on:error="ShowError($event)"></Auth>
    <Table v-on:success="ShowSuccess($event)" v-on:error="ShowError($event)"></Table>
    <el-collapse accordion>
      <el-collapse-item title="Raw Data" name="1">
        <div><span style="white-space: pre-wrap;">{{ this.info }}</span></div>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import Auth from "@/components/LoginAndProject";
import Table from "@/components/Table";


export default {
  name: "Info",
  components: {
    Auth,
    Table,
  },
  data() {
    return {
      form: {
        offset: 0,
        limit: 10,
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