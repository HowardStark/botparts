<template>
  <div :class="a ? 'active centerer' : 'centerer'">
  <div :class="activeCell!=undefined ? 'active sections ' + activeCell : 'sections'">
    <div @click="setCell('one')" :class="activeCell==='one' ? 'active row one' : 'row one'">
      <div class="heading">
        SYSTEMS
      </div>
    </div>
    <div @click="setCell('two')" :class="activeCell==='two' ? 'active row two' : 'row two'">
      <div class="heading">
        USER MGMT (soon™)
      </div>
      <div class="container">
      </div>
    </div>
    <div @click="setCell('three')" :class="activeCell==='three' ? 'active row three' : 'row three'">
      <div class="heading">
        ORDERS (soon™)
      </div>
      <div class="container">
      </div>
    </div>
  </div>
  </div>
</template>

<script>
export default {
  name: 'BoilerplateSplash',
  data() {
    return {
      activeCell: undefined,
      a: false
    }
  },
  methods: {
    setCell: function(cell) {
      if (cell!='one') return;
      this.activeCell = cell
      var translation = {
        one: 'Systems'
      }
      setTimeout(() => {this.$router.push(translation[cell])}, 500)
    }
  },
  mounted() {
    setTimeout(() => {
      this.a = true
    }, 600)
  }
}
</script>

<style scoped>
.centerer {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  overflow: visible;
  align-items: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}
.centerer.active {
  opacity: 1;
}
.sections {
  height: 120vh;
  width: 120vw;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-flow: column nowrap;
  position: fixed;
  will-change: transform;
  transform: rotate(5deg);
  overflow: visible;
  transition: transform 0.4s ease;
}
.active.sections.one {
  transform: translateY(10vh);
}
.active.sections.two {
  transform: translateY(0vh);
}
.active.sections.three {
  transform: translateY(-10vh);
}
.row {
  height: calc(120vh/3);
  width: 120vw;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-flow: column nowrap;
  color: #EDF2F4;
  transition: height 0.3s ease;
  flex-shrink: 1;
  will-change: height;
  font-weight: 900;
  position: relative;
  font-size: 10vmin;
}
.row.one {
  background: #2b2d42;
}
.row.two {
  background: #EF233C;
}
.row.three {
  background: #8D99AE;
}
.row:hover {
  height:calc(45vh);
}
.row:not(.active) {
  cursor: pointer;
}
.row.active {
  height: 400vh;
}
.row.two.active {
  height: 400vh;
}
.heading {
  transition: all 0.4s ease;
  position: absolute;
  top: 50%;
  margin-top: -5vmin;
}
.row.active .heading {
  top: 0;
  opacity: 0;
  margin-top: 0;
}
</style>
