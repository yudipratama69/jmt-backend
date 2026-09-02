<template>
  <div class="p-4 space-y-5">
    
    <!-- Header Page -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-xl font-black theme-text-main tracking-wide flex items-center gap-2">
          <Icon name="ph:users-three-bold" class="text-orange-500 text-2xl" /> Squad Pemain
        </h2>
        <p class="text-xs theme-text-muted mt-0.5">Daftar pemain resmi per sesi jadwal pertandingan.</p>
      </div>
      
      <!-- Total Lunas Badge -->
      <span class="bg-emerald-500/10 text-emerald-500 text-xs font-black px-3 py-1 rounded-full border border-emerald-500/30 flex items-center gap-1.5 shadow-sm shrink-0">
        <span class="w-1.5 h-1.5 rounded-full bg-emerald-500 animate-pulse"></span>
        <span v-if="selectedEventId === 'all'">{{ totalPemainLunas }} Pemain Lunas</span>
        <span v-else-if="currentEvent">{{ currentEventLunas.length }} / {{ currentEvent.quota_max }} Lunas</span>
        <span v-else>0 Pemain Lunas</span>
      </span>
    </div>

    <!-- Loading State Awal -->
    <div v-if="pendingEvents || pendingRegs" class="theme-bg-card rounded-3xl p-10 text-center border theme-border space-y-2">
      <Icon name="ph:spinner-gap-bold" class="text-3xl text-orange-500 animate-spin mx-auto mb-2" />
      <p class="text-xs theme-text-muted">Memuat jadwal & squad pemain...</p>
    </div>

    <!-- Jika Belum Ada Jadwal Sama Sekali -->
    <div v-else-if="events.length === 0" class="theme-bg-card rounded-3xl p-10 text-center border theme-border space-y-3">
      <Icon name="ph:calendar-x-bold" class="text-5xl theme-text-muted mx-auto" />
      <h3 class="text-base font-black theme-text-main">Belum Ada Jadwal Pertandingan</h3>
      <p class="text-xs theme-text-muted max-w-xs mx-auto">
        Pengurus belum membuat jadwal pertandingan. Cek kembali nanti ya!
      </p>
    </div>

    <!-- Konten Squad Per Sesi -->
    <div v-else class="space-y-4">
      
      <!-- Selector Pill Sesi Jadwal (Horizontal Scrollable) -->
      <div class="space-y-1.5">
        <div class="flex items-center justify-between px-0.5">
          <span class="text-[10px] font-black uppercase tracking-wider theme-text-muted flex items-center gap-1">
            <Icon name="ph:calendar-star-bold" class="text-orange-500 text-xs" /> Pilih Sesi Jadwal
          </span>
          <span class="text-[10px] theme-text-dim font-bold">{{ events.length }} Sesi Tersedia</span>
        </div>

        <div class="flex items-center gap-2 overflow-x-auto pb-1.5 pt-0.5 -mx-4 px-4 scrollbar-none">
          
          <!-- Tombol Semua Sesi -->
          <button 
            type="button"
            @click="selectedEventId = 'all'"
            :class="selectedEventId === 'all' 
              ? 'bg-gradient-to-r from-red-600 to-orange-500 text-white shadow-md shadow-orange-600/30 border-orange-500 font-black' 
              : 'theme-bg-card theme-border theme-text-muted hover:theme-text-main font-bold'"
            class="px-3.5 py-2 rounded-2xl text-xs shrink-0 border transition-all duration-200 flex items-center gap-2 active:scale-95">
            <Icon name="ph:squares-four-bold" class="text-sm shrink-0" />
            <span>Semua Sesi</span>
            <span 
              class="text-[10px] px-1.5 py-0.2 rounded-full font-mono font-bold"
              :class="selectedEventId === 'all' ? 'bg-white/20 text-white' : 'theme-bg-surface theme-text-muted'">
              {{ totalPemainLunas }}
            </span>
          </button>

          <!-- Tombol Masing-Masing Sesi Jadwal -->
          <button 
            v-for="event in events" 
            :key="event.id"
            type="button"
            @click="selectedEventId = event.id"
            :class="selectedEventId === event.id 
              ? 'bg-gradient-to-r from-red-600 to-orange-500 text-white shadow-md shadow-orange-600/30 border-orange-500 font-black' 
              : 'theme-bg-card theme-border theme-text-muted hover:theme-text-main font-bold'"
            class="px-3.5 py-2 rounded-2xl text-xs shrink-0 border transition-all duration-200 flex items-center gap-2 active:scale-95">
            <Icon name="ph:soccer-ball-bold" class="text-sm shrink-0 text-amber-300" v-if="selectedEventId === event.id" />
            <Icon name="ph:calendar-blank-bold" class="text-sm shrink-0 text-orange-500" v-else />
            <div class="flex flex-col text-left">
              <span class="max-w-[130px] truncate leading-tight">{{ event.title }}</span>
              <span class="text-[9px] opacity-80 font-normal leading-tight mt-0.5">
                {{ formatShortDate(event.match_date) }}
              </span>
            </div>
            <span 
              class="text-[10px] px-1.5 py-0.2 rounded-full font-mono font-bold shrink-0 ml-0.5"
              :class="selectedEventId === event.id ? 'bg-white/20 text-white' : 'theme-bg-surface text-emerald-500 border theme-border-subtle'">
              {{ getPemainLunas(event.id).length }}/{{ event.quota_max }}
            </span>
          </button>

        </div>
      </div>

      <!-- TAMPILAN 1: JIKA SATU SESI DIPILIH -->
      <div v-if="selectedEventId !== 'all' && currentEvent" class="space-y-4 animate-fade-in">
        
        <!-- Info Banner Sesi Pertandingan -->
        <div class="theme-bg-card backdrop-blur-md rounded-3xl p-4 border theme-border shadow-xl space-y-3 transition-colors duration-300">
          
          <div class="flex justify-between items-start gap-2">
            <div class="flex items-center gap-1.5 text-xs text-orange-500 font-bold">
              <Icon name="ph:map-pin-bold" class="text-sm shrink-0" />
              <span class="truncate">{{ currentEvent.location || 'Lapangan JMT' }}</span>
            </div>

            <!-- Status Badge Jadwal -->
            <span 
              v-if="currentEvent.status === 'OPEN'" 
              class="bg-emerald-500/10 text-emerald-500 border border-emerald-500/30 text-[10px] font-black px-2.5 py-0.5 rounded-full uppercase flex items-center gap-1">
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-500 animate-pulse"></span> BUKA
            </span>
            <span 
              v-else-if="currentEvent.status === 'FULL'" 
              class="bg-red-500/10 text-red-500 border border-red-500/30 text-[10px] font-black px-2.5 py-0.5 rounded-full uppercase">
              PENUH
            </span>
            <span 
              v-else 
              class="theme-bg-surface theme-text-muted text-[10px] font-black px-2.5 py-0.5 rounded-full uppercase">
              {{ currentEvent.status }}
            </span>
          </div>

          <div>
            <h3 class="font-black theme-text-main text-base leading-snug">{{ currentEvent.title }}</h3>
          </div>

          <!-- Waktu & Jam Main -->
          <div class="grid grid-cols-2 gap-2 theme-bg-surface p-2.5 rounded-2xl border theme-border-subtle text-xs transition-colors duration-300">
            <div class="flex items-center gap-1.5 theme-text-main">
              <Icon name="ph:calendar-blank-bold" class="text-orange-500 text-sm shrink-0" />
              <span class="font-bold truncate">{{ formatFullDate(currentEvent.match_date) }}</span>
            </div>
            <div class="flex items-center gap-1.5 theme-text-main">
              <Icon name="ph:clock-bold" class="text-orange-500 text-sm shrink-0" />
              <span class="font-bold">{{ formatTime(currentEvent.match_date) }} WIB</span>
            </div>
          </div>

          <!-- Progress Kuota Bar Sesi -->
          <div class="space-y-1.5 pt-1">
            <div class="flex justify-between items-center text-xs">
              <span class="theme-text-muted font-bold flex items-center gap-1 text-[11px]">
                <Icon name="ph:users-three-bold" class="text-sm theme-text-main" />
                <span>{{ currentEventLunas.length }} / {{ currentEvent.quota_max }} Pemain Lunas</span>
              </span>
              <span 
                :class="(currentEvent.quota_max - currentEventLunas.length) <= 3 ? 'text-red-500 font-black' : 'text-emerald-500 font-bold'"
                class="text-[11px]">
                {{ (currentEvent.quota_max - currentEventLunas.length) > 0 ? `Sisa ${currentEvent.quota_max - currentEventLunas.length} Slot` : 'Kuota Penuh' }}
              </span>
            </div>
            
            <div class="w-full theme-bg-surface rounded-full h-2 overflow-hidden p-0.5 border theme-border">
              <div 
                class="h-full rounded-full transition-all duration-500"
                :class="(currentEventLunas.length / currentEvent.quota_max) >= 0.8 ? 'bg-gradient-to-r from-orange-500 to-red-500' : 'bg-gradient-to-r from-amber-500 to-emerald-400'"
                :style="{ width: `${Math.min(100, Math.round((currentEventLunas.length / currentEvent.quota_max) * 100))}%` }">
              </div>
            </div>
          </div>

        </div>

        <!-- Squad List Box Sesi Ini -->
        <div class="theme-bg-card backdrop-blur-md rounded-3xl border theme-border shadow-xl overflow-hidden transition-colors duration-300">
          
          <!-- Header List -->
          <div class="p-4 theme-bg-surface border-b theme-border-subtle flex justify-between items-center transition-colors duration-300">
            <div>
              <span class="text-[10px] font-black uppercase tracking-wider theme-text-muted">LINE-UP RESMI SESI INI</span>
            </div>
            <span class="text-[10px] font-bold text-orange-500 truncate max-w-[160px]">{{ currentEvent.title }}</span>
          </div>

          <!-- Jika Kosong -->
          <div v-if="currentEventLunas.length === 0" class="p-10 text-center theme-text-muted text-xs space-y-2">
            <Icon name="ph:user-minus-bold" class="text-4xl theme-text-dim mx-auto" />
            <p class="font-bold theme-text-main">Belum Ada Pemain Lunas di Sesi Ini</p>
            <p class="text-[11px] theme-text-muted">Pemain yang telah melakukan pembayaran untuk sesi ini akan otomatis muncul di sini secara realtime.</p>
            <div class="pt-2">
              <NuxtLink to="/player" class="inline-flex items-center gap-1.5 px-4 py-2 bg-gradient-to-r from-red-600 to-orange-500 hover:from-red-700 hover:to-orange-600 text-white rounded-xl text-xs font-black shadow-md transition active:scale-95">
                <Icon name="ph:ticket-bold" class="text-sm" /> Daftar Sekarang
              </NuxtLink>
            </div>
          </div>

          <!-- List Pemain Resmi Sesi Ini -->
          <ul v-else class="divide-y theme-border-subtle">
            <li 
              v-for="(pemain, index) in currentEventLunas" 
              :key="pemain.id" 
              class="p-4 flex items-center justify-between hover:bg-orange-500/5 transition">
              
              <div class="flex items-center gap-3.5">
                <!-- Nomor Jersey Sesi Ini -->
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
                    <span class="text-[9px] theme-bg-surface theme-text-main font-mono px-1.5 py-0.2 rounded border theme-border uppercase">
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

        <!-- Daftar Pemain Waiting List / Menunggu Verifikasi (Jika Ada) -->
        <div v-if="currentEventWaiting.length > 0" class="theme-bg-card backdrop-blur-md rounded-3xl border theme-border shadow-md overflow-hidden transition-colors duration-300">
          <div class="p-3.5 theme-bg-surface border-b theme-border-subtle flex justify-between items-center">
            <span class="text-[10px] font-black uppercase tracking-wider theme-text-muted flex items-center gap-1.5">
              <Icon name="ph:clock-countdown-bold" class="text-amber-500 text-xs" /> Waiting List / Verifikasi
            </span>
            <span class="text-[10px] font-bold text-amber-500 font-mono">{{ currentEventWaiting.length }} Pemain</span>
          </div>

          <ul class="divide-y theme-border-subtle">
            <li 
              v-for="pemain in currentEventWaiting" 
              :key="pemain.id" 
              class="p-3.5 flex items-center justify-between opacity-85 hover:opacity-100 transition">
              <div class="flex items-center gap-3">
                <div class="w-8 h-8 rounded-xl theme-bg-surface border theme-border flex items-center justify-center text-amber-500 font-bold text-xs shrink-0">
                  ⏳
                </div>
                <div>
                  <p class="text-xs font-bold theme-text-main">{{ pemain.user_name }}</p>
                  <p class="text-[10px] theme-text-muted">
                    {{ pemain.polling_status === 'WAITING' ? 'Waiting List' : 'Menunggu Konfirmasi Pembayaran' }}
                  </p>
                </div>
              </div>

              <span 
                v-if="pemain.payment_status === 'VERIFYING'" 
                class="px-2.5 py-0.5 bg-amber-500/10 text-amber-500 border border-amber-500/30 rounded-full text-[9px] font-black flex items-center gap-1">
                <Icon name="ph:hourglass-bold" class="text-[10px] animate-spin" /> VERIFIKASI
              </span>
              <span 
                v-else 
                class="px-2.5 py-0.5 bg-red-500/10 text-red-500 border border-red-500/30 rounded-full text-[9px] font-black">
                BELUM BAYAR
              </span>
            </li>
          </ul>
        </div>

      </div>

      <!-- TAMPILAN 2: JIKA PILIH "SEMUA SESI" (Dikelompokkan Terpisah Per Kartu Sesi) -->
      <div v-else class="space-y-6 animate-fade-in">
        
        <div 
          v-for="event in events" 
          :key="event.id"
          class="theme-bg-card backdrop-blur-md rounded-3xl border theme-border shadow-xl overflow-hidden transition-colors duration-300">
          
          <!-- Header Sesi Card -->
          <div class="p-4 theme-bg-surface border-b theme-border-subtle flex flex-col gap-2">
            <div class="flex justify-between items-start">
              <div>
                <span class="text-[10px] font-bold text-orange-500 flex items-center gap-1 uppercase tracking-wider">
                  <Icon name="ph:map-pin-bold" class="text-xs" /> {{ event.location || 'Lapangan JMT' }}
                </span>
                <h3 class="font-black theme-text-main text-sm mt-0.5">{{ event.title }}</h3>
              </div>
              <span class="text-[10px] font-black bg-emerald-500/10 text-emerald-500 border border-emerald-500/30 px-2.5 py-0.5 rounded-full shrink-0">
                {{ getPemainLunas(event.id).length }} / {{ event.quota_max }} Lunas
              </span>
            </div>

            <div class="flex items-center gap-3 text-[11px] theme-text-muted">
              <span class="flex items-center gap-1 font-medium">
                <Icon name="ph:calendar-blank-bold" class="text-orange-500" />
                {{ formatFullDate(event.match_date) }}
              </span>
              <span class="flex items-center gap-1 font-medium">
                <Icon name="ph:clock-bold" class="text-orange-500" />
                {{ formatTime(event.match_date) }} WIB
              </span>
            </div>
          </div>

          <!-- Empty State Sesi -->
          <div v-if="getPemainLunas(event.id).length === 0" class="p-6 text-center theme-text-muted text-xs space-y-1.5">
            <Icon name="ph:user-minus-bold" class="text-3xl theme-text-dim mx-auto" />
            <p class="font-bold theme-text-main text-xs">Belum Ada Pemain Lunas di Sesi Ini</p>
            <p class="text-[10px] theme-text-muted">Jadwal ini belum memiliki pemain yang lunas.</p>
          </div>

          <!-- Line-up Pemain Per Sesi -->
          <ul v-else class="divide-y theme-border-subtle">
            <li 
              v-for="(pemain, index) in getPemainLunas(event.id)" 
              :key="pemain.id" 
              class="p-3.5 flex items-center justify-between hover:bg-orange-500/5 transition">
              
              <div class="flex items-center gap-3">
                <!-- Nomor Urut Jersey Sesi Ini -->
                <div class="w-7 h-7 rounded-lg theme-bg-surface border theme-border flex items-center justify-center text-orange-500 font-black text-xs shrink-0 shadow-inner">
                  #{{ index + 1 }}
                </div>
                
                <!-- Avatar Inisial -->
                <div class="w-9 h-9 rounded-xl bg-gradient-to-br from-red-600 to-orange-500 text-white font-black text-xs flex items-center justify-center shadow-sm shrink-0 border border-white/20">
                  {{ pemain.user_name ? pemain.user_name.substring(0, 2).toUpperCase() : 'JM' }}
                </div>
                
                <!-- Info Nama Pemain -->
                <div>
                  <p class="text-xs font-black theme-text-main leading-tight">{{ pemain.user_name }}</p>
                  <div class="flex items-center gap-1 mt-0.5">
                    <span class="text-[9px] theme-text-muted font-medium">Masuk Kuota</span>
                    <span class="text-[8px] theme-bg-surface theme-text-main font-mono px-1 py-0.2 rounded border theme-border uppercase">
                      {{ pemain.payment_method === 'deposit' ? 'DEPOSIT' : 'TRANSFER' }}
                    </span>
                  </div>
                </div>
              </div>

              <!-- Status Lunas Badge -->
              <div class="px-2.5 py-0.5 bg-emerald-500/10 border border-emerald-500/30 text-emerald-500 rounded-full text-[9px] font-black tracking-wider flex items-center gap-1 shadow-sm shrink-0">
                <Icon name="ph:check-circle-fill" class="text-[10px]" />
                LUNAS
              </div>
            </li>
          </ul>

        </div>

      </div>

    </div>
    
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'

definePageMeta({ layout: 'mobile' })

// Tarik data jadwal (events) dan pendaftaran (registrations)
const { data: eventsData, pending: pendingEvents, refresh: refreshEvents } = await useApiFetch('/events')
const { data: regData, pending: pendingRegs, refresh: refreshRegistrations } = await useApiFetch('/registrations')

const { useAutoRefresh } = useRealtime()

// Pasang Auto-Refresh Realtime
useAutoRefresh(['REGISTRATION_UPDATED', 'PAYMENT_UPDATED', 'EVENT_UPDATED'], () => {
  refreshEvents()
  refreshRegistrations()
})

const events = computed(() => eventsData.value?.data || [])
const registrations = computed(() => regData.value?.data || [])

// State ID Sesi yang dipilih (default memilih sesi pertama yang ada)
const selectedEventId = ref('all')

// Set default sesi ke event pertama jika data tersedia
onMounted(() => {
  if (events.value.length > 0) {
    selectedEventId.value = events.value[0].id
  }
})

watch(events, (newEvents) => {
  if (newEvents.length > 0 && selectedEventId.value === 'all') {
    // Jika awalnya kosong atau belum dipilih
    selectedEventId.value = newEvents[0].id
  }
}, { immediate: true })

// Helper: Filter pendaftaran lunas untuk satu event tertentu
const getPemainLunas = (eventId) => {
  if (!registrations.value || !eventId) return []
  return registrations.value.filter(
    (reg) => String(reg.event_id) === String(eventId) && reg.polling_status === 'JOIN' && reg.payment_status === 'PAID'
  )
}

