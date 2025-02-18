<script setup lang="ts">
import {useStore} from "@/stores";
const store = useStore()
store.init()
</script>

<template>
  <div class="m_body">
    <router-view/>
  </div>
</template>

<style scoped lang="scss">
.m_body {
  width: 100%;
  height: 100vh;
  overflow-y: hidden;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
