<template>
  <div class="navcontainer">
    <div class="navbar">
      <div class="heading">
        4904PARTS
      </div>
    </div>
  </div>
</template>

<script>
import anime from 'animejs'

export default {
  name: 'navbar'
}
</script>

<style scoped>
.navcontainer {
  width: 100%;
  display: flex;
  justify-content: center;
}
.navbar {
  background: #333333;
  height: 60px;
  width: 600px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding-left: 15px;
}
.heading {
  font-weight: 600;
  font-family: Avenir Next;
  color: #eb;
  font-size: 1.5rem;
}
</style>
