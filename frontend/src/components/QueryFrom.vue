<template>
  <div>
    <p v-show="isAdmin">你好,管理员{{user.UserName}}</p>

    <el-collapse style="border-color: white">
      <el-collapse-item style="width: 200px;margin-left: 30px;margin-right: 30px;margin-top: 10px" title="高级搜索">
        <el-row>
            <el-row type="flex" style="margin-top:10px" justify="center">
              <el-col>
                <el-input  clearable v-model="filterPara.Info" placeholder="搜索描述">
                </el-input>
              </el-col>
          </el-row>
          <el-row>
              <el-col >
                <el-input  style="margin-top:10px" clearable v-model="filterPara.BiliBiliID" placeholder="输入b站号">
                </el-input>
              </el-col>
            </el-row>
        </el-row>
        <el-row>
          <el-row type="flex" style="margin-top:10px" justify="center">
            <el-col >
              <el-input clearable  v-model="filterPara.YoutubeLink" placeholder="输入油管链接">
              </el-input>
            </el-col>
          </el-row>
          <el-row>
            <el-col>
              <el-input  style="margin-top:10px" clearable  v-model="filterPara.OtherLink" placeholder="输入其他链接">
              </el-input>
            </el-col>
          </el-row>
        </el-row>
      </el-collapse-item>
    </el-collapse>
    <el-row>
      <el-row type="flex" style="margin-top:10px" justify="center">
        <span style="margin-top:10px;font-size: small">类型：</span>
        <el-radio-group v-model="filterPara.InfoType" size="mini">
          <el-radio-button label="" >全部</el-radio-button>
          <el-radio-button label="0" >视频</el-radio-button>
          <el-radio-button label="1" >图片</el-radio-button>
          <el-radio-button label="2" >同人</el-radio-button>
        </el-radio-group>
      </el-row>
    </el-row>
    <el-row>
      <el-row  type="flex" style="margin-top:10px" justify="center">
        <span style="margin-top:10px;font-size: small">所属ip：</span>
        <el-select
          v-model="filterPara.IntellectualPropertyName"
          @change="getIntellectualPropertyName()"
          clearable
          filterable
          placeholder="所属ip">
          <el-option
            v-for="item in intellectualPropertyNames"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-row>
    </el-row>
    <el-row>
      <el-row v-show="notShowIfNull(filterPara.IntellectualPropertyName)" type="flex" style="margin-top:10px" justify="center">
        <span style="margin-top:10px;font-size: small">关键词：</span>
        <el-select
          v-model="filterPara.keyword"
          @focus="getKeyWordsToShow()"
          clearable
          filterable
          placeholder="关键词">
          <el-option
            v-for="item in keywords"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-row>
    </el-row>

