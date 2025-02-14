<script setup lang="ts">
import {reactive, ref} from "vue";
import Animation_grid from "@/components/animation/animation_grid.vue";
import {ApiAuth, typeAuth} from "@/api/api_auth.ts";
import {ElMessage, type FormRules} from "element-plus";
import {useStore} from "@/stores/stores.ts";
const store = useStore()
const form = reactive(typeAuth.loginRequest())
const formRef = ref()
const rules = reactive<FormRules>({
  name: [
    {required: true, message: '用户名不能为空', trigger: 'blur'},
  ],
  password: [
    {required: true, message: '密码不能为空', trigger: 'blur'},
  ],
})

async function login() {
  if (!await formRef.value.validate()) return;

  let response = await ApiAuth.Login(form)
  if (response.code) {
    ElMessage.error(response.message)
    return
  }
  store.setToken(response.data.token);
}



</script>

<template>
  <div class="m_login">
    <Animation_grid style="width: 100%; height: 140px; background-color: #d9d9d9;"/>
    <div class="form">
      <el-form :model="form" ref="formRef" :rules="rules">
        <el-form-item prop="name"><el-input v-model="form.name" placeholder="用户名"><template #prefix><i class="iconfont icon-yonghuming"/></template></el-input></el-form-item>
        <el-form-item prop="password" class="item"><el-input v-model="form.password" type="password" placeholder="密码"><template #prefix><i class="iconfont icon-mima"/></template></el-input></el-form-item>
        <el-form-item class="item"><el-checkbox>记住密码</el-checkbox></el-form-item>
        <el-form-item ><el-button style="width: 100%;" type="primary" @click="login">登录</el-button></el-form-item>
      </el-form>
      <div class="open_login">
        <div class="line-space-between">第三方登陆</div>
        <div class="icon">
          <i class="iconfont icon-QQ"/>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.m_login {
  width: 500px;
  height: 406px;
  background-color: white;
  border-radius: 5px;
  overflow: hidden;
  box-shadow: 0 0 5px 5px rgb(0, 0, 0, 0.1);

  .form {
    padding: 20px 80px;

    .item {
      margin-bottom: 10px;
    }

    .open_login {
      display: flex;
      flex-direction: column;
      align-items: center;


      @include line-space-between();

      .icon {
        margin-top: 5px;

        i {
          font-size: 36px;
          cursor: pointer;
        }
      }

    }

  }
}


</style>