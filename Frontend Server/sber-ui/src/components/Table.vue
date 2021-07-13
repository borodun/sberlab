<template>
  <el-form :inline="false" :model="form" label-width="120px">
    <el-form-item :label="'Query ' + type">
      <el-button type="primary" v-on:click="showInfo">Query</el-button>
    </el-form-item>
    <el-form-item label="List of type">
      <el-table
          :data="entity[elName]"
          style="width: 100%"
          max-height="500"
          @selection-change="handleSelectionChange"
          border>
        <el-table-column
            type="selection"
            width="55">
        </el-table-column>
        <el-table-column
            prop="name"
            label="Name"
            width="100">
        </el-table-column>
        <el-table-column
            fixed
            prop="id"
            label="ID"
            width="155">
        </el-table-column>
        <el-table-column
            prop="status"
            label="Status"
            width="120">
        </el-table-column>
      </el-table>
    </el-form-item>
    <el-form-item>
      <el-popconfirm title="Are you sure to delete this?" v-on:confirm="deleteElements()">
        <template #reference>
          <el-button>Delete</el-button>
        </template>
      </el-popconfirm>
    </el-form-item>
  </el-form>
</template>

<script>
import axios from "axios";

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_IP,
});

export default {
  name: "Table",
  props: {
    elName: String,
    type: String,
    form: Object,
  },
  data() {
    return {
      entity: {
        error: "",
        count: 4,
        [this.elName]: [
          {
            name: "",
            id: "",
            status: ""
          }
        ]
      },
      multipleSelection: []
    }
  },

  methods: {
    showInfo() {
      console.log("Offset: " + this.form.offset + " Limit: " + this.form.limit)
      axios_instance.get("/" + this.type + "?",
          {
            params: {
              offset: this.form.offset,
              limit: this.form.limit,
            }
          }).then(result => {
        this.entity = result.data
        if (this.entity.error.length !== 0) {
          console.log("Error: " + this.entity.error)
          this.$emit('error', result.data.error)
        }
        console.log(result)
      }, error => {
        console.error(error);
      });
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    deleteElements() {
      for (var i = 0; i < this.multipleSelection.length; i++) {
        console.log(this.multipleSelection[i].id)
      }
    }
  }
}
</script>

<style scoped>
</style>