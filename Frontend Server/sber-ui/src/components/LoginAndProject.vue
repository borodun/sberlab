<template>
  <el-form ref="form" :model="keys" label-width="120px">
    <el-form-item label="Choose project">
      <el-select v-model="value" placeholder="Select" v-on:change="saveProjID">
        <el-option
            v-for="item in projects"
            :key="item.id"
            :label="item.name"
            :value="item.id">
          <span style="float: left">{{ item.name }}</span>
          <span style="float: right; color: #8492a6; font-size: 13px">{{ item.id }}</span>
        </el-option>
      </el-select>
    </el-form-item>
  </el-form>
</template>

<script>
import axios from "axios"

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_IP,
});

export default {
  name: "auth",
  data() {
    return {
      projects: [{
        id: '',
        name: ''
      }],

      value: '',

      disabled: false,
      keys: {
        accessKey: "",
        secretKey: "",
        domain: "fitnsu"
      },
    }
  },
  created: function () {
    axios_instance.get("/projects").then(function (response) {
      console.log(response.data)
      if (response.data.error.length === 0) {
        this.disabled = true
        this.$emit('success', "Successfully authenticated")
        console.log(response.data.projects)
        this.projects = response.data.projects
      } else {
        this.$emit('error', response.data.error)
      }
    }.bind(this)).catch(function (error) {
      console.log(error);
    });
  },
  methods: {
    saveProjID() {
      axios_instance.post(
          "/projid",
          {
            projectID: this.value
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