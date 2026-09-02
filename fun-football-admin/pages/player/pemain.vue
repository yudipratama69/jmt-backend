<template>
  <div class="p-4 space-y-5">
    
    <!-- Header Page -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-xl font-black theme-text-main tracking-wide flex items-center gap-2">
          <Icon name="ph:users-three-bold" class="text-orange-500 text-2xl" /> Squad Pemain
        </h2>
        <p class="text-xs theme-text-muted mt-0.5">Daftar pemain yang sudah masuk kuota dan lunas.</p>
      </div>
      <span class="bg-emerald-500/10 text-emerald-500 text-xs font-black px-3 py-1 rounded-full border border-emerald-500/30 flex items-center gap-1.5 shadow-sm">
        <span class="w-1.5 h-1.5 rounded-full bg-emerald-500 animate-pulse"></span>
        {{ pemainLunas.length }} Pemain Lunas
      </span>
    </div>

    <!-- Squad List Box -->
    <div class="theme-bg-card backdrop-blur-md rounded-3xl border theme-border shadow-xl overflow-hidden transition-colors duration-300">
      
      <!-- Header List -->
      <div class="p-4 theme-bg-surface border-b theme-border-subtle flex justify-between items-center transition-colors duration-300">
        <span class="text-[10px] font-black uppercase tracking-wider theme-text-muted">LINE-UP PEMAIN RESMI</span>
        <span class="text-[10px] font-bold text-orange-500">JMT Fun Football</span>
      </div>

      <div v-if="pending" class="p-10 text-center theme-text-muted text-xs">
        <Icon name="ph:spinner-gap-bold" class="text-3xl text-orange-500 animate-spin mx-auto mb-2" />
        Memuat squad pemain...
      </div>
      
      <div v-else-if="pemainLunas.length === 0" class="p-10 text-center theme-text-muted text-xs space-y-2">
        <Icon name="ph:user-minus-bold" class="text-4xl theme-text-dim mx-auto" />
        <p class="font-bold theme-text-main">Belum Ada Pemain Lunas</p>
        <p class="text-[11px] theme-text-muted">Pemain yang telah melakukan pembayaran akan otomatis muncul di sini secara realtime.</p>
      </div>

      <ul v-else class="divide-y theme-border-subtle">
        <li 
          v-for="(pemain, index) in pemainLunas" 
          :key="pemain.id" 
          class="p-4 flex items-center justify-between hover:bg-orange-500/5 transition">
          
          <div class="flex items-center gap-3.5">
            <!-- Nomor Jersey / Urut -->
            <div class="w-8 h-8 rounded-xl theme-bg-surface border theme-border flex items-center justify-center text-orange-500 font-black text-xs shrink-0 shadow-inner">
              #{{ index + 1 }}
            </div>
            
            <!-- Avatar Inisial Bulat -->
            <div class="w-11 h-11 rounded-2xl bg-gradient-to-br from-red-600 to-orange-500 text-white font-black text-sm flex items-center justify-center shadow-md shadow-orange-600/20 shrink-0 border border-white/20">
              {{ pemain.user_name ? pemain.user_name.substring(0, 2).toUpperCase() : 'JM' }}
            </div>
            
            <!-- Info Nama Pemain -->
            <div>
              <p class="text-sm font-black theme-text-main leading-tight">{{ pemain.user_name }}</p>
              <div class="flex items-center gap-1.5 mt-0.5">
                <span class="text-[10px] theme-text-muted font-medium">Masuk Kuota</span>
                <span class="text-[9px] theme-bg-surface theme-text-main font-mono px-1.5 py-0.2 rounded border theme-border">
                  {{ pemain.payment_method === 'deposit' ? 'DEPOSIT' : 'TRANSFER' }}
                </span>
              </div>
            </div>
          </div>

          <!-- Status Lunas Badge -->
          <div class="px-3 py-1 bg-emerald-500/10 border border-emerald-500/30 text-emerald-500 rounded-full text-[10px] font-black tracking-wider flex items-center gap-1 shadow-sm shrink-0">
            <Icon name="ph:check-circle-fill" class="text-xs" />
            LUNAS
          </div>
        </li>
      </ul>
    </div>
    
  </div>
</template>

<script setup>
import { computed } from 'vue'

definePageMeta({ layout: 'mobile' })

// Tarik data semua pendaftaran
const { data: regData, pending, refresh: refreshRegistrations } = await useApiFetch('/registrations')

const { useAutoRefresh } = useRealtime()

// Pasang Auto-Refresh Realtime
useAutoRefresh(['REGISTRATION_UPDATED', 'PAYMENT_UPDATED'], () => {
  refreshRegistrations()
})

// Filter hanya pemain yang Polling = JOIN dan Bayar = PAID
const pemainLunas = computed(() => {
  if (!regData.value?.data) return []
  
  return regData.value.data.filter(
    (reg) => reg.polling_status === 'JOIN' && reg.payment_status === 'PAID'
  )
})
</script>