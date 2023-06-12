<template>
  <div style="display: inline" align="center">
    <div style="z-index:999;">
        <el-collapse style="border-color: white">
          <el-collapse-item style="width: 200px" title="考古频道" name="1">
            <el-row style="margin-top: 10px">
<!--              <el-col :span="12">-->
              <el-col :span="12">
                <el-button
                  circle
                  type="warning"
                  icon="el-icon-collection"
                  @click="handleSelect('油管考古频道')">
                </el-button>
                <p>油管考古频道</p>
              </el-col>
              <el-col :span="12">
                <el-button
                  circle
                  type="bilibili"
                  icon="el-icon-collection"
                  @click="handleSelect('b站考古频道')">
                </el-button>
                <p>b站考古频道</p>
              </el-col>
            </el-row>

          </el-collapse-item>
        </el-collapse>
        <el-menu
          type="success"
          class="el-menu-vertical-demo"
          @select="handleSelect"
        >
          <el-menu-item index="1">
            <i class="el-icon-star-off"></i>
            <span slot="title">查看萌点</span>
          </el-menu-item>
          <el-menu-item index="2">
            <i class="el-icon-plus"></i>
            <span slot="title">新增萌点</span>
          </el-menu-item>
          <el-button v-show="!IsLogin" size="mini" type="primary"  round @click="openLoginDialog = true">登录</el-button>
          <el-button v-show="IsLogin" size="mini"  round @click="openLogoutDialog = true">登出</el-button>
        </el-menu>
      </div>
    <el-dialog title="登出" :visible.sync="openLogoutDialog" width="80%">
      <strong style="font-size:15px" >当前登录：{{user.UserName}}。</strong>
      <el-row>
        <el-col :span="24">
          <el-button  @click="onChangePassWord(user.UserName)" size="mini" plain>修改密码</el-button>
        </el-col>
      </el-row>
      <div slot="footer" class="dialog-footer">
        <el-button @click="openLogoutDialog = false">取 消</el-button>
        <el-button type="primary" @click="logout">登 出</el-button>
      </div>
    </el-dialog>

    <el-dialog title="登录" :visible.sync="openLoginDialog" width="80%">
      <el-form :model="user">
        <el-form-item label="用户名">
          <el-input v-model="user.UserName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input
            :key="passwordType"
            ref="password"
            v-model="user.PassWord"
            :type="passwordType"
            placeholder="Password"
            name="password"
            auto-complete="on"
          >
            <i slot="suffix" @click="showPwd" :class="passwordType === 'password' ? 'el-icon-view' : 'el-icon-edit'" />
          </el-input>
        </el-form-item>
      </el-form>

      <el-row>
        <el-col :span="12">
          <el-button type="success" @click="openRegister" size="mini" plain>注 册</el-button>
        </el-col>
        <el-col :span="12">
          <el-button  @click="onChangePassWord()" size="mini" plain>修改密码</el-button>
        </el-col>
      </el-row>

      <div slot="footer" class="dialog-footer">

          <el-button @click="openLoginDialog = false">取 消</el-button>
          <el-button :disabled="loginDisable" type="primary" @click="login">确 定</el-button>

      </div>
    </el-dialog>

    <el-dialog title="注册" :visible.sync="openRegisterDialog" width="80%">
      <el-form :model="registerUser">
        <el-form-item label="用户名">
          <el-input v-model="registerUser.UserName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="邮箱(非必填)">
          <el-input v-model="registerUser.Email" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input
            :key="passwordType"
            ref="password"
            v-model="registerUser.PassWord"
            :type="passwordType"
            placeholder="Password"
            name="password"
            auto-complete="on"
          >
            <i slot="suffix" @click="showPwd" :class="passwordType === 'password' ? 'el-icon-view' : 'el-icon-edit'" />
          </el-input>
        </el-form-item>
          <el-form-item label="确认密码">
            <el-input
              :key="passwordType"
              ref="password"
              v-model="registerUser.PassWordAgain"
              :type="passwordType"
              placeholder="Password"
              name="password"
              auto-complete="on"
            >
              <i slot="suffix" @click="showPwd" :class="passwordType === 'password' ? 'el-icon-view' : 'el-icon-edit'" />
            </el-input>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">

        <el-button @click="openRegisterDialog = false">取 消</el-button>
        <el-button :disabled="registerDisable" type="primary" @click="register">确 定</el-button>

      </div>
    </el-dialog>

    <el-dialog title="修改密码" :visible.sync="openChangePassWordDialog" width="80%">
      <el-form :model="changePassWordUser">
        <el-form-item label="用户名">
          <el-input v-model="changePassWordUser.UserName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input
            :key="passwordType"
            ref="password"
            v-model="changePassWordUser.PassWord"
            :type="passwordType"
            placeholder="Password"
            name="password"
            auto-complete="on"
          >
            <i slot="suffix" @click="showPwd" :class="passwordType === 'password' ? 'el-icon-view' : 'el-icon-edit'" />
          </el-input>
        </el-form-item>
        <el-form-item label="新密码">
          <el-input
            :key="passwordType"
            ref="password"
            v-model="changePassWordUser.newPassWord"
            :type="passwordType"
            placeholder="Password"
            name="password"
            auto-complete="on"
          >
            <i slot="suffix" @click="showPwd" :class="passwordType === 'password' ? 'el-icon-view' : 'el-icon-edit'" />
          </el-input>
        </el-form-item>
        <el-form-item label="确认新密码">
          <el-input
            :key="passwordType"
            ref="password"
            v-model="changePassWordUser.newPassWordAgain"
            :type="passwordType"
            placeholder="Password"
            name="password"
            auto-complete="on"
          >
            <i slot="suffix" @click="showPwd" :class="passwordType === 'password' ? 'el-icon-view' : 'el-icon-edit'" />
          </el-input>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">

        <el-button @click="openChangePassWordDialog = false">取 消</el-button>
        <el-button :disabled="changePassWordDisable" type="primary" @click="changePassWord">确 定</el-button>

      </div>
    </el-dialog>

    <div style="z-index:-1">
      <QueryFrom v-if="curIndex === '1'" :currentCom="currentCom" ></QueryFrom>
      <addForm v-if="curIndex === '2'" ></addForm>
      <YoutubeList v-if="curIndex === '油管考古频道'"></YoutubeList>
      <BilibiliList v-if="curIndex === 'b站考古频道'"></BilibiliList>
    </div>

    <el-footer style="font-size: x-small; margin-top:20px">备案号：<a href="javascript:;" @click="openUrl('https://beian.miit.gov.cn/')">
      粤ICP备2020125255号</a></el-footer>
  </div>
