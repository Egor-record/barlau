<template>
  <main class="mt-5">
    <div class="text-center">
      <img src="../assets/logo.png" class="text-center">
      <h1 class="mt-5">{{ msg }}</h1>
    </div>
    <div class="container mt-5">
       <div class="row">
          <div class="col-12 col-md-6 mt-3 mt-md-0 justify-content-md-end d-md-flex">
            <div class="item-menu text-white rounded mr-md-5 text-center p-5 bg-blue">
              <router-link class="nav-link" to="/checker">
                <b-icon-check font-scale="5"></b-icon-check>
                <h4 class="text-white mt-4">Play your manifest</h4>
              </router-link>
            </div>
          </div>
           <div class="col-12 col-md-6  mt-3 mt-md-0">
             <div class="item-menu text-white ml-md-5 rounded text-center p-5 bg-blue">
               <router-link class="nav-link" to="/generator">
                 <b-icon-file-code font-scale="5"></b-icon-file-code>
                 <h4 class="text-white mt-4">Create new app</h4>
               </router-link>
             </div>
           </div>
       </div>
    </div>
  </main>
</template>

<script>
export default {
  name: 'home',
  data() {
    return {
      msg: 'webOS App Development Helper',
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}

.item-menu {
  max-width: 260px;
  max-height: 250px;
  height: 250px;
  width: 100%;
}
</style>
