<template>
  <div align="center">
    <strong v-show="!IsLogin" >请先登录</strong>
    <div v-show="IsLogin">
      <p><strong style="font-size:12px" >{{user.UserName}}你好。</strong></p>
      <el-form ref="form" :model="form" label-width="130px" label-position="left">

        <el-form-item label="是否显示指引">
          <div style="text-align: left">
            <el-switch
              v-model="showTips">
            </el-switch>
          </div>
        </el-form-item>

        <el-form-item label="萌点类型" >
          <div style="text-align: left">
            <el-radio-group v-model="form.InfoType" size="mini">
              <el-radio-button label="0" >视频</el-radio-button>
              <el-radio-button label="1" >图片</el-radio-button>
              <el-radio-button label="2" >同人</el-radio-button>
            </el-radio-group>
          </div>
        </el-form-item>

        <p v-show="showTips" style="font-size:10px" > ↓可选择萌点关键词,没有的话双击箭头手动输入</p>
        <el-form-item label="关键词">
          <div style="text-align: left">
            <el-select
              v-model="keyword"
              @focus="getKayWordsToShow()"
              clearable
              filterable
              allow-create
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </div>
        </el-form-item>

        <p v-show="showTips" style="font-size:10px" > ↓描述不可为空</p>
        <el-form-item label="详细描述">
          <el-input
            type="textarea"
            :autosize="true"
            v-model="form.Info">
          </el-input>
        </el-form-item>

        <el-form-item v-show="form.InfoType === '2'" label="同人作品外链">
          <el-input
            v-model="form.FanFictionLink">
          </el-input>
        </el-form-item>
        <p v-show="showTips&&form.InfoType === '0'" style="font-size:10px" > ↓强烈建议尽量带上【在该视频的哪个时段】的描述，会极大方便定位萌点</p>
        <el-form-item v-show="form.InfoType === '0'" label="出现在视频的时段">
          <el-input
            type="textarea"
            v-model="form.TheTime">
          </el-input>
        </el-form-item>
        <p v-show="showTips&&form.InfoType === '0'" style="font-size:10px" > ↓链接必须填至少一个</p>
        <el-form-item v-show="form.InfoType === '0'" label="b站链接">
          <el-input v-model="form.BiliBiliLink"></el-input>
        </el-form-item>
        <p v-show="showTips&&form.InfoType === '0'" style="font-size:10px" > 手机端可以在界面长按B站号，就复制到剪贴板了</p>
        <p v-show="showTips&&form.InfoType === '0'" style="font-size:10px" > 电脑端获取B站号方法</p>
        <p v-show="showTips&&form.InfoType === '0'" style="font-size:10px">
          例如电脑端的B站链接：https://www.bilibili.com/video/<strong>BV1eb411m7GR</strong>/?spm_id_from=333.788.recommend_more_video.0</p>
        <p v-show="showTips&&form.InfoType === '0'" style="font-size:10px" ><strong>BV1eb411m7GR</strong>就是对应的b站号</p>
        <el-form-item v-show="form.InfoType === '0'" label="b站号">
          <el-input v-model="form.BiliBiliID"></el-input>
        </el-form-item>
        <el-form-item v-show="form.InfoType === '0'" label="油管链接">
          <el-input v-model="form.YoutubeLink"></el-input>
        </el-form-item>
        <div v-show="showTips" style="display: inline-block;text-align: center;margin-bottom: 10px">
          <span>可以用这个网站生成本地视频的链接:</span>
          <a href="javascript:;" @click="openUrl('https://streamja.com/')">https://streamja.com/</a>
        </div>
        <el-form-item v-show="form.InfoType === '0'" label="其余渠道的链接">
          <el-input v-model="form.OtherLink"></el-input>
        </el-form-item>

        <p v-show="showTips" style="font-size:10px" > ↓描述需要附图的，请提供外链</p>
        <p v-show="showTips" style="font-size:10px" > ↓外链可在应用市场搜素关键词【图床】，有各种app提供支持</p>

        <div v-show="showTips" style="display: inline-block;text-align: center">
          <span>在线图床sm.ms(推荐):</span>
          <a href="javascript:;" @click="openUrl('https://sm.ms/')">https://sm.ms/</a>
        </div>

        <p v-show="showTips" style="font-size:10px" > sm.ms挑网络，二狗家用流量能上，可是用wifi不行</p>

        <div v-show="showTips" style="display: inline-block;text-align: center">
          <span>在线图床imgchr:</span>
          <a href="javascript:;" @click="openUrl('https://imgchr.com/')">https://imgchr.com/</a>
        </div>

        <p v-show="showTips" style="font-size:10px" > 注意用imgchr拿到图链后往下划拉，提交【<strong>图片URL链接</strong>】而不是【图片链接】,否则不支持在线浏览。 </p>
        <el-button style="margin-top: 20px; margin-bottom: 20px" @click="addImgLink">新增图链</el-button>

        <el-form-item
          v-for="(imglink, index) in form.ImgLinks"
          :label="'图链'"
          :key="imglink.key"
          :prop="'ImgLinks.' + index + '.value'"
        >
          <el-input  v-model="imglink.value">
            <i slot="suffix" class="el-icon-delete" @click="removeImgLink(imglink)"></i>
          </el-input>
        </el-form-item>
      </el-form>
      <strong style="font-size:12px" >请确保提交的都是琴琴相关有效信息</strong>
      <el-button type="primary" :disabled="submitDisable" @click="AddTwoSetVideoInfo">提交</el-button>
    </div>
  </div>
