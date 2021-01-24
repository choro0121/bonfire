<template>
  <div>

    <h1>
      {{ title }}
    </h1>

    <div>
      {{ body }}
    </div>

    <code-editor v-model="code" :lang="lang" readonly />

    <b-button @click="clicked" />

  </div>
</template>

<script>
import CodeEditor from '@/components/CodeEditor.vue'

export default {
  layout: 'main',
  components: {
    CodeEditor
  },
  props: {
    title: {
      type: String,
      default: 'hogefuga'
    },
    body: {
      type: String,
      default: 'piyoooooooooooooo'
    },
    lang: {
      type: String,
      default: 'js'
    },
    code: {
      type: String,
      default: 'console.log("aaa")'
    }
  },
  methods: {
    clicked () {
      console.log(this.$axios.$get(`/api/v1/posts/${this.$route.params.postId}`))
    }
  }
}
</script>
