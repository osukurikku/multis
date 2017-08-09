import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    clickedMatch: null
  },
  mutations: {
    setClickedMatch (state, match) {
      state.clickedMatch = match
    }
  }
})

export default store
