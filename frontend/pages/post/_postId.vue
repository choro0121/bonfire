<template>
  <div>
    <b-button @click="clicked" />
  </div>
</template>

<script>
export default {
  methods: {
    clicked () {
      console.log(this.$axios.$get(`/api/v1/posts/${this.$route.params.postId}`))
    }
  }
}
</script>
