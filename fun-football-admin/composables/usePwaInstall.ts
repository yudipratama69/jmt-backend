import { ref, computed, onMounted } from 'vue'

const deferredPrompt = ref<any>(null)
const isInstallable = ref(false)
const isInstalled = ref(false)
const isIOS = ref(false)
const isDesktop = ref(false)
const showModal = ref(false)

export function usePwaInstall() {
  const syncPrompt = () => {
    if (typeof window === 'undefined') return
    if ((window as any).__deferredPwaPrompt) {
      deferredPrompt.value = (window as any).__deferredPwaPrompt
      isInstallable.value = true
      console.log('⚡ [usePwaInstall] Prompt synced from window.__deferredPwaPrompt!')
    }
  }

  const checkInstallStatus = () => {
    if (typeof window === 'undefined') return
    const isStandalone = window.matchMedia('(display-mode: standalone)').matches 
      || (window.navigator as any).standalone === true
      || document.referrer.includes('android-app://')
      || localStorage.getItem('pwa_installed') === 'true'

    if (isStandalone) {
      isInstalled.value = true
      isInstallable.value = false
    }
  }

  // Cek saat composable dipanggil
  syncPrompt()

  onMounted(() => {
    checkInstallStatus()
    syncPrompt()

    // Deteksi Device
    const userAgent = window.navigator.userAgent.toLowerCase()
    isIOS.value = /iphone|ipad|ipod/.test(userAgent)
    isDesktop.value = !/android|iphone|ipad|ipod|mobile/.test(userAgent)

    // Tangkap Custom Event jika script head menangkap duluan
    window.addEventListener('pwa-prompt-ready', (e: any) => {
      deferredPrompt.value = e.detail || (window as any).__deferredPwaPrompt
      isInstallable.value = true
      console.log('⚡ [usePwaInstall] pwa-prompt-ready event received!')
    })

    // Tangkap event install PWA langsung jika belum sempat tertangkap
    window.addEventListener('beforeinstallprompt', (e) => {
      e.preventDefault()
      deferredPrompt.value = e
      ;(window as any).__deferredPwaPrompt = e
      isInstallable.value = true
      console.log('⚡ [usePwaInstall] beforeinstallprompt captured directly!')
    })

    window.addEventListener('appinstalled', () => {
      isInstalled.value = true
      isInstallable.value = false
      deferredPrompt.value = null
      ;(window as any).__deferredPwaPrompt = null
      showModal.value = false
      localStorage.setItem('pwa_installed', 'true')
      console.log('⚡ [usePwaInstall] Aplikasi berhasil di-install!')
    })

    try {
      window.matchMedia('(display-mode: standalone)').addEventListener('change', (evt) => {
        if (evt.matches) {
          isInstalled.value = true
          isInstallable.value = false
          localStorage.setItem('pwa_installed', 'true')
        }
      })
    } catch (e) {
      // Ignore
    }
  })

  const hasPrompt = computed(() => {
    return !!deferredPrompt.value || (typeof window !== 'undefined' && !!(window as any).__deferredPwaPrompt)
  })

  const triggerInstall = async () => {
    if (isInstalled.value) {
      return
    }

    // Sync ulang sebelum prompt
    syncPrompt()
    const promptEvent = deferredPrompt.value || (typeof window !== 'undefined' ? (window as any).__deferredPwaPrompt : null)

    if (promptEvent) {
      try {
        promptEvent.prompt()
        const { outcome } = await promptEvent.userChoice
        console.log('⚡ [usePwaInstall] Outcome user pilihan install:', outcome)
        if (outcome === 'accepted') {
          isInstallable.value = false
          isInstalled.value = true
          showModal.value = false
          localStorage.setItem('pwa_installed', 'true')
        }
        deferredPrompt.value = null
        if (typeof window !== 'undefined') (window as any).__deferredPwaPrompt = null
      } catch (err) {
        console.warn('⚡ [usePwaInstall] Gagal memicu install otomatis:', err)
        showModal.value = true
      }
    } else {
      // Jika browser belum melempar prompt otomatis, munculkan panduan instruksi
      showModal.value = true
    }
  }

  const closeModal = () => {
    showModal.value = false
  }

  return {
    isInstallable,
    isInstalled,
    isIOS,
    isDesktop,
    hasPrompt,
    showModal,
    triggerInstall,
    closeModal
  }
}
