<template>
  <main>
    <navigation></navigation>
    <div class="container">
      <h1 class="mt-5">{{msg}}</h1>
    </div>
  </main>
</template>

<script>
import axios from 'axios';
import navigation from './navigation';

export default {
  name: 'generator',
  components: { navigation },
  data() {
    return {
      msg: 'App Generator',
    };
  },
  mounted() {
    axios
      .get('/api/v1/packager', {
        method: 'init',
        data: '',
      }).then((response) => {
        console.log(response);
      });
  },
};
</script>

<style scoped>

</style>