<!--    <el-row>-->
<!--      <div style="text-align: left">-->
<!--        只看我的-->
<!--        <el-switch-->
<!--          v-model="onlyMe">-->
<!--        </el-switch>-->
<!--      </div>-->
<!--    </el-row>-->
    <el-row type="flex" justify="center">
     <el-button style="margin-top:10px" type="primary" @click="getInfo">查询</el-button>
      <!--        <el-button @click="onCancle">清空</el-button>-->
      <el-button style="margin-top:10px;margin-right: 7px;" @click="searchReset">重置</el-button>
    </el-row>
    <el-table
      :data="showList"
      v-loading="loading"
      header-align = "center"
      style="width: 100%">
      <el-table-column type="expand" fixed = "left" label="详情" width="80">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item v-show="notShowIfNull(props.row.User)" label="提交人">
              <span>{{ props.row.User }}</span>
            </el-form-item>
            <el-form-item v-show="notShowIfNull(props.row.TheTime)" label="在该视频的哪个时段">
              <span>{{ props.row.TheTime }}</span>
            </el-form-item>
            <el-form-item v-show="notShowIfNull(props.row.BiliBiliLink)" label="b站链接">
              <div slot="label" style="display: inline-block;text-align: center">
                <i style="margin-right: 10px"
                   class = "el-icon-copy-document"
                   v-clipboard:copy="props.row.BiliBiliLink"
                   v-clipboard:success="copy"></i>
                <span>b站链接</span>
              </div>
              <a href="javascript:;" @click="openUrl(props.row.BiliBiliLink)">
                {{ props.row.BiliBiliLink | showFullLink }}
              </a>
            </el-form-item>

            <el-form-item v-show="notShowIfNull(props.row.BiliBiliID)" label="">
              <div slot="label" style="display: inline-block;text-align: center">
                <i style="margin-right: 10px"
                   class = "el-icon-copy-document"
                   v-clipboard:copy="props.row.BiliBiliID"
                   v-clipboard:success="copy"></i>
                <span>b站号</span>
              </div>
              <a href="javascript:;" @click="openUrl(props.row.BiliBiliLink)">
                {{ props.row.BiliBiliID }}
              </a>
            </el-form-item>

            <el-form-item v-show="notShowIfNull(props.row.YoutubeLink)"  label="油管链接">
              <div slot="label" style="display: inline-block;text-align: center">
                <i style="margin-right: 10px"
                   class = "el-icon-copy-document"
                   v-clipboard:copy="props.row.YoutubeLink"
                   v-clipboard:success="copy"></i>
                <span>油管链接</span>
              </div>
              <a href="javascript:;" @click="openUrl(props.row.YoutubeLink)">
                {{ props.row.YoutubeLink | showFullLink }}
              </a>
            </el-form-item>
            <el-form-item v-show="notShowIfNull(props.row.OtherLink)" label="其余渠道的链接">
              <div slot="label" style="display: inline-block;text-align: center">
                <i style="margin-right: 10px"
                   class = "el-icon-copy-document"
                   v-clipboard:copy="props.row.OtherLink"
                   v-clipboard:success="copy"></i>
                <span>其余渠道的链接</span>
              </div>
              <a href="javascript:;" @click="openUrl(props.row.OtherLink)">
                {{ props.row.OtherLink | showFullLink }}
              </a>
            </el-form-item>

            <el-form-item v-show="notShowIfNull(props.row.FanFictionLink)"  label="同人链接">
              <div slot="label" style="display: inline-block;text-align: center">
                <i style="margin-right: 10px"
                   class = "el-icon-copy-document"
                   v-clipboard:copy="props.row.FanFictionLink"
                   v-clipboard:success="copy"></i>
                <span>同人链接</span>
              </div>
              <a href="javascript:;" @click="openUrl(props.row.FanFictionLink)">
                {{ props.row.FanFictionLink | showFullLink }}
              </a>
            </el-form-item>

            <el-form-item
              label="图链"
              v-for="(item,index) in props.row.ShowImgLinks"
              v-if="settingShowImg(item.value,'img')"
              :key="index"
              :prop="item.value">
              <div slot="label" style="display: inline-block;text-align: center">
                <i style="margin-right: 10px"
                   class = "el-icon-copy-document"
                   v-clipboard:copy="item.value"
                   v-clipboard:success="copy"></i>
                <span>←复制图链</span>
              </div>
              <div style="display: inline-block;text-align: center">
                <el-image
                  style="width: 100px; height: 100px"
                  :src="item.value"
                  fit = "scale-down"
                  lazy
                  :preview-src-list="item.srcList">
                  <div slot="placeholder" class="image-slot">
                    <img style="width: 100px; height: 100px" class="el-icon-picture-outline"/>
                  </div>
                </el-image>
              </div>
            </el-form-item>

            <el-form-item
                v-for="(item,index) in props.row.ShowImgLinks"
                v-if="settingShowImg(item.value,'info')"
                :key="index"
                :prop="item.value">
                <div slot="label" style="display: inline-block;text-align: center">
                  <i style="margin-right: 10px"
                     class = "el-icon-copy-document"
                     v-clipboard:copy="item.value"
                     v-clipboard:success="copy"></i>
                  <span>图链</span>
                </div>
                <div style="display: inline-block;text-align: center">
                  <a href="javascript:;" @click="openUrl(item.value)">
                    {{ item.value | showFullLink }}
                  </a>
                </div>

            </el-form-item>

            <el-form-item label="添加时间">
              <span>
                {{ props.row.AddTime | dateFormat}}
              </span>
            </el-form-item>
            <el-form-item label="更新时间">
              <span>
                {{ props.row.UpdateTime | dateFormat}}
              </span>
            </el-form-item>
            <el-form-item v-show="isAdmin" label="管理员操作">
              <el-button size="mini" @click="disableThisInfo(props.row.Id)">屏蔽此条</el-button>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
        fixed = "left"
        label="类型"
        width="60">
        <template slot-scope="scope">
          <el-tag size="mini" :type="infoTypeStyle(scope.row.InfoType)">{{ scope.row.InfoType | infoTypeName }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if = "MobileIsPc"
        prop="IntellectualPropertyName"
        label="所属IP"
        width="150">
      </el-table-column>
      <el-table-column
        fixed = "left"
        label="详细描述"
        width="230">
        <template slot-scope="scope">
          <div style="white-space: pre-line">
            <p>{{ scope.row.Info }}</p>
          </div>
        </template>
      </el-table-column>
      <el-table-column
        v-if = "MobileIsPc"
        prop="TheTime"
        label="在该视频的哪个时段"
        width="150">
      </el-table-column>
      <el-table-column
        v-if = "MobileIsPc"
        prop="BiliBiliLink"
        label="b站链接"
        width="120">
      </el-table-column>
      <el-table-column
        v-if = "MobileIsPc"
        prop="YoutubeLink"
        label="油管链接"
        width="120">
      </el-table-column>
      <el-table-column
        v-if = "MobileIsPc"
        prop="OtherLink"
        label="其余渠道的链接"
        width="120">
      </el-table-column>
      <el-table-column
        v-if = "MobileIsPc"
        prop="FanFictionLink"
        label="同人链接"
        width="120">
      </el-table-column>
      <el-table-column
        v-if = "MobileIsPc"
        label="创建时间"
        width="120">
        <template slot-scope="scope">
          <span>
            {{ scope.row.AddTime | dateFormat}}
          </span>
        </template>
      </el-table-column>
      <el-table-column
        v-if = "MobileIsPc"
        prop="UpdateTime"
        label="更新时间"
        width="120">
        <template slot-scope="scope">
          <span>
            {{ scope.row.UpdateTime | dateFormat}}
          </span>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
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
import moment from 'moment'
import Clipboard from 'clipboard'
import storage from 'good-storage'

export default {
  name: 'queryFrom',
  filters: {
    infoTypeName (val) {
      const dist = {
        '': '全部',
        '0': '视频',
        '1': '图片',
        '2': '同人'
      }
      // console.log(val)
      if (dist.hasOwnProperty(val)) {
        return dist[val]
      } else {
        return '其他'
      }
    },
    dateFormat (dateStr) {
      const pattern = 'YYYY-MM-DD HH:mm:ss'
      return moment(dateStr).add(8, 'hours').format(pattern)
    },
    showFullLink (val) {
      if (val.length <= 27) {
        return val
      } else {
        return '点击跳转或复制获取链接'
      }
    }
  },
  data () {
    return {
      intellectualPropertyNames: [],
      keywords: [],
      allkeywords: [],
      isAdmin: false,
      lenternUrl: 'https://github.com/getlantern/lantern',
      loading: true,
      filterPara: {
        IntellectualPropertyName: '',
        keyword: '',
        Info: '',
        BiliBiliID: '',
        YoutubeLink: '',
        OtherLink: '',
        InfoType: ''
      },
      totalList: [],
      showList: [],
      page: {
        totalNum: 0,
        currentPage: 1,
        pageSize: 10
      },
      PagerCount: 7,
      PagerIsSmall: false,
      MobileIsPc: true,
      user: {
        UserName: '',
        PassWord: ''
      }
    }
  },
  created () {
    this.MobileIsPc = !this.isMobile()
    if (!this.MobileIsPc) {
      this.PagerCount = 5
      this.PagerIsSmall = true
    }
    this.getIntellectualPropertyNames()
    this.user.UserName = storage.get('two-set-info-username')
    this.filterPara.IntellectualPropertyName = storage.get('awesomeProject-intellectualPropertyNames')
    console.log('jfiudfhgi' + this.filterPara.IntellectualPropertyName)
    if (this.filterPara.IntellectualPropertyName === undefined) {
      this.filterPara.IntellectualPropertyName = ''
    }
    this.getKeyWords()
    console.log(this.user.UserName)
    if (this.user.UserName !== undefined || this.user.UserName !== '') {
      this.IsLogin = true
    }
    this.getInfo()
    // this.loading = false
  },
  mounted () {
    if (storage.get('two-set-info-isAdmin') === '1') {
      this.isAdmin = true
    } else {
      this.isAdmin = false
    }
  },
  methods: {
    getIntellectualPropertyName () {
      this.allkeywords = []
      storage.set('awesomeProject-intellectualPropertyNames', this.filterPara.IntellectualPropertyName)
      this.getKeyWords()
    },
    getIntellectualPropertyNames () {
      this.$http.post('/api/get_intellectual_property_names')
        .then((response) => {
          // console.log(response)
          let res = response.data
          this.intellectualPropertyNames = []
          if (res.error_num === 0) {
            for (let indexOfIntellectualPropertyNames = 0; indexOfIntellectualPropertyNames < res.intellectual_property_name_list.length; indexOfIntellectualPropertyNames++) {
              this.intellectualPropertyNames.push({
                value: res.intellectual_property_name_list[indexOfIntellectualPropertyNames].IntellectualPropertyNames,
                label: res.intellectual_property_name_list[indexOfIntellectualPropertyNames].IntellectualPropertyNames
              })
            }
            console.log('this.intellectualPropertyNames', this.intellectualPropertyNames)
            console.log('res.intellectual_property_name_list', res.intellectual_property_name_list)
          } else {
            this.$message.error(res['msg'])
          }
        })
    },
    getKeyWordsToShow () {
      this.keywords = []
      this.getKeyWords()
      console.log('执行')
      this.keywords = this.allkeywords
    },
    getKeyWords () {
      if (this.filterPara.IntellectualPropertyName !== undefined && this.filterPara.IntellectualPropertyName !== '') {
        this.$http.post('/api/get_keywords',
          {
            'IntellectualPropertyName': this.filterPara.IntellectualPropertyName
          })
          .then((response) => {
            // console.log(response)
            let res = response.data
            // this.keywords = []
            this.allkeywords = []
            if (res.error_num === 0) {
              for (let indexOfkeyWord = 0; indexOfkeyWord < res.keyWord_list.length; indexOfkeyWord++) {
                this.allkeywords.push({
                  value: res.keyWord_list[indexOfkeyWord].KeyWord,
                  label: res.keyWord_list[indexOfkeyWord].KeyWord,
                  InfoType: res.keyWord_list[indexOfkeyWord].InfoType
                })
              }
              // console.log('this.allkeywords', this.allkeywords)
              // console.log('res.keyWord_list', res.keyWord_list)
            } else {
              this.$message.error(res['msg'])
            }
          })
      }
    },
    infoTypeStyle (val) {
      const dist = {
        '0': '',
        '1': 'success',
        '2': 'warning'
      }
      // console.log(val)
      if (dist.hasOwnProperty(val)) {
        return dist[val]
      } else {
        return 'info'
      }
    },
    settingShowImg (val, setting) {
      if (setting === 'img') {
        return (val.length !== 0 && this.verifyImg(val))
      } else if (setting === 'info') {
        return (val.length !== 0 && !this.verifyImg(val))
      } else {
        // 默认img模式
        return (val.length !== 0 && this.verifyImg(val))
      }
    },
    notShowIfNull (val) {
      return val.length !== 0
    },
    isMobile () {
      return navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)
    },
    searchReset () {
      this.paraReset()
      this.getInfo()
    },
    paraReset () {
      this.filterPara.Info = ''
      this.filterPara.BiliBiliID = ''
      this.filterPara.YoutubeLink = ''
      this.filterPara.OtherLink = ''
      this.filterPara.InfoType = ''
      this.filterPara.keyword = ''
      this.filterPara.IntellectualPropertyName = ''
      storage.set('awesomeProject-intellectualPropertyNames', this.filterPara.IntellectualPropertyName)
    },
    disableThisInfo (id) {
      let canDisable = true
      this.user.UserName = storage.get('two-set-info-username')
      if (this.user.UserName === undefined || this.user.UserName === '') {
        canDisable = false
        this.$message({type: 'waring', message: '没有获取到用户名，请找二狗'})
      }
      this.user.PassWord = storage.get('two-set-info-password')
      if (this.user.PassWord === undefined || this.user.PassWord === '') {
        canDisable = false
        this.$message({type: 'waring', message: '用户验证没通过，请找二狗'})
      }
      if (storage.get('two-set-info-isAdmin') !== '1') {
        this.$message({type: 'waring', message: '管理员检查没通过，请找二狗', duration: 500})
        canDisable = false
        this.isAdmin = false
      }
      if (canDisable) {
        this.$http.post('/api/set_disable_info',
          {
            'UserName': this.user.UserName,
            'PassWord': this.user.PassWord,
            'InfoId': id,
            'Enable': '0'
          },
          {emulateJSON: true})
          .then((response) => {
            console.log(response)
            let res = response.data
            if (res.error_num === 0) {
              this.$message({type: 'success', message: '屏蔽成功，刷新后生效', duration: 500})
            } else {
              this.$message.error(res['msg'])
            }
          })
      }
    },
    getInfo () {
      this.loading = true
      this.page.totalNum = 0
      this.page.currentPage = 1
      this.$http.post('/api/query_video_info',
        {
          'KeyWord': this.filterPara.keyword,
          'Info': this.filterPara.Info,
          'BiliBiliID': this.filterPara.BiliBiliID,
          'YoutubeLink': this.filterPara.YoutubeLink,
          'OtherLink': this.filterPara.OtherLink,
          'InfoType': this.filterPara.InfoType,
          'IntellectualPropertyName': this.filterPara.IntellectualPropertyName
        })
        .then((response) => {
          console.log(response)
          let res = response.data
          if (res.error_num === 0) {
            this.totalList = res.info_list
            this.showList = res.info_list
            this.page.totalNum = res.total_num
            this.updateCurrentPageInfo()
            this.loading = false
            // this.$message({type: 'success', message: '过滤器已经重置', duration: 100})
            // this.paraReset()
            console.log('this.page.totalNum:', this.page.totalNum)
          } else {
            this.$message.error(res['msg'])
          }
        })
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
    handleSizeChange (val) {
      console.log(`每页 ${val} 条`)
      this.page.pageSize = val
      this.updateCurrentPageInfo()
    },
    handleCurrentChange (val) {
      this.page.currentPage = val
      console.log(`当前页: ${val}`)
      this.updateCurrentPageInfo()
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
    // 校验图片
    verifyImg (val) {
      if (val !== undefined && val !== '') {
        let valArray = val.split('.')
        let endWith = valArray[valArray.length - 1].toLowerCase()
        // console.log(endWith)
        return endWith === 'png' || endWith === 'jpg' || endWith === 'jpeg' || endWith === 'gif'
      } else {
        return false
      }
    },

    updateCurrentPageInfo () {
      const startPosition = (this.page.currentPage - 1) * this.page.pageSize
      let endPosition = this.page.currentPage * this.page.pageSize
      endPosition = endPosition > this.page.totalNum ? this.page.totalNum : endPosition
      this.showList = this.totalList.slice(startPosition, endPosition)
      for (let indexOfShowList = 0; indexOfShowList < this.showList.length; indexOfShowList++) {
        let ImgLinksArray = this.showList[indexOfShowList].ImgLinks.split('\n')
        this.showList[indexOfShowList].ShowImgLinks = []
        // 可以显示的图片组
        const ShowImgLinks = []
        for (let indexOfImgLinksArray = 0; indexOfImgLinksArray < ImgLinksArray.length; indexOfImgLinksArray++) {
          if (this.verifyImg(ImgLinksArray[indexOfImgLinksArray])) {
            ShowImgLinks.push(ImgLinksArray[indexOfImgLinksArray])
          }
        }
        for (let indexOfImgLinksArray = 0; indexOfImgLinksArray < ImgLinksArray.length; indexOfImgLinksArray++) {
          const enableShow = this.verifyImg(ImgLinksArray[indexOfImgLinksArray])
          this.showList[indexOfShowList].ShowImgLinks.push(
            {
              enableShow: enableShow,
              value: ImgLinksArray[indexOfImgLinksArray],
              srcList: ShowImgLinks
            })
        }
      }
    }
  }
}
</script>

<style scoped>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 100%;
  }
</style>
