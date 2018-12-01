import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    expenses: []
  },
  mutations: {
    setExpenses (state, expenses) {
      state.expenses = expenses
    }
  },
  actions: {

  }
})
