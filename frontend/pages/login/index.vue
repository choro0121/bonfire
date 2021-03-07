<template>
  <div class="d-flex flex-column justify-content-center" style="width: 100%;">
    <div class="page-title mt-2 mb-3 text-center">
      ログイン
    </div>

    <form-textfield
      v-model="username"
      class="my-2"
      label="ユーザー名"
      formtype="text"
    />
    <form-textfield
      v-model="password"
      class="my-2"
      label="パスワード"
      formtype="password"
    />
    <form-button
      class="my-2"
      label="ログイン"
      @click="login"
    />

    <hr class="my-3" width="100%">

    <svg-button
      class="my-2"
      label="Googleアカウントでログイン"
      src="assets/svg/Google"
      color="#FFFFFF"
      font="var(--secondary)"
      border="#E4E4E4"
      @click="clicked('google')"
    />
    <svg-button
      class="my-2"
      label="GitHubアカウントでログイン"
      src="assets/svg/GitHub"
      color="#444444"
      font="#FFFFFF"
      border="#363636"
      @click="clicked('github')"
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
  },
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    login () {
      const forms = new FormData()
      forms.append('username', this.username)
      forms.append('password', this.password)

      this.$axios.$post('/login', forms)
        .then((res) => {
          localStorage.setItem('token', res)
          if (this.$route.query.continue === undefined) {
            this.$router.push('/home')
          } else {
            this.$router.push(this.$route.query.continue)
          }
        })
    },
    clicked (provider) {
      // console.log(process.env.HOST_URL)
      location.href = `/api/auth/${provider}`
      // this.$axios.$get(`/auth/${provider}`)
      //   .then((res) => {
      //     console.log(res)
      //   })
    }
  }
}
</script>
