<template>
  <div :class="a ? 'active inner' : 'inner'">
    <div class="column left">
      <div class="node">
        <div class="itemhead">
          SYSTEMS
        </div>
        <div class="subhead">
          ROOT NODE
        </div>
      </div>
      <div class="up">
        <div class="up1">
          <i class="icon ion-ios-arrow-up"/>
        </div>
        <div class="up2">
          <i class="icon ion-ios-arrow-up"/>
          <i class="icon ion-ios-arrow-up shift"/>
        </div>
      </div>
      <div class="add">
        <i class="icon ion-ios-plus-empty"/>
      </div>
    </div>
    <div class="column center">
      <div class="addpart">
        <i class="icon ion-ios-plus-empty"/>
      </div>
    </div>
    <div class="column right">
      <div @click="menu()" class="item">
        <i class="icon ion-ios-more-outline"/>
      </div>
      <div class="item">
        <i class="icon ion-ios-search"/>
      </div>
    </div>
  </div>
</template>

<script>
import anime from 'animejs'

export default {
  name: 'partspanel',
  data() {
    return {
      a: false
    }
  },
  mounted() {
    setTimeout(() => {
      this.a = true
    }, 600)
  },
  methods: {
    menu: function() {
      this.a = false;
      setTimeout(()=>{this.$router.push('/')},500)
    }
  }
}
</script>

<style scoped>
.inner {
  min-height: 100%;
  width: 100%;
  display: flex;
  flex-grow: 1;
  justify-content: center;
  align-items: center;
  position: relative;
  opacity: 0;
  transition: opacity 0.2s ease;
}
.active.inner {
  opacity: 1;
}
.add {
  background: rgba(255,255,255,0.6);
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;
  transition: background-color 0.2s ease;
  border-radius: 2px;
  cursor: pointer;
}
.add:hover {
  background: white;
}
.node {
  background: rgba(255,255,255,0.6);
  width: 100%;
  padding: 20px;
  height: 130px;
  border-radius: 2px;
  margin-bottom: 10px;
  display: flex;
  justify-content: center;
  flex-flow: column nowrap;
  align-items: center;
}
.left {
  flex-grow: 2;
  width: 200px;
}
.column {
  height: 100%;
  flex-shrink: 0;
  height: fill-available;
  height: -webkit-fill-available;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  flex-flow: column wrap;
}
.right {
  flex-grow: 1;
}
.center {
  flex-grow: 5;
  margin-left: 10px;
  width: 800px;
  margin-right: 10px;
}
.icon {
  font-size: 40px;
}
.item:hover {
  background: white;
}
.item {
  display: flex;
  position: relative;
  justify-content: center;
  align-items: center;
  flex-flow: column nowrap;
  cursor: pointer;
  width: 100%;
  padding: 30px;
  border-radius: 2px;
  background: rgba(255,255,255,0.6);
  transition: background-color 0.2s ease,
              color 0.2s ease;
}
.item:not(:last-child) {
  margin-bottom: 10px;
}
.itemhead {
  font-size: 30px;
  font-weight: 600;
}
.ion-ios-more-outline {
  height: 25px;
  margin-top: -25px;
}
.ion-ios-search {
  height: 25px;
  width: 25px;
  margin-top: -25px;
  font-size: 34px;
  margin-bottom: -8px;
}
.part {
  padding: 30px;
  width: 100%;
  background: rgba(255,255,255,0.6);
  border-radius: 2px;
}
.addpart {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background: rgba(255,255,255,0.6);
  border-radius: 2px;
  height: 60px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}
.addpart:hover {
  background: white;
}
.add {
  padding: 30px;
  width: 100%;
  background: rgba(255,255,255,0.6);
  border-radius: 2px;
}
.part:not(:last-child) {
  margin-bottom: 10px;
}
.up {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  margin-bottom: 10px;
}
.up1 {
  background: rgba(255,255,255,0.6);
  flex-grow: 1;
  height: 60px;
  margin-right: 10px;
  transition: background-color 0.2s ease;
  border-radius: 2px;
}
.up1:hover {
  background: white;
  cursor: pointer;
}
.up2:hover {
  background: white;
  cursor: pointer;
}
.up2 {
  background: rgba(255,255,255,0.6);
  flex-grow: 1;
  height: 60px;
  transition: background-color 0.2s ease;
  border-radius: 2px;
}
.up div {
  display: flex;
  justify-content: center;
  align-items: center;
}
.subhead {
  font-size: 15px;
  font-weight: 700;
}
.shift {
  margin-left: -2px;
}
</style>
