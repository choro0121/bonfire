<template>
  <b-row align-v="center" style="font-size: 0;">
    <svg-loader :size="48" src="assets/svg/SnippetHub" />
    <span class="logo-font" style="color: var(--orange);">Snippet</span>
    <span class="logo-font" style="color: var(--yellow);">Hub</span>
  </b-row>
</template>

<script>
import SvgLoader from '@/components/SvgLoader.vue'

export default {
  components: {
    SvgLoader
  }
}
</script>

<style lang="scss" scoped>
.logo-font {
  font-size: 14px;
}
</style>
