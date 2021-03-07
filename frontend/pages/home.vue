<template>
  <div>
    <post
      v-for="(post, idx) in posts"
      :key="idx"
      :post="post"
    />
  </div>
</template>

<script>
import Post from '@/components/Post.vue'

export default {
  layout: 'main',
  components: {
    Post
  },
  data () {
    return {
      posts: {}
    }
  },
  created () {
    this.$axios.$get('/api/v1/posts')
      .then((res) => {
        this.posts = res.reverse()
      })
  }
}
</script>
