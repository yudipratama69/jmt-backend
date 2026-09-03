export function getApiBaseUrl(): string {
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
  const config = useRuntimeConfig()
  return config.public.apiBase || 'http://localhost:8080'
}

export function useApiFetch<T = any>(url: string | (() => string), options: any = {}) {
  const baseURL = getApiBaseUrl()

  return useFetch<T>(url, {
    baseURL,
    onRequest({ options: reqOptions }) {
      const headers = new Headers(reqOptions.headers || {})
      headers.set('ngrok-skip-browser-warning', 'true')
      headers.set('Accept', 'application/json')
      reqOptions.headers = headers
    },
    ...options
  })
}

export function useApiUrl(path?: string | null): string {
  if (!path) return ''
  const base = getApiBaseUrl().replace(/\/+$/, '')

  // Jika sudah merupakan URL lengkap
  if (path.startsWith('http://') || path.startsWith('https://')) {
    if (!path.includes('ngrok-skip-browser-warning')) {
      const separator = path.includes('?') ? '&' : '?'
      return `${path}${separator}ngrok-skip-browser-warning=true`
    }
    return path
  }

  const cleanPath = path.startsWith('/') ? path : `/${path}`
  const separator = cleanPath.includes('?') ? '&' : '?'
  return `${base}${cleanPath}${separator}ngrok-skip-browser-warning=true`
}

export function useApiClient() {
  const { $api } = useNuxtApp()
  return $api
}

export async function copyToClipboard(text: string): Promise<boolean> {
  if (typeof window === 'undefined') return false
  if (navigator.clipboard && window.isSecureContext) {
    try {
      await navigator.clipboard.writeText(text)
      return true
    } catch (e) {
      // Continue to fallback
    }
  }

  // Fallback for non-HTTPS or legacy mobile webviews
  try {
    const textArea = document.createElement('textarea')
    textArea.value = text
    textArea.style.position = 'fixed'
    textArea.style.left = '-999999px'
    textArea.style.top = '-999999px'
    document.body.appendChild(textArea)
    textArea.focus()
    textArea.select()
    const successful = document.execCommand('copy')
    document.body.removeChild(textArea)
    return successful
  } catch (err) {
    console.error('Gagal menyalin:', err)
    return false
  }
}
