<template>
  <el-form ref="form" :model="keys" label-width="120px" :disabled="disabled">
    <el-form-item label="Choose project">
      <el-select v-model="value" placeholder="Select" v-on:change="saveProjID" v-on:focus="ShowProjects">
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

      disabled: true,
      gotProjects: false,
      keys: {
        accessKey: "",
        secretKey: "",
        domain: "fitnsu"
      },
    }
  },
  created: function () {
    this.CheckToken()
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
    },
    ShowProjects(){
      if (!this.gotProjects) {
        axios_instance.get("/projects").then(function (response) {
          console.log(response.data)
          if (response.data.error.length === 0) {
            this.gotProjects = true
            this.projects = response.data.projects
          } else {
            this.$emit('error', response.data.error)
          }
        }.bind(this)).catch(function (error) {
          console.log(error);
        })}
    },
    CheckToken(){
      axios_instance.get("/token").then(function (response) {
        console.log(response.data)
        if (response.data.length === 0) {
          this.disabled = false
          this.$emit('success', "Token accepted")
          console.log(response.data.projects)
          this.projects = response.data.projects
        } else {
          this.$emit('error', response.data)
        }
      }.bind(this)).catch(function (error) {
        console.log(error);
      });
    }
  }
}
</script>

<style scoped>

</style>