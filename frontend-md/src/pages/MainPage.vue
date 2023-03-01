<template>
  <div>
    <NavBar/>
    <v-container fluid>
      <v-row>
        <v-col
            class="d-flex"
            cols="12"
            sm="6"
      >
    <InputField/>
        </v-col>
        <v-col
            class="d-flex"
            cols="12"
            sm="6"
        >
    <OutputField/>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
import InputField from "@/components/fields/InputField";
import OutputField from "@/components/fields/OutputField";

export default {
  name: 'MainPage',
  components: {OutputField, InputField, NavBar},
  data () {
    return {
    }
  },
}
</script>

<style scoped>

</style>
