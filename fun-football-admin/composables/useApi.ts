export function useApiFetch<T = any>(url: string | (() => string), options: any = {}) {
  const config = useRuntimeConfig()

  return useFetch<T>(url, {
    baseURL: config.public.apiBase,
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
  const config = useRuntimeConfig()
  const base = (config.public.apiBase || '').replace(/\/+$/, '')

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
