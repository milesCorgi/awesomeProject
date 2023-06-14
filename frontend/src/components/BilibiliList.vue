<template>
  <div>
    <el-collapse>
      <el-collapse-item style="width: 360px" title="搜索" name="1">
        <el-row>
          <el-input size="mini" v-model="filterPara.keyword" placeholder="请输入标题关键字"></el-input>
        </el-row>
        <el-row style="margin-top: 10px">
          <el-switch
            v-model="filterPara.ifDesc"
            inactive-color="#13ce66"
            active-text="新到旧"
            inactive-text="旧到新">
          </el-switch>
          <span style="font-size: small;margin-left: 10px ">从</span>
          <el-button type="text"  @click="dateTimeShow('from')">{{filterPara.from}}</el-button>
          <span style="font-size: small">到</span>
          <el-button type="text"  @click="dateTimeShow('to')">{{filterPara.to}}</el-button>
        </el-row>
        <el-row>
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-refresh"
            circle
            @click="paraReset"></el-button>
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-search"
            circle
            @click="getInfo"></el-button>
        </el-row>
      </el-collapse-item>
    </el-collapse>
    <date-time ref="dateTime"
               type="date"
               format="yyyy-MM-dd"
               @confirm="dateTimeConfirm"></date-time>
    <el-table
      highlight-current-row
      :row-class-name="getTableRowIndex"
      @row-click="handleHighlightCurrentChange"
      :data="showList"
      ref="aTable"
      v-loading="loading"
      header-align = "center"
      style="width: 100%">

      <el-table-column
        v-if = "!MobileIsPc"
        fixed = "left"
        width="30">
        <template slot-scope="scope">
          <i style="margin-right: 10px"
             class = "el-icon-copy-document"
             v-clipboard:copy="scope.row.VideoUrl"
             v-clipboard:success="copy"></i>
        </template>
      </el-table-column>
      <el-table-column
        v-if = "!MobileIsPc"
        fixed = "left"
        label="标题"
        width="200">
        <template slot-scope="scope">
          <div>
            <span style="white-space: normal;word-break: keep-all;"
                  href="javascript:;" @click="openUrl(scope.row.VideoUrl)">
              {{ scope.row.Title }}
            </span>
          </div>
        </template>
      </el-table-column>

      <el-table-column
        v-if = "MobileIsPc"
        prop="Title"
        label="标题"
        width="250">
        <template slot-scope="scope">
          <div>
            <a style="white-space: normal;word-break: keep-all;"
               href="javascript:;" @click="openUrl(scope.row.VideoUrl)">
              {{ scope.row.Title }}
            </a>
          </div>
        </template>
      </el-table-column>
      <el-table-column
        label="发布时间"
        width="160">
        <template slot-scope="scope">
          <span>
            {{ scope.row.Created | dateFormat}}
          </span>
        </template>
      </el-table-column>

      <el-table-column
        v-if = "MobileIsPc"
        label="点击保存"
        width="120">
        <template slot-scope="scope">
          <i style="margin-right: 10px"
             class = "el-icon-copy-document"
             v-clipboard:copy="scope.row.VideoUrl"
             v-clipboard:success="copy"></i>
        </template>
      </el-table-column>

    </el-table>

    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentPageChange"
      :small="PagerIsSmall"
      :current-page="page.currentPage"
      :pager-count="PagerCount"
      :page-sizes="[1,10, 20, 50, 100]"
      :page-size="page.pageSize"
      layout="total, sizes, prev, pager, next"
      :total="page.totalNum">
    </el-pagination>
  </div>
</template>

