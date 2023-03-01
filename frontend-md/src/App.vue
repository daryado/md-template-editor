<template>
    <v-app app>
      <v-main id="main">
          <router-view></router-view>
      </v-main>
    </v-app>
</template>

<script>
import 'vuetify/dist/vuetify.min.css'

export default {
  name: 'App',
  mounted() {
    let storageStringConfiguration = localStorage.getItem('configuration');
    let storageConfiguration = JSON.parse(storageStringConfiguration);
    if(storageConfiguration !== null) {
      for (let storageElement in storageConfiguration) {
        this.$store.commit(storageElement + '/setCurrent', storageConfiguration[storageElement]);
      }
    }
  }
}
</script>

<style>
{
font-family: Avenir, Helvetica, Arial, sans-serif;
-webkit-font-smoothing: antialiased;
-moz-osx-font-smoothing: grayscale;
text-align: center;
color: #2c3e50;
}
body {
  margin: 0;
}
</style>
