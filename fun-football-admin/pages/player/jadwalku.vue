<template>
  <div class="p-5 space-y-5">
    <h2 class="text-xl font-bold text-gray-800">Jadwal Saya</h2>
    <p class="text-xs text-gray-500 -mt-3">Daftar pertandingan yang sudah Anda ikuti.</p>

    <!-- Looping Riwayat Pendaftaran -->
    <div v-for="reg in myRegistrations" :key="reg.id" class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100 relative space-y-3">
      
      <!-- Status Badge -->
      <div class="flex justify-between items-center">
        <span class="text-xs font-bold text-gray-400">ID: {{ reg.id.slice(-6) }}</span>
        
        <!-- Status Pembayaran -->
        <div>
          <span v-if="reg.payment_status === 'PAID'" class="bg-green-100 text-green-700 text-[10px] font-extrabold px-2.5 py-1 rounded-full">LUNAS (PAID)</span>
          <span v-else-if="reg.payment_status === 'VERIFYING'" class="bg-amber-100 text-amber-700 text-[10px] font-extrabold px-2.5 py-1 rounded-full">MENUNGGU VERIFIKASI</span>
          <span v-else class="bg-red-100 text-red-700 text-[10px] font-extrabold px-2.5 py-1 rounded-full">BELUM BAYAR (UNPAID)</span>
        </div>
      </div>

      <div class="border-t border-gray-100 pt-2">
        <p class="text-xs text-gray-400 font-semibold">Status Polling</p>
        <p class="text-sm font-bold text-gray-800">{{ reg.polling_status === 'JOIN' ? '⚽ Berhasil Masuk Kuota' : '⏳ Masuk Waiting List' }}</p>
      </div>

      <!-- Jika belum bayar / ingin upload ulang bukti -->
      <div v-if="reg.payment_status === 'UNPAID'" class="pt-2 border-t border-gray-100 flex justify-between items-center">
        <span class="text-xs text-red-500 font-medium">Segera unggah bukti pembayaran!</span>
        <button @click="navigateTo('/player')" class="bg-gradient-to-r from-red-600 to-orange-600 text-white text-xs font-bold py-2 px-4 rounded-xl active:scale-95 transition">
          Bayar Sekarang
        </button>
      </div>
    </div>

    <!-- Jika kosong -->
    <div v-if="myRegistrations.length === 0" class="text-center py-12">
      <Icon name="ph:ticket" class="text-5xl text-gray-300 mx-auto mb-2" />
      <p class="text-sm text-gray-500 font-medium">Anda belum mendaftar di jadwal manapun.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
definePageMeta({ layout: 'mobile' })

const myRegistrations = ref([])

onMounted(async () => {
  const userId = localStorage.getItem('user_id')
  if (!userId) return navigateTo('/player/login')

  try {
    const res = await $fetch(`http://localhost:8080/my-registrations?user_id=${userId}`)
    myRegistrations.value = res.data || []
  } catch (error) {
    console.error("Gagal memuat jadwal saya")
  }
})
</script>