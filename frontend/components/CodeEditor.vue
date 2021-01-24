<template>
  <client-only>
    <prism-editor
      :value="value"
      class="my-editor"
      :highlight="highlighter"
      :readonly="readonly"
      style="height: 300px"
      line-numbers
      @input="$emit('input', $event)"
    />
  </client-only>
</template>

<script>
import { PrismEditor } from 'vue-prism-editor'
import 'vue-prism-editor/dist/prismeditor.min.css'

export default {
  components: {
    PrismEditor
  },
  props: {
    value: {
      type: String,
      default: ''
    },
    lang: {
      type: String,
      default: ''
    },
    readonly: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    highlighter (code) {
      if (process.client) {
        if (this.$Prism.languages[this.lang] === undefined) {
          return code
        } else {
          return this.$Prism.highlight(code, this.$Prism.languages[this.lang])
        }
      }
    }
  }
}
</script>

<style>
.my-editor {
  background-color: #F6F6F6;

  font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
  font-size: 14px;
  padding: 6px 12px;
  border-radius: 5px;
  border: 1px solid #ced4da;
}

.prism-editor__textarea, .prism-editor__editor {
  caret-color: #495057;
  color: #495057;
}

.prism-editor__textarea:focus {
  outline: none;
}
</style>
