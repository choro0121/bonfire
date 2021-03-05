<template>
  <div>
    <b-input-group>
      <datalist id="language-list">
        <option
          v-for="(lang, idx) in languages"
          :key="idx"
        >
          {{ lang }}
        </option>
      </datalist>

      <b-form-input
        list="language-list"
        placeholder="言語"
        :value="value"
        @input="$emit('input', $event)"
      />

      <b-input-group-append>
        <b-button
          @click="$emit('input', '')"
        >
          <b-icon icon="x" />
        </b-button>
      </b-input-group-append>
    </b-input-group>
  </div>
</template>

<script>
const components = require('prismjs/components')

export default {
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      selected: ''
    }
  },
  computed: {
    languages () {
      // const langs = Array.from(Object.values(components.languages), (val) => {
      //   return val.title
      // })
      const langs = Array.from(Object.keys(components.languages), (key) => {
        return key
      })

      return langs.sort()
    }
  }
}
</script>

<style>

</style>
