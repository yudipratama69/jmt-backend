<template>
  <div class="min-h-screen bg-gray-900 flex items-center justify-center p-4">
    <div class="text-center">
      <div class="w-12 h-12 border-4 border-orange-500 border-t-transparent rounded-full animate-spin mx-auto mb-3"></div>
      <p class="text-xs text-gray-400 font-medium">Membuka JMT Sport...</p>
    </div>
  </div>
</template>

<script setup>
// Redirect otomatis ke halaman Player
definePageMeta({
  layout: false,
  middleware: [
    () => {
      return navigateTo('/player', { replace: true })
    }
  ]
})

// Fallback jika middleware dilewati di client
onMounted(() => {
  navigateTo('/player', { replace: true })
})
</script>