<template>
  <div class="output-f" v-html='compiledMarkdown'></div>
</template>

<script>

import {mapState} from "vuex";
import markdown from '../../api/markdown';

export default {
  data(){
    return{
      scrollTrigger: false
    }
  },
  computed: {
    ...mapState([
      'editorText',
    ]),
    compiledMarkdown() {
      return markdown.markdownToHTML(this.editorText);
    }
  },
  methods: {

  }
}
</script>

<style scoped>
.output-f{
  box-sizing: border-box;
  height: 600px;
  width: 100%;
  padding: 15px;
  resize: none;
  border: none ;
  background-color: #f5f5f5;
  outline: none;
  font-family: inherit;
  font-size: 18px;
  color: #616161;
  box-shadow: 4px 5px 3px #aaa;
  white-space: normal;
  overflow-y: scroll;
  word-wrap: break-word;
}
</style>