</template>

<script>
// import AddForm from './AddForm'
// import QueryFrom from './QueryFrom'
// import YoutubeList from './YoutubeList'
// import BilibiliList from './BilibiliList'
import storage from 'good-storage'
const AddForm = () => import('./AddForm')
const QueryFrom = () => import('./QueryFrom')
const YoutubeList = () => import('./YoutubeList')
const BilibiliList = () => import('./BilibiliList')
export default {
  name: 'HelloWorld',
  components: {
    AddForm,
    QueryFrom,
    YoutubeList,
    BilibiliList
  },
  data () {
    return {
      IsLogin: false,
      loginDisable: false,
      registerDisable: false,
      changePassWordDisable: false,
      openRegisterDialog: false,
      openChangePassWordDialog: false,
      openLoginDialog: false,
      openLogoutDialog: false,
      passwordType: 'password',
      msg: '展开',
      isCollapse: true,
      fullHeight: document.documentElement.clientHeight,
      currentCom: '',
      canShowDocument: false,
      curIndex: '1',
      setName: '',
      commitName: '',
      user: {
        UserName: '',
        PassWord: ''
      },
      registerUser: {
        UserName: '',
        Email: '',
        PassWord: '',
        PassWordAgain: ''
      },
      changePassWordUser: {
        UserName: '',
        PassWord: '',
        newPassWord: '',
        newPassWordAgain: ''
      }
    }
  },
  watch: {
    fullHeight (val) {
      // 监控浏览器高度变化
      if (!this.timer) {
        this.fullHeight = val - 20
        this.timer = true
        let that = this
        setTimeout(function () {
          that.timer = false
        }, 400)
      }
    }
  },
  created () {
    this.curIndex = storage.get('two-set-info-curIndex')
    if (this.curIndex === undefined || this.curIndex === '') {
      this.curIndex = '1'
    }
    this.user.UserName = storage.get('two-set-info-username')
    console.log(this.user.UserName)
    if (this.user.UserName !== undefined && this.user.UserName !== '') {
      this.IsLogin = true
    }
  },
  mounted () {
    this.get_bodyHeight()
  },
  methods: {
    showPwd () {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    openUrl (val) {
      window.open(val)
    },
    get_bodyHeight () {
      // 动态获取浏览器高度
      const that = this
      window.onresize = () => {
        return (() => {
          window.fullHeight = document.documentElement.clientHeight
          that.fullHeight = window.fullHeight
          console.log(that.fullHeight)
        })()
      }
    },

    handleSelect (key) {
      this.curIndex = key
      storage.set('two-set-info-curIndex', this.curIndex)
      console.log(key)
    },

    handleClose (key, keyPath) {
      console.log(key, keyPath)
    },
    openRegister () {
      this.openLoginDialog = false
      this.openRegisterDialog = true
    },
    onChangePassWord (val) {
      this.changePassWordUser.UserName = val
      this.openLoginDialog = false
      this.openRegisterDialog = false
      this.openChangePassWordDialog = true
    },
    logout () {
      storage.remove('two-set-info-username')
      storage.remove('two-set-info-password')
      storage.remove('two-set-info-isAdmin')
      this.openLogoutDialog = false
      this.IsLogin = false
      this.$message({type: 'success', message: '登出成功', duration: 500})
      location.reload()
    },
    login () {
      let canLogin = true
      if (this.user.UserName === '') {
        canLogin = false
        this.$message.warning('用户名不可为空')
      }
      if (this.user.PassWord === '') {
        canLogin = false
        this.$message.warning('密码不可为空')
      }

      if (canLogin) {
        this.loginDisable = true
        // this.$http.post('/api/add_two_set_video_info',
        this.$http.post('/api/login',
          {
            'UserName': this.user.UserName,
            'PassWord': this.user.PassWord
          },
          {emulateJSON: true})
          .then((response) => {
            console.log(response)
            let res = response.data
            if (res.error_num === 0) {
              this.loginDisable = false
              this.openLoginDialog = false
              this.IsLogin = true
              storage.set('two-set-info-username', this.user.UserName)
              storage.set('two-set-info-password', this.user.PassWord)
              storage.set('two-set-info-isAdmin', res['IsAdmin'])
              this.$message({type: 'success', message: '登录成功', duration: 500})
              location.reload()
            } else {
              this.IsLogin = false
              this.loginDisable = false
              this.$message.error(res['msg'])
            }
          })
      }
    },

    register () {
      let canRegister = true
      if (this.registerUser.UserName === '') {
        canRegister = false
        this.$message.warning('用户名不可为空')
      }
      if (this.registerUser.PassWord === '') {
        canRegister = false
        this.$message.warning('密码不可为空')
      }
      if (this.registerUser.PassWord !== this.registerUser.PassWordAgain) {
        canRegister = false
        this.$message.warning('密码和确认密码要一致')
      }

      // eslint-disable-next-line no-useless-escape
      const reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/
      if (this.registerUser.Email !== '') {
        if (!reg.test(this.registerUser.Email)) {
          canRegister = false
          this.$message.warning(new Error('请输入正确的邮箱地址'))
        }
      }

      if (canRegister) {
        this.registerDisable = true
        // this.$http.post('/api/add_two_set_video_info',
        this.$http.post('/api/register',
          {
            'UserName': this.registerUser.UserName,
            'PassWord': this.registerUser.PassWord,
            'Email': this.registerUser.Email
          },
          {emulateJSON: true})
          .then((response) => {
            console.log(response)
            let res = response.data
            if (res.error_num === 0) {
              this.registerDisable = false
              this.openLoginDialog = false
              this.openRegisterDialog = false
              this.IsLogin = true
              this.user.UserName = this.registerUser.UserName
              storage.set('two-set-info-username', this.registerUser.UserName)
              storage.set('two-set-info-password', this.registerUser.PassWord)
              this.$message({type: 'success', message: '注册成功', duration: 500})
              location.reload()
            } else {
              this.IsLogin = false
              this.registerDisable = false
              this.$message.error(res['msg'])
            }
          })
      }
    },

    changePassWord () {
      let canChangePassWord = true
      if (this.changePassWordUser.UserName === '') {
        canChangePassWord = false
        this.$message.warning('用户名不可为空')
      }
      if (this.changePassWordUser.PassWord === '') {
        canChangePassWord = false
        this.$message.warning('密码不可为空')
      }
      if (this.changePassWordUser.newPassWord === '') {
        canChangePassWord = false
        this.$message.warning('新密码不可为空')
      }
      if (this.changePassWordUser.newPassWord !== this.changePassWordUser.newPassWordAgain) {
        canChangePassWord = false
        this.$message.warning('新密码和新确认密码要一致')
      }

      if (canChangePassWord) {
        this.changePassWordDisable = true
        this.$http.post('/api/change_password',
          {
            'UserName': this.changePassWordUser.UserName,
            'PassWord': this.changePassWordUser.PassWord,
            'NewPassWord': this.changePassWordUser.newPassWord
          },
          {emulateJSON: true})
          .then((response) => {
            console.log(response)
            let res = response.data
            if (res.error_num === 0) {
              this.changePassWordDisable = false
              this.openLoginDialog = false
              this.openRegisterDialog = false
              this.openChangePassWordDialog = false
              this.IsLogin = false
              storage.remove('two-set-info-username')
              storage.remove('two-set-info-password')
              storage.remove('two-set-info-isAdmin')
              this.$message({type: 'success', message: '密码修改成功，请重新登录', duration: 500})
              location.reload()
            } else {
              this.IsLogin = false
              this.changePassWordDisable = false
              this.$message.error(res['msg'])
            }
          })
      }
    }

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: auto;
    min-width: 100px;
    border: white;
    height: complex;
  }
  .el-menu-item {
    padding-left: 10px;
    padding: 10px 10px 10px 10px;
  }
  .el-menu--collapse {
    border: white;
    margin-left: 10px;
  }

  .iconCollapse {
    width: 20px;
    height: 20px;
    vertical-align: -0.15em;
    fill: currentColor;
  }

  .el-radio-button{
    border: black;
  }

  .el-submenu{
    margin-top: 30px;
  }

  h1, h2 {
    font-weight: normal;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
    margin: 0 10px;
  }
  a {
    color: #42b983;
  }
  .el-scrollbar {
    overflow-y: hidden;
    overflow-x: hidden;
  }

  .el-button--bilibili.is-active,
  .el-button--bilibili:active {
    background: #FB7299;
    border-color: #FB7299;
    color: #fff;
  }

  .el-button--bilibili:focus,
  .el-button--bilibili:hover {
    background: #FB83AD;
    border-color: #FB83AD;
    color: #fff;
  }

  .el-button--bilibili {
    color: #FFF;
    background-color: #FB7299;
    border-color: #FB7299;
  }
</style>
