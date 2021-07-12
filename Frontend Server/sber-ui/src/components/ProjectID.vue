<template>
  <el-form ref="form" :model="projectID" label-width="120px">
    <el-form-item label="Project ID">
      <el-input v-model="projectID"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" v-on:click="saveKeys">Save</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import axios from "axios"

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_IP,
});

export default {
  name: "ProjectID",
  data() {
    return {
      projectID: "",
    }
  },
  methods: {
    saveKeys() {
      axios_instance.post(
          "/projid",
          {
            projectID: this.projectID
          }).then(function (response) {
        console.log(response.data)
        this.$emit('success', response.data)
      }.bind(this)).catch(function (error) {
        console.log(error);
      });
    }
  }
}
</script>

<style scoped>

</style>