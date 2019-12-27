<template>
  <div id="app">
    <img alt="Wails logo" src="./assets/images/logo.png" style="width: 100px; height: 100px;" class="logo zoomIn">
    <CPUUsage/>
  </div>
</template>

<script>
import CPUUsage from "./components/CPUUsage.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    CPUUsage
  }
};
</script>
