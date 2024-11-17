<template>
  <div v-if="bandsInfo">
    <!-- Iteramos sobre las bandas divididas en grupos de 3 -->
    <div v-for="(group, groupIndex) in chunkedBandsInfo" :key="groupIndex">
      <BRow class="g-0">
        <!-- Iteramos sobre cada banda dentro de ese grupo -->
        <div
          v-for="(bandInfo, index) in group"
          :key="index"
          class="col-md-4 p-2"
        >
          <BCard no-body class="overflow-hidden band-info-bcard h-100">
            <BRow class="g-0">
              <BCol md="12">
                <BCardBody>
                  <h5 class="card-title">{{ bandInfo.name }}</h5>
                  <BCardText>
                    {{ bandInfo.bio.summary }}
                  </BCardText>
                </BCardBody>
              </BCol>
            </BRow>
          </BCard>
        </div>
      </BRow>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { BCard, BCardBody, BCardText, BCol, BRow } from "bootstrap-vue-next";

export default {
  components: {
    BCard,
    BCardBody,
    BCardText,
    BCol,
    BRow,
  },
  data() {
    return {
      bandsInfo: [],
    };
  },
  mounted() {
    this.fetchBandInfo();
  },
  computed: {
    // Método para dividir el array en grupos de 3
    chunkedBandsInfo() {
      const chunkSize = 3;
      let result = [];
      for (let i = 0; i < this.bandsInfo.length; i += chunkSize) {
        result.push(this.bandsInfo.slice(i, i + chunkSize));
      }
      return result;
    },
  },
  methods: {
    async fetchBandInfo() {
      const bands = [
        "melvins",
        "mastodon",
        "neurosis",
        "crowbar",
        "eyehategod",
        "crowbar",
        "down",
        "kylesa",
        "sleep",
      ];
      try {
        const requests = bands.map((band) =>
          axios.get(`http://localhost:8080/api/bands/${band}/lastfm`)
        );
        const responses = await Promise.all(requests);
        this.bandsInfo = responses.map((response) => response.data.artist);
      } catch (error) {
        console.error("Error fetching band info:", error);
      }
    },
  },
};
</script>
