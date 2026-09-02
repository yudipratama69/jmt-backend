<template>
  <div :class="isDark ? 'light bg-gray-50 text-gray-900' : 'dark bg-slate-950 text-slate-100'" class="min-h-screen transition-colors duration-300">
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
    <ToastContainer />
  </div>
</template>

<script setup>
const { isDark, initTheme } = useTheme()

onMounted(() => {
  initTheme()
})
</script>