<template>
  <div class="p-4 space-y-5">
    
    <!-- Header Page -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-xl font-black theme-text-main tracking-wide flex items-center gap-2">
          <Icon name="ph:ticket-bold" class="text-orange-500 text-2xl" /> Tiket & Jadwalku
        </h2>
        <p class="text-xs theme-text-muted mt-0.5">Pertandingan yang Anda ikuti.</p>
      </div>
      <span class="theme-bg-card theme-text-main text-xs font-bold px-3 py-1 rounded-full border theme-border">
        {{ myRegistrations.length }} Tiket
      </span>
    </div>

    <!-- Looping Tiket Pertandingan -->
    <div class="space-y-4">
      <div 
        v-for="reg in myRegistrations" 
        :key="reg.id" 
        class="theme-bg-card backdrop-blur-md rounded-3xl border theme-border shadow-xl overflow-hidden relative transition-colors duration-300">
        
        <!-- Header Tiket Strip -->
        <div class="bg-gradient-to-r from-red-600 to-orange-600 px-5 py-2.5 flex justify-between items-center text-white">
          <span class="text-[10px] font-black uppercase tracking-wider flex items-center gap-1">
            <Icon name="ph:soccer-ball-bold" class="text-sm" /> MATCH PASS
          </span>
          <span class="text-[10px] font-mono opacity-90">PASS #{{ reg.id.slice(-6).toUpperCase() }}</span>
        </div>

        <div class="p-5 space-y-4">
          <!-- Status Polling & Status Bayar -->
          <div class="flex justify-between items-center">
            <div>
              <p class="text-[10px] theme-text-muted uppercase tracking-wider font-bold">Status Polling</p>
              <span 
                :class="reg.polling_status === 'JOIN' ? 'text-emerald-500 font-black' : 'text-amber-500 font-black'"
                class="text-sm flex items-center gap-1 mt-0.5">
                <Icon :name="reg.polling_status === 'JOIN' ? 'ph:check-circle-fill' : 'ph:clock-countdown-fill'" class="text-base" />
                {{ reg.polling_status === 'JOIN' ? 'MASUK KUOTA (MAIN)' : 'WAITING LIST' }}
              </span>
            </div>

            <!-- Status Pembayaran Badge -->
            <div>
              <span v-if="reg.payment_status === 'PAID'" class="bg-emerald-500/10 text-emerald-500 border border-emerald-500/30 text-[10px] font-black px-3 py-1.5 rounded-full flex items-center gap-1 shadow-sm">
                <Icon name="ph:check-bold" class="text-xs" /> LUNAS
              </span>
              <span v-else-if="reg.payment_status === 'VERIFYING'" class="bg-amber-500/10 text-amber-500 border border-amber-500/30 text-[10px] font-bold px-3 py-1.5 rounded-full flex items-center gap-1">
                <Icon name="ph:hourglass-bold" class="text-xs animate-spin" /> VERIFIKASI
              </span>
              <span v-else class="bg-red-500/10 text-red-500 border border-red-500/30 text-[10px] font-black px-3 py-1.5 rounded-full flex items-center gap-1">
                <Icon name="ph:warning-circle-bold" class="text-xs" /> BELUM BAYAR
              </span>
            </div>
          </div>

          <!-- Waktu Pendaftaran -->
          <div class="theme-bg-surface p-3 rounded-2xl border theme-border-subtle flex justify-between items-center text-xs theme-text-muted transition-colors duration-300">
            <span>Terdaftar Pada:</span>
            <span class="font-bold theme-text-main">{{ new Date(reg.registered_at).toLocaleString('id-ID', { dateStyle: 'medium', timeStyle: 'short' }) }}</span>
          </div>

          <!-- Action jika belum lunas -->
          <div v-if="reg.payment_status === 'UNPAID'" class="pt-2 border-t theme-border-subtle flex items-center justify-between gap-3">
            <span class="text-xs text-amber-500 font-semibold flex items-center gap-1">
              <Icon name="ph:warning-bold" class="text-sm shrink-0" /> Segera bayar sebelum kick-off!
            </span>
            <NuxtLink to="/player" class="bg-gradient-to-r from-orange-500 to-red-600 hover:from-orange-600 hover:to-red-700 text-white text-xs font-black py-2 px-4 rounded-xl shadow-lg transition active:scale-95 shrink-0">
              Bayar di Home
            </NuxtLink>
          </div>
        </div>

      </div>

      <!-- Jika Kosong -->
      <div v-if="myRegistrations.length === 0" class="theme-bg-card rounded-3xl p-10 text-center border theme-border space-y-3">
        <Icon name="ph:ticket-bold" class="text-6xl theme-text-muted mx-auto" />
        <h3 class="text-base font-black theme-text-main">Belum Ada Tiket Terdaftar</h3>
        <p class="text-xs theme-text-muted max-w-xs mx-auto">
          Kamu belum mendaftar di jadwal manapun. Yuk cek jadwal pertandingan dan amankan posisimu!
        </p>
        <NuxtLink to="/player" class="inline-block bg-gradient-to-r from-red-600 to-orange-600 text-white font-black text-xs py-3 px-6 rounded-2xl shadow-lg active:scale-95 transition">
          Cari Jadwal Sekarang ⚽
        </NuxtLink>
      </div>

    </div>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
definePageMeta({ layout: 'mobile' })

const { $api } = useNuxtApp()
const { useAutoRefresh } = useRealtime()
const myRegistrations = ref([])

const ambilJadwalku = async () => {
  const userId = localStorage.getItem('user_id')
  if (!userId) return navigateTo('/player/login')

  try {
    const res = await $api(`/my-registrations?user_id=${userId}`)
    myRegistrations.value = res.data || []
  } catch (error) {
    console.error("Gagal memuat jadwal saya")
  }
}

// Pasang Auto-Refresh Realtime
useAutoRefresh(['REGISTRATION_UPDATED', 'PAYMENT_UPDATED', 'EVENT_UPDATED'], () => {
  ambilJadwalku()
})

onMounted(() => {
  ambilJadwalku()
})
</script>