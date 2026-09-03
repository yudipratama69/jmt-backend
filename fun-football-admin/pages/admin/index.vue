<template>
  <div class="space-y-8 w-full max-w-7xl mx-auto px-2 sm:px-4 pb-16 transition-colors duration-300">
    
    <!-- ================================================================= -->
    <!-- HEADER & GREETING COMMAND CENTER                                  -->
    <!-- ================================================================= -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 border-b theme-border-subtle pb-6">
      <div>
        <div class="flex items-center gap-2 mb-1">
          <span class="w-8 h-8 rounded-xl bg-gradient-to-tr from-red-600 to-orange-500 text-white flex items-center justify-center shadow-md shadow-orange-600/30 text-base">
            <Icon name="ph:squares-four-bold" />
          </span>
          <h1 class="text-2xl font-black theme-text-main">Command Center Pengurus</h1>
          <span class="text-[10px] uppercase tracking-wider font-extrabold px-2.5 py-0.5 rounded-full bg-orange-500/10 text-orange-500 border border-orange-500/30">
            JMT Sport Live
          </span>
        </div>
        <p class="text-xs theme-text-muted">
          Pantau seluruh aktivitas fun football, keterisian kuota jadwal, antrean verifikasi, dan arus kas komunitas.
        </p>
      </div>

      <!-- Quick Action Buttons -->
      <div class="flex items-center gap-2.5 flex-wrap">
        <NuxtLink 
          to="/verifikasi" 
          class="px-3.5 py-2 rounded-xl theme-bg-surface border theme-border hover:border-orange-500 text-xs font-bold theme-text-main transition shadow-sm flex items-center gap-1.5 active:scale-95">
          <Icon name="ph:shield-check-bold" class="text-base text-orange-500" />
          <span>Pusat Verifikasi</span>
          <span v-if="totalPendingQueue > 0" class="w-2 h-2 rounded-full bg-orange-500 animate-pulse"></span>
        </NuxtLink>

        <NuxtLink 
          to="/admin-keuangan" 
          class="px-3.5 py-2 rounded-xl theme-bg-surface border theme-border hover:border-emerald-500 text-xs font-bold theme-text-main transition shadow-sm flex items-center gap-1.5 active:scale-95">
          <Icon name="ph:wallet-bold" class="text-base text-emerald-500" />
          <span>Buku Kas</span>
        </NuxtLink>

        <NuxtLink 
          to="/jadwal" 
          class="px-4 py-2 rounded-xl bg-gradient-to-r from-red-600 to-orange-500 hover:from-red-700 hover:to-orange-700 text-white text-xs font-black shadow-md shadow-orange-600/20 transition flex items-center gap-1.5 active:scale-95">
          <Icon name="ph:plus-bold" class="text-base" />
          <span>+ Buat Jadwal</span>
        </NuxtLink>
      </div>
    </div>

    <!-- ================================================================= -->
    <!-- 4 KARTU STATISTIK UTAMA (EXECUTIVE KPI CARDS)                     -->
    <!-- ================================================================= -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      
      <!-- Kartu 1: Jadwal Pertandingan Aktif -->
      <NuxtLink 
        to="/jadwal" 
        class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm relative overflow-hidden flex flex-col justify-between group hover:border-orange-500/50 transition-all duration-300">
        <div class="flex justify-between items-start">
          <div class="space-y-1">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-orange-500 bg-orange-500/10 px-2.5 py-0.5 rounded-full border border-orange-500/20">
              📅 Agenda Match
            </span>
            <p class="text-xs theme-text-muted mt-1 font-medium">Jadwal Terdaftar</p>
          </div>
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-red-600 to-orange-500 text-white flex items-center justify-center shadow-md group-hover:scale-110 transition-transform">
            <Icon name="ph:calendar-blank-bold" class="text-xl" />
          </div>
        </div>
        <div class="mt-4">
          <p class="text-2xl font-black theme-text-main tracking-tight">
            {{ stats?.total_events || 0 }} <span class="text-xs font-semibold theme-text-muted">Sesi Jadwal</span>
          </p>
          <p class="text-[11px] text-orange-500 font-bold mt-1 flex items-center gap-1">
            <span>Kelola Sesi & Kuota →</span>
          </p>
        </div>
      </NuxtLink>

      <!-- Kartu 2: Partisipasi Pemain -->
      <NuxtLink 
        to="/member" 
        class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm relative overflow-hidden flex flex-col justify-between group hover:border-blue-500/50 transition-all duration-300">
        <div class="flex justify-between items-start">
          <div class="space-y-1">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-blue-500 bg-blue-500/10 px-2.5 py-0.5 rounded-full border border-blue-500/20">
              👥 Member & Squad
            </span>
            <p class="text-xs theme-text-muted mt-1 font-medium">Total Pemain (Ikut)</p>
          </div>
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-blue-600 to-indigo-500 text-white flex items-center justify-center shadow-md group-hover:scale-110 transition-transform">
            <Icon name="ph:users-three-bold" class="text-xl" />
          </div>
        </div>
        <div class="mt-4">
          <p class="text-2xl font-black theme-text-main tracking-tight">
            {{ stats?.total_players || 0 }} <span class="text-xs font-semibold theme-text-muted">Pemain Terdata</span>
          </p>
          <p class="text-[11px] text-blue-500 font-bold mt-1 flex items-center gap-1">
            <span>Lihat Member & Squad →</span>
          </p>
        </div>
      </NuxtLink>

      <!-- Kartu 3: Antrean Verifikasi -->
      <NuxtLink 
        to="/verifikasi" 
        class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm relative overflow-hidden flex flex-col justify-between group hover:border-amber-500/50 transition-all duration-300">
        <div class="flex justify-between items-start">
          <div class="space-y-1">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-amber-500 bg-amber-500/10 px-2.5 py-0.5 rounded-full border border-amber-500/20">
              ⏳ Butuh Verifikasi
            </span>
            <p class="text-xs theme-text-muted mt-1 font-medium">Tiket & Deposit</p>
          </div>
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-amber-500 to-yellow-500 text-white flex items-center justify-center shadow-md group-hover:scale-110 transition-transform">
            <Icon name="ph:receipt-bold" class="text-xl" />
          </div>
        </div>
        <div class="mt-4">
          <p class="text-2xl font-black tracking-tight" :class="totalPendingQueue > 0 ? 'text-amber-500' : 'theme-text-main'">
            {{ totalPendingQueue }} <span class="text-xs font-semibold theme-text-muted">Antrean Menunggu</span>
          </p>
          <p class="text-[11px] font-bold mt-1 flex items-center gap-1" :class="totalPendingQueue > 0 ? 'text-amber-500' : 'text-emerald-500'">
            <Icon :name="totalPendingQueue > 0 ? 'ph:warning-circle-bold' : 'ph:check-circle-bold'" />
            <span>{{ totalPendingQueue > 0 ? 'Perlu Di-Approve' : 'Semua Terverifikasi' }}</span>
          </p>
        </div>
      </NuxtLink>

      <!-- Kartu 4: Saldo Kas Komunitas -->
      <NuxtLink 
        to="/admin-keuangan" 
        class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm relative overflow-hidden flex flex-col justify-between group hover:border-emerald-500/50 transition-all duration-300">
        <div class="flex justify-between items-start">
          <div class="space-y-1">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-emerald-500 bg-emerald-500/10 px-2.5 py-0.5 rounded-full border border-emerald-500/20">
              💰 Kas Komunitas
            </span>
            <p class="text-xs theme-text-muted mt-1 font-medium">Saldo Kas Berjalan</p>
          </div>
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-emerald-600 to-teal-500 text-white flex items-center justify-center shadow-md group-hover:scale-110 transition-transform">
            <Icon name="ph:wallet-bold" class="text-xl" />
          </div>
        </div>
        <div class="mt-4">
          <p class="text-2xl font-black text-emerald-500 tracking-tight">
            Rp {{ (totalSaldoKas || 0).toLocaleString('id-ID') }}
          </p>
          <p class="text-[11px] text-emerald-500 font-bold mt-1 flex items-center gap-1">
            <span>Buka Laporan Kas →</span>
          </p>
        </div>
      </NuxtLink>

    </div>

    <!-- ================================================================= -->
    <!-- SOROTAN JADWAL TERDEKAT (UPCOMING MATCH BANNER)                   -->
    <!-- ================================================================= -->
    <div v-if="upcomingEvent" class="theme-bg-surface rounded-3xl border theme-border p-6 shadow-sm relative overflow-hidden transition-colors duration-300">
      
      <div class="flex flex-col lg:flex-row lg:items-center justify-between gap-6">
        
        <!-- Info Sesi Terdekat -->
        <div class="space-y-2 flex-1">
          <div class="flex items-center gap-2 flex-wrap">
            <span class="px-3 py-1 bg-gradient-to-r from-red-600 to-orange-500 text-white rounded-full text-[10px] font-black uppercase tracking-wider flex items-center gap-1.5 shadow-sm">
              <span class="w-1.5 h-1.5 rounded-full bg-white animate-pulse"></span>
              Sesi Terdekat
            </span>
            <span 
              v-if="upcomingEvent.status === 'OPEN'" 
              class="px-2.5 py-0.5 bg-emerald-500/10 text-emerald-500 rounded-full text-[10px] font-black border border-emerald-500/30">
              PENDAFTARAN BUKA
            </span>
            <span 
              v-else 
              class="px-2.5 py-0.5 theme-bg-card theme-text-muted rounded-full text-[10px] font-bold border theme-border">
              {{ upcomingEvent.status }}
            </span>
          </div>

          <h3 class="text-xl font-black theme-text-main">
            {{ upcomingEvent.title }}
          </h3>

          <div class="flex flex-wrap items-center gap-4 text-xs theme-text-muted pt-1">
            <div class="flex items-center gap-1.5">
              <Icon name="ph:calendar-blank-bold" class="text-orange-500 text-base" />
              <span class="font-bold theme-text-main">{{ formatFullDate(upcomingEvent.match_date) }}</span>
            </div>
            <div class="flex items-center gap-1.5">
              <Icon name="ph:clock-bold" class="text-orange-500 text-base" />
              <span class="font-bold theme-text-main">{{ formatTime(upcomingEvent.match_date) }} WIB</span>
            </div>
            <div class="flex items-center gap-1.5">
              <Icon name="ph:map-pin-bold" class="text-orange-500 text-base" />
              <span>{{ upcomingEvent.location }}</span>
            </div>
            <div class="flex items-center gap-1.5">
              <Icon name="ph:tag-bold" class="text-emerald-500 text-base" />
              <span class="font-black text-emerald-500">Rp {{ (upcomingEvent.price_per_person || 0).toLocaleString('id-ID') }} / org</span>
            </div>
          </div>
        </div>

        <!-- Progress Kuota & Action Shortcuts -->
        <div class="lg:w-80 space-y-3 theme-bg-card p-4 rounded-2xl border theme-border-subtle">
          <div class="flex justify-between items-center text-xs font-bold">
            <span class="theme-text-main flex items-center gap-1.5">
              <Icon name="ph:users-three-bold" class="text-orange-500" />
              <span>{{ upcomingEvent.registered_count || 0 }} / {{ upcomingEvent.quota_max }} Slot Terisi</span>
            </span>
            <span class="text-emerald-500">{{ upcomingEvent.paid_count || 0 }} Lunas</span>
          </div>

          <!-- Progress Bar -->
          <div class="w-full theme-bg-surface rounded-full h-2.5 overflow-hidden p-0.5 border theme-border">
            <div 
              class="h-full rounded-full bg-gradient-to-r from-amber-500 to-emerald-400 transition-all duration-300"
              :style="{ width: `${Math.min(100, Math.round(((upcomingEvent.registered_count || 0) / upcomingEvent.quota_max) * 100))}%` }">
            </div>
          </div>

          <div class="flex items-center gap-2 pt-1">
            <NuxtLink 
              :to="'/member?tab=squad&event_id=' + upcomingEvent.id" 
              class="flex-1 text-center py-2 bg-gradient-to-r from-red-600 to-orange-500 hover:from-red-700 hover:to-orange-700 text-white rounded-xl text-xs font-black shadow-sm transition active:scale-95">
              Lihat Squad
            </NuxtLink>
            <NuxtLink 
              to="/jadwal" 
              class="px-3.5 py-2 theme-bg-surface border theme-border hover:border-orange-500 theme-text-main rounded-xl text-xs font-bold transition">
              Edit Jadwal
            </NuxtLink>
          </div>
        </div>

      </div>

    </div>

    <!-- ================================================================= -->
    <!-- GRID 2 KOLOM: DAFTAR JADWAL (KIRI) & FEED VERIFIKASI & KAS (KANAN)-->
    <!-- ================================================================= -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-6 items-start">
      
      <!-- KOLOM KIRI (7/12): DAFTAR SELURUH JADWAL PERTANDINGAN -->
      <div class="lg:col-span-7 theme-bg-surface rounded-3xl shadow-sm border theme-border overflow-hidden transition-colors duration-300">
        
        <div class="p-5 border-b theme-border theme-bg-card flex justify-between items-center">
          <div>
            <h3 class="font-black theme-text-main text-base flex items-center gap-2">
              <Icon name="ph:calendar-check-bold" class="text-orange-500 text-lg" />
              Daftar Jadwal Pertandingan
            </h3>
            <p class="text-xs theme-text-muted mt-0.5">Pantau status pendaftaran dan partisipasi kuota setiap match.</p>
          </div>
          <NuxtLink to="/jadwal" class="text-xs font-bold text-orange-500 hover:underline flex items-center gap-1">
            Kelola Semua →
          </NuxtLink>
        </div>

        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="theme-bg-card theme-text-muted text-[11px] uppercase tracking-wider font-extrabold">
                <th class="p-4 border-b theme-border">Sesi & Lokasi</th>
                <th class="p-4 border-b theme-border">Waktu</th>
                <th class="p-4 border-b theme-border">Keterisian</th>
                <th class="p-4 border-b theme-border">Tarif</th>
                <th class="p-4 border-b theme-border text-center">Status</th>
              </tr>
            </thead>
            <tbody class="text-xs divide-y theme-border-subtle">
              <tr 
                v-for="event in eventsData?.data" 
                :key="event.id" 
                class="hover:bg-orange-500/5 transition">
                
                <!-- Judul & Lokasi -->
                <td class="p-4">
                  <p class="font-black theme-text-main text-xs">{{ event.title }}</p>
                  <p class="text-[11px] theme-text-muted flex items-center gap-1 mt-0.5">
                    <Icon name="ph:map-pin-bold" class="text-orange-500" />
                    <span class="truncate max-w-[180px]">{{ event.location }}</span>
                  </p>
                </td>

                <!-- Tanggal & Jam -->
                <td class="p-4 whitespace-nowrap">
                  <p class="font-bold theme-text-main">{{ formatShortDate(event.match_date) }}</p>
                  <p class="text-[10px] theme-text-muted">{{ formatTime(event.match_date) }} WIB</p>
                </td>

                <!-- Progress Kuota -->
                <td class="p-4 min-w-[120px]">
                  <div class="space-y-1">
                    <div class="flex justify-between text-[10px] font-bold">
                      <span class="theme-text-main">{{ event.registered_count || 0 }}/{{ event.quota_max }}</span>
                      <span class="text-emerald-500">{{ event.paid_count || 0 }} Lunas</span>
                    </div>
                    <div class="w-full theme-bg-card rounded-full h-1.5 overflow-hidden border theme-border">
                      <div 
                        class="h-full rounded-full bg-gradient-to-r from-amber-500 to-emerald-400"
                        :style="{ width: `${Math.min(100, Math.round(((event.registered_count || 0) / event.quota_max) * 100))}%` }">
                      </div>
                    </div>
                  </div>
                </td>

                <!-- Harga -->
                <td class="p-4 font-black theme-text-main whitespace-nowrap">
                  Rp {{ (event.price_per_person || 0).toLocaleString('id-ID') }}
                </td>

                <!-- Status -->
                <td class="p-4 text-center whitespace-nowrap">
                  <span 
                    v-if="event.status === 'OPEN'" 
                    class="px-2.5 py-0.5 bg-emerald-500/10 text-emerald-500 rounded-full text-[10px] font-black border border-emerald-500/30">
                    OPEN
                  </span>
                  <span 
                    v-else-if="event.status === 'FULL'" 
                    class="px-2.5 py-0.5 bg-red-500/10 text-red-500 rounded-full text-[10px] font-black border border-red-500/30">
                    FULL
                  </span>
                  <span 
                    v-else 
                    class="px-2.5 py-0.5 theme-bg-card theme-text-muted rounded-full text-[10px] font-bold border theme-border">
                    {{ event.status }}
                  </span>
                </td>

              </tr>

              <!-- Jika Kosong -->
              <tr v-if="!eventsData?.data || eventsData.data.length === 0">
                <td colspan="5" class="p-10 text-center theme-text-muted space-y-2">
                  <Icon name="ph:calendar-blank-bold" class="text-3xl mx-auto opacity-40" />
                  <p class="font-bold theme-text-main text-xs">Belum ada jadwal yang dibuat.</p>
                  <NuxtLink to="/jadwal" class="text-xs text-orange-500 font-bold hover:underline inline-block">
                    + Buat Jadwal Pertama
                  </NuxtLink>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

      </div>

      <!-- KOLOM KANAN (5/12): FEED VERIFIKASI & RINGKASAN KAS -->
      <div class="lg:col-span-5 space-y-6">
        
        <!-- KOTAK 1: FEED ANTREAN VERIFIKASI CEPAT -->
        <div class="theme-bg-surface rounded-3xl shadow-sm border theme-border p-5 space-y-4 transition-colors duration-300">
          <div class="flex justify-between items-center border-b theme-border-subtle pb-3">
            <h3 class="font-black theme-text-main text-sm flex items-center gap-2">
              <Icon name="ph:shield-check-bold" class="text-orange-500 text-base" />
              Antrean Verifikasi Terbaru
            </h3>
            <NuxtLink to="/verifikasi" class="text-xs font-bold text-orange-500 hover:underline">
              Buka Semua →
            </NuxtLink>
          </div>

          <!-- List Pending Queue Items -->
          <div class="space-y-2.5">
            
            <!-- Pending Registrations (Tiket) -->
            <div 
              v-for="reg in pendingTicketsList.slice(0, 3)" 
              :key="'pending-reg-' + reg.id"
              class="p-3 rounded-2xl theme-bg-card border theme-border-subtle flex items-center justify-between gap-3">
              <div class="flex items-center gap-2.5 overflow-hidden">
                <span class="w-8 h-8 rounded-xl bg-orange-500/10 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0">
                  🎫
                </span>
                <div class="truncate">
                  <p class="font-bold theme-text-main text-xs truncate">{{ reg.user_name || 'Member' }}</p>
                  <p class="text-[10px] theme-text-muted">Tiket Transfer Main</p>
                </div>
              </div>
              <NuxtLink 
                to="/verifikasi?tab=tiket" 
                class="px-2.5 py-1 bg-orange-500/10 text-orange-500 hover:bg-orange-500/20 rounded-xl text-[10px] font-black transition shrink-0">
                Periksa
              </NuxtLink>
            </div>

            <!-- Pending Topups (Deposit) -->
            <div 
              v-for="topup in pendingTopupsList.slice(0, 3)" 
              :key="'pending-topup-' + topup._id"
              class="p-3 rounded-2xl theme-bg-card border theme-border-subtle flex items-center justify-between gap-3">
              <div class="flex items-center gap-2.5 overflow-hidden">
                <span class="w-8 h-8 rounded-xl bg-teal-500/10 text-teal-500 flex items-center justify-center font-bold text-xs shrink-0">
                  💳
                </span>
                <div class="truncate">
                  <p class="font-bold theme-text-main text-xs truncate">{{ topup.user_name || 'Member' }}</p>
                  <p class="text-[10px] theme-text-muted">Deposit Rp {{ (topup.amount || 0).toLocaleString('id-ID') }}</p>
                </div>
              </div>
              <NuxtLink 
                to="/verifikasi?tab=deposit" 
                class="px-2.5 py-1 bg-teal-500/10 text-teal-500 hover:bg-teal-500/20 rounded-xl text-[10px] font-black transition shrink-0">
                Periksa
              </NuxtLink>
            </div>

            <!-- Jika Semua Bersih -->
            <div v-if="totalPendingQueue === 0" class="p-6 text-center space-y-1.5 theme-bg-card rounded-2xl border theme-border-subtle">
              <Icon name="ph:seal-check-bold" class="text-3xl text-emerald-500 mx-auto" />
              <p class="text-xs font-bold theme-text-main">Semua Antrean Bersih!</p>
              <p class="text-[10px] theme-text-muted">Tidak ada tiket atau deposit yang menunggu persetujuan.</p>
            </div>

          </div>
        </div>

        <!-- KOTAK 2: FINANCIAL ARUS KAS SNAPSHOT -->
        <div class="theme-bg-surface rounded-3xl shadow-sm border theme-border p-5 space-y-4 transition-colors duration-300">
          <div class="flex justify-between items-center border-b theme-border-subtle pb-3">
            <h3 class="font-black theme-text-main text-sm flex items-center gap-2">
              <Icon name="ph:chart-pie-slice-bold" class="text-emerald-500 text-base" />
              Ringkasan Arus Kas
            </h3>
            <NuxtLink to="/admin-keuangan" class="text-xs font-bold text-orange-500 hover:underline">
              Buku Kas →
            </NuxtLink>
          </div>

          <div class="grid grid-cols-2 gap-3 text-xs">
            <div class="p-3 rounded-2xl theme-bg-card border theme-border-subtle">
              <p class="text-[10px] theme-text-muted uppercase font-bold tracking-wider">Total Kas Masuk</p>
              <p class="text-sm font-black text-emerald-500 mt-0.5">+ Rp {{ (totalMasukKas || 0).toLocaleString('id-ID') }}</p>
            </div>
            <div class="p-3 rounded-2xl theme-bg-card border theme-border-subtle">
              <p class="text-[10px] theme-text-muted uppercase font-bold tracking-wider">Total Beban Keluar</p>
              <p class="text-sm font-black text-rose-500 mt-0.5">- Rp {{ (totalKeluarKas || 0).toLocaleString('id-ID') }}</p>
            </div>
          </div>

          <div class="p-3.5 rounded-2xl bg-gradient-to-r from-orange-500/10 via-red-500/10 to-amber-500/10 border border-orange-500/20 flex justify-between items-center">
            <div>
              <p class="text-[10px] uppercase tracking-wider font-extrabold theme-text-muted">Saldo Kas Tersedia</p>
              <p class="text-base font-black text-orange-500">Rp {{ (totalSaldoKas || 0).toLocaleString('id-ID') }}</p>
            </div>
            <NuxtLink 
              to="/admin-keuangan" 
              class="px-3 py-1.5 bg-orange-500 text-white rounded-xl text-xs font-bold shadow transition active:scale-95">
              Kelola Kas
            </NuxtLink>
          </div>
        </div>

      </div>

    </div>

  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'

