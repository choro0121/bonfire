<template>
  <section class="py-2">
    <b-container class="d-flex flex-column align-items-center">
      <border-heading>
        <h2>Editor.</h2>
      </border-heading>
    </b-container>
  </section>
</template>

<script>
import BorderHeading from '@/components/BorderHeading.vue'

export default {
  components: {
    BorderHeading
  }
}
</script>

<style lang="scss" scoped>
section {
  background-color: $navyblue;
}

h2 {
  color: $white;
}
</style>
