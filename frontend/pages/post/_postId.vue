<template>
  <div>
    <h1>
      {{ title }}
    </h1>

    <div>
      {{ desc }}
    </div>

    <code-editor v-model="code" :lang="lang" readonly />

  </div>
</template>

<script>
import CodeEditor from '@/components/CodeEditor.vue'

export default {
  layout: 'main',
  components: {
    CodeEditor
  },
  data () {
    return {
      title: '',
      desc: '',
      lang: '',
      code: ''
    }
  },
  created () {
    this.$axios.$get(`/api/v1/posts/${this.$route.params.postId}`)
      .then((res) => {
        console.log(res)
        this.title = res.title
        this.desc = res.description
        this.lang = res.language
        this.code = res.code
      })
  }
}
</script>
