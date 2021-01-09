<template>
  <component :is="svg" :style="{width: `${size}px`}" />
</template>

<script>
export default {
  props: {
    src: {
      type: String,
      default: ''
    },
    size: {
      type: Number,
      default: 128
    }
  },
  computed: {
    svg() {
      return () => import(`~/${this.src}.svg`)
    }
  }
}
</script>
