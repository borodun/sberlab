<template>
  <el-table
      :data="content"
      style="width: 100%"
      max-height="500">
    <el-table-column
        prop="name"
        label="Name"
        width="200">
    </el-table-column>
    <el-table-column
        fixed
        prop="id"
        label="ID"
        width="300">
    </el-table-column>
    <el-table-column
        prop="status"
        label="Status"
        width="120">
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "Table",
  props: {
    title: String,
    count: Number,
    content: Object,
    callback: Function,
  }
}
</script>

<style scoped>

</style>