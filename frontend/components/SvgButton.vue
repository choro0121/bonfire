<template>
  <div>
    <b-button class="label" block style="height: 48px;" :style="{background: color, color: font, borderColor: border}">
      <b-row>
        <div style="width: 48px;">
          <svg-loader class="mx-3" :size="20" :src="src" />
        </div>
        {{ label }}
      </b-row>
    </b-button>
  </div>
</template>

<script>
import SvgLoader from '@/components/SvgLoader.vue'

export default {
  components: {
    SvgLoader
  },
  props: {
    label: {
      type: String,
      default: ''
    },
    src: {
      type: String,
      default: ''
    },
    color: {
      type: String,
      default: '#000000'
    },
    font: {
      type: String,
      default: '#FFFFFF'
    },
    border: {
      type: String,
      default: '#E4E4E4'
    }
  }
}
</script>

<style scoped>
.label {
  font-size: 14px;
}
</style>
