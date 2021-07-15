<template>
  <el-form :inline="false" :model="form" label-width="120px">
    <el-form-item label="Info list">
      <el-table
          :data="projectData.entities_array"
          style="width: 100%"
          @selection-change="handleSelectionChange"
          stripe
          fit
          border>
        <el-table-column
            type="selection"
            width="40">
        </el-table-column>
        <el-table-column
            prop="name"
            label="Name"
            width="100"
            sortable>
        </el-table-column>
        <el-table-column
            prop="type"
            label="Type"
            width="100"
            sortable>
        </el-table-column>
        <el-table-column
            fixed
            prop="id"
            label="ID"
            width="155"
            sortable>
        </el-table-column>
        <el-table-column
            prop="status"
            label="Status"
            width="120"
            sortable>
        </el-table-column>
        <el-table-column
            label="Operations"
            fit>
          <template #header>
            <el-button type="primary" v-on:click="showInfo" :loading="loading">Query</el-button>
          </template>
          <template #default="scope">
            <el-button
                size="mini"
                @click="showDetail(scope.$index)">Details
            </el-button>
            <el-popconfirm title="Are you sure to delete this?" v-on:confirm="deleteElement(scope.$index)">
              <template #reference>
                <el-button
                    size="mini"
                    type="danger">Delete
                </el-button>
              </template>
            </el-popconfirm>
          </template>
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
    form: Object,
  },
  data() {
    return {
      loading: false,
      projectData: {
        error: "",
        entities_array: [{
          name: "",
          id: "",
          status: "",
          type: ""
        }],
        multipleSelection: []
      }
    }
  },

  methods: {
    showInfo() {
      //console.log("Offset: " + this.form.offset + " Limit: " + this.form.limit)
      this.loading = true
      axios_instance.get("/entities",
          {
            params: {
              offset: this.form.offset,
              limit: this.form.limit,
            }
          }).then(result => {
        this.loading = false
        this.projectData = result.data
        if (this.projectData.error.length !== 0) {
          console.log("Error: " + this.projectData.error)
          this.$emit('error', this.projectData.error)
        }
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
    },
    showDetail(row) {
      console.log(this.entity[this.elName][row].id)
    },
    deleteElement(row) {
      console.log(this.entity[this.elName][row].id)
    }
  }
}
</script>

<style scoped>
</style>