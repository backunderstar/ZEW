<template>
  <div id="home">
    <el-container>
      <el-affix>
        <el-aside width="200px">
          <div class="title">
            <h2>朝而往</h2>
          </div>
          <AdminMenu />
        </el-aside>
      </el-affix>
      <el-container>
        <el-header>
          <AdminHeader />
        </el-header>
        <el-main>
          <router-view />
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import AdminHeader from '@/components/layout/AdminHeader.vue'
import AdminMenu from '@/components/layout/AdminMenu.vue'
</script>

<style lang="scss" scoped>
#home {
  min-height: 100vh;

  .el-aside {
    background-color: #d0e7de;
    min-height: 100vh;

    .title {
      display: flex;
      height: 3.75rem;
      align-items: center;
      justify-content: center;
    }
  }
  .el-header {
    height: 3.75rem;
    background-color: #d8eee5;
  }

  .el-footer {
    height: 2rem;
  }
}
</style>
