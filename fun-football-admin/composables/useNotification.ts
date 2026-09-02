import { ref, computed } from 'vue'

export interface AppNotification {
  id?: string
  title: string
  message: string
  type: 'INFO' | 'JADWAL' | 'URGENT' | 'PROMO' | string
  sender?: string
  created_at: string
  isRead?: boolean
}

const notifications = ref<AppNotification[]>([])
const isDrawerOpen = ref(false)
const hasPermission = ref(false)

// Fungsi sintesis audio notifikasi sporty (Dua nada chime peluit lembut)
export function playNotificationSound() {
  if (typeof window === 'undefined') return
  try {
    const AudioContextClass = window.AudioContext || (window as any).webkitAudioContext
    if (!AudioContextClass) return
    const ctx = new AudioContextClass()

    const now = ctx.currentTime
    // Nada 1 (587 Hz - D5)
    const osc1 = ctx.createOscillator()
    const gain1 = ctx.createGain()
    osc1.type = 'sine'
    osc1.frequency.setValueAtTime(587.33, now)
    gain1.gain.setValueAtTime(0.15, now)
    gain1.gain.exponentialRampToValueAtTime(0.001, now + 0.18)
    osc1.connect(gain1)
    gain1.connect(ctx.destination)
    osc1.start(now)
    osc1.stop(now + 0.18)

    // Nada 2 (880 Hz - A5)
    const osc2 = ctx.createOscillator()
    const gain2 = ctx.createGain()
    osc2.type = 'sine'
    osc2.frequency.setValueAtTime(880, now + 0.12)
    gain2.gain.setValueAtTime(0.2, now + 0.12)
    gain2.gain.exponentialRampToValueAtTime(0.001, now + 0.38)
    osc2.connect(gain2)
    gain2.connect(ctx.destination)
    osc2.start(now + 0.12)
    osc2.stop(now + 0.38)
  } catch (e) {
    // Audio restricted before user gesture
  }
}

// Fungsi getar HP
export function triggerHaptic() {
  if (typeof window !== 'undefined' && 'vibrate' in navigator) {
    try {
      navigator.vibrate([100, 50, 100])
    } catch (e) {
      // Ignore
    }
  }
}

export function useNotification() {
  const { $api } = useNuxtApp()
  const toast = useToast()

  // Cek izin notifikasi browser
  const checkPermission = () => {
    if (typeof window !== 'undefined' && 'Notification' in window) {
      hasPermission.value = Notification.permission === 'granted'
    }
  }

  // Minta izin notifikasi browser
  const requestPermission = async () => {
    if (typeof window === 'undefined' || !('Notification' in window)) return false
    try {
      const permission = await Notification.requestPermission()
      hasPermission.value = permission === 'granted'
      return hasPermission.value
    } catch (e) {
      return false
    }
  }

  // Ambil daftar notifikasi dari backend & gabungkan status isRead lokal
  const fetchNotifications = async () => {
    try {
      const data: any = await $api('/notifications')
      if (Array.isArray(data)) {
        const readIds: string[] = JSON.parse(localStorage.getItem('jmt_read_notifications') || '[]')
        notifications.value = data.map(item => ({
          ...item,
          isRead: readIds.includes(item.id)
        }))
      }
    } catch (err) {
      console.warn('Gagal memuat notifikasi dari server:', err)
    }
  }

  // Tampilkan notifikasi push ke sistem browser / status bar HP
  const triggerPushNotification = (title: string, body: string, icon = '/logo-jmt.png') => {
    playNotificationSound()
    triggerHaptic()

    if (typeof window !== 'undefined' && 'Notification' in window && Notification.permission === 'granted') {
      try {
        new Notification(title, {
          body,
          icon,
          badge: '/logo-jmt.png'
        })
      } catch (e) {
        // Fallback service worker
        if ('serviceWorker' in navigator && navigator.serviceWorker.controller) {
          navigator.serviceWorker.ready.then(registration => {
            registration.showNotification(title, {
              body,
              icon,
              badge: '/logo-jmt.png'
            })
          })
        }
      }
    }
  }

  // Handler saat ada notifikasi masuk dari WebSocket
  const handleIncomingNotification = (notif: AppNotification) => {
    console.log('🔔 [useNotification] Menerima notifikasi baru:', notif)
    
    // Cegah duplikasi
    const exists = notifications.value.some(n => n.id === notif.id)
    if (!exists) {
      notifications.value.unshift({
        ...notif,
        isRead: false
      })
    }

    // Munculkan toast sporty di layar
    toast.info(notif.message, `📢 ${notif.title}`, 6000)

    // Trigger audio, getar, dan push notification
    triggerPushNotification(`⚽ ${notif.title}`, notif.message)
  }

  // Tandai notifikasi sebagai telah dibaca
  const markAsRead = (id?: string) => {
    if (!id) return
    const readIds: string[] = JSON.parse(localStorage.getItem('jmt_read_notifications') || '[]')
    if (!readIds.includes(id)) {
      readIds.push(id)
      localStorage.setItem('jmt_read_notifications', JSON.stringify(readIds))
    }
    const target = notifications.value.find(n => n.id === id)
    if (target) {
      target.isRead = true
    }
  }

  // Tandai semua dibaca
  const markAllAsRead = () => {
    const allIds = notifications.value.map(n => n.id).filter(Boolean) as string[]
    localStorage.setItem('jmt_read_notifications', JSON.stringify(allIds))
    notifications.value.forEach(n => { n.isRead = true })
  }

  // Hapus dari list lokal
  const removeLocalNotification = (id?: string) => {
    if (!id) return
    notifications.value = notifications.value.filter(n => n.id !== id)
  }

  const unreadCount = computed(() => {
    return notifications.value.filter(n => !n.isRead).length
  })

  return {
    notifications,
    isDrawerOpen,
    hasPermission,
    unreadCount,
    checkPermission,
    requestPermission,
    fetchNotifications,
    triggerPushNotification,
    handleIncomingNotification,
    markAsRead,
    markAllAsRead,
    removeLocalNotification
  }
}
