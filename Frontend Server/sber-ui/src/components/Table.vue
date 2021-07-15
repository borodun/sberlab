<template>
  <el-form ref="form" :model="form" label-width="120px" :inline="true">
    <el-form-item label="Offset">
      <el-input-number v-model="form.offset" :min="0"></el-input-number>
    </el-form-item>
    <el-form-item label="Limit">
      <el-input-number v-model="form.limit" :min="1"></el-input-number>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" v-on:click="showInfo" :loading="loading">{{ buttonText }}</el-button>
    </el-form-item>
  </el-form>
  <el-form :inline="false" :model="form" label-width="120px">
    <el-form-item>
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
            width="130"
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
            width="130"
            sortable>
        </el-table-column>
        <el-table-column
            label="Operations"
            fit>
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
  <el-dialog
      title="Details"
      v-model="dialogVisible"
      width="80%">
    <div>
      <vue-json-pretty :data="detail.details" :showDoubleQuotes="false" :showLength="true" :deep="2"></vue-json-pretty>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="dialogVisible = false">OK</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import axios from "axios";
import VueJsonPretty from 'vue-json-pretty';
import 'vue-json-pretty/lib/styles.css';

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_IP,
});

export default {
  name: "Table",
  data() {
    return {
      dialogVisible: false,
      detail: {
        error: "",
        details: {}
      },
      buttonText: "Query",
      loading: false,
      form: {
        offset: 0,
        limit: 10,
      },
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

  components: {
    VueJsonPretty,
  },

  methods: {
    showInfo() {
      //console.log("Offset: " + this.form.offset + " Limit: " + this.form.limit)
      this.loading = true
      this.buttonText = "Querying"
      axios_instance.get("/info/entities",
          {
            params: {
              offset: this.form.offset,
              limit: this.form.limit,
            }
          }).then(result => {
        this.loading = false
        this.buttonText = "Query"
        this.projectData = result.data
        if (this.projectData.error.length !== 0) {
          console.log("Error: " + this.projectData.error)
          this.$emit('error', this.projectData.error)
        }
      }, error => {
        console.error(error);
        this.loading = false
        this.buttonText = "Query"
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
      console.log(this.projectData.entities_array[row].id)
      axios_instance.get("/info/detail",
          {
            params: {
              id: this.projectData.entities_array[row].id,
              type: this.projectData.entities_array[row].type,
            }
          }).then(result => {
        this.detail = result.data
        if (this.detail.error.length !== 0) {
          console.log("Error: " + this.detail.error)
          this.$emit('error', this.detail.error)
        }
        this.detail.details = JSON.parse(this.detail.details)
        this.dialogVisible = true
      }, error => {
        console.error(error);
      });
    },
    deleteElement(row) {
      console.log(this.projectData.entities_array[row].id)
    }
  }
}
</script>

<style scoped>
</style>