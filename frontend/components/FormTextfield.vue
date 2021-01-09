<template>
  <div>
    <div class="label mb-1">
      {{ label }}<span v-if="need" style="color: #E74C3C;">*</span>
    </div>
    <b-form-input
      v-model="data"
      :state="state"
      :type="formtype"
      size="sm"
    />
    <b-form-invalid-feedback>
      {{ feedback }}
    </b-form-invalid-feedback>
  </div>
</template>

<script>
export default {
  props: {
    label: {
      type: String,
      default: ''
    },
    formtype: {
      type: String,
      default: 'text'
    },
    need: {
      type: Boolean,
      default: false
    },
    validation: {
      type: Boolean,
      default: true
    },
    input: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      state: null
    }
  },
  computed: {
    data: {
      get () {
        return this.input
      },
      set (value) {
        if (this.validation !== true) {
          this.state = null
        } else if (value === '') {
          this.state = null
        } else if (this.formtype === 'email') {
          this.state = this.validate(value, /^[A-Za-z0-9]{1}[A-Za-z0-9_.-]*@{1}[A-Za-z0-9_.-]{1,}\.[A-Za-z0-9]{1,}$/)
        } else if (this.formtype === 'password') {
          this.state = this.validate(value, /^(?=.*?[a-z])(?=.*?\d)[a-z\d]{8,100}$/i)
        } else {
          this.state = null
        }

        this.$emit('input', value)
        this.$emit('state', this.state)
      }
    },
    feedback () {
      if (this.formtype === 'email') {
        return '無効なメール形式です'
      } else if (this.formtype === 'password') {
        return 'パスワードは8文字以上の英数字にしてください'
      } else {
        return ''
      }
    }
  },
  methods: {
    validate (text, regex) {
      if (text) {
        if (!text.match(regex)) {
          return false
        } else {
          return true
        }
      } else {
        return false
      }
    }
  }
}
</script>

<style scoped>
.label {
  font-size: 14px;
}
</style>
