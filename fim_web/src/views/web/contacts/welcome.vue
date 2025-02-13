<script setup lang="ts">
import {useStore} from "@/stores/stores.ts";
const store = useStore()

</script>


<template>
  <div v-for="item  in store.config.menuIcon">
    <i :class="item"/>
  </div>
</template>