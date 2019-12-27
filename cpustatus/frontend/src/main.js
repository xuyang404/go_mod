import Vue from 'vue';
import App from './App.vue';
import VueApexCharts from 'vue-apexcharts'

Vue.config.productionTip = false;
Vue.config.devtools = true;
Vue.use(VueApexCharts)
Vue.component('apexchart', VueApexCharts)

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
