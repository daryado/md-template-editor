<template>
  <div>
    <v-toolbar elevation="3">
      <v-toolbar-items>
        <HeadingsBtn/>
        <InsertBtn title="Bold" to_insert="**"></InsertBtn>
        <InsertBtn title="Italic" to_insert="*"></InsertBtn>
        <InsertBtn title="Link" to_insert="[Link]"></InsertBtn>
        <InsertBtn title="Img" to_insert="![Img]"></InsertBtn>
        <InsertBtn title="Code" to_insert="`"></InsertBtn>
        <InsertBtn title="Quote" to_insert="
         > "></InsertBtn>
        <InsertBtn title="List" to_insert="- "></InsertBtn>
        <InsertBtn title="Table"
                   to_insert=" Column1 | Column2
------ | ------
Cell1   | Cell2   "></InsertBtn>
        <InsertBtn title="Separator"
                   to_insert="---"></InsertBtn>
          <v-btn :href='mdDataUrl'
                 :download='titleMd'
                 @mouseenter='createUrl("md")'
                 color="#BBDEFB">
            Download
          </v-btn>
      </v-toolbar-items>
      <v-spacer></v-spacer>
        <modal-win/>
        <SaveForm/>
        <SignIn/>
    </v-toolbar>
  </div>
</template>

<script>

import ModalWin from "@/components/modal/ChooseTemplateWin";
import SaveForm from "@/components/modal/SaveFormWin";
import HeadingsBtn from "@/components/btns/HeadingsBtn";
import insert from '../plugins/insert';
import InsertBtn from "@/components/btns/InsertBtn";
import SignIn from "@/components/modal/SignInWin";
import {mapState} from "vuex";

export default {
  components: { InsertBtn, HeadingsBtn, SaveForm, ModalWin, SignIn},
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
        case '[Link](http://example.com/)':
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
        case '`code`':
          newContent = insert.upDateContent(inputer, oldContent, newContent, startPosition, endPosition, '`', '`');
          break;
        case '[Link](http://example.com/)':
          newContent = insert.upDateContent(inputer, oldContent, newContent, startPosition, endPosition, '[', '](http://example.com/)');
          break;
        case '![Img](http://example.com/)':
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
    createUrl (mode) {
      let self = this
      let val = ''
      if (mode === 'md') {
        val = self.$store.getters.articleRaw
        let blobObj = new Blob([val])
        let objectURL = URL.createObjectURL(blobObj)
        self.mdDataUrl = objectURL
      } else {
        val = self.editorText
        let blobObj = new Blob([val])
        let objectURL = URL.createObjectURL(blobObj)
        self.htmlDataUrl = objectURL
      }
    }
  },
  computed: {
    ...mapState([
      'editorText',
      'logged',
    ]),
    titleMd () {
      return this.$store.state.editorText.split('\n')[0] + '.md'
    }
  },
  data(){
    return {
      htmlDataUrl: '',
      mdDataUrl: ''
    }
  },
  // redirect(url) {
  //   window.open(url, '_blank');
  // },
};
</script>

<style>

</style>

