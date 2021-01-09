<template>
  <div class="d-flex flex-column justify-content-center" style="width: 100%;">
    <div class="page-title mt-2 mb-3 text-center">
      アカウントを登録
    </div>

    <form-textfield
      class="my-2"
      label="ユーザー名"
      formtype="text"
      :need="true"
    />
    <form-textfield
      class="my-2"
      label="メールアドレス"
      formtype="email"
      :need="true"
    />
    <form-textfield
      class="my-2"
      label="パスワード"
      formtype="password"
      :need="true"
    />
    <form-button
      class="my-2"
      label="アカウントを作成"
    />

    <hr class="my-3" width="100%">

    <svg-button
      class="my-2"
      label="Googleアカウントで登録"
      src="assets/svg/Google"
      color="#FFFFFF"
      font="var(--secondary)"
      border="#E4E4E4"
    />
    <svg-button
      class="my-2"
      label="GitHubアカウントで登録"
      src="assets/svg/GitHub"
      color="#444444"
      font="#FFFFFF"
      border="#363636"
    />

    <b-button
      class="my-2"
      variant="link"
      style="color: var(--secondary); font-size: 14px; text-decoration: underline;"
      to="/signin"
    >
      アカウントをお持ちの方
    </b-button>
  </div>
</template>

<script>
import FormTextfield from '@/components/FormTextfield.vue'
import FormButton from '@/components/FormButton.vue'
import SvgButton from '@/components/SvgButton.vue'

export default {
  layout: 'auth',
  components: {
    FormTextfield,
    FormButton,
    SvgButton
  }
}
</script>
