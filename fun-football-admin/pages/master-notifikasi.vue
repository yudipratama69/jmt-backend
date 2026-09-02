<template>
  <div class="space-y-6">
    <!-- Header Page -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-2xl font-black theme-text-main flex items-center gap-2">
          <Icon name="ph:broadcast-bold" class="text-orange-500 text-3xl" />
          Master Notifikasi & Broadcast
        </h1>
        <p class="text-xs theme-text-muted mt-1">
          Kirim pengumuman dan peringatan langsung ke seluruh pengguna aplikasi (PWA & Web) secara realtime.
        </p>
      </div>

      <div class="flex items-center gap-2">
        <span class="bg-emerald-500/15 text-emerald-500 border border-emerald-500/30 text-xs font-black px-3 py-1.5 rounded-xl flex items-center gap-1.5 shadow-sm">
          <span class="w-2 h-2 rounded-full bg-emerald-500 animate-ping"></span>
          Realtime WebSocket Hub Active
        </span>
      </div>
    </div>

    <!-- Grid: Form Composer & Live Preview -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      
      <!-- Kolom 1 & 2: Form Composer Broadcast -->
      <div class="lg:col-span-2 theme-bg-card backdrop-blur-md border theme-border rounded-3xl p-6 shadow-xl space-y-5 transition-colors duration-300">
        <h2 class="text-base font-black theme-text-main flex items-center gap-2 border-b theme-border-subtle pb-3">
          <Icon name="ph:paper-plane-tilt-bold" class="text-orange-500 text-lg" />
          Buat Pesan Pengumuman Baru
        </h2>

        <form @submit.prevent="kirimBroadcast" class="space-y-4">
          <!-- Kategori Notifikasi Selector -->
          <div>
            <label class="block text-xs font-bold theme-text-main mb-2">Kategori Notifikasi</label>
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-2">
              <button 
                type="button"
                @click="form.type = 'INFO'"
                :class="form.type === 'INFO' ? 'bg-blue-600 text-white font-black shadow-md' : 'theme-bg-surface theme-text-muted hover:theme-text-main border theme-border'"
                class="py-2.5 px-3 rounded-2xl text-xs transition flex items-center justify-center gap-1.5 active:scale-95">
                <Icon name="ph:info-bold" class="text-sm" /> 📢 Info Umum
              </button>
              <button 
                type="button"
                @click="form.type = 'JADWAL'"
                :class="form.type === 'JADWAL' ? 'bg-emerald-600 text-white font-black shadow-md' : 'theme-bg-surface theme-text-muted hover:theme-text-main border theme-border'"
                class="py-2.5 px-3 rounded-2xl text-xs transition flex items-center justify-center gap-1.5 active:scale-95">
                <Icon name="ph:soccer-ball-bold" class="text-sm" /> ⚽ Jadwal Main
              </button>
              <button 
                type="button"
                @click="form.type = 'URGENT'"
                :class="form.type === 'URGENT' ? 'bg-rose-600 text-white font-black shadow-md' : 'theme-bg-surface theme-text-muted hover:theme-text-main border theme-border'"
                class="py-2.5 px-3 rounded-2xl text-xs transition flex items-center justify-center gap-1.5 active:scale-95">
                <Icon name="ph:warning-circle-bold" class="text-sm" /> ⚠️ Penting
              </button>
              <button 
                type="button"
                @click="form.type = 'PROMO'"
                :class="form.type === 'PROMO' ? 'bg-amber-600 text-white font-black shadow-md' : 'theme-bg-surface theme-text-muted hover:theme-text-main border theme-border'"
                class="py-2.5 px-3 rounded-2xl text-xs transition flex items-center justify-center gap-1.5 active:scale-95">
                <Icon name="ph:gift-bold" class="text-sm" /> 🎉 Promo/Event
              </button>
            </div>
          </div>

          <!-- Judul Pengumuman -->
          <div>
            <label class="block text-xs font-bold theme-text-main mb-1.5">Judul Pengumuman</label>
            <input 
              v-model="form.title"
              type="text" 
              required
              placeholder="Contoh: Jadwal Mini Soccer Sabtu Ini Resmi Dibuka!" 
              class="w-full theme-bg-surface border theme-border rounded-2xl px-4 py-3 text-xs theme-text-main focus:outline-none focus:border-orange-500 transition" />
          </div>

          <!-- Isi Pesan -->
          <div>
            <label class="block text-xs font-bold theme-text-main mb-1.5">Isi Pesan / Keterangan</label>
            <textarea 
              v-model="form.message"
              rows="4" 
              required
              placeholder="Tuliskan detail informasi, jam kick-off, kuota pemain, atau info penting lainnya..." 
              class="w-full theme-bg-surface border theme-border rounded-2xl p-4 text-xs theme-text-main focus:outline-none focus:border-orange-500 transition resize-none"></textarea>
          </div>

          <!-- Tombol Kirim Broadcast -->
          <div class="pt-2">
            <button 
              type="submit" 
              :disabled="loading"
              class="w-full py-3.5 px-6 bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black text-sm rounded-2xl shadow-lg shadow-orange-600/30 transition active:scale-95 flex items-center justify-center gap-2">
              <Icon v-if="loading" name="ph:spinner-gap-bold" class="text-lg animate-spin" />
              <Icon v-else name="ph:broadcast-bold" class="text-lg" />
              <span>{{ loading ? 'Mengirim Broadcast...' : 'Kirim Broadcast ke Seluruh Pengguna Sekarang' }}</span>
            </button>
          </div>
        </form>
      </div>

      <!-- Kolom 3: Live Preview Notifikasi di HP -->
      <div class="theme-bg-card backdrop-blur-md border theme-border rounded-3xl p-6 shadow-xl space-y-4 flex flex-col justify-between transition-colors duration-300">
        <div>
          <h2 class="text-base font-black theme-text-main flex items-center gap-2 border-b theme-border-subtle pb-3">
            <Icon name="ph:device-mobile-bold" class="text-orange-500 text-lg" />
            Live Preview di HP Pemain
          </h2>

          <div class="mt-4 p-4 rounded-2xl theme-bg-surface border-2 border-orange-500/30 shadow-md space-y-3 relative overflow-hidden">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-2">
                <div class="w-7 h-7 rounded-lg bg-white p-0.5 shadow flex items-center justify-center shrink-0">
                  <img :src="'/logo-jmt.png'" alt="Logo" class="w-full h-full object-contain" />
                </div>
                <div>
                  <p class="font-black text-[11px] theme-text-main">JMT Sport</p>
                  <p class="text-[9px] theme-text-muted">Baru Saja</p>
                </div>
              </div>
              <span 
                :class="getTypeBadgeClass(form.type)"
                class="text-[9px] font-black px-2 py-0.5 rounded-full uppercase">
                {{ form.type || 'INFO' }}
              </span>
            </div>

            <div>
              <p class="font-black text-xs theme-text-main leading-snug">
                {{ form.title || 'Judul Pengumuman Akan Muncul di Sini...' }}
              </p>
              <p class="text-[11px] theme-text-muted mt-1 leading-relaxed whitespace-pre-line">
                {{ form.message || 'Isi pesan lengkap yang Anda tulis akan disiarkan dan langsung berbunyi di HP seluruh member komunitas.' }}
              </p>
            </div>
          </div>
        </div>

        <div class="text-[11px] theme-text-dim p-3 rounded-2xl theme-bg-surface border theme-border space-y-1">
          <p class="font-bold theme-text-muted flex items-center gap-1">
            <Icon name="ph:sparkle-bold" class="text-amber-500" /> Tips Broadcast:
          </p>
          <p>Gunakan judul yang ringkas dan informatif agar menarik perhatian pemain untuk segera membuka jadwal.</p>
        </div>
      </div>

    </div>

    <!-- Tabel Riwayat Broadcast Terkirim -->
    <div class="theme-bg-card backdrop-blur-md border theme-border rounded-3xl p-6 shadow-xl space-y-4 transition-colors duration-300">
      <div class="flex justify-between items-center border-b theme-border-subtle pb-3">
        <h2 class="text-base font-black theme-text-main flex items-center gap-2">
          <Icon name="ph:clock-counter-clockwise-bold" class="text-orange-500 text-lg" />
          Riwayat Pengumuman Terkirim
        </h2>
        <span class="text-xs font-bold theme-text-muted">
          {{ history.length }} Total Pengumuman
        </span>
      </div>

      <div v-if="loadingFetch" class="py-10 text-center text-xs theme-text-muted">
        <Icon name="ph:spinner-gap-bold" class="text-2xl text-orange-500 animate-spin mx-auto mb-2" />
        Memuat riwayat pengumuman...
      </div>

      <div v-else-if="history.length === 0" class="py-12 text-center text-xs theme-text-muted space-y-2">
        <Icon name="ph:broadcast-slash-bold" class="text-3xl theme-text-dim mx-auto" />
        <p class="font-bold theme-text-main">Belum Ada Riwayat Broadcast</p>
        <p class="text-[11px] theme-text-muted">Setiap pengumuman yang Anda kirim akan tercatat di sini.</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="w-full text-left text-xs">
          <thead>
            <tr class="theme-bg-surface border-b theme-border-subtle theme-text-muted uppercase text-[10px] tracking-wider">
              <th class="p-3 font-bold rounded-l-xl">Kategori</th>
              <th class="p-3 font-bold">Judul & Pesan</th>
              <th class="p-3 font-bold">Pengirim</th>
              <th class="p-3 font-bold">Waktu</th>
              <th class="p-3 font-bold text-center rounded-r-xl">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y theme-border-subtle">
            <tr v-for="item in history" :key="item.id" class="hover:bg-orange-500/5 transition">
              <td class="p-3 whitespace-nowrap">
                <span 
                  :class="getTypeBadgeClass(item.type)"
                  class="text-[10px] font-black px-2.5 py-1 rounded-full uppercase flex items-center gap-1 w-fit">
                  <Icon :name="getTypeIcon(item.type)" class="text-xs" />
                  {{ item.type || 'INFO' }}
                </span>
              </td>
              <td class="p-3">
                <p class="font-bold theme-text-main text-xs">{{ item.title }}</p>
                <p class="text-[11px] theme-text-muted line-clamp-2 mt-0.5">{{ item.message }}</p>
              </td>
              <td class="p-3 whitespace-nowrap font-medium theme-text-main">
                {{ item.sender || 'Admin' }}
              </td>
              <td class="p-3 whitespace-nowrap theme-text-muted font-mono text-[11px]">
                {{ formatTanggal(item.created_at) }}
              </td>
              <td class="p-3 whitespace-nowrap text-center">
                <button 
                  @click="hapusNotifikasi(item.id)" 
                  class="p-2 rounded-xl bg-red-500/10 hover:bg-red-500 text-red-500 hover:text-white transition active:scale-90"
                  title="Hapus Pengumuman">
                  <Icon name="ph:trash-bold" class="text-sm" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'

