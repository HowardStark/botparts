<template>
  <div id="app">
    <transition>
      <router-view/>
    </transition>
  </div>
</template>

<script>
import NavBar from '@/components/UIComponents/NavBar'

export default {
  components: {NavBar},
  name: 'app'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
body {
  margin: 0;
  font-size: 14px;
  background: #2b2d42;
  user-select: none;
}
* {
  box-sizing: border-box;
}
</style>
