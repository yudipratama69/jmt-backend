import { ref, onMounted } from 'vue'

const isDark = ref(true)

export function useTheme() {
  const applyTheme = (dark: boolean) => {
    isDark.value = dark
    if (import.meta.client) {
      localStorage.setItem('jmt_theme', dark ? 'dark' : 'light')
      if (dark) {
        document.documentElement.classList.add('dark')
        document.documentElement.classList.remove('light')
      } else {
        document.documentElement.classList.remove('dark')
        document.documentElement.classList.add('light')
      }
    }
  }

  const toggleTheme = () => {
    applyTheme(!isDark.value)
  }

  const initTheme = () => {
    if (import.meta.client) {
      const saved = localStorage.getItem('jmt_theme')
      if (saved) {
        applyTheme(saved === 'dark')
      } else {
        // Default ke Dark Mode untuk tampilan sporty modern
        applyTheme(true)
      }
    }
  }

  onMounted(() => {
    initTheme()
  })

  return {
    isDark,
    toggleTheme,
    applyTheme,
    initTheme
  }
}
