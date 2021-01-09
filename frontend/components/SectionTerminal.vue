<template>
  <section class="py-2">
    <b-container class="d-flex flex-column align-items-center">
      <border-heading>
        <h2>CLI.</h2>
      </border-heading>
      <terminal style="max-width: 720px;" />
    </b-container>
  </section>
</template>

<script>
import Terminal from '@/components/Terminal.vue'
import BorderHeading from '@/components/BorderHeading.vue'

export default {
  components: {
    Terminal,
    BorderHeading
  }
}
</script>
