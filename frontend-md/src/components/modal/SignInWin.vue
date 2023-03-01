<template>
  <div class="mx-3">
    <v-dialog
        v-model="dialog_in"
        persistent
        max-width="600px"
        v-if="!isLogged"
    >
      <v-card>
        <v-card-title>
          <span class="text-h5">Account login</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                    label="Username*"
                    v-model="form.username"
                    :rules="rules.username"
                    counter
                    minlength="8"
                    required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                    label="Password*"
                    v-model="form.password"
                    :append-icon="show_pass ? 'mdi-eye' : 'mdi-eye-off'"
                    :type="show_pass ? 'text' : 'password'"
                    :rules="rules.password"
                    name="input-10-1"
                    hint="At least 8 characters"
                    counter
                    @click:append="show_pass = !show_pass"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
          <small style="color: red" v-if="notCorrect">Unable to sign in</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>

          <SignUp></SignUp>

          <v-btn
              color="blue darken-1"
              text
              @click="dialog_in = false"
              :disabled="!isLogged"
          >
            Close
          </v-btn>
          <v-btn
              color="blue darken-1"
              text
              @click="sendSignInForm"
          >
            Sign In
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import {mapState} from "vuex";
import SignUp from "@/components/modal/SignUpWin";

export default {
  components: {SignUp},
  computed:{
    ...mapState([
      'editorText',
      'logged',
      'sign_errors',
      'jwt'
    ]),

    isLogged(){
      return this.logged
    },
    notCorrect(){
      return this.sign_errors.length > 0
    },
  },

  methods:{
    sendSignInForm(){
      let data = new FormData();
      data.append('username', this.form.username);
      data.append('password', this.form.password);
      this.$store.dispatch('signIn', {username: this.form.username, password: this.form.password});
      if (this.isLogged){
        this.dialog_in = false;
      }
    },
  },
  data: () => ({
    form: {
      username: "",
      password: "",
    },
    dialog_in: true,
    show_pass: false,
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
