<template>
  <div>
    <h1>Band Information</h1>
    <div v-if="bandInfo">
      <pre>{{ bandInfo }}</pre>
    </div>
    <div v-else>
      <p>Loading...</p>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      bandInfo: null,
    };
  },
  mounted() {
    this.fetchBandInfo("melvins");
  },
  methods: {
    async fetchBandInfo(bandName) {
      try {
        const response = await axios.get(
          `http://localhost:8080/api/bands/${bandName}/lastfm`
        );
        this.bandInfo = response.data;
      } catch (error) {
        console.error("Error fetching band info:", error);
      }
    },
  },
};
</script>

<style scoped>
/* Estilos aquí */
</style>
