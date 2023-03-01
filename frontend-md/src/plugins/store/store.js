import Vue from 'vue'
import Vuex from 'vuex'
import ApiService from "@/services/ApiService";

const highlight = require('highlight.js')
const marked = require('marked')

marked.setOptions({
  highlight: function (code) {
    return highlight.highlightAuto(code).value
  }
})

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    name: '',
    filename: '',

    path:"",
    editorText: '',
    renderedHTML: '',
    offset: 0,
    logged: false,
    jwt: "",
    templates: [
      {
        template_id: 0,
        owner_id: 0,
        name: '',
        filename: '',
      }
    ],
    sign_errors: [],
  },
  mutations: {
    setProductsLoading(state, payload) {
      state.productsLoading = payload;
    },
    currentPath(state, path) {
      state.path = path;
    },
    updateEditorText(state, text) {
      state.editorText = text;
    },
    pushCurrentTemplate(state, data){
      state.name = data.name;
      state.filename = data.filename;
      state.editorText = data.text;
    },
    clearCurrentTemplate(state){
      state.name = '';
      state.filename = '';
    },
    updateOffsetText(state, off) {
      state.offset = off;
    },
    searchText(state, text) {
      state.currentQuery = text;
    },
    pushResult(state, g){
      state.gifs = g;
    },
    pushTemps(state, t){
      state.templates = t;
    },
    userLogged(state, jwt){
      state.jwt = jwt;
      state.logged = true;
    },
    clearSignInErrors(state){
      state.sign_errors = [];
    },
    pushSignInErrors(state, error){
      state.sign_errors.push(error);
    },
    deleteTemplate(state, t_id){
      state.templates.filter(function(f) { return f.template_id !== t_id })
    },
  },
  actions: {
    inputText ({ commit }, txt){
      commit("updateEditorText", txt)
    },
    pushCurrentTemplate({ commit }, data){
      ApiService.getTemplateById(data.template_id, data.jwt).
      then(result => {
        commit("pushCurrentTemplate", result.data)
        }, error => {
        console.error(error);
      });
    },
    createTemplate({ commit }, data){
      ApiService.createTemplate({
            name: data.name,
            filename: data.filename,
            text: data.text}, data.jwt).then(response => {
        if (response.status === 200) {
          commit('pushCurrentTemplate', data);
        }
      }, error => {
        console.error(error);
      });
    },
    updateTemplate({commit}, data){
      ApiService.updateTemplate(data.templateId, {name : data.name,
                                                       filename: data.filename,
                                                       text: data.text}, data.jwt)
          .then(response => {
            if (response.status === 200){
              commit('updateEditorText', data.text);
            }
          })
    },
    updateOffset({ commit }, off){
      commit("updateOffsetText", off)
    },
    pushTemplates ({ commit }, data){
      ApiService.getAllTemplates(data.jwt).then(result => {
        let ts = [];
        result.data.data.forEach(element => {
          ts.push(element)
          console.log(element)
        })
        commit("pushTemps", ts)
      }, error => {
        console.error(error);
      });
    },
    signIn({commit}, data){
      ApiService.signInUser(data).then(response => {
        if (response.status === 200) {
          commit('userLogged', response.data.token);
          commit('clearSignInErrors');
        }
      }, error => {
        commit('pushSignInErrors', error)
      });
    },
    signUp({commit}, data){
      ApiService.signUpUser(data).then(response => {
        if (response.status === 200) {
          commit('clearSignInErrors');
        }
      }, error => {
        console.error(error);
      });
    },
    deleteTemplate({commit}, data){
      ApiService.deleteTemplate(data.template_id, data.jwt).
      then( result => {
            if (result.status === 200) {
              commit("deleteTemplate", data.template_id)
            }
          },
          error => { console.error(error); });
    }
  },
  getters: {
    getController : state => {
      return state.controller;
    },
    getProductsLoading : state => {
      return state.productsLoading;
    },
    getPath : state => {
      return state.path;
    },
  },
})
