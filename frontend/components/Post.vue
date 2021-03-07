<template>
  <div class="py-2 outline">
    <b-row>
      <b-avatar />
      <b-col>
        <div>@hogehoge</div>
        <div class="date">
          {{post.CreatedAt}}
        </div>
      </b-col>
    </b-row>
    <b-btn
      variant="link"
      style="color: var(--secondary); font-size: 14px; padding: 0px;"
      @click="open"
    >
      <h3 class="title">
        {{post.title}}
      </h3>
    </b-btn>

    <b-row align-v="center">
      <b-icon icon="suit-heart-fill" class="mx-1" />
      {{ good }}
      <b-icon icon="bookmark-fill" class="mx-1" />
      {{ bookmark }}
      <b-icon icon="tags-fill" class="mx-1" />
      <b-button class="mx-1" variant="outline-secondary" size="sm">
        {{post.language}}
      </b-button>
    </b-row>
  </div>
</template>

<script>
export default {
  props: {
    post: {
      type: Object,
      default: () => ({})
    },
    good: {
      type: Number,
      default: 0
    },
    bookmark: {
      type: Number,
      default: 0
    }
  },
  methods: {
    open () {
      this.$router.push(`/post/${this.post.post_id}`)
    }
  }
}
</script>

<style lang="scss" scoped>
.outline {
  border-bottom: 1px solid $lightgray;
}

.date {
  color: $lightgray;
  font-size: 14px;
}

.title {
  font-weight: bold;
}
</style>
