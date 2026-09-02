import { ref } from 'vue'

export interface ToastItem {
  id: string
  type: 'success' | 'error' | 'info' | 'warning'
  title?: string
  message: string
  duration?: number
}

export interface ConfirmOptions {
  title?: string
  message: string
  confirmText?: string
  cancelText?: string
  type?: 'danger' | 'warning' | 'info'
  onConfirm: () => void | Promise<void>
  onCancel?: () => void
}

const toasts = ref<ToastItem[]>([])
const confirmModal = ref<{
  show: boolean
  options: ConfirmOptions | null
}>({
  show: false,
  options: null
})

export function useToast() {
  const show = (type: ToastItem['type'], message: string, title?: string, duration = 3500) => {
    const id = Date.now().toString() + Math.random().toString(36).substring(2, 5)
    toasts.value.push({ id, type, title, message, duration })

    if (duration > 0) {
      setTimeout(() => {
        remove(id)
      }, duration)
    }
  }

  const success = (message: string, title = 'Berhasil!') => show('success', message, title)
  const error = (message: string, title = 'Gagal!') => show('error', message, title)
  const info = (message: string, title = 'Info') => show('info', message, title)
  const warning = (message: string, title = 'Perhatian') => show('warning', message, title)

  const remove = (id: string) => {
    const index = toasts.value.findIndex(t => t.id === id)
    if (index !== -1) {
      toasts.value.splice(index, 1)
    }
  }

  const confirm = (options: ConfirmOptions) => {
    confirmModal.value = {
      show: true,
      options
    }
  }

  const handleConfirm = async () => {
    if (confirmModal.value.options?.onConfirm) {
      await confirmModal.value.options.onConfirm()
    }
    confirmModal.value.show = false
    confirmModal.value.options = null
  }

  const handleCancel = () => {
    if (confirmModal.value.options?.onCancel) {
      confirmModal.value.options.onCancel()
    }
    confirmModal.value.show = false
    confirmModal.value.options = null
  }

  return {
    toasts,
    confirmModal,
    show,
    success,
    error,
    info,
    warning,
    remove,
    confirm,
    handleConfirm,
    handleCancel
  }
}
