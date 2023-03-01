<template>
  <div class="mx-3">
    <v-dialog
        v-model="dialog"
        fullscreen
        hide-overlay
        transition="dialog-bottom-transition"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
            color="primary"
            dark
            v-bind="attrs"
            v-on="on"
            :disabled="!isLogged"
            @click="getAllTemplates"
        >
          Template
        </v-btn>
      </template>
      <v-card>
        <v-toolbar
            dark
            color="primary"
        >
          <v-btn
              icon
              dark
              @click="dialog = false"
          >
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>Saved Templates</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn
                class="mx-2"
                dark
                text
                :disabled = 'isEmptyTemplatesList'
                @click="deleteTemplate"
            >Delete</v-btn>
            <v-btn
                dark
                text
                :disabled = 'isEmptyTemplatesList'
                @click="chooseItem"
            >
              Open
            </v-btn>
          </v-toolbar-items>
        </v-toolbar>
        <v-list
            three-line
            subheader
        ></v-list>
        <v-main>
          <v-card>
            <v-list three-line>
              <v-list-item-group
                  v-model="selectedItem"
                  active-class="blue--text"
              >
                <v-list-item
                    v-for="t in getTemplates"
                    :key="t.template_id"
                    @click="selectedItem = t.template_id"
                >
                  <v-list-item-content class="template-row">
                    <v-list-item-title v-html="t.name" ></v-list-item-title>
                    <v-list-item-subtitle v-html="t.filename"></v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-divider></v-divider>
              </v-list-item-group>
            </v-list>
          </v-card>
        </v-main>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import {mapState} from "vuex";


export default {
  computed: {
    ...mapState([
      'logged',
      'templates',
      'jwt',
      'editorText',
    ]),
    isLogged(){
      return this.logged
    },
    getTemplates(){
      return this.templates
    },
    isEmptyTemplatesList(){
      return this.templates.length === 0
    },
  },
  methods:{
    getAllTemplates(){

      this.$store.dispatch('pushTemplates', {jwt: this.jwt});
    },
    chooseItem(){
      this.$store.dispatch('pushCurrentTemplate', {template_id : this.templates[this.selectedItem].template_id,
                                                                jwt: this.jwt})

      this.$store.dispatch('updateOffset', this.editorText.length);
      this.dialog = false;
    },
    deleteTemplate(){
      this.$store.dispatch('deleteTemplate', {template_id : this.templates[this.selectedItem].template_id,
                                                          jwt: this.jwt})
      this.dialog = false;
    },
  },
  data () {
    return {
      selectedItem: 0,
      selectedTemplate: "",
      dialog: false,
      notifications: false,
      rowSelected: false,
      widgets: false,
    }
  },
}
</script>

<style>

.template-row{
  display: inline-flex;
}
</style>
