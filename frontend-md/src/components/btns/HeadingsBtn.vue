<template>
  <v-menu
      left
      bottom
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn
          v-bind="attrs"
          v-on="on"
      >
        Headings
      </v-btn>
    </template>

    <v-list>
      <InsertItem title="level 1" to_insert="# "></InsertItem>
      <InsertItem title="level 2" to_insert="## "></InsertItem>
      <InsertItem title="level 3" to_insert="### "></InsertItem>
      <InsertItem title="level 4" to_insert="#### "></InsertItem>
      <InsertItem title="level 5" to_insert="##### "></InsertItem>
    </v-list>
  </v-menu>
</template>

<script>
import InsertItem from "@/components/InsertItem";
export default {
  name: "HeadingsBtn",
  components: {InsertItem}
}
</script>

<style scoped>

</style>
