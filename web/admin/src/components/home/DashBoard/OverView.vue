<template>
  <div id="overview">
    <div class="overview-list">
      <div class="overview-item">
        <div class="title">123</div>
        <div class="number">123</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style lang="scss" scoped>
#overview {
  .overview-list {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 1fr 1fr;
    .overview-item {
      
      display: flex;
      flex-direction: column;
    }
  }
}
</style>
