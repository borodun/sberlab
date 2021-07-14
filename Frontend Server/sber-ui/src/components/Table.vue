<template>
  <el-form :inline="false" :model="form" label-width="120px">
    <el-form-item label="Info list" >
      <el-table
          :data="ents[0].entities"
          style="width: 100%"
          max-height="500"
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
            width="100">
        </el-table-column>
        <el-table-column
            prop="type"
            label="Type"
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
        <el-table-column
            label="Operations"
            fit>
          <template #header>
            <el-button type="primary" v-on:click="showInfo">Query</el-button>
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
                    type="danger">Delete</el-button>
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
    elName: String,
    type: String,
    form: Object,
  },
  data() {
    return {
      typeName: [this.type],

      ents: [{
        error: "",
        entities: [
          {
            name: "",
            id: "",
            status: "",
            type:""
          }
        ]
      }],
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
        if (this.entity[0].error.length !== 0) {
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