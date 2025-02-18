<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import {useStore} from "@/stores";
import {Message} from "@/utils/element-ui/message.ts";
import {ApiUser} from "@/api/api_user.ts";
import {baseResponse} from "@/api/response.ts";
import {m_type_el} from "@/utils/element-ui/type.ts";
import {ElInput} from "element-plus";
import MElAvatarCropper from "@/utils/element-ui/m-el-avatar-cropper.vue";


const store = useStore()

interface typeForm {
  avatar: m_type_el.form.Avatar[]
  input: m_type_el.form.Input[]
}

const form = reactive<typeForm>({
  avatar: [{label: "头像", src: store.userInfo.avatar},],
  input: [
    {label: "昵称", max: 36, value: store.userInfo.username, type: "text",},
    {label: "签名", max: 99, value: store.userInfo.sign, type: "textarea",},
  ],
});
const source = JSON.parse(JSON.stringify(form));

class m_api {
  Reset = () => {
    Object.assign(form, JSON.parse(JSON.stringify(source)));
    Message.success("恢复默认");
  }
  Update = async () => {
    if (JSON.stringify(form) === JSON.stringify(source)) {
      Message.success("配置相同");
      return;
    }
    let response = await ApiUser.Update({
      user_info: {
        avatar: form.avatar[0].src,
        username: form.input[0].value,
        sign: form.input[1].value,
      },
    })
    if (!baseResponse(response)) return;
    await store.loadUserInfo(true);
    Message.success("更新成功");
  }
}

const m = new m_api()

const clipperRef = ref()
const clipperData = ref<m_type_el.props.AvatarCropper>({
  type: 'browserLogo', // 该参数可根据实际要求修改类型
  allowTypeList: ['png', 'jpg'], // 允许上传的图片格式
  limitSize: 1, // 限制的大小
  fixedNumber: [1, 1],  // 截图比例，可根据实际情况进行修改
  previewWidth: 100, // 预览宽度
});
const showCropper = () => clipperRef.value.uploadFile()
const onConfirm = (value: any) => form.avatar[0].src = value as string;
</script>

<template>
  <div class="m_view_info">
    <m-el-avatar-cropper
        ref="clipperRef"
        :type="clipperData.type"
        :allow-type-list="clipperData.allowTypeList"
        :limit-size="clipperData.limitSize"
        :fixed-number="clipperData.fixedNumber"
        :preview-width="clipperData.previewWidth"
        @confirm="onConfirm"
    />
    <el-form :model="form">
      <el-form-item :label="item.label" v-for="item in form.avatar"><el-avatar :src="item.src" @click="showCropper"/></el-form-item>
      <el-form-item :label="item.label" v-for="item in form.input"><el-input v-model="item.value" :type="item.type" :maxlength="item.max" autosize show-word-limit :disabled="item.disabled"/></el-form-item>
    </el-form>

    <el-button plain @click="m.Reset">恢复默认</el-button>
    <el-button type="primary" plain @click="m.Update">更新设置</el-button>
  </div>
</template>

<style>
</style>