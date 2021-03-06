<template>
  <div>
    <b-form-input
      v-model="title"
      placeholder="投稿のタイトル"
      required
    />

    <b-form-textarea
      v-model="desc"
      placeholder="説明"
    />

    <language-select v-model="lang" />

    <code-editor v-model="code" :lang="lang" />

    <b-btn @click="clicked">
      投稿
    </b-btn>
    <b-btn variant="outline-secondary">
      下書き保存
    </b-btn>

  </div>
</template>

<script>
import CodeEditor from '@/components/CodeEditor.vue'
import LanguageSelect from '@/components/LanguageSelect.vue'

export default {
  layout: 'main',
  components: {
    CodeEditor,
    LanguageSelect
  },
  data () {
    return {
      title: '',
      desc: '',
      lang: '',
      code: ''
    }
  },
  methods: {
    clicked () {
      localStorage.setItem('token', 'super')
      this.$axios.$post('/api/v1/posts', {
        title: this.title,
        description: this.desc,
        language: this.lang,
        code: this.code
      })
        .then((res) => {
          this.$router.push(`/post/${res.post_id}`)
        })
    }
  }
}
</script>
