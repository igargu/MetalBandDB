<template>
  <div class="p-3">
    <h1>MetalBandDB</h1>
    <BandInfo />
  </div>
</template>

<script>
import BandInfo from "./components/BandInfo.vue";

export default {
  name: "App",
  components: {
    BandInfo,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
