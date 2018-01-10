<template>
  <div class="centerer">
  <div :class="activeCell!=undefined ? 'active sections ' + activeCell : 'sections'">
    <div @click="setCell('one')" :class="activeCell==='one' ? 'active row one' : 'row one'">
      SYSTEMS
    </div>
    <div @click="setCell('two')" :class="activeCell==='two' ? 'active row two' : 'row two'">
      USER MGMT
    </div>
    <div @click="setCell('three')" :class="activeCell==='three' ? 'active row three' : 'row three'">
      ORDERS (soon™)
    </div>
  </div>
  </div>
</template>

<script>
export default {
  name: 'BoilerplateSplash',
  data() {
    return {
      activeCell: undefined
    }
  },
  methods: {
    setCell: function(cell) {
      if (this.activeCell === undefined) {
        this.activeCell = cell
        return
      }
      this.activeCell = undefined
    }
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
  color: #EDF2F4;
  transition: height 0.3s ease;
  cursor: pointer;
  flex-shrink: 1;
  will-change: height;
  font-weight: 900;
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
.row.active {
  height: 375vh;
}
.row.two.active {
  height: 340vh;
}
</style>