</template>

<script>
import storage from 'good-storage'
export default {
  name: 'addForm',
  data () {
    return {
      IsLogin: false,
      showTips: true,
      submitDisable: false,
      keyword: '',
      passwordType: 'password',
      form: {
        Info: '',
        TheTime: '',
        BiliBiliLink: '',
        BiliBiliID: '',
        YoutubeLink: '',
        OtherLink: '',
        EditPassWord: '',
        InfoType: '0',
        FanFictionLink: '',
        ImgLinks: [{
          value: ''
        }]
      },
      user: {
        UserName: ''
      },
      imgLinksStr: '',
      options: [],
      allOptions: []
    }
  },
  created () {
    this.user.UserName = storage.get('two-set-info-username')
    console.log(this.user.UserName)
    if (this.user.UserName !== undefined && this.user.UserName !== '') {
      this.IsLogin = true
    }
    this.getKeyWords()
  },
  methods: {
    isMobile () {
      let flag = navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)
      return flag
    },
    getKayWordsToShow () {
      console.log('执行')
      this.options = []
      for (let indexOfkeyWord = 0; indexOfkeyWord < this.allOptions.length; indexOfkeyWord++) {
        console.log(this.allOptions[indexOfkeyWord])
        if (this.allOptions[indexOfkeyWord].InfoType === this.form.InfoType) {
          this.options.push(this.allOptions[indexOfkeyWord])
        }
      }
    },
    addTheKeyWord () {
      var addTheKeyWord = true
      for (let indexOfkeyWord = 0; indexOfkeyWord < this.options.length; indexOfkeyWord++) {
        if (this.options[indexOfkeyWord].value === this.keyword) {
          addTheKeyWord = false
          break
        }
      }
      if (this.keyword === '') {
        addTheKeyWord = false
      }
      console.log('addTheKeyWord', addTheKeyWord)
      if (addTheKeyWord) {
        this.$http.post('/api/add_keyword',
          {
            'KeyWord': this.keyword,
            'InfoType': this.form.InfoType,
            'User': this.user.UserName
          },
          {emulateJSON: true})
          .then((response) => {
            let res = response.data
            if (res.error_num === 0) {
              this.$message({type: 'success', message: '新关键词提交成功', duration: 500})
            } else {
              this.$message.error(res['msg'])
            }
          })
      }
    },
    getKeyWords () {
      this.$http.post('/api/get_keywords')
        .then((response) => {
          console.log(response)
          let res = response.data
          this.options = []
          this.allOptions = []
          if (res.error_num === 0) {
            for (let indexOfkeyWord = 0; indexOfkeyWord < res.keyWord_list.length; indexOfkeyWord++) {
              this.allOptions.push({
                value: res.keyWord_list[indexOfkeyWord].KeyWord,
                label: res.keyWord_list[indexOfkeyWord].KeyWord,
                InfoType: res.keyWord_list[indexOfkeyWord].InfoType
              })
            }
            // console.log('this.allOptions', this.allOptions)
            // console.log('res.keyWord_list', res.keyWord_list)
          } else {
            this.$message.error(res['msg'])
          }
        })
    },
    openUrl (val) {
      window.open(val)
    },
    removeImgLink (item) {
      var index = this.form.ImgLinks.indexOf(item)
      if (index !== -1) {
        this.form.ImgLinks.splice(index, 1)
      }
    },
    addImgLink () {
      this.form.ImgLinks.push({
        value: '',
        key: Date.now()
      })
    },
    AddTwoSetVideoInfo () {
      for (let indexOfLink = 0; indexOfLink < this.form.ImgLinks.length; indexOfLink++) {
        this.imgLinksStr = this.imgLinksStr + this.form.ImgLinks[indexOfLink].value
        if (indexOfLink !== (this.form.ImgLinks.length - 1)) {
          this.imgLinksStr = this.imgLinksStr + '\n'
        }
      }
      console.log(this.imgLinksStr)
      let canAdd = true
      if (this.user.UserName === undefined || this.user.UserName === '') {
        canAdd = false
        this.$message.warning('请重新登录')
      }
      if (this.form.Info === '') {
        canAdd = false
        this.$message.warning('描述不可为空')
      }
      if (this.form.BiliBiliLink === '' &&
        this.form.YoutubeLink === '' &&
        this.form.OtherLink === '' &&
        this.imgLinksStr === '' &&
        this.form.FanFictionLink === '' &&
        this.form.BiliBiliID === '') {
        canAdd = false
        this.$message.warning('链接必须填至少一个,或者起码提供个B站号')
      }

      if (this.form.BiliBiliLink !== '' &&
        this.form.BiliBiliID === '') {
        canAdd = false
        this.$message.warning('提供B站链接同时，需要提供B站号用于检索')
      }

      if (canAdd) {
        this.addTheKeyWord()
        this.submitDisable = true
        // this.$http.post('/api/add_two_set_video_info',
        // console.log(this.keyword === '')
        // console.log(this.keyword === '' ? this.form.Info : '【' + this.keyword + '】' + this.form.Info)
        this.$http.post('/api/add_two_set_video_info',
          {
            'Info': this.keyword === '' ? this.form.Info : '【' + this.keyword + '】' + this.form.Info,
            'TheTime': this.form.TheTime,
            'KeyWord': this.keyword,
            'BiliBiliLink': this.form.BiliBiliLink,
            'BiliBiliID': this.form.BiliBiliID,
            'YoutubeLink': this.form.YoutubeLink,
            'OtherLink': this.form.OtherLink,
            'EditPassWord': this.form.EditPassWord,
            'ImgLinks': this.imgLinksStr,
            'InfoType': this.form.InfoType,
            'FanFictionLink': this.form.FanFictionLink,
            'User': this.user.UserName
          },
          {emulateJSON: true})
          .then((response) => {
            console.log(response)
            let res = response.data
            if (res.error_num === 0) {
              this.submitDisable = false
              this.form.Info = ''
              this.form.TheTime = ''
              this.form.BiliBiliLink = ''
              this.form.BiliBiliID = ''
              this.form.YoutubeLink = ''
              this.form.OtherLink = ''
              this.form.EditPassWord = ''
              this.form.FanFictionLink = ''
              this.form.InfoType = '0'
              this.form.ImgLinks = [ {value: ''} ]
              this.form.imgLinksStr = ''
              this.keyword = ''
              this.$message({type: 'success', message: '提交成功', duration: 500})
            } else {
              this.submitDisable = false
              this.$message.error(res['msg'])
            }
          })
      }
    }
  }
}
</script>

<style scoped>
</style>
