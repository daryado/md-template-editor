<template>
  <v-btn
      @click="insert(to_insert)">{{title}}</v-btn>
</template>

<script>

import insert from "@/plugins/insert";

export default {
  name: "InsertBtn",
  props: {
    title: String,
    to_insert: String
  },
  methods: {
    insert(content) {
      const inputer = document.querySelector('#inputer');
      const startPosition = inputer.selectionStart;
      const endPosition = inputer.selectionEnd;
      const oldContent = inputer.value;

      inputer.focus();

      let newContent = '';

      if (startPosition === endPosition) {
        switch (content) {
          case '**':
            newContent = insert.setContent(inputer, oldContent, newContent, content, endPosition, 2, 2);
            break;
          case '*':
            newContent = insert.setContent(inputer, oldContent, newContent, content, endPosition, 1, 1);
            break;
          case '[Link]':
            newContent = insert.setContent(inputer, oldContent, newContent, content, endPosition, 7, 1);
            break;
          default:
            newContent = oldContent.substring(0, endPosition) + content + oldContent.substring(endPosition, oldContent.length);
            inputer.value = newContent;
            break;
        }
      } else {
        switch (content) {
          case '**':
            newContent = insert.upDateContent(inputer, oldContent, newContent, startPosition, endPosition, '**', '**');
            break;
          case '*':
            newContent = insert.upDateContent(inputer, oldContent, newContent, startPosition, endPosition, '*', '*');
            break;
          case '`':
            newContent = insert.upDateContent(inputer, oldContent, newContent, startPosition, endPosition, '`', '`');
            break;
          case '[Link]':
            newContent = insert.upDateContent(inputer, oldContent, newContent, startPosition, endPosition, '[', '](http://example.com/ "Необязательная подсказка")');
            break;
          case '![Img]':
            newContent = insert.upDateContent(inputer, oldContent, newContent, startPosition, endPosition, '![', '](http://example.com/)');
            break;
          default:
            return false;
        }
      }

      if (newContent.length) {
        this.$store.dispatch('inputText', newContent);
      }
    },
  },
}
</script>

<style scoped>

</style>
