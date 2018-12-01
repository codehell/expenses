<template>
  <div class="flex justify-center container mx-auto">
    <div class="w-full max-w-xs">
      <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" @submit.prevent="send">
        <div class="mb-4">
          <label class="block text-grey-darker text-sm font-bold mb-2" for="email">
            Email
          </label>
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker leading-tight focus:outline-none focus:shadow-outline"
            id="email" type="email" placeholder="email" v-model="email" required>
        </div>
        <div class="mb-6">
          <label class="block text-grey-darker text-sm font-bold mb-2" for="password">
            Password
          </label>
          <input
            class="shadow appearance-none rounded w-full py-2 px-3 text-grey-darker mb-3 leading-tight focus:outline-none focus:shadow-outline"
            id="password" type="password" placeholder="******************" v-model="password" required
            maxlength="16" minlength="6">
        </div>
        <div class="flex items-center bg-red text-white text-sm font-bold px-4 py-3 mb-2" role="alert" v-if="error">
          <p>{{ error }}</p>
        </div>
        <div class="flex items-center bg-green text-white text-sm font-bold px-4 py-3 mb-2" role="alert" v-if="message">
          <p>{{ message }}</p>
        </div>
        <div class="flex items-center justify-between">
          <button
            class="bg-blue hover:bg-blue-dark text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit">
            Login
          </button>
          <a class="inline-block align-baseline font-bold text-sm text-blue hover:text-blue-darker" href="#">
            Forgot Password?
          </a>
        </div>
      </form>
      <p class="text-center text-grey text-xs">
        ©2018 Acme Corp. All rights reserved.
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'
import Config from '../../config'

@Component
export default class Login extends Vue {
  email: string = ''
  password: string = ''
  error: string = ''
  message: string = ''

  send () {
    let headers = new Headers()
    headers.set('Content-type', 'text/plain')
    headers.set('Authorization', 'Basic ' + btoa(this.email + ':' + this.password))
    fetch(Config.baseUrl + 'auth/login', {
      method: 'POST',
      headers: headers
    })
      .then(res => {
        this.error = ''
        this.message = ''
        if (res.status === 200) {
          res.json()
            .then((json) => {
              window.localStorage.setItem('token', json.token)
              this.$router.push({ name: 'dashboard/expenses' })
            })
          this.message = this.$t('messages.login_ok') as string
        } else if (res.status === 401) {
          this.error = this.$t('errors.user_not_found') as string
        } else {
          throw new Error(res.statusText)
        }
      })
      .catch(() => {
        this.error = this.$t('errors.undefined') as string
      })
  }
}
</script>

<i18n>
  {
    "en": {
      "errors": {
        "user_not_found": "Incorrect user or password"
      }
    },
    "messages": {
      "login_ok": "Correct login"
    }
  }
</i18n>
