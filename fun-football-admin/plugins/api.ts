export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig()

  const getBaseUrl = () => {
    if (typeof window !== 'undefined') {
      const host = window.location.hostname
      const port = window.location.port
      const proto = window.location.protocol

      // Jika diakses via Nuxt Dev Server di port 3000, arahkan ke backend Go di port 8080 pada host yang sama
      if (port === '3000') {
        return `${proto}//${host}:8080`
      }
      // Jika diakses via Docker / Ngrok / Reverse Proxy langsung
      return `${proto}//${window.location.host}`
    }
    return config.public.apiBase || 'http://localhost:8080'
  }

  const api = $fetch.create({
    baseURL: getBaseUrl(),
    onRequest({ options }) {
      const headers = new Headers(options.headers || {})
      headers.set('ngrok-skip-browser-warning', 'true')
      headers.set('Accept', 'application/json')
      options.headers = headers
    }
  })

  return {
    provide: {
      api
    }
  }
})
