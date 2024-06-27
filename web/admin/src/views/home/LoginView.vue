<template>
  <div id="login">
    <div class="login-box">
      <div class="logo">
        <el-image :src="loginImg" alt="logo" fit="cover" :lazy="true" />
      </div>
      <el-form :model="form" ref="formRef" :rules="loginRules">
        <el-form-item prop="username">
          <el-input
            class="input"
            v-model="form.username"
            placeholder="用户名"
            :prefix-icon="User"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            class="input"
            v-model="form.password"
            placeholder="密码"
            :prefix-icon="Lock"
          />
        </el-form-item>
        <div class="btns">
          <el-button class="btn" type="primary" @click="login">登录</el-button>
          <el-button class="btn" @click="reset">重置</el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import loginImg from '@/assets/img/login.jpg'
import { User, Lock } from '@element-plus/icons-vue'
import type { FormInstance } from 'element-plus'
import { ref } from 'vue'
import { loginRules } from '@/rules/loginInfo'
import { loginApi } from '@/apis/login'
import { useUserInfoStore } from '@/stores/userInfo.store'
import router from '@/router'

const userInfoStore = useUserInfoStore()

const form = ref<{
  username: string
  password: string
}>({
  username: 'admin',
  password: 'password123'
})

const formRef = ref<FormInstance>()

const login = async () => {
  const res = await loginApi(form.value)
  console.log(res)
  userInfoStore.setAuth(res.data.token)
  ElMessage.success('登录成功')
  router.push('/home')
}

const reset = () => {
  formRef.value?.resetFields()
}
</script>

<style lang="scss" scoped>
#login {
  height: 100vh;
  width: 100%;
  background-image: linear-gradient(#c1dfc4, #deecdd 100%);
  display: flex;
  align-items: center;
  justify-content: center;

  .login-box {
    width: 22rem;
    height: 26rem;
    background-color: #fff;
    border-radius: 0.5rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;
    padding: 1rem;

    .logo {
      width: 8rem;
      height: 8rem;
      border-radius: 50%;
      box-shadow: 0 0 4px 2px #9fa1a5;
      margin: 1rem;
      overflow: hidden;
      display: flex;
    }

    .el-form {
      width: 90%;

      .input {
        height: 2rem;

        :deep(.el-input__wrapper) {
          border-radius: 0.5rem;
        }
      }

      .btns {
        display: flex;
        justify-content: center;
        align-items: center;

        .btn {
          width: 100%;
          height: 2rem;
          border-radius: 0.3rem;
          margin-bottom: 1rem;
        }
      }
    }
  }
}
</style>
