// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import ELEMENTUI from 'element-ui'
import VueResource from 'vue-resource'
import App from './App'
import router from './router'
import 'element-ui/lib/theme-chalk/index.css'
import VueClipboard from 'vue-clipboard2'
import dateTime from 'vue-date-time-m'
Vue.config.productionTip = false
Vue.use(ELEMENTUI)
Vue.use(VueResource)
Vue.use(VueClipboard)
Vue.component('data-time', dateTime)
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})

router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})
