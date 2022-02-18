<template>
  <main>
    <navigation></navigation>
    <div class="container">
      <h1 class="mt-5">{{msg}}</h1>
      <codemirror v-model="this.code" class="mt-5" :options="cmOption" ></codemirror>
    </div>
  </main>
</template>

<script>

import navigation from './navigation';
import codemirror from 'vue-codemirror'
import 'codemirror/mode/javascript/javascript.js'
import 'codemirror/theme/base16-dark.css'

export default {
  name: 'generator',
  components: { navigation, codemirror },
  data() {
    return {
      msg: 'Apps Generator',
      code: `<!DOCTYPE html>
<html>
<head>
<title>Example Web App</title>
<style>
    body {
        width: 100%;
        height: 100%;
        background-color:#202020;
    }
</style>

</head>
<body>
    <div>
        <video src="{{ . }}" autoplay>
        </video>
    </div>
</body>
</html>`,
      cmOption: {
        tabSize: 4,
        styleActiveLine: true,
        lineNumbers: true,
        line: true,
        foldGutter: true,
        styleSelectedText: true,
        mode: 'text/javascript',
        keyMap: "sublime",
        matchBrackets: true,
        showCursorWhenSelecting: true,
        theme: "monokai",
        extraKeys: { "Ctrl": "autocomplete" },
        hintOptions:{
          completeSingle: false
        },
      }
    };
  },
  computed: {
    codemirror() {
      return this.$refs.myCm.codemirror
    }
  },
  mounted() {

  },
};
</script>

<style scoped>

</style>
