<template>
  <div class="mx-3">
    <v-dialog
        v-model="dialog_up"
        persistent
        max-width="600px"
        v-if="!isLogged"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
            color="primary"
            dark
            v-bind="attrs"
            v-on="on"
        >
          Sign up
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="text-h5">Registration</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                    label="Username*"
                    :rules="rules.username"
                    v-model="form.username"
                    required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                    label="Password*"
                    v-model="form.password"
                    :rules="rules.password"
                    type="password"
                    hint="At least 8 characters"
                    counter
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
              @click="dialog_up = false"
          >
            Close
          </v-btn>
          <v-btn
              color="blue darken-1"
              text
              @click="sendSignUpForm"
          >
            Sign Up
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import {mapState} from "vuex";

export default {
  computed: {
    ...mapState([
      'editorText',
      'logged',
    ]),
    isLogged() {
      return this.logged
    },
  },
  methods: {
    sendSignUpForm() {
      this.$store.dispatch('signUp', {username: this.form.username, password: this.form.password})
      this.dialog_up = false;
    },
  },
    data: () => ({
      form: {
        username: "",
        password: "",
      },
      dialog_up: false,
      rules: {
        username: [
          val => (val || '').length > 0 || 'This field is required'
        ],
        password: [
          val => (val || '').length > 0 || 'This field is required'
        ]
      }
    }),
}
</script>

<style scoped>

</style>
