<template>
  <div id="home">
    <HeaderComponent></HeaderComponent>
    <h2>首页</h2>
    <main>
      <router-view :key="$route.path"></router-view>
    </main>
  </div>
</template>

<script setup lang="ts">
import HeaderComponent from '@/components/layout/HeaderComponent.vue'
</script>

<style lang="scss" scoped>
#home {
  min-height: 100vh;
  text-align: center;
  display: flex;
  flex-direction: column;
  title{
    height: 60px;

  }
  main{
    flex: 1;
  }
}
</style>
