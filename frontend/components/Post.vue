<template>
  <div class="py-2 outline">
    <b-row>
      <b-avatar />
      <b-col>
        <div>@hogehoge</div>
        <div class="date">
          2020-12-20 03:00
        </div>
      </b-col>
    </b-row>
    <h3 class="title">
      JavaScriptのforEachの書き方
    </h3>
    <b-row align-v="center">
      <b-icon icon="suit-heart-fill" class="mx-1" />
      {{ good }}
      <b-icon icon="bookmark-fill" class="mx-1" />
      {{ bookmark }}
      <b-icon icon="tags-fill" class="mx-1" />
      <b-button class="mx-1" variant="outline-secondary" size="sm">
        JavaScript
      </b-button>
      <b-button class="mx-1" variant="outline-secondary" size="sm">
        jQuery
      </b-button>
    </b-row>
  </div>
</template>

<script>
export default {
  props: {
    good: {
      type: Number,
      default: 0
    },
    bookmark: {
      type: Number,
      default: 0
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
