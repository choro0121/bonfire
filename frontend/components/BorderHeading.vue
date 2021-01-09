<template>
  <b-row class="mb-2" align-h="center" :style="{'border-bottom': `${width}px solid var(--orange)`}">
    <slot />
  </b-row>
</template>

<script>
export default {
  props: {
    width: {
      type: Number,
      default: 6
    }
  }
}
</script>

<style lang="scss" scoped>
.row {
  width: 200px;
}

h2 {
  font-weight: 700;
  letter-spacing: 3px;
}
</style>