const { $api } = useNuxtApp()
const toast = useToast()

const form = reactive({
  title: '',
  message: '',
  type: 'INFO',
  sender: 'Admin JMT Sport'
})

const loading = ref(false)
const loadingFetch = ref(false)
const history = ref([])

const fetchHistory = async () => {
  loadingFetch.value = true
  try {
    const res = await $api('/notifications')
    if (Array.isArray(res)) {
      history.value = res
    }
  } catch (err) {
    console.error('Gagal mengambil riwayat broadcast:', err)
  } finally {
    loadingFetch.value = false
  }
}

onMounted(() => {
  fetchHistory()
})

const kirimBroadcast = async () => {
  if (!form.title || !form.message) return
  loading.value = true

  try {
    const res = await $api('/broadcast-notification', {
      method: 'POST',
      body: form
    })

    if (res?.data) {
      history.value.unshift(res.data)
      form.title = ''
      form.message = ''
      toast.success('Pengumuman berhasil disiarkan ke seluruh aplikasi secara realtime!')
    }
  } catch (err) {
    toast.error('Gagal mengirim broadcast: ' + (err?.data?.error || err.message))
  } finally {
    loading.value = false
  }
}

const hapusNotifikasi = async (id) => {
  toast.confirm({
    title: 'Hapus Pengumuman',
    message: 'Apakah Anda yakin ingin menghapus notifikasi ini dari riwayat?',
    confirmText: 'Ya, Hapus',
    cancelText: 'Batal',
    type: 'danger',
    onConfirm: async () => {
      try {
        await $api(`/notifications/${id}`, {
          method: 'DELETE'
        })
        history.value = history.value.filter(item => item.id !== id)
        toast.info('Notifikasi berhasil dihapus')
      } catch (err) {
        toast.error('Gagal menghapus: ' + (err?.data?.error || err.message))
      }
    }
  })
}

const getTypeBadgeClass = (type) => {
  switch (type) {
    case 'JADWAL':
      return 'bg-emerald-500/15 text-emerald-500 border border-emerald-500/30'
    case 'URGENT':
      return 'bg-rose-500/15 text-rose-500 border border-rose-500/30'
    case 'PROMO':
      return 'bg-amber-500/15 text-amber-500 border border-amber-500/30'
    case 'INFO':
    default:
      return 'bg-blue-500/15 text-blue-500 border border-blue-500/30'
  }
}

const getTypeIcon = (type) => {
  switch (type) {
    case 'JADWAL':
      return 'ph:soccer-ball-bold'
    case 'URGENT':
      return 'ph:warning-circle-bold'
    case 'PROMO':
      return 'ph:gift-bold'
    case 'INFO':
    default:
      return 'ph:info-bold'
  }
}

const formatTanggal = (isoDate) => {
  if (!isoDate) return '-'
  try {
    const d = new Date(isoDate)
    return d.toLocaleDateString('id-ID', {
      day: 'numeric',
      month: 'short',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    })
  } catch (e) {
    return isoDate
  }
}
</script>
