<script setup lang="ts">
import {useStore} from "@/stores";
import {reactive} from "vue";
import {Message} from "@/utils/element-ui/message.ts";
import {ApiUser} from "@/api/api_user.ts";
import {baseResponse} from "@/api/response.ts";
import {type m_type_el} from "@/utils/element-ui/type.ts";
import {ElInput} from "element-plus";

const store = useStore()

interface typeForm {
  input: m_type_el.form.Input[]
  select: m_type_el.form.Select[]
  issue: m_type_el.form.Input[]
  switch: m_type_el.form.Switch[]
}

const form = reactive<typeForm>({
  input: [
    {label: "撤回消息提示", max: 32, value: store.userInfo.user_config_model.recall_message, type: "text",},
  ],
  select: [
    {
      label: "加我为好友的验证方式", value: store.userInfo.user_config_model.valid, option: [
        {value: 0, label: '禁止加我为好友',},
        {value: 1, label: '允许任何人添加',},
        {value: 2, label: '需要验证',},
        {value: 3, label: '需要正确回答问题',},
      ],
    },
  ],
  issue: [
    {label: "问题", value: store.userInfo.user_config_model.valid_info.issue,},
    {label: "答案", value: store.userInfo.user_config_model.valid_info.answer,},
  ],
  switch: [
    {label: "保存密码", value: store.userInfo.user_config_model.save_password,},
    {label: "安全链接", value: store.userInfo.user_config_model.secure_link,},
    {label: "好友上线提醒", value: store.userInfo.user_config_model.friend_online,},
    {label: "关闭所有声音", value: store.userInfo.user_config_model.sound,},
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
      user_config: {
        recall_message: form.input[0].value,
        valid: form.select[0].value as number,
        valid_info: {issue: form.issue[0].value, answer: form.issue[1].value,},
        save_password: form.switch[0].value,
        secure_link: form.switch[1].value,
        friend_online: form.switch[2].value,
        sound: form.switch[3].value,
      },
    })
    if (!baseResponse(response)) return;
    await store.loadUserInfo(true);
    Message.success("更新成功");
  }
}

const m = new m_api()
</script>

<template>
  <div class="m_view_base_info">
    <el-form :model="form">
      <el-form-item :label="item.label" v-for="item in form.input"><el-input v-model="item.value" :type="item.type" :maxlength="item.max" autosize show-word-limit :disabled="item.disabled"/></el-form-item>
      <el-form-item :label="item.label" v-for="item in form.select">
        <el-select v-model="item.value" style="width: 240px">
          <el-option v-for="opt in item.option" :key="opt.value" :label="opt.label" :value="opt.value"/>
        </el-select>
      </el-form-item>
      <el-form-item v-if="form.select[0].value === 3" :label="item.label" v-for="item in form.issue">
        <el-input v-model="item.value" :type="item.type" :maxlength="item.max" autosize show-word-limit :disabled="item.disabled"/>
      </el-form-item>
      <el-form-item :label="item.label" v-for="item in form.switch"><el-switch v-model="item.value"/></el-form-item>
    </el-form>

    <el-button plain @click="m.Reset">恢复默认</el-button>
    <el-button type="primary" plain @click="m.Update">更新设置</el-button>
  </div>
</template>