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
            class="shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker mb-3 leading-tight focus:outline-none focus:shadow-outline"
            id="password" type="password" placeholder="******************" v-model="password" required
            maxlength="16" minlength="6">
        </div>
        <div class="flex items-center bg-red text-white text-sm font-bold px-4 py-3 mb-2" role="alert" v-if="error">
          <svg class="fill-current w-4 h-4 mr-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
            <path
              d="M12.432 0c1.34 0 2.01.912 2.01 1.957 0 1.305-1.164 2.512-2.679 2.512-1.269 0-2.009-.75-1.974-1.99C9.789 1.436 10.67 0 12.432 0zM8.309 20c-1.058 0-1.833-.652-1.093-3.524l1.214-5.092c.211-.814.246-1.141 0-1.141-.317 0-1.689.562-2.502 1.117l-.528-.88c2.572-2.186 5.531-3.467 6.801-3.467 1.057 0 1.233 1.273.705 3.23l-1.391 5.352c-.246.945-.141 1.271.106 1.271.317 0 1.357-.392 2.379-1.207l.6.814C12.098 19.02 9.365 20 8.309 20z"/>
          </svg>
          <p>{{ error }}</p>
        </div>
        <div class="flex items-center bg-green text-white text-sm font-bold px-4 py-3 mb-2" role="alert" v-if="message">
          <svg class="fill-current w-4 h-4 mr-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
            <path
              d="M12.432 0c1.34 0 2.01.912 2.01 1.957 0 1.305-1.164 2.512-2.679 2.512-1.269 0-2.009-.75-1.974-1.99C9.789 1.436 10.67 0 12.432 0zM8.309 20c-1.058 0-1.833-.652-1.093-3.524l1.214-5.092c.211-.814.246-1.141 0-1.141-.317 0-1.689.562-2.502 1.117l-.528-.88c2.572-2.186 5.531-3.467 6.801-3.467 1.057 0 1.233 1.273.705 3.23l-1.391 5.352c-.246.945-.141 1.271.106 1.271.317 0 1.357-.392 2.379-1.207l.6.814C12.098 19.02 9.365 20 8.309 20z"/>
          </svg>
          <p>{{ message }}</p>
        </div>
        <div class="flex items-center justify-between">
          <button
            class="bg-blue hover:bg-blue-dark text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit">
            Sign In
          </button>
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
import Config from '@/config'

@Component
export default class Register extends Vue {
  email: string = ''
  password: string = ''
  error: string = ''
  message: string = ''

  send () {
    this.error = ''
    this.message = ''
    let data = {
      'email': this.email,
      'password': this.password
    }
    let headers = new Headers()
    headers.set('Content-Type', 'application/json')
    fetch(Config.baseUrl + 'auth/register', {
      method: 'POST',
      body: JSON.stringify(data),
      headers: headers
    })
      .then(res => {
        if (res.status === 201) {
          this.message = this.$t('messages.user_created') as string
        } else {
          throw new Error(res.statusText)
        }
      })
      .catch(err => {
        if (err.message === 'Conflict') {
          this.error = this.$t('errors.unique_violation') as string
        } else {
          this.error = this.$t('errors.undefined') as string
        }
      })
  }
}
</script>

<i18n>
  {
    "en": {
      "errors": {
        "unique_violation": "This email is already registered",
        "undefined": "Sorry, something is going wrong"
      },
      "messages": {
        "user_created": "User was successfully created"
      }
    }
  }
</i18n>
