<template>
  <div id="app">
    <transition name="fade" mode="out-in">
      <router-view class="mb-5"/>
    </transition>
    <div class="container mt-5">
      <footer class="mt-5">
        <div class="text-center">
         <p class="mt-1">
           Powered by <a href="https://gofiber.io/">Fiber</a> and <a href="https://github.com/webosose/ares-cli"> ares-cli</a>
           • Made by <a href="https://github.com/Egor-record">George Siaskov</a>
         </p>
        </div>
      </footer>
    </div>
  </div>
</template>

<script>

export default {
  name: 'App',
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
  .fade-enter-active, .fade-leave-active {
  transition: opacity .30s ease;
}
.fade-enter, .fade-leave-active {
  opacity: 0;
}
.child-view {
  position: absolute;
  transition: all .75s cubic-bezier(.55,0,.1,1);
}
.slide-left-enter, .slide-right-leave-active {
  opacity: 0;
  -webkit-transform: translate(30px, 0);
  transform: translate(30px, 0);
}
.slide-left-leave-active, .slide-right-enter {
  opacity: 0;
  -webkit-transform: translate(-30px, 0);
  transform: translate(-30px, 0);
}
:root {
  --dark-blue: #2e425c;
}

  .bg-blue {
      background-color: var(--dark-blue);
  }

footer {
  border-top: 1px dotted #000;
  font-size: 75%;
}
</style>
