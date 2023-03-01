<template>
        <v-textarea
            solo
            id="inputer"
            name="input-7-4"
            height="600px"
            color="#f5f5f5"
            v-model="textAreaContent"
            label="Start typing here..."
            class="inputter-field"
        ></v-textarea>
</template>

<script>
import {mapMutations, mapState} from "vuex";

export default {
  computed: {
    ...mapState([
      'editorText',
    ]),
    textAreaContent: {
      get() {
        return this.editorText;
      },
      set(val) {
        this.$store.dispatch('inputText', val)
        this.$store.dispatch('updateOffset', val.length)
      },
    },
  },
  methods: {
    ...mapMutations([
      'inputText',
    ]),
  },
name: 'InputField',

data () {
return {
    }
  },
}
</script>

<style scoped>
.inputter-field{
  box-sizing: border-box;
  resize: none;
  border: none ;
  outline: none;
}
</style>
