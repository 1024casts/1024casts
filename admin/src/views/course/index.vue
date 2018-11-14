<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input :placeholder="$t('course.name')" v-model="listQuery.name" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
      <el-select v-model="listQuery.update_status" :placeholder="$t('course.updateStatus')" clearable class="filter-item" style="width: 130px">
        <el-option v-for="item in updateStatusOptions" :key="item.key" :label="item.display_name+'('+item.key+')'" :value="item.key"/>
      </el-select>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">{{ $t('table.search') }}</el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-plus" @click="handleCreate">{{ $t('table.add') }}</el-button>
      <el-button v-waves :loading="downloadLoading" class="filter-item" type="primary" icon="el-icon-download" @click="handleDownload">{{ $t('table.export') }}</el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      stripe
      highlight-current-row
      height="550"
      style="width: 100%;">
      <el-table-column fixed :label="$t('table.id')" align="center" width="65">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('course.coverImage')" width="110px" align="center">
        <template slot-scope="scope">
          <img class="user-avatar" style="height:40px;" :src=scope.row.cover_image>
        </template>
      </el-table-column>
      <el-table-column :label="$t('course.name')" width="280px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.type')" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.type }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('course.isPublish')" class-name="status-col" width="100">
        <template slot-scope="scope">
          <el-tag :type="scope.row.is_publish">{{ scope.row.is_publish | publishFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="$t('course.updateStatus')" class-name="status-col" width="100">
        <template slot-scope="scope">
          <el-tag :type="scope.row.update_status">{{ scope.row.update_status | updateStatusFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.createTime')" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.created_at }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.updateTime')" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.updated_at }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" :label="$t('table.actions')" align="center" width="400" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" icon="el-icon-edit" @click="handleUpdate(scope.row)">{{ $t('table.edit') }}</el-button>
          <el-button type="primary" size="small" icon="el-icon-tickets" @click="handleSection(scope.row)">{{ $t('course.sectionManager') }}</el-button>
          <el-button type="primary" size="small" icon="el-icon-tickets" @click="handleVideo(scope.row)">{{ $t('course.videoManager') }}</el-button>
          <el-button v-if="scope.row.status!='deleted'" size="mini" icon="el-icon-delete" type="danger" @click="handleModifyStatus(scope.row,'deleted')">{{ $t('table.delete') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container">
      <el-pagination v-show="total>0" :current-page="listQuery.page" :page-sizes="[10,20,30, 50]" :page-size="listQuery.limit" :total="total" background layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange" @current-change="handleCurrentChange"/>
    </div>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="80px" style="width: 400px; margin-left:50px;">
        <el-form-item :label="$t('course.type')" prop="type">
          <el-select v-model="temp.type" class="filter-item" placeholder="Please select">
            <el-option v-for="item in typeOptions" :key="item.key" :label="item.display_name" :value="item.key"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('course.name')" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item :label="$t('course.updateStatus')">
          <el-select v-model="temp.update_status" class="filter-item" placeholder="Please select">
            <el-option v-for="item in updateStatusOptions" :key="item.key" :label="item.display_name" :value="item.key"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('table.description')">
          <el-input :autosize="{ minRows: 2, maxRows: 4}" v-model="temp.description" type="textarea" placeholder="Please input"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">{{ $t('table.confirm') }}</el-button>
      </div>
    </el-dialog>

    <el-dialog :title="sectiontTitle" :visible.sync="dialogSectionVisible">
      <el-table
        v-loading="sectionListLoading"
        :key="sectionTableKey"
        :data="sectionList"
        border
        style="width: 100%">
        <el-table-column prop="id" label="ID" align="center" width="60">
        </el-table-column>
        <el-table-column prop="name" label="名称" align="center" width="200">
          <template slot-scope="scope">
            <template v-if="scope.row.edit">
              <el-input v-model="scope.row.name" class="edit-input" size="small"/>
              <el-button class="cancel-btn" size="small" icon="el-icon-refresh" type="warning" @click="cancelEdit(scope.row)">cancel</el-button>
            </template>
            <span v-else>{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column
        prop="order"
        label="排序" align="center"
        width="60">
        </el-table-column>
        <el-table-column
        prop="created_at"
        label="创建时间" align="center"
        width="150">
        </el-table-column>
        <el-table-column prop="updated_at" label="更新时间" align="center" width="150">
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="100">
          <template slot-scope="scope">
            <el-button v-if="scope.row.edit" type="success" size="small" icon="el-icon-circle-check-outline" @click="confirmEdit(scope.row)">Ok</el-button>
            <el-button v-else type="primary" size="small" icon="el-icon-edit" @click="scope.row.edit=!scope.row.edit">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="addSection(temp.id)">{{ $t('table.add') }}</el-button>
      </div>
    </el-dialog>

    <el-dialog :title="videoTitle" :visible.sync="dialogVideoVisible">
      <el-table
        v-loading="videoListLoading"
        :key="videoTableKey"
        :data="videoList"
        border
        height="600"
        style="width: 100%">
        <el-table-column prop="id" label="ID" align="center" width="60">
        </el-table-column>
        <el-table-column prop="name" label="名称" align="center" width="200">
          <template slot-scope="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="duration" label="时长" align="center" width="50">
          <template slot-scope="scope">
            <span>{{ scope.row.duration }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="section" label="section" align="center" width="150">
          <template slot-scope="scope">
            <span>{{ scope.row.section }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="published_at" label="发布时间" align="center" width="150">
          <template slot-scope="scope">
            <span>{{ scope.row.published_at }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" align="center" width="150">
        </el-table-column>
        <el-table-column prop="updated_at" label="更新时间" align="center" width="150">
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="100">
          <template slot-scope="scope">
            <el-button v-if="scope.row.edit" type="success" size="small" icon="el-icon-circle-check-outline" @click="confirmEdit(scope.row)">Ok</el-button>
            <el-button v-else type="primary" size="small" icon="el-icon-edit" @click="scope.row.edit=!scope.row.edit">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="addSection(temp.id)">{{ $t('table.add') }}</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import { fetchList, fetchSectionList, createCourse, updateCourse, updateSection } from '@/api/course'
import { fetchVideoList } from '@/api/video'
import waves from '@/directive/waves' // 水波纹指令
import { parseTime } from '@/utils'
const typeOptions = [
  { key: 'backend', display_name: '后端' },
  { key: 'frontend', display_name: '前端' },
  { key: 'service', display_name: '服务' },
  { key: 'tool', display_name: '工具' }
]

const publishOptions = [
  {key: '0', display_name: '否'},
  {key: '1', display_name: '是'}
]

const publishKeyValue = publishOptions.reduce((acc, cur) => {
  acc[cur.key] = cur.display_name
  return acc
}, {})

const updateStatusOptions = [
  { key: '-1', display_name: '全部' },
  { key: '0', display_name: '初始化' },
  { key: '1', display_name: '预告' },
  { key: '2', display_name: '更新中' },
  { key: '3', display_name: '已完结' }
]
// arr to obj ,such as { CN : "China", US : "USA" }
const updateStatusKeyValue = updateStatusOptions.reduce((acc, cur) => {
  acc[cur.key] = cur.display_name
  return acc
}, {})
export default {
  name: 'ComplexTable',
  directives: {
    waves
  },
  filters: {
    statusFilter (status) {
      const statusMap = {
        1: 'success',
        0: 'danger'
      }
      return statusMap[status]
    },
    publishFilter (type) {
      return publishKeyValue[type]
    },
    updateStatusFilter (updateStatus) {
      return updateStatusKeyValue[updateStatus]
    }
  },
  data () {
    return {
      tableKey: 0,
      list: null,
      total: null,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        title: undefined,
        update_status: -1,
        sort: '+id'
      },
      sectionTableKey: 0,
      sectionList: null,
      sectionTotal: null,
      sectionListLoading: true,
      sectiontTitle: 'section',
      videoTableKey: 0,
      videoList: null,
      videoTotal: null,
      videoListLoading: true,
      videoTitle: 'video',
      typeOptions,
      updateStatusOptions,
      showReviewer: false,
      temp: {
        id: undefined,
        importance: 1,
        description: '',
        timestamp: new Date(),
        name: '',
        type: '',
        update_status: '-1'
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      dialogSectionVisible: false,
      dialogVideoVisible: false,
      pvData: [],
      rules: {
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      downloadLoading: false
    }
  },
  created () {
    this.getList()
  },
  methods: {
    getList () {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        console.log('get user list', response)
        response = response.data
        this.list = response.data.list
        this.total = response.data.totalCount
        this.listLoading = false
      })
    },
    getSectionList (row) {
      this.sectionListLoading = true
      console.log('get section list start...')
      fetchSectionList(row.id).then(response => {
        console.log('get section list', response)
        response = response.data
        var items = []
        items = response.data.list
        this.sectionList = items
        this.sectionTotal = response.data.totalCount

        this.sectionList = items.map(v => {
          this.$set(v, 'edit', false) // https://vuejs.org/v2/guide/reactivity.html
          v.originalName = v.name //  will be used when user click the cancel botton
          return v
        })
        this.sectionListLoading = false
      })
    },
    getVideoList (row) {
      this.videoListLoading = true
      console.log('get video list start...')
      fetchVideoList(row.id).then(response => {
        console.log('get video list', response)
        response = response.data
        var items = []
        items = response.data.list
        this.videoList = items
        this.videoTotal = response.data.totalCount

        this.videoList = items.map(v => {
          this.$set(v, 'edit', false) // https://vuejs.org/v2/guide/reactivity.html
          v.originalName = v.name //  will be used when user click the cancel botton
          return v
        })
        this.videoListLoading = false
      })
    },
    handleFilter () {
      this.listQuery.page = 1
      this.getList()
    },
    handleSizeChange (val) {
      this.listQuery.limit = val
      this.getList()
    },
    handleCurrentChange (val) {
      this.listQuery.page = val
      this.getList()
    },
    handleModifyStatus (row, status) {
      this.$message({
        message: '操作成功',
        type: 'success'
      })
      row.status = status
    },
    resetTemp () {
      this.temp = {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        status: 'published',
        type: ''
      }
    },
    handleCreate () {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.temp.id = parseInt(Math.random() * 100) + 1024 // mock a id
          this.temp.author = 'vue-element-admin'
          createCourse(this.temp).then(() => {
            this.list.unshift(this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleUpdate (row) {
      this.temp = Object.assign({}, row) // copy obj
      this.temp.timestamp = new Date(this.temp.timestamp)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleSection (row) {
      this.temp = Object.assign({}, row) // copy obj
      this.temp.timestamp = new Date(this.temp.timestamp)
      this.dialogSectionVisible = true
      this.getSectionList(row)
    },
    handleSectionEdit (row) {
      this.temp = Object.assign({}, row) // copy obj
      this.temp.timestamp = new Date(this.temp.timestamp)
      this.dialogSectionVisible = true
    },
    handleVideo (row) {
      this.temp = Object.assign({}, row) // copy obj
      this.temp.timestamp = new Date(this.temp.timestamp)
      this.dialogVideoVisible = true
      this.getVideoList(row)
    },
    updateData () {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          tempData.timestamp = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
          tempData.update_status = parseInt(tempData.update_status)
          updateCourse(tempData).then(() => {
            for (const v of this.list) {
              if (v.id === this.temp.id) {
                const index = this.list.indexOf(v)
                this.list.splice(index, 1, this.temp)
                break
              }
            }
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleDelete (row) {
      this.$notify({
        title: '成功',
        message: '删除成功',
        type: 'success',
        duration: 2000
      })
      const index = this.list.indexOf(row)
      this.list.splice(index, 1)
    },
    handleDownload () {
      this.downloadLoading = true
    //   import('@/vendor/Export2Excel').then(excel => {
    //     const tHeader = ['timestamp', 'title', 'type', 'importance', 'status']
    //     const filterVal = ['timestamp', 'title', 'type', 'importance', 'status']
    //     const data = this.formatJson(filterVal, this.list)
    //     excel.export_json_to_excel({
    //       header: tHeader,
    //       data,
    //       filename: 'table-list'
    //     })
    //     this.downloadLoading = false
    //   })
    },
    formatJson (filterVal, jsonData) {
      return jsonData.map(v => filterVal.map(j => {
        if (j === 'timestamp') {
          return parseTime(v[j])
        } else {
          return v[j]
        }
      }))
    },
    cancelEdit (row) {
      row.name = row.originalName
      row.edit = false
      this.$message({
        message: 'The title has been restored to the original value',
        type: 'warning'
      })
    },
    confirmEdit (row) {
      row.edit = false
      updateSection(row).then(() => {
        this.$notify({
          title: '成功',
          message: '更新成功',
          type: 'success',
          duration: 2000
        })
      })
    }
  }
}
</script>

<style scoped>
.edit-inputedit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
