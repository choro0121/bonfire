<template>
  <div class="d-flex flex-column justify-content-center" style="width: 100%;">
    <div class="page-title mt-2 mb-3 text-center">
      アカウントを登録
    </div>

    <form-textfield
      v-model="username"
      class="my-2"
      label="ユーザー名"
      formtype="text"
      :need="true"
    />
    <form-textfield
      v-model="email"
      class="my-2"
      label="メールアドレス"
      formtype="email"
      :need="true"
    />
    <form-button
      class="my-2"
      label="アカウントを作成"
      @click="clicked"
    />
  </div>
</template>

<script>
import FormTextfield from '@/components/FormTextfield.vue'
import FormButton from '@/components/FormButton.vue'

export default {
  layout: 'auth',
  components: {
    FormTextfield,
    FormButton
  },
  data () {
    return {
      username: '',
      email: ''
    }
  },
  validate ({ params, redirect }) {
    if (params.userId === undefined) {
      redirect('/signup')
    }
    return true
  },
  methods: {
    clicked () {
      const forms = new FormData()
      forms.append('username', this.username)
      forms.append('email', this.email)

      this.$axios.$post(`/signup/${this.$route.params.provider}/${this.$route.params.userId}`, forms)
        .then((res) => {
          localStorage.setItem('token', res)
          if (this.$route.query.continue === undefined) {
            this.$router.push('/home')
          } else {
            this.$router.push(this.$route.query.continue)
          }
        })
    }
  }
}
</script>
