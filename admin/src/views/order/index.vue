<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="listQuery.idType" clearable style="width: 90px" class="filter-item">
        <el-option v-for="item in idTypeOptions" :key="item.key" :label="item.label" :value="item.key"/>
      </el-select>
      <el-input :placeholder="$t('order.orderIdOrTradeId')" v-model="listQuery.id" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
      <el-select v-model="listQuery.status" class="filter-item" placeholder="Please select">
        <el-option v-for="item in statusOptions" :key="item.key" :label="item.display_name" :value="item.key"/>
      </el-select>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">{{ $t('table.search') }}</el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      fit
      highlight-current-row
      height="600"
      style="width: 100%;">
      <el-table-column :label="$t('order.orderId')" align="center" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.order_id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('order.orderAmount')" width="90px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.order_amount }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('order.payAmount')" width="90px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.pay_amount }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('order.payMethod')" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.pay_method }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('order.paidAt')" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.paid_at }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('order.status')" width="110px" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.userId')" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.user_id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('order.qrcodeId')" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.qrcode_id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('order.tradeId')" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.trade_id }}</span>
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
    </el-table>

    <div class="pagination-container">
      <el-pagination v-show="total>0" :current-page="listQuery.page" :page-sizes="[10,20,30, 50]" :page-size="listQuery.limit" :total="total" background layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange" @current-change="handleCurrentChange"/>
    </div>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item :label="$t('order.orderId')" prop="type">
          <el-select v-model="temp.type" class="filter-item" placeholder="Please select">
            <el-option v-for="item in calendarTypeOptions" :key="item.key" :label="item.display_name" :value="item.key"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('table.date')" prop="timestamp">
          <el-date-picker v-model="temp.timestamp" type="datetime" placeholder="Please pick a date"/>
        </el-form-item>
        <el-form-item :label="$t('table.title')" prop="title">
          <el-input v-model="temp.title"/>
        </el-form-item>
        <el-form-item :label="$t('table.status')">
          <el-select v-model="temp.status" class="filter-item" placeholder="Please select">
            <el-option v-for="item in statusOptions" :key="item" :label="item" :value="item"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('table.importance')">
          <el-rate v-model="temp.importance" :colors="['#99A9BF', '#F7BA2A', '#FF9900']" :max="3" style="margin-top:8px;"/>
        </el-form-item>
        <el-form-item :label="$t('table.remark')">
          <el-input :autosize="{ minRows: 2, maxRows: 4}" v-model="temp.remark" type="textarea" placeholder="Please input"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">{{ $t('table.confirm') }}</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import { fetchList } from '@/api/order'
import waves from '@/directive/waves' // 水波纹指令
import { parseTime } from '@/utils'
const statusOptions = [
  { key: '', display_name: '全部' },
  { key: 'pending', display_name: '待支付' },
  { key: 'paid', display_name: '已支付' },
  { key: 'completed', display_name: '已完成' },
  { key: 'canceled', display_name: '已取消' }
]
// arr to obj ,such as { CN : "China", US : "USA" }
const statusOptionsKeyValue = statusOptions.reduce((acc, cur) => {
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
        'pending': 'info',
        'paid': 'success',
        'compleded': 'success',
        'canceled': 'warning'
      }
      return statusMap[status]
    },
    typeFilter (type) {
      return statusOptionsKeyValue[type]
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
        status: '',
        sort: '+id'
      },
      statusOptions,
      idTypeOptions: [{ label: '订单号', key: 'orderId' }, { label: '交易号', key: 'tradeId' }],
      showReviewer: false,
      temp: {
        id: undefined,
        timestamp: new Date(),
        idType: '',
        status: ''
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      dialogPvVisible: false,
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
    }
  }
}
</script>
