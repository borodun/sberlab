<template>
  <div id="app">
    <Carcass/>
  </div>
</template>

<script>
import Carcass from "@/components/Carcass";

export default {
  components: {
    Carcass
  },
}
</script>

<style>
body,html {
  height: 100%;
  margin: 0;
  padding: 0px;
  min-height: 100vh;
  width: 100%;
}

#app {
  display: flex;
  height: 100%;
  width: 100%;
  position: center;
}
</style>