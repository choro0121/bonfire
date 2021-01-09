<template>
  <div>
    <b-button class="label mt-2" variant="primary" block>
      {{ label }}
    </b-button>
  </div>
</template>

<script>
export default {
  props: {
    label: {
      type: String,
      default: ''
    }
  }
}
</script>

<style scoped>
.label {
  font-size: 14px;
}
</style>
