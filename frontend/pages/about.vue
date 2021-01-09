<template>
  <div style="height: 1280px;">
    <div class="top-view">
      <div class="header">
        <b-row align-v="center" align-h="end" style="height: 100%;">
          <b-button class="button-style" variant="link">
            ダウンロード
          </b-button>
          <b-button class="button-style" variant="link" to="/signin">
            ログイン
          </b-button>
          <b-button class="button-style" variant="outline-light" to="signup">
            登録
          </b-button>
        </b-row>
      </div>

      <div class="d-flex flex-column align-items-center justify-content-center" style="height: 100%;">
        <svg-loader :size="128" src="assets/svg/SnippetHub" />
        <h1 class="font-shadow">
          “探す”を“早く”
        </h1>
        <h4 class="font-shadow">
          スニペットは登録しない時代へ
        </h4>
      </div>
    </div>

    <scroll-fadein>
      <section-browser style="height: 1080px;" />
    </scroll-fadein>
    <scroll-fadein>
      <section-editor style="height: 1080px;" />
    </scroll-fadein>
    <scroll-fadein>
      <section-terminal style="height: 1080px;" />
    </scroll-fadein>
  </div>
</template>

<script>
import SvgLoader from '@/components/SvgLoader.vue'
import ScrollFadein from '@/components/ScrollFadein.vue'
import SectionBrowser from '@/components/SectionBrowser.vue'
import SectionEditor from '@/components/SectionEditor.vue'
import SectionTerminal from '@/components/SectionTerminal.vue'

export default {
  components: {
    SvgLoader,
    ScrollFadein,
    SectionBrowser,
    SectionEditor,
    SectionTerminal
  }
}
</script>

<style lang="scss" scoped>
.top-view {
  width: 100%;
  height: 866px;
  background: url('~@/static/macbook-photo.jpg') no-repeat center;
  background-size: cover;
  position: relative;
  overflow: hidden;
  z-index: 0;
}

.top-view:before {
  background: inherit;
  content: '';
  filter: blur(5px);
  position: absolute;
  top: -5px;
  right: -5px;
  bottom: -5px;
  left: -5px;
  z-index: -1;
}

.top-view > .header {
  position: absolute;
  width: 100%;
  height: 48px;
}

h1 {
  color: $white;
  font-family: 'Noto Serif JP';
  font-size: 64px;
}

h4 {
  color: $white;
  letter-spacing: 3px;
  font-weight: 100;
}

.row {
  margin-left: 0px;
  margin-right: 0px;
}

.font-shadow {
  text-shadow: $navyblue 1px 0 10px;
}

.button-style {
  color: $white;
  font-size: 14px;
}
</style>