const { data: eventsData, refresh: refreshEvents } = await useApiFetch('/events')
const { data: stats, refresh: refreshStats } = await useApiFetch('/dashboard-stats')
const { data: regData, refresh: refreshReg } = await useApiFetch('/registrations')
const { data: topupData, refresh: refreshTopup } = await useApiFetch('/approved-topups')
const { data: pendingTopupsData, refresh: refreshPendingTopup } = await useApiFetch('/pending-topups')

const { useAutoRefresh } = useRealtime()

// Pasang Auto-Refresh Realtime
useAutoRefresh(['EVENT_UPDATED', 'REGISTRATION_UPDATED', 'PAYMENT_UPDATED', 'TOPUP_UPDATED'], () => {
  refreshEvents()
  refreshStats()
  refreshReg()
  refreshTopup()
  refreshPendingTopup()
})

const manualTransactions = ref([])
onMounted(() => {
  manualTransactions.value = JSON.parse(localStorage.getItem('jmt_manual_tx') || '[]')
})

// Antrean Tiket Menunggu
const pendingTicketsList = computed(() => {
  if (!regData.value?.data) return []
  return regData.value.data.filter(r => r.payment_status === 'VERIFYING')
})

// Antrean Deposit Menunggu
const pendingTopupsList = computed(() => {
  return pendingTopupsData.value?.data || []
})

