<template>
  <div class="space-y-8 max-w-7xl mx-auto pb-12 transition-colors duration-300">
    
    <!-- Header Page -->
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-black theme-text-main flex items-center gap-2.5">
          <span class="w-9 h-9 rounded-xl bg-gradient-to-tr from-red-600 to-orange-500 text-white flex items-center justify-center shadow-md shadow-orange-600/30 text-lg">
            <Icon name="ph:calendar-plus-bold" />
          </span>
          Manajemen Jadwal Pertandingan
        </h1>
        <p class="text-xs theme-text-muted mt-1">
          Buat jadwal baru dengan template cepat, atur batas waktu bayar, dan pantau keterisian slot kuota pemain.
        </p>
      </div>

      <div class="flex items-center gap-2">
        <span class="px-3.5 py-1.5 theme-bg-surface border theme-border rounded-xl text-xs font-bold theme-text-main flex items-center gap-1.5 shadow-sm">
          <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
          {{ events.length }} Jadwal Terdaftar
        </span>
      </div>
    </div>

    <!-- Form Input Jadwal (Full Width & Luas) -->
    <div class="theme-bg-surface rounded-3xl shadow-sm border theme-border p-6 sm:p-8 space-y-6 transition-colors duration-300">
      
      <!-- Form Header -->
      <div class="flex justify-between items-center border-b theme-border-subtle pb-4">
        <div>
          <h2 class="text-lg font-black theme-text-main flex items-center gap-2">
            <Icon :name="isEditMode ? 'ph:note-pencil-bold' : 'ph:sparkle-bold'" class="text-orange-500 text-xl" />
            {{ isEditMode ? 'Edit Informasi Jadwal' : 'Form Jadwal Main Baru' }}
          </h2>
          <p class="text-xs theme-text-muted mt-0.5">Pilih template instan di bawah atau atur data pertandingan secara kustom.</p>
        </div>

        <button 
          v-if="isEditMode" 
          @click="resetForm" 
          type="button"
          class="px-3.5 py-2 bg-red-500/10 text-red-500 hover:bg-red-500/20 rounded-xl text-xs font-bold transition flex items-center gap-1.5 active:scale-95">
          <Icon name="ph:x-bold" /> Batal Edit
        </button>
      </div>

      <!-- Template Pilihan Cepat (Quick Presets) -->
      <div class="space-y-2">
        <label class="block text-[11px] font-bold uppercase tracking-wider theme-text-muted">
          ⚡ Template Pilihan Cepat (Klik untuk Otomatis Isi):
        </label>
        <div class="grid grid-cols-2 sm:grid-cols-4 gap-3">
          <button 
            v-for="preset in presets" 
            :key="preset.name"
            type="button"
            @click="applyPreset(preset)"
            class="p-3.5 rounded-2xl border theme-border theme-bg-card hover:border-orange-500/60 text-left transition active:scale-95 group shadow-sm">
            <span class="text-2xl block mb-1.5">{{ preset.icon }}</span>
            <p class="text-xs font-bold theme-text-main group-hover:text-orange-500 transition leading-tight">{{ preset.name }}</p>
            <p class="text-[10px] theme-text-muted mt-1">{{ preset.timeDesc }}</p>
          </button>
        </div>
      </div>

      <!-- Form Input -->
      <form @submit.prevent="submitForm" class="space-y-5">
        
        <!-- Baris 1: Nama Jadwal & Lokasi Lapangan (2 Kolom) -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          
          <!-- Nama Jadwal -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Nama Jadwal Pertandingan <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <Icon name="ph:soccer-ball-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 text-orange-500 text-base" />
              <input 
                v-model="form.title" 
                type="text" 
                placeholder="Misal: Fun Football Pagi - Mini Soccer" 
                class="w-full pl-10 pr-4 py-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-medium transition" 
                required 
              />
            </div>
            
            <!-- Quick Title Tag Pills -->
            <div class="flex flex-wrap gap-1.5 pt-0.5">
              <button 
                v-for="tag in titleTags" 
                :key="tag"
                type="button"
                @click="form.title = tag"
                class="text-[10px] px-2.5 py-1 rounded-lg border theme-border-subtle theme-bg-card theme-text-muted hover:theme-text-main hover:border-orange-500/40 transition">
                + {{ tag }}
              </button>
            </div>
          </div>

          <!-- Lokasi Lapangan -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Lokasi Lapangan <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <Icon name="ph:map-pin-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 text-orange-500 text-base" />
              <input 
                v-model="form.location" 
                type="text" 
                placeholder="Misal: Lapangan Cozy Infinity - Cikutra" 
                class="w-full pl-10 pr-4 py-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-medium transition" 
                required 
              />
            </div>

            <!-- Quick Location Pills -->
            <div class="flex flex-wrap gap-1.5 pt-0.5">
              <button 
                v-for="loc in locationPresets" 
                :key="loc"
                type="button"
                @click="form.location = loc"
                class="text-[10px] px-2.5 py-1 rounded-lg border theme-border-subtle theme-bg-card theme-text-muted hover:theme-text-main hover:border-orange-500/40 transition">
                📍 {{ loc }}
              </button>
            </div>
          </div>

        </div>

        <!-- Baris 2: Tanggal Main & Batas Waktu Bayar (2 Kolom) -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          
          <!-- Tanggal & Jam Main -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Tanggal & Jam Kick-off <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <input 
                v-model="form.match_date" 
                type="datetime-local" 
                class="w-full p-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-medium transition" 
                required 
              />
            </div>
          </div>

          <!-- Batas Waktu Bayar -->
          <div class="space-y-2">
            <div class="flex justify-between items-center">
              <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
                Batas Waktu Bayar <span class="text-red-500">*</span>
              </label>
              <button 
                type="button" 
                @click="autoSetDeadline" 
                class="text-[11px] text-orange-500 font-bold hover:underline">
                ⚡ Auto H-1 Jam 18:00
              </button>
            </div>
            <div class="relative">
              <input 
                v-model="form.payment_deadline" 
                type="datetime-local" 
                class="w-full p-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-medium transition" 
                required 
              />
            </div>
          </div>

        </div>

        <!-- Baris 3: Kuota Maksimal & Harga Patungan (2 Kolom) -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          
          <!-- Kuota Pemain -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Kuota Maksimal Pemain <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <Icon name="ph:users-three-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 text-orange-500 text-base" />
              <input 
                v-model.number="form.quota_max" 
                type="number" 
                min="2" 
                max="100" 
                class="w-full pl-10 pr-4 py-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-bold transition" 
                required 
              />
            </div>
            <!-- Quick Quota Pills -->
            <div class="flex gap-1.5 pt-0.5">
              <button 
                v-for="q in [14, 16, 18, 20, 22, 28]" 
                :key="q"
                type="button"
                @click="form.quota_max = q"
                :class="form.quota_max === q ? 'bg-orange-500 text-white font-bold' : 'theme-bg-card theme-text-muted border theme-border-subtle'"
                class="text-[11px] px-2.5 py-0.5 rounded-lg transition">
                {{ q }} Pemain
              </button>
            </div>
          </div>

          <!-- Harga Patungan -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Harga Patungan / Orang <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <span class="absolute left-3.5 top-1/2 -translate-y-1/2 text-orange-500 font-bold text-xs">Rp</span>
              <input 
                v-model.number="form.price_per_person" 
                type="number" 
                step="5000" 
                class="w-full pl-10 pr-4 py-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-bold transition" 
                required 
              />
            </div>
            <!-- Quick Price Pills -->
            <div class="flex flex-wrap gap-1.5 pt-0.5">
              <button 
                v-for="p in [35000, 40000, 50000, 65000, 75000]" 
                :key="p"
                type="button"
                @click="form.price_per_person = p"
                :class="form.price_per_person === p ? 'bg-orange-500 text-white font-bold' : 'theme-bg-card theme-text-muted border theme-border-subtle'"
                class="text-[11px] px-2.5 py-0.5 rounded-lg transition">
                Rp {{ (p).toLocaleString('id-ID') }}
              </button>
            </div>
          </div>

        </div>

        <!-- Baris 4: Status Pendaftaran -->
        <div class="space-y-2 pt-1">
          <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
            Status Jadwal Pertandingan
          </label>
          <div class="grid grid-cols-3 gap-3">
            <button 
              type="button"
              @click="form.status = 'OPEN'"
              :class="form.status === 'OPEN' ? 'bg-emerald-500 text-white font-black shadow-md' : 'theme-bg-card border theme-border theme-text-muted'"
              class="py-3 px-4 rounded-2xl text-xs flex items-center justify-center gap-2 transition active:scale-95">
              <span class="w-2 h-2 rounded-full bg-emerald-300 animate-pulse" v-if="form.status === 'OPEN'"></span>
              <span>BUKA (OPEN)</span>
            </button>

            <button 
              type="button"
              @click="form.status = 'FULL'"
              :class="form.status === 'FULL' ? 'bg-red-500 text-white font-black shadow-md' : 'theme-bg-card border theme-border theme-text-muted'"
              class="py-3 px-4 rounded-2xl text-xs flex items-center justify-center gap-2 transition active:scale-95">
              <span>PENUH (FULL)</span>
            </button>

            <button 
              type="button"
              @click="form.status = 'CLOSED'"
              :class="form.status === 'CLOSED' ? 'bg-gray-600 text-white font-black shadow-md' : 'theme-bg-card border theme-border theme-text-muted'"
              class="py-3 px-4 rounded-2xl text-xs flex items-center justify-center gap-2 transition active:scale-95">
              <span>TUTUP (CLOSED)</span>
            </button>
          </div>
        </div>

        <!-- Tombol Simpan / Publikasikan -->
        <div class="pt-4">
          <button 
            type="submit" 
            class="w-full bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black py-4 px-6 rounded-2xl shadow-lg shadow-orange-600/30 transition active:scale-95 text-sm flex items-center justify-center gap-2">
            <Icon :name="isEditMode ? 'ph:check-circle-bold' : 'ph:paper-plane-tilt-bold'" class="text-lg" />
            <span>{{ isEditMode ? 'Simpan Perubahan Jadwal' : 'Simpan & Publikasikan Jadwal' }}</span>
          </button>
        </div>

      </form>

    </div>

    <!-- ================================================================= -->
    <!-- BAGIAN TABEL & DAFTAR SELURUH JADWAL PERTANDINGAN                 -->
    <!-- ================================================================= -->
    <div class="theme-bg-surface rounded-3xl shadow-sm border theme-border overflow-hidden transition-colors duration-300">
      
      <!-- Table Header & Search Filter -->
      <div class="p-5 sm:p-6 border-b theme-border theme-bg-card flex flex-col sm:flex-row sm:items-center justify-between gap-4">
        <div>
          <h3 class="font-black theme-text-main text-base flex items-center gap-2">
            <Icon name="ph:list-dashes-bold" class="text-orange-500 text-lg" />
            Daftar Seluruh Jadwal
          </h3>
          <p class="text-xs theme-text-muted mt-0.5">Kelola, edit, atau pantau partisipasi pemain setiap jadwal.</p>
        </div>

        <div class="flex items-center gap-2">
          <div class="relative w-full sm:w-64">
            <Icon name="ph:magnifying-glass-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 theme-text-muted text-xs" />
            <input 
              v-model="searchQuery" 
              type="text" 
              placeholder="Cari jadwal / lokasi..." 
              class="w-full pl-9 pr-3.5 py-2 border theme-border rounded-xl text-xs theme-text-main theme-bg-surface focus:border-orange-500 outline-none" 
            />
          </div>
        </div>
      </div>
      
      <!-- Table Content -->
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="theme-bg-card theme-text-muted text-[11px] uppercase tracking-wider font-bold">
              <th class="p-4 border-b theme-border">Jadwal & Lokasi</th>
              <th class="p-4 border-b theme-border">Tanggal & Jam</th>
              <th class="p-4 border-b theme-border">Keterisian Kuota</th>
              <th class="p-4 border-b theme-border">Harga</th>
              <th class="p-4 border-b theme-border">Status</th>
              <th class="p-4 border-b theme-border text-center">Aksi Pengurus</th>
            </tr>
          </thead>
          <tbody class="text-xs divide-y theme-border-subtle">
            <tr 
              v-for="event in filteredEvents" 
              :key="event.id" 
              class="hover:bg-orange-500/5 transition">
              
              <!-- Nama & Lokasi -->
              <td class="p-4">
                <div class="space-y-0.5">
                  <p class="theme-text-main font-black text-sm">{{ event.title }}</p>
                  <p class="text-[11px] theme-text-muted flex items-center gap-1">
                    <Icon name="ph:map-pin-bold" class="text-orange-500" />
                    {{ event.location }}
                  </p>
                </div>
              </td>

              <!-- Tanggal & Jam -->
              <td class="p-4">
                <div class="space-y-0.5">
                  <p class="font-bold theme-text-main">{{ formatFullDate(event.match_date) }}</p>
                  <p class="text-[11px] theme-text-muted font-medium">{{ formatTime(event.match_date) }} WIB</p>
                </div>
              </td>

              <!-- Keterisian Kuota Bar -->
              <td class="p-4 min-w-[150px]">
                <div class="space-y-1">
                  <div class="flex justify-between text-[10px] font-bold">
                    <span class="theme-text-main">{{ event.registered_count || 0 }}/{{ event.quota_max }} Terisi</span>
                    <span class="text-emerald-500">{{ event.paid_count || 0 }} Lunas</span>
                  </div>
                  <div class="w-full theme-bg-card rounded-full h-2 overflow-hidden p-0.5 border theme-border">
                    <div 
                      class="h-full rounded-full bg-gradient-to-r from-amber-500 to-emerald-400 transition-all duration-300"
                      :style="{ width: `${Math.min(100, Math.round(((event.registered_count || 0) / event.quota_max) * 100))}%` }">
                    </div>
                  </div>
                </div>
              </td>

              <!-- Harga -->
              <td class="p-4 theme-text-main font-black text-sm">
                Rp {{ (event.price_per_person || 0).toLocaleString('id-ID') }}
              </td>

              <!-- Status -->
              <td class="p-4">
                <span 
                  v-if="event.status === 'OPEN'" 
                  class="px-2.5 py-1 bg-emerald-500/10 text-emerald-500 rounded-full text-[10px] font-black border border-emerald-500/30">
                  OPEN
                </span>
                <span 
                  v-else-if="event.status === 'FULL'" 
                  class="px-2.5 py-1 bg-red-500/10 text-red-500 rounded-full text-[10px] font-black border border-red-500/30">
                  FULL
                </span>
                <span 
                  v-else 
                  class="px-2.5 py-1 theme-bg-card theme-text-muted rounded-full text-[10px] font-bold border theme-border">
                  {{ event.status }}
                </span>
              </td>

              <!-- Tombol Aksi -->
              <td class="p-4 text-center">
                <div class="flex items-center justify-center gap-2">
                  <button 
                    @click="pilihEdit(event)" 
                    title="Edit Jadwal"
                    class="px-3.5 py-1.5 bg-amber-500/10 text-amber-500 hover:bg-amber-500/20 rounded-xl text-xs font-bold transition flex items-center gap-1 active:scale-95">
                    <Icon name="ph:pencil-simple-bold" /> Edit
                  </button>
                  <button 
                    @click="hapusJadwal(event.id)" 
                    title="Hapus Jadwal"
                    class="px-3.5 py-1.5 bg-red-500/10 text-red-500 hover:bg-red-500/20 rounded-xl text-xs font-bold transition flex items-center gap-1 active:scale-95">
                    <Icon name="ph:trash-bold" /> Hapus
                  </button>
                </div>
              </td>

            </tr>

            <!-- Jika Kosong -->
            <tr v-if="filteredEvents.length === 0">
              <td colspan="6" class="p-12 text-center theme-text-muted space-y-2">
                <Icon name="ph:calendar-blank-bold" class="text-4xl mx-auto opacity-40" />
                <p class="font-bold theme-text-main text-sm">Belum Ada Jadwal Pertandingan</p>
                <p class="text-xs">Gunakan formulir di atas untuk membuat jadwal baru dengan mudah.</p>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

    </div>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const { $api } = useNuxtApp()
