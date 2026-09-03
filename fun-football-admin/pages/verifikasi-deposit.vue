<template>
  <div class="p-10 text-center theme-text-muted">
    <Icon name="ph:spinner-gap-bold" class="text-3xl text-orange-500 animate-spin mx-auto mb-2" />
    <p class="text-sm">Mengalihkan ke Pusat Verifikasi...</p>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'

onMounted(() => {
  navigateTo('/verifikasi?tab=deposit')
})
</script>