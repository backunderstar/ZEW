<template>
  <div id="header">
    <div class="header-left"></div>

    <div class="header-right">
      <div class="avatar">
        <el-avatar
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style lang="scss" scoped>
#header {
  display: flex;
  justify-content: space-between;
  height: 100%;
  width: 100%;

  .header-left {
    width: 200px;
    height: 100%;
  }

  .header-right{
    display: flex;
    align-items: center;
    justify-content: center;

    height: 100%;
  }
}
</style>