const { useAutoRefresh } = useRealtime()
const toast = useToast()

const { data: eventsData, refresh } = await useApiFetch('/events')

// Pasang Auto-Refresh Realtime
useAutoRefresh(['EVENT_UPDATED', 'REGISTRATION_UPDATED'], () => {
  refresh()
})

const events = computed(() => eventsData.value?.data || [])

// Preset Templates
const presets = [
  {
    name: 'Sesi Pagi',
    icon: '🌅',
    timeDesc: '07:00 - 09:00',
    title: 'Fun Football Pagi - Mini Soccer',
    location: 'Lapangan Cozy Infinity - Cikutra',
    timeHour: 7,
    quota: 18,
    price: 45000
  },
  {
    name: 'Sesi Sore',
    icon: '🌇',
    timeDesc: '16:00 - 18:00',
    title: 'Fun Football Sore - Sunset Match',
    location: 'Lapangan Cozy Infinity - Cikutra',
    timeHour: 16,
    quota: 18,
    price: 50000
  },
  {
    name: 'Sesi Malam',
    icon: '🌙',
    timeDesc: '19:30 - 21:30',
    title: 'Night Mini Soccer Match',
    location: 'Lapangan Cozy Infinity - Cikutra',
    timeHour: 19,
    timeMin: 30,
    quota: 20,
    price: 55000
  },
  {
    name: 'Weekend Big Match',
    icon: '🏆',
    timeDesc: 'Sabtu/Minggu',
    title: 'Weekend Super Mini Soccer',
    location: 'Lapangan Cozy Infinity - Cikutra',
    timeHour: 8,
    quota: 22,
    price: 50000
  }
]

const titleTags = [
  'Fun Football Pagi (7v7)',
  'Sunset Match Sore',
  'Night Mini Soccer (8v8)',
  'Weekend Match Super League'
]

const locationPresets = [
  'Lapangan Cozy Infinity - Cikutra',
  'Stadion Gelora Bung Karno',
  'Lapangan JMT Arena'
]

const isEditMode = ref(false)
const selectedEventId = ref(null)
const searchQuery = ref('')

const form = ref({
  title: '',
  location: 'Lapangan Cozy Infinity - Cikutra',
  match_date: '',
  quota_max: 18,
  price_per_person: 50000,
  payment_deadline: '',
  status: 'OPEN'
})

// Fungsi Pasang Preset Instan
const applyPreset = (preset) => {
  form.value.title = preset.title
  form.value.location = preset.location
  form.value.quota_max = preset.quota
  form.value.price_per_person = preset.price

  // Buat tanggal default besok dengan jam preset
  const now = new Date()
  const targetDate = new Date(now.getFullYear(), now.getMonth(), now.getDate() + 1, preset.timeHour, preset.timeMin || 0)
  
  // Format to local ISO string (YYYY-MM-DDTHH:mm)
  const pad = (n) => String(n).padStart(2, '0')
  const formattedMatch = `${targetDate.getFullYear()}-${pad(targetDate.getMonth() + 1)}-${pad(targetDate.getDate())}T${pad(targetDate.getHours())}:${pad(targetDate.getMinutes())}`
  form.value.match_date = formattedMatch

  // Batas bayar H-1 jam 18:00
  const deadlineDate = new Date(targetDate.getFullYear(), targetDate.getMonth(), targetDate.getDate() - 1, 18, 0)
  const formattedDeadline = `${deadlineDate.getFullYear()}-${pad(deadlineDate.getMonth() + 1)}-${pad(deadlineDate.getDate())}T${pad(deadlineDate.getHours())}:${pad(deadlineDate.getMinutes())}`
  form.value.payment_deadline = formattedDeadline

  toast.info(`Template ${preset.name} berhasil diterapkan!`)
}

// Otomatis atur deadline pembayaran (H-1 jam 18:00)
const autoSetDeadline = () => {
  if (!form.value.match_date) {
    toast.error('Silakan pilih Tanggal & Jam Main terlebih dahulu!')
    return
  }
  const matchDate = new Date(form.value.match_date)
  const deadlineDate = new Date(matchDate.getFullYear(), matchDate.getMonth(), matchDate.getDate() - 1, 18, 0)
  
  const pad = (n) => String(n).padStart(2, '0')
  form.value.payment_deadline = `${deadlineDate.getFullYear()}-${pad(deadlineDate.getMonth() + 1)}-${pad(deadlineDate.getDate())}T${pad(deadlineDate.getHours())}:${pad(deadlineDate.getMinutes())}`
  toast.success('Batas bayar diatur ke H-1 jam 18:00!')
}

const pilihEdit = (event) => {
  isEditMode.value = true
  selectedEventId.value = event.id
  
  const formatDateTime = (dateStr) => {
    if (!dateStr) return ''
    const d = new Date(dateStr)
    const pad = (n) => String(n).padStart(2, '0')
    return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}T${pad(d.getHours())}:${pad(d.getMinutes())}`
  }

  form.value = {
    title: event.title,
    location: event.location,
    match_date: formatDateTime(event.match_date),
    payment_deadline: formatDateTime(event.payment_deadline),
    quota_max: event.quota_max,
    price_per_person: event.price_per_person,
    status: event.status || 'OPEN'
  }
  
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const resetForm = () => {
  isEditMode.value = false
  selectedEventId.value = null
  form.value = {
    title: '',
    location: 'Lapangan Cozy Infinity - Cikutra',
    match_date: '',
    quota_max: 18,
    price_per_person: 50000,
    payment_deadline: '',
    status: 'OPEN'
  }
}

const submitForm = async () => {
  try {
    const payload = {
      ...form.value,
      match_date: new Date(form.value.match_date).toISOString(),
      payment_deadline: new Date(form.value.payment_deadline).toISOString(),
      status: form.value.status || 'OPEN'
    }

    if (isEditMode.value) {
      await $api(`/events/${selectedEventId.value}`, {
        method: 'PUT',
        body: payload
      })
      toast.success('Jadwal pertandingan berhasil diperbarui!', 'Sukses Update')
    } else {
      await $api('/events', {
        method: 'POST',
        body: payload
      })
      toast.success('Jadwal baru berhasil dipublikasikan!', 'Jadwal Rilis')
    }

    resetForm()
    refresh()
  } catch (error) {
    toast.error('Terjadi kesalahan saat menyimpan data.', 'Error')
  }
}

const hapusJadwal = (id) => {
  toast.confirm({
    title: 'Hapus Jadwal Pertandingan',
    message: 'Apakah Anda yakin ingin menghapus jadwal pertandingan ini? Seluruh data tiket yang terkait akan dibersihkan.',
    confirmText: 'Ya, Hapus',
    cancelText: 'Batal',
    onConfirm: async () => {
      try {
        await $api(`/events/${id}`, {
          method: 'DELETE'
        })
        toast.success('Jadwal pertandingan berhasil dihapus!', 'Terhapus')
        refresh()
      } catch (error) {
        toast.error('Gagal menghapus jadwal.', 'Error')
      }
    }
  })
}

// Filter Pencarian
const filteredEvents = computed(() => {
  if (!searchQuery.value) return events.value
  const q = searchQuery.value.toLowerCase()
  return events.value.filter((ev) => {
    return (ev.title && ev.title.toLowerCase().includes(q)) || 
           (ev.location && ev.location.toLowerCase().includes(q))
  })
})

// Format Helpers
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