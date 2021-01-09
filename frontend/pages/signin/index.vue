<template>
  <div class="d-flex flex-column justify-content-center" style="width: 100%;">
    <div class="page-title mt-2 mb-3 text-center">
      ログイン
    </div>

    <form-textfield
      class="my-2"
      label="ユーザー名 または メールアドレス"
      formtype="text"
    />
    <form-textfield
      class="my-2"
      label="パスワード"
      formtype="password"
    />
    <form-button
      class="my-2"
      label="ログイン"
    />

    <hr class="my-3" width="100%">

    <svg-button
      class="my-2"
      label="Googleアカウントでログイン"
      src="assets/svg/Google"
      color="#FFFFFF"
      font="var(--secondary)"
      border="#E4E4E4"
    />
    <svg-button
      class="my-2"
      label="GitHubアカウントでログイン"
      src="assets/svg/GitHub"
      color="#444444"
      font="#FFFFFF"
      border="#363636"
    />

    <b-button
      class="my-2"
      variant="link"
      style="color: var(--secondary); font-size: 14px; text-decoration: underline;"
      to="/signup"
    >
      アカウントを作成する
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
