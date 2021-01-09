<template>
  <footer class="my-2">
    <b-row align-h="center">
      <a href="/">
        <svg-loader id="snippethub-svg" :size="32" src="assets/svg/SnippetHub" />
      </a>
      <a href="https://twitter.com/snippethub">
        <svg-loader id="twitter-svg" :size="32" src="assets/svg/Twitter" />
      </a>
      <b-button to="/terms" variant="link">
        利用規約
      </b-button>
      <b-button to="/privacy" variant="link">
        プライバシーポリシー
      </b-button>
      <b-button to="/api" variant="link">
        API
      </b-button>
      <b-button to="/feedback" variant="link">
        ご意見
      </b-button>
      <b-button to="/support" variant="link">
        お問い合わせ
      </b-button>
    </b-row>
  </footer>
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
a {
  font-size: 12px;
  color: $navyblue;
}
</style>

<style lang="scss">
#twitter-svg > path {
  fill: $navyblue !important;
}
#snippethub-svg > path {
  fill: $navyblue !important;
}
</style>
