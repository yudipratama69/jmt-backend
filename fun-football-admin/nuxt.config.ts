// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  app: {
    head: {
      title: 'JMT Sport',
      link: [
        { rel: 'icon', type: 'image/png', href: '/logo-jmt.png' }
      ]
    }
  },
  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxt/icon',
    '@vite-pwa/nuxt'
  ],
  pwa: {
    strategies: 'generateSW',
    registerType: 'autoUpdate',
    manifest: {
      name: 'JMT Sport',
      short_name: 'JMT Sport',
      description: 'Aplikasi Komunitas Futsal & Mini Soccer JMT Sport',
      theme_color: '#ea580c',
      background_color: '#ffffff',
      display: 'standalone',
      start_url: '/',
      icons: [
        {
          src: '/logo-jmt.png',
          sizes: '192x192',
          type: 'image/png',
          purpose: 'any maskable'
        },
        {
          src: '/logo-jmt.png',
          sizes: '512x512',
          type: 'image/png',
          purpose: 'any maskable'
        }
      ]
    },
    workbox: {
      navigateFallback: undefined
    }
  }
})