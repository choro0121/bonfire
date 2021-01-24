<template>
  <div>
    <b-form-input
      placeholder="投稿のタイトル"
      required
    />

    <b-form-textarea
      placeholder="説明"
    />

    <language-select v-model="lang" />

    <code-editor v-model="code" :lang="lang" />

    <b-btn>
      投稿
    </b-btn>
    <b-btn variant="outline-secondary">
      下書き保存
    </b-btn>

    <div>{{ lang }}</div>

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
      code: '',
      lang: ''
    }
  }
}
</script>
