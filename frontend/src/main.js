import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from './service/api'
import ElementUI from 'element-ui'

Vue.use(ElementUI)
Vue.prototype.Free = window.Free
Vue.config.productionTip = false
Vue.prototype.$axios = axios;

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')