<script>
import Clipboard from 'clipboard'
import moment from 'moment'
import storage from 'good-storage'
import DateTime from 'vue-date-time-m'
export default {
  name: 'BilibiliList',
  components: {
    DateTime
  },
  data () {
    return {
      setData: 'from',
      loading: true,
      PagerIsSmall: false,
      PagerCount: 7,
      currentRow: '',
      page: {
        totalNum: 0,
        currentPage: 1,
        pageSize: 10
      },
      totalList: [],
      showList: [],
      MobileIsPc: true,
      filterPara: {
        from: '2014-01-01',
        to: '2050-01-01',
        keyword: '',
        ifDesc: false
      }
    }
  },
  filters: {
    dateFormat (dateStr) {
      const pattern = 'YYYY-MM-DD HH:mm:ss'
      return moment(dateStr).add(8, 'hours').format(pattern)
    }
  },
  created () {
    this.MobileIsPc = !this.isMobile()
    if (!this.MobileIsPc) {
      this.PagerCount = 5
      this.PagerIsSmall = true
    }
    let filterPara = storage.get('two-set-info-youtube-list-filterPara')
    if (filterPara !== undefined) {
      this.filterPara = filterPara
    }
    this.page.pageSize = storage.get('two-set-info-bilibili-list-pageSize')
    if (this.page.pageSize === undefined || this.page.pageSize === '') {
      this.page.pageSize = 10
    }

    this.page.currentPage = storage.get('two-set-info-bilibili-list-currentPage')
    if (this.page.currentPage === undefined || this.page.currentPage === '') {
      this.page.currentPage = 1
    }

    this.getInfo()
    this.currentRow = storage.get('two-set-info-bilibili-list-currentRow')
    // console.log('this.currentRow' + this.currentRow)
  },
  methods: {
    dateTimeShow (val) {
      this.setData = val
      this.$refs.dateTime.show()
    },
    dateTimeConfirm (val) {
      if (this.setData === 'to') {
        this.filterPara.to = val
        if (this.filterPara.from > this.filterPara.to) {
          this.$message({type: 'warn', message: '输入日期【从】要早于【到】'})
          this.filterPara.to = '2050-01-01'
        }
      } else {
        this.filterPara.from = val
        if (this.filterPara.from > this.filterPara.to) {
          this.$message({type: 'warn', message: '输入日期【从】要早于【到】'})
          this.filterPara.from = '2014-01-01'
        }
      }
      console.log('get date' + val)
    },
    getTableRowIndex ({row, rowIndex}) {
      // 把每一行的索引放进row
      row.index = rowIndex
    },
    openUrl (val) {
      let reg = /^(https?|ftp|file):\/\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]/
      if (!reg.test(val)) {
        this.$message({message: 'url格式不合法', duration: 500})
      } else {
        window.open(val)
      }
      console.log(val)
    },
    handleHighlightCurrentChange (row) {
      this.currentRow = row.index
      storage.set('two-set-info-bilibili-list-currentRow', this.currentRow)
      console.log(`当前行:` + this.currentRow)
    },
    handleSizeChange (val) {
      this.currentRow = ''
      storage.set('two-set-info-bilibili-list-currentRow', this.currentRow)
      console.log(`每页 ${val} 条`)
      this.page.pageSize = val
      storage.set('two-set-info-bilibili-list-pageSize', val)
      this.updateCurrentPageInfo()
    },
    handleCurrentPageChange (val) {
      this.currentRow = ''
      storage.set('two-set-info-bilibili-list-currentRow', this.currentRow)
      this.page.currentPage = val
      storage.set('two-set-info-bilibili-list-currentPage', val)
      console.log(`当前页: ${val}`)
      this.updateCurrentPageInfo()
    },
    isMobile () {
      return navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)
    },
    copy () {
      var clipboard = new Clipboard('.tag-read')
      clipboard.on('success', e => {
        // 释放内存
        clipboard.destroy()
      })
      clipboard.on('error', e => {
        // 释放内存
        clipboard.destroy()
      })
      this.$message({type: 'success', message: '已复制', duration: 500})
    },
    paraReset () {
      this.filterPara.from = '2014-01-01'
      this.filterPara.to = '2050-01-01'
      this.filterPara.keyword = ''
      this.filterPara.ifDesc = false
    },
    getInfo () {
      this.loading = true
      this.page.totalNum = 0
      this.$http.post('/api/query_bilibili_video_list',
        {
          'from': this.filterPara.from,
          'to': this.filterPara.to,
          'keyword': this.filterPara.keyword,
          'ifDesc': this.filterPara.ifDesc ? '1' : '0'
        })
        // {emulateJSON: true})
        .then((response) => {
          // console.log(response)
          let res = response.data
          if (res.error_num === 0) {
            storage.set('two-set-info-youtube-list-filterPara', this.filterPara)
            this.totalList = res.info_list
            this.showList = res.info_list
            this.page.totalNum = res.total_num
            this.updateCurrentPageInfo()
            this.loading = false
            // 不主动paraReset了
            // this.paraReset()
            console.log('this.page.totalNum:', this.page.totalNum)
          } else {
            this.$message.error(res['msg'])
          }
        })
    },
    updateCurrentPageInfo () {
      const startPosition = (this.page.currentPage - 1) * this.page.pageSize
      let endPosition = this.page.currentPage * this.page.pageSize
      endPosition = endPosition > this.page.totalNum ? this.page.totalNum : endPosition
      this.showList = this.totalList.slice(startPosition, endPosition)
      // console.log(this.showList)
      for (let indexOfShowList = 0; indexOfShowList < this.showList.length; indexOfShowList++) {
        this.showList[indexOfShowList].Title = this.showList[indexOfShowList].Title.replace(/&lt;/g, '<')
          .replace(/&gt;/g, '>')
          .replace(/&amp;/g, '&')
          .replace(/&quot;/g, '"')
          .replace(/&#39;/g, '\'')
          .replace(/&nbsp;/g, '')
          .replace(/\n/g, '<br/>')
          .replace(/\r/g, '<br/>')
      }
      if (this.currentRow !== undefined && this.currentRow !== '') {
        console.log('高亮指定行')
        this.$nextTick(() => {
          this.$refs.aTable.setCurrentRow(this.showList[this.currentRow])
        })
      }
    }
  }
}

</script>

<style scoped>
  a:link{color: #080504;}

  a:visited{color:black;}

  a:hover{color:#3a8ee6;}

  a:active{color:#ccc;}
</style>
