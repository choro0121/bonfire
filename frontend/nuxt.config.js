export default {
  // Target (https://go.nuxtjs.dev/config-target)
  target: 'static',

  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: {
    title: 'snippethub',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: [
    '~assets/scss/app.scss'
  ],

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [
    { src: '~/plugins/prism.js', mode: 'client' }
  ],

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: true,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    '@nuxtjs/eslint-module'
  ],

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
    // https://go.nuxtjs.dev/bootstrap
    ['bootstrap-vue/nuxt', { css: false }],
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    '@nuxtjs/proxy',
    '@nuxtjs/style-resources',

    'nuxt-svg-loader',
    'nuxt-webfontloader'
  ],

  bootstrapVue: {
    icons: true
  },
 
  // Axios module configuration (https://go.nuxtjs.dev/config-axios)
  axios: {
    prefix: '/api',
    proxy: true
  },

  // proxy
  proxy: {
    '/api': {
      target: process.env.HOST_URL || 'undefined',
      pathRewrite: {
        '^/api': '/'
      }
    }
  },

  // style resource
  styleResources: {
    scss: [
      '~assets/scss/variable.scss'
    ]
  },

  // Nuxt-WebFontLoader module
  webfontloader: {
    google: {
      families: ['Noto+Sans+JP:100,300,400,700', 'Noto+Serif+JP:700']
    }
  },

  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {
  }
}
