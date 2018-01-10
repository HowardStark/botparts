<template>
  <div class="container">
    <div class="spinner">
    </div>
  </div>
</template>

<script>
export default {
  name: 'loader',
  mounted() {
    this.$router.push('Menu')
  }
}
</script>

<style scoped>
@keyframes load {
  0% {
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}

.container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100vw;
  height: 100vh;
}
.spinner,
.spinner:after {
  border-radius: 50%;
  width: 20vmin;
  height: 20vmin;
}
.spinner {
  margin: 60px auto;
  font-size: 10px;
  position: relative;
  text-indent: -9999em;
  border-top: 1.1em solid rgba(255, 255, 255, 0.2);
  border-right: 1.1em solid rgba(255, 255, 255, 0.2);
  border-bottom: 1.1em solid rgba(255, 255, 255, 0.2);
  border-left: 1.1em solid rgba(255,255,255,0.8);
  transform: translateZ(0);
  animation: load 1.1s infinite linear;
}
</style>
