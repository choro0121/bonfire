<template>
  <header>
    <b-navbar class="pr-2" type="dark" variant="dark" style="padding: 0px; height: 100%">
      <b-navbar-brand to="/home">
        <app-logo />
      </b-navbar-brand>

      <b-navbar-nav class="ml-auto">
        <b-row align-v="center">
          <b-button variant="primary" size="sm">
            <b-icon icon="pencil-square" aria-hidden="true" />
            ポスト
          </b-button>
          <b-icon icon="bell-fill" variant="light" font-scale="1.2" />
          <b-avatar />
        </b-row>
      </b-navbar-nav>
    </b-navbar>
  </header>
</template>

<script>
import AppLogo from '@/components/AppLogo.vue'

export default {
  components: {
    AppLogo
  }
}
</script>
