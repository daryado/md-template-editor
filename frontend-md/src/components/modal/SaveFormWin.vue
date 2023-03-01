<template>
  <div class="mx-3">
    <v-dialog
        v-model="dialog"
        persistent
        max-width="600px"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
            color="primary"
            dark
            v-bind="attrs"
            v-on="on"
        >
          Save Template
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="text-h5">Save as template</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                    label="Name*"
                    v-model="form.name"
                    required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                    label="Filename*"
                    v-model="form.filename"
                    required
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
              color="blue darken-1"
              text
              @click="dialog = false"
          >
            Close
          </v-btn>
          <v-btn
              color="blue darken-1"
              text
              @click="saveTemplate"
          >
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import {mapState} from "vuex";

export default {
  computed:{
    ...mapState([
      'editorText',
      'templates',
      'jwt',
    ]),
  },
  methods:{
    templateExist(){
      let exist = false
      this.templates.forEach(element => {
        console.log(element.filename)
        console.log(this.form.filename)
        if (element.filename === this.form.filename){
          this.templateId = element.template_id;
          exist = true;
        }
      })
      return exist;
    },
    saveTemplate(){
      if (this.templateExist()){
        this.$store.dispatch('updateTemplate', {
                            templateId: this.templateId,
                            name : this.form.name,
                            filename: this.form.filename,
                            text: this.editorText,
                            jwt: this.jwt})
      } else {
        this.$store.dispatch('createTemplate', {
          name: this.form.name,
          filename: this.form.filename,
          text: this.editorText,
          jwt: this.jwt})
        this.dialog_up = false;

      }
      this.name = '';
      this.filename = '';

      this.dialog = false;
    },
  },
  data: () => ({
    form: {
      name: '',
      filename: '',
    },
    dialog: false,
    templateId: 0,
  }),
}
</script>
