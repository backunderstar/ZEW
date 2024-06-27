<template>
  <div id="dashboard">
    <overView class="item" />
  </div>
</template>

<script setup lang="ts">
import overView from './DashBoard/OverView.vue'
</script>

<style lang="scss" scoped>
#dashboard {
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  width: 100%;
  height: 100%;

  overflow: hidden;
  box-shadow: var(--el-box-shadow-light);
  border-radius: 0.5rem;

  .item {
    min-height: 20rem;
    padding: 1rem;
    border: 1px solid #f2f2f2;
  }
}
</style>