// Helper: Filter pendaftaran waiting / belum lunas untuk satu event tertentu
const getPemainWaiting = (eventId) => {
  if (!registrations.value || !eventId) return []
  return registrations.value.filter(
    (reg) => String(reg.event_id) === String(eventId) && (reg.polling_status !== 'JOIN' || reg.payment_status !== 'PAID')
  )
}

// Event yang sedang aktif dipilih
const currentEvent = computed(() => {
  if (selectedEventId.value === 'all') return null
  return events.value.find((e) => String(e.id) === String(selectedEventId.value)) || null
})

// Pemain lunas pada event yang sedang dipilih
const currentEventLunas = computed(() => {
  if (!currentEvent.value) return []
  return getPemainLunas(currentEvent.value.id)
})

// Pemain waiting / belum lunas pada event yang sedang dipilih
const currentEventWaiting = computed(() => {
  if (!currentEvent.value) return []
  return getPemainWaiting(currentEvent.value.id)
})

// Total seluruh pemain lunas di semua sesi
const totalPemainLunas = computed(() => {
  if (!registrations.value) return 0
  return registrations.value.filter(
    (reg) => reg.polling_status === 'JOIN' && reg.payment_status === 'PAID'
  ).length
})

// Helper format tanggal & jam
const formatShortDate = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { weekday: 'short', day: 'numeric', month: 'short' })
}

const formatFullDate = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { weekday: 'short', day: 'numeric', month: 'short', year: 'numeric' })
}

const formatTime = (dateStr) => {
  if (!dateStr) return '--:--'
  const d = new Date(dateStr)
  return d.toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
}
</script>

<style scoped>
.scrollbar-none::-webkit-scrollbar {
  display: none;
}
.scrollbar-none {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.animate-fade-in {
  animation: fadeIn 0.25s ease-out forwards;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(4px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>