<template>
  <div>
    <b-navbar toggleable="lg" type="dark" class="bg-blue">
      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
      <b-collapse id="nav-collapse" is-nav class="container">
        <b-navbar-nav>
           <router-link to="/" class="nav-link">Home</router-link>
          <router-link to="/generator" class="nav-link">Create the app</router-link>
          <router-link to="/checker" class="nav-link">Play manifest</router-link>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-item-dropdown right>
            <!-- Using 'button-content' slot -->
            <template v-slot:button-content>
              <em>User</em>
            </template>
            <b-dropdown-item href="#">Profile</b-dropdown-item>
            <b-dropdown-item href="#">Sign Out</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: 'navigation',
};
</script>

<style scoped>

</style>