const totalPendingQueue = computed(() => {
  return pendingTicketsList.value.length + pendingTopupsList.value.length
})

// Sesi Terdekat (Upcoming Event)
const upcomingEvent = computed(() => {
  if (!eventsData.value?.data || eventsData.value.data.length === 0) return null
  // Urutkan jadwal yang match_date nya paling mendekati masa sekarang
  const sorted = [...eventsData.value.data].sort((a, b) => new Date(a.match_date) - new Date(b.match_date))
  return sorted[0]
})

// Perhitungan Kas Ringkas
const autoRegistrations = computed(() => {
  if (!regData.value?.data || !eventsData.value?.data) return []
  return regData.value.data
    .filter(r => r.payment_status === 'PAID' && r.payment_method !== 'deposit')
    .map(r => {
      const evt = eventsData.value.data.find(e => e.id === r.event_id)
      return {
        amount: evt ? evt.price_per_person : 0
      }
    })
})

const autoTopups = computed(() => topupData.value?.data || [])

const totalMasukKas = computed(() => {
  const fromReg = autoRegistrations.value.reduce((acc, curr) => acc + (curr.amount || 0), 0)
  const fromTopup = autoTopups.value.reduce((acc, curr) => acc + (curr.amount || 0), 0)
  const fromManualIN = manualTransactions.value
    .filter(t => t.type === 'IN')
    .reduce((acc, curr) => acc + Number(curr.amount || 0), 0)
  return fromReg + fromTopup + fromManualIN
})

const totalKeluarKas = computed(() => {
  return manualTransactions.value
    .filter(t => t.type === 'OUT')
    .reduce((acc, curr) => acc + Number(curr.amount || 0), 0)
})

const totalSaldoKas = computed(() => totalMasukKas.value - totalKeluarKas.value)

// Format Helpers
const formatFullDate = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' })
}

const formatShortDate = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

const formatTime = (dateStr) => {
  if (!dateStr) return '--:--'
  const d = new Date(dateStr)
  return d.toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
}
</script>

