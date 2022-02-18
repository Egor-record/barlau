<template>
  <main>
    <navigation></navigation>
    <div class="container">
        <div class="mt-5">
            <div class="d-flex">
              <div class="col-6">
                <h1>
                  Manifest Player
                </h1>
                <p>
                  Run <strong>Sample Player</strong> to play your stream on TV
                </p>
                <ul>
                  <li>
                    Select your Streaming Protocol
                  </li>
                  <li>
                    Paste link to the manifest
                  </li>
                  <li>
                    Download generated .ipk file with sample player in it
                  </li>
                  <li>
                    Run on your TV
                  </li>
                </ul>
              </div>
              <div class="mt-5 mt-md-0 col-6">
                <div v-if="steps.choose">
                  <h3 class="mb-3">Select player you want to use</h3>
                  <button class="btn btn-outline-secondary" v-on:click="changeStatus('native')">
                    Native Player
                  </button>
                  <button class="btn btn-outline-secondary" v-on:click="changeStatus('shaka')">
                    Shaka Player
                  </button>
                </div>
                <div v-if="steps.native">
                  <p><button class="btn-sm btn"  v-on:click="changeStatus('choose')">< Back</button></p>
                  <h4>Native Player</h4>
                  <p>Paste link to manifest. It would be inserted into video tag</p>
                  <form @submit="generateApp" method="post">
                    <input type="text" name="stream" class="w-100" v-model="manifestLink" required="required">
                    <input type="submit" class="btn btn-success mt-3" value="Generate App">
                    <a v-if="downloadUrl"  :href="'./static/ipk' + this.downloadUrl + '/com.sample.player_1.0.0_all.ipk'" download="com.sample.player_1.0.0_all.ipk" class="btn btn-outline-success mt-3">Download ipk</a>
                  </form>
                </div>
                <div v-if="steps.shaka">
                  <p><button class="btn-sm btn"  v-on:click="changeStatus('choose')">< Back</button></p>
                  <h4>Shaka Player</h4>
                  <p>Paste link to manifest. It would be inserted into Shaka Player</p>
                  <form @submit="generateApp" method="post">
                    <input type="text" name="stream" class="w-100" v-model="manifestLink" required="required">
                    <input type="submit" class="btn btn-success mt-3" value="Generate App">
                    <a v-if="downloadUrl" :href="'./static/ipk' + this.downloadUrl + '/com.sample.player_1.0.0_all.ipk'" download="com.sample.player_1.0.0_all.ipk" class="btn btn-outline-success mt-3">Download ipk</a>
                  </form>
                </div>
              </div>
            </div>
        </div>
    </div>
  </main>
</template>

<script>
import navigation from './navigation';

export default {
  name: 'checker',
  components: { navigation },
  data() {
    return {
      steps: {
        choose: true,
        native: false,
        shaka: false,
      },
      manifestLink : "",
      downloadUrl: 0
    };
  },
  mounted() {

  },
  methods: {
    changeStatus(item) {
      for (let key in this.steps) {
        if (this.steps.hasOwnProperty(key)) {
          this.steps[key] = false
        }
      }
      this.steps[item] = true;
    },
    generateApp(e){
      e.preventDefault();
      fetch("api/v1/createPlayer", {
        method: "post",
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          url: this.manifestLink,
          native: this.steps.native
        })
      }).then(response=>response.json())
      .then(data=>{
        if (data.path) {
          this.downloadUrl = data.path
        }
      });
    }
  },
};
</script>

<style scoped>

</style>
