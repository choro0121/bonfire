<template>
  <div>
    <my-header />

    <b-container>
      <Nuxt />
    </b-container>

    <my-footer />
  </div>
</template>

<script>
import MyHeader from '@/components/MyHeader.vue'
import MyFooter from '@/components/MyFooter.vue'

export default {
  components: {
    MyHeader,
    MyFooter
  }
}
</script>

<style lang="scss" scoped>
header {
  position: fixed;
  width: 100%;
  height: 64px;
  z-index: 1;
}

.container {
  max-width: 720px;
  height: 100vh;
  background-color: white;
  padding-top: 64px;
}
</style>
