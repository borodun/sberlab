<template>
  <el-form ref="form" :model="keys" label-width="120px" :disabled="disabled">
    <el-form-item label="Domain Name">
      <el-input v-model="keys.domain"></el-input>
    </el-form-item>
    <el-form-item label="Login">
      <el-input v-model="keys.accessKey"></el-input>
    </el-form-item>
    <el-form-item label="Password">
      <el-input v-model="keys.secretKey" show-password></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" v-on:click="saveKeys">Login</el-button>
    </el-form-item>
  </el-form>
  <el-form ref="form" :model="keys" label-width="120px" :disabled="!disabled">
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
  methods: {
    saveKeys() {
      axios_instance.post(
          "/keys",
          {
            aKey: this.keys.accessKey,
            sKey: this.keys.secretKey,
            domain: this.keys.domain
          }).then(function (response) {
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