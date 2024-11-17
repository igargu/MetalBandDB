<template>
  <div class="p-2">
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
