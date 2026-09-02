import { ref, onMounted, onUnmounted } from 'vue'

type MessageHandler = (payload: any) => void

const socket = ref<WebSocket | null>(null)
const isConnected = ref(false)
const listeners = new Map<string, Set<MessageHandler>>()
let reconnectTimeout: any = null

function getWsUrl(): string {
  if (typeof window !== 'undefined') {
    const proto = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
    const host = window.location.hostname
    const port = window.location.port

    // Jika dibuka di dev server port 3000
    if (port === '3000') {
      return `${proto}//${host}:8080/ws`
    }
    // Jika dibuka via Docker port 7777 atau tunnel ngrok langsung
    return `${proto}//${window.location.host}/ws`
  }

  const config = useRuntimeConfig()
  const apiBase = config.public.apiBase || 'http://localhost:8080'
  let wsUrl = apiBase.replace(/^http(s)?:\/\//, (match, p1) => (p1 ? 'wss://' : 'ws://')).replace(/\/$/, '')
  return `${wsUrl}/ws`
}

function initWebSocket() {
  if (typeof window === 'undefined') return
  if (socket.value && (socket.value.readyState === WebSocket.OPEN || socket.value.readyState === WebSocket.CONNECTING)) {
    return
  }

  const url = getWsUrl()
  try {
    const ws = new WebSocket(url)

    ws.onopen = () => {
      console.log('⚡ [Realtime] Terhubung ke WebSocket server:', url)
      isConnected.value = true
      if (reconnectTimeout) {
        clearTimeout(reconnectTimeout)
        reconnectTimeout = null
      }
    }

    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data)
        if (data && data.type) {
          console.log('⚡ [Realtime] Event diterima:', data.type, data.payload)
          const handlers = listeners.get(data.type)
          if (handlers) {
            handlers.forEach((handler) => {
              try {
                handler(data.payload)
              } catch (err) {
                console.error('Error executing realtime listener for', data.type, err)
              }
            })
          }
          // Panggil juga wildcard listener '*'
          const allHandlers = listeners.get('*')
          if (allHandlers) {
            allHandlers.forEach((handler) => handler(data))
          }
        }
      } catch (err) {
        console.error('Gagal mem-parsing pesan WebSocket:', err)
      }
    }

    ws.onclose = () => {
      isConnected.value = false
      socket.value = null
      // Reconnect otomatis setelah 3 detik
      if (!reconnectTimeout) {
        reconnectTimeout = setTimeout(() => {
          reconnectTimeout = null
          initWebSocket()
        }, 3000)
      }
    }

    ws.onerror = (err) => {
      console.warn('⚠️ [Realtime] WebSocket error:', err)
      ws.close()
    }

    socket.value = ws
  } catch (err) {
    console.error('Gagal menginisialisasi WebSocket:', err)
  }
}

export function useRealtime() {
  onMounted(() => {
    initWebSocket()
  })

  const on = (eventType: string, handler: MessageHandler) => {
    if (!listeners.has(eventType)) {
      listeners.set(eventType, new Set())
    }
    listeners.get(eventType)!.add(handler)

    // Return cleanup function
    return () => {
      const handlers = listeners.get(eventType)
      if (handlers) {
        handlers.delete(handler)
      }
    }
  }

  const onMessage = (handler: (data: { type: string, payload: any }) => void) => {
    return on('*', handler)
  }

  const useAutoRefresh = (eventTypes: string | string[], refreshCallback: () => void | Promise<any>) => {
    const types = Array.isArray(eventTypes) ? eventTypes : [eventTypes]
    const unsubs: (() => void)[] = []

    onMounted(() => {
      initWebSocket()
      types.forEach((type) => {
        const unsub = on(type, () => {
          console.log(`⚡ [Realtime Auto-Refresh] Memicu refresh untuk event: ${type}`)
          refreshCallback()
        })
        unsubs.push(unsub)
      })
    })

    onUnmounted(() => {
      unsubs.forEach((unsub) => unsub())
    })
  }

  return {
    isConnected,
    on,
    onMessage,
    useAutoRefresh
  }
}
