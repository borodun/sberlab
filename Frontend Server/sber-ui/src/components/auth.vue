<template>
  <el-form ref="form" :model="keys" label-width="120px">
    <el-form-item label="Access Key">
      <el-input v-model="keys.accessKey"></el-input>
    </el-form-item>
    <el-form-item label="Secret Key">
      <el-input v-model="keys.secretKey" show-password></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" v-on:click="saveKeys">Save keys</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import axios from "axios"

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_IP,
  //baseURL: process.env.BACKEND,
  // baseURL: 'http://localhost:9999/v1',
  // baseURL: 'http://37.230.196.108/v1/'
});

export default {
  name: "auth",
  data() {
    return {
      keys: {
        accessKey: '123',
        secretKey: '123',
      }
    }
  },
  methods: {
    saveKeys() {
      axios_instance.post(
          "/keys",
          {
            aKey: this.keys.accessKey,
            sKey: this.keys.secretKey,
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