<template>
  <div id="app">
    <nav class="nav-top">
      <img class="logo" src="./assets/chlogo.png" alt="chlogo">
      <router-link to="/">Home</router-link>
      <router-link to="/about">About</router-link>
      <router-link to="/auth/register">Register</router-link>
      <router-link to="/auth/login">Login</router-link>
      <router-link :to="{ name: 'dashboard/expenses' }">Dashboard</router-link>
    </nav>
    <router-view/>
  </div>
</template>

<style>
  body {
    background-color: #f7f7f7;
  }
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    background-color: #f7f7f7;
  }

  .nav-top {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-evenly;
    align-items: center;
  }

  .logo {
    height: 50px;
  }

  nav a {
    font-weight: bold;
    color: #2c3e50;
  }

  nav a.router-link-exact-active {
    @apply text-indigo-dark;
  }
</style>
