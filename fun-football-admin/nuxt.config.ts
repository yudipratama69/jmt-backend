// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  devServer: {
    host: '0.0.0.0',
    port: 3000
  },

  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'https://jmt-sport.transmedic.co.id'
    }
  },

  app: {
    head: {
      title: 'JMT Sport - Komunitas Fun Football',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no, viewport-fit=cover' },
        { name: 'theme-color', content: '#ea580c' },
        { name: 'mobile-web-app-capable', content: 'yes' },
        { name: 'apple-mobile-web-app-capable', content: 'yes' },
        { name: 'apple-mobile-web-app-status-bar-style', content: 'black-translucent' },
        { name: 'apple-mobile-web-app-title', content: 'JMT Sport' },
        { name: 'description', content: 'Aplikasi Komunitas Futsal & Mini Soccer JMT Sport' }
      ],
      link: [
        { rel: 'icon', type: 'image/png', href: '/logo-jmt.png' },
        { rel: 'apple-touch-icon', href: '/logo-jmt.png' },
        { rel: 'manifest', href: '/manifest.webmanifest' }
      ],
      script: [
        {
          innerHTML: `
            (function() {
              window.__deferredPwaPrompt = null;
              window.addEventListener('beforeinstallprompt', function(e) {
                e.preventDefault();
                window.__deferredPwaPrompt = e;
                window.dispatchEvent(new CustomEvent('pwa-prompt-ready', { detail: e }));
                console.log('⚡ [PWA-Boot] beforeinstallprompt captured early!');
              });
            })();
          `,
          type: 'text/javascript'
        }
      ]
    }
  },

  css: [
    '~/assets/css/theme.css'
  ],

  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxt/icon',
    '@vite-pwa/nuxt'
  ],

  nitro: {
    prerender: {
      routes: [
        '/',
        '/player',
        '/player/login',
        '/player/register',
        '/player/jadwalku',
        '/player/keuangan',
        '/player/pemain',
        '/player/profile',
        '/admin',
        '/admin-login',
        '/jadwal',
        '/verifikasi',
        '/verifikasi-deposit',
        '/admin-keuangan',
        '/profil-admin',
        '/master-notifikasi'
      ]
    }
  },

  pwa: {
    strategies: 'generateSW',
    registerType: 'autoUpdate',
    devOptions: {
      enabled: false
    },
    client: {
      installPrompt: true
    },
    manifest: {
      id: '/player',
      name: 'JMT Sport',
      short_name: 'JMT Sport',
      description: 'Aplikasi Komunitas Fun Football & Mini Soccer JMT Sport',
      theme_color: '#ea580c',
      background_color: '#0f172a',
      display: 'standalone',
      display_override: ['standalone', 'window-controls-overlay', 'minimal-ui'],
      orientation: 'portrait',
      start_url: '/player',
      scope: '/',
      categories: ['sports', 'utilities'],
      lang: 'id',
      dir: 'ltr',
      prefer_related_applications: false,
      icons: [
        {
          src: '/logo-jmt.png',
          sizes: '192x192 512x512',
          type: 'image/png',
          purpose: 'any maskable'
        },
        {
          src: '/pwa-192x192.png',
          sizes: '192x192',
          type: 'image/png',
          purpose: 'any'
        },
        {
          src: '/pwa-512x512.png',
          sizes: '512x512',
          type: 'image/png',
          purpose: 'any'
        },
        {
          src: '/pwa-maskable-512x512.png',
          sizes: '512x512',
          type: 'image/png',
          purpose: 'maskable'
        }
      ]
    },
    workbox: {
      navigateFallback: undefined,
      globPatterns: ['**/*.{js,css,html,png,svg,ico}']
    }
  }
})