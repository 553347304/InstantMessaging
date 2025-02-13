<script setup lang="ts">
</script>

<template>
  <div class="m_web">
    <div class="m_slide">
      <div class="avatar">
        <img src="https://vip.123pan.cn/1821560246/website/image/icon/role/1.jpg" alt="">
      </div>
      <div class="m_menu">
        <router-link class="icon" :to="{name: 'session_welcome'}"/>
        <div class="icon">
        </div>
        <div class="icon">通知</div>
        <div class="icon">其他</div>
        <div class="other">其他</div>
      </div>
    </div>
    <div class="m_main">
      <router-view/>
    </div>
  </div>
</template>

<style scoped lang="scss">
.m_web {
  @include bg();

  .m_slide {
    width: 80px;
    height: 100%;
    background-color: #e0e000;
    position: relative;

    display: flex;
    flex-direction: column;
    align-items: center;

    @include avatar();

    .m_menu {
      margin-top: 20px;
      display: flex;
      flex-direction: column;
      //align-items: center;


      > div {
        margin-bottom: 20px;
      }


    }

    .icon {
      cursor: pointer;
      font-size: 20px;
      width: 40px;
      height: 40px;
      border-radius: 10px;
      //display: flex;
      //justify-content: center;
      //align-items: center;
      transition: all .2s;

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
}
</style>