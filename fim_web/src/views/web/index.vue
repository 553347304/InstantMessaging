<script setup lang="ts">
import {useRoute} from "vue-router";
import {useStore} from "@/stores";

const route = useRoute()
const store = useStore()
console.log(route)
</script>

<template>
  <div class="m_web">
    <div class="m_slide">
      <div class="avatar" style="margin-top: 20px">
        <img :src="store.userInfo.avatar" alt="">
      </div>
      <div class="m_menu">
        <router-link class="icon" :to="{name: 'session_welcome'}"><i class="iconfont icon-liaotian"/></router-link>
        <router-link class="icon" :to="{name: 'contacts'}"><i class="iconfont icon-yonghu1"/></router-link>
        <router-link class="icon" :to="{name: 'notice'}"><i class="iconfont icon-mail"/></router-link>
      </div>
      <router-link class="other icon" :to="{name: 'info'}"><i class="iconfont icon-shezhi1"/></router-link>
    </div>
    <router-view class="m_main"/>
  </div>
</template>

<style scoped lang="scss">
.m_web {
  width: 1200px;
  height: 600px;
  background-color: var(--el-bg-color);
  border-radius: 5px;
  overflow: hidden;
  box-shadow: 0 0 5px 3px rgba(0, 0, 0, 0.05);
  display: flex;

  .m_slide {
    width: 60px;
    height: 100%;
    background-color: var(--nav_bg);
    position: relative;

    display: flex;
    flex-direction: column;
    align-items: center;

    @include avatar();

    .m_menu {
      margin-top: 20px;
      display: flex;
      flex-direction: column;
      align-items: center;


      > a {
        margin-bottom: 20px;
      }
    }

    .icon {
      cursor: pointer;
      font-size: 20px;
      width: 40px;
      height: 40px;
      border-radius: 10px;
      display: flex;
      justify-content: center;
      align-items: center;
      transition: all .2s;
      text-decoration: none;

      > i {
        font-size: 30px;
      }

      &:hover {
        background-color: #d9d9d9;
      }

      &:active {
        color: #1684fc;
        background-color: #d9d9d9;
      }

    }

    .other {
      position: absolute;
      bottom: 10px;
    }

  }

  .m_main {
    background-color: var(--info_bg);
    width: calc(100% - 80px);
  }
}
</style>