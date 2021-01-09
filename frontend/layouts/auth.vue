<template>
  <b-container class="d-flex flex-column align-items-center justify-content-center">
    <svg-loader :size="48" src="assets/svg/SnippetHub" />
    <Nuxt />
  </b-container>
</template>

<script>
import SvgLoader from '@/components/SvgLoader.vue'

export default {
  components: {
    SvgLoader
  }
}
</script>

<style>
.page-title {
  font-size: 28px;
  font-weight: 700;
}
</style>

<style scoped>
.container {
  max-width: 360px;
  height: 100vh;
}
</style>
