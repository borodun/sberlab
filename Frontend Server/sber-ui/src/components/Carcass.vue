<template>
  <el-tabs id="h" type="border-card">
    <el-tab-pane label="ECS"><ListECS></ListECS></el-tab-pane>
    <el-tab-pane label="VPC">Nothing there</el-tab-pane>
  </el-tabs>
</template>

<script>
import ListECS from './listECS.vue'

export default {
  name: "List",
  components: {
    ListECS,
  }
}
</script>

<style>
#h {
  width: 60%;
  margin: 100px auto 0;
}
</style>
