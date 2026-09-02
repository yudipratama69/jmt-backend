<template>
  <div class="space-y-8 w-full px-4 md:px-8 pb-10 transition-colors duration-300">
    
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-extrabold theme-text-main">Ringkasan Dashboard</h1>
        <p class="text-xs theme-text-muted mt-0.5">Statistik utama aktivitas dan partisipasi JMT Sport.</p>
      </div>
      <NuxtLink to="/jadwal" class="inline-flex items-center gap-2 px-4 py-2.5 bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white rounded-xl text-xs font-black shadow-md transition active:scale-95 w-max">
        <Icon name="ph:plus-bold" class="text-sm" /> Buat Jadwal Baru
      </NuxtLink>
    </div>
    
    <!-- Kartu Statistik Interaktif -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      
      <!-- Kartu 1: Jadwal Terdaftar -->
      <NuxtLink to="/jadwal" style="background: linear-gradient(135deg, #ea580c, #dc2626);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40 group hover:shadow-2xl transition-all">
        <div class="absolute right-4 -bottom-4 text-white/15 pointer-events-none group-hover:scale-110 transition-transform">
          <Icon name="ph:calendar-blank-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            📅 Jadwal
          </div>
          <h3 class="text-orange-100 text-xs font-bold uppercase tracking-wider">Jadwal Terdaftar</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">{{ stats?.total_events || 0 }} Agenda</p>
      </NuxtLink>

      <!-- Kartu 2: Total Pemain -->
      <NuxtLink to="/verifikasi" style="background: linear-gradient(135deg, #2563eb, #1d4ed8);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40 group hover:shadow-2xl transition-all">
        <div class="absolute right-4 -bottom-4 text-white/15 pointer-events-none group-hover:scale-110 transition-transform">
          <Icon name="ph:users-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            👥 Partisipasi
          </div>
          <h3 class="text-blue-100 text-xs font-bold uppercase tracking-wider">Total Pemain (Ikut)</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">{{ stats?.total_players || 0 }} Orang</p>
      </NuxtLink>

      <!-- Kartu 3: Menunggu Verifikasi -->
      <NuxtLink to="/verifikasi" style="background: linear-gradient(135deg, #d97706, #b45309);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40 group hover:shadow-2xl transition-all">
        <div class="absolute right-4 -bottom-4 text-white/15 pointer-events-none group-hover:scale-110 transition-transform">
          <Icon name="ph:clock-countdown-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            ⏳ Verifikasi
          </div>
          <h3 class="text-amber-100 text-xs font-bold uppercase tracking-wider">Menunggu Verifikasi</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">{{ stats?.pending_verification || 0 }} Resi</p>
      </NuxtLink>

    </div>

    <!-- Tabel Daftar Jadwal Main -->
    <div class="theme-bg-surface rounded-2xl shadow-sm border theme-border overflow-hidden transition-colors duration-300">
      <div class="p-5 border-b theme-border flex justify-between items-center theme-bg-card">
        <h2 class="font-bold theme-text-main text-base">Daftar Jadwal Main</h2>
        <NuxtLink to="/jadwal" class="text-xs font-bold text-orange-500 hover:underline">Kelola Semua Jadwal →</NuxtLink>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="theme-bg-card theme-text-muted text-xs uppercase tracking-wider">
              <th class="p-4 font-bold border-b theme-border">Nama Jadwal</th>
              <th class="p-4 font-bold border-b theme-border">Tanggal</th>
              <th class="p-4 font-bold border-b theme-border">Lokasi</th>
              <th class="p-4 font-bold border-b theme-border">Harga</th>
              <th class="p-4 font-bold border-b theme-border">Status</th>
            </tr>
          </thead>
          <tbody class="text-sm">
            <tr v-for="event in eventsData?.data" :key="event.id" class="border-b theme-border-subtle hover:bg-orange-500/5 transition">
              <td class="p-4 theme-text-main font-bold">{{ event.title }}</td>
              <td class="p-4 theme-text-muted">{{ new Date(event.match_date).toLocaleDateString('id-ID') }}</td>
              <td class="p-4 theme-text-muted">{{ event.location }}</td>
              <td class="p-4 theme-text-main font-semibold">Rp {{ (event.price_per_person || 0).toLocaleString('id-ID') }}</td>
              <td class="p-4">
                <span class="px-3 py-1 bg-orange-500/10 text-orange-500 rounded-full text-xs font-extrabold border border-orange-500/20">
                  {{ event.status }}
                </span>
              </td>
            </tr>
            <tr v-if="!eventsData?.data || eventsData.data.length === 0">
              <td colspan="5" class="p-8 text-center theme-text-muted text-sm">Belum ada jadwal yang dibuat.</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

  </div>
</template>

<script setup>
const { data: eventsData, refresh: refreshEvents } = await useApiFetch('/events')
const { data: stats, refresh: refreshStats } = await useApiFetch('/dashboard-stats')

const { useAutoRefresh } = useRealtime()

// Pasang Auto-Refresh Realtime
useAutoRefresh(['EVENT_UPDATED', 'REGISTRATION_UPDATED', 'PAYMENT_UPDATED', 'TOPUP_UPDATED'], () => {
  refreshEvents()
  refreshStats()
})
</script>
