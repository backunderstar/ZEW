<template>
  <div id="menu">
    <el-menu
      active-text-color="#008040"
      text-color="#000000"
      background-color="#d0e7de"
      :default-active="defaultActive"
      class="el-menu-vertical-demo"
      router
    >
      <component
        :is="item.children ? ElSubMenu : ElMenuItem"
        v-for="item in menulist"
        :key="item.id"
        :index="item.index"
      >
        <template v-if="item.children" #title>
          <el-icon v-if="item.icon">
            <component :is="item.icon"> </component>
          </el-icon>
          <span>{{ item.name }}</span>
        </template>
        <span v-if="!item.children">
          <el-icon v-if="item.icon">
            <component :is="item.icon"> </component>
          </el-icon>
          <span>{{ item.name }}</span>
        </span>
        <el-menu-item
          v-for="subItem in item.children"
          :key="subItem.id"
          :index="subItem.index"
        >
          {{ subItem.name }}
        </el-menu-item>
      </component>
    </el-menu>
  </div>
</template>

<script setup lang="ts">
import router from '@/router'
import { House, Notebook, User } from '@element-plus/icons-vue'
import { ElSubMenu, ElMenuItem } from 'element-plus'
import type { Component } from 'vue'
import { ref } from 'vue'

const defaultActive = ref<string>(router.currentRoute.value.path)

interface MenuItem {
  name: string
  id: number
  index: string
  icon?: Component
  children?: MenuItem[]
}
//菜单假数据
const menulist: MenuItem[] = [
  {
    id: 1,
    name: '首页',
    index: '/home/dashboard',
    icon: House
  },
  {
    id: 2,
    name: '用户管理',
    index: '/home/user',
    icon: User
  },
  {
    id: 3,
    name: '文章管理',
    index: '/home/article',
    icon: Notebook,
    children: [
      {
        id: 301,
        name: '文章',
        index: '/home/article'
      },
      {
        id: 302,
        name: '分类',
        index: '/home/category'
      },
      {
        id: 303,
        name: '标签',
        index: '/home/label'
      }
    ]
  }
]
</script>

<style lang="scss" scoped></style>
