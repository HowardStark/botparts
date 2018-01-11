<template>
  <div class="container">
    <PartsPanel/>
  </div>
</template>

<script>
import anime from 'animejs'
import PartsPanel from '@/components/UIComponents/PartsPanel'

export default {
  components: {PartsPanel},
  name: 'systems'
}
</script>

<style scoped>
.container {
  color: #2b2d42;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
  min-height: 100vh;
  width: 100vw;
}
</style>
