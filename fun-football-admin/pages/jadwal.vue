<template>
  <div class="space-y-8 max-w-6xl mx-auto pb-10 transition-colors duration-300">
    
    <!-- Bagian Form (Dibuat Grid 2 Kolom agar Rapi & Tidak Ada Ruang Kosong) -->
    <div class="theme-bg-surface rounded-2xl shadow-sm border theme-border p-8 transition-colors duration-300">
      <div class="flex justify-between items-center mb-6 border-b theme-border-subtle pb-4">
        <div>
          <h2 class="text-xl font-bold theme-text-main">
            {{ isEditMode ? 'Edit Jadwal Main' : 'Buat Jadwal Baru' }}
          </h2>
          <p class="text-xs theme-text-muted mt-0.5">Kelola informasi pertandingan dengan akurat.</p>
        </div>
        <button v-if="isEditMode" @click="resetForm" class="px-3 py-1.5 bg-red-500/10 text-red-500 rounded-xl text-xs font-bold hover:bg-red-500/20 transition">
          Batal Edit
        </button>
      </div>
      
      <form @submit.prevent="submitForm" class="space-y-5">
        
        <!-- Baris 1: Nama Jadwal & Lokasi (2 Kolom) -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          <div>
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted mb-2">Nama Jadwal</label>
            <input v-model="form.title" type="text" placeholder="Misal: Fun Football Pagi" class="w-full border theme-border rounded-xl p-3.5 focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card placeholder-slate-400 font-medium" required>
          </div>
          <div>
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted mb-2">Lokasi Lapangan</label>
            <input v-model="form.location" type="text" placeholder="Misal: Stadion Gelora" class="w-full border theme-border rounded-xl p-3.5 focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card placeholder-slate-400 font-medium" required>
          </div>
        </div>

        <!-- Baris 2: Tanggal Main & Batas Waktu Bayar (2 Kolom) -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          <div>
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted mb-2">Tanggal & Jam Main</label>
            <input v-model="form.match_date" type="datetime-local" class="w-full border theme-border rounded-xl p-3.5 focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-medium" required>
          </div>
          <div>
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted mb-2">Batas Waktu Bayar</label>
            <input v-model="form.payment_deadline" type="datetime-local" class="w-full border theme-border rounded-xl p-3.5 focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-medium" required>
          </div>
        </div>

        <!-- Baris 3: Kuota Maksimal & Harga Patungan (2 Kolom) -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          <div>
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted mb-2">Kuota Maksimal (Orang)</label>
            <input v-model="form.quota_max" type="number" class="w-full border theme-border rounded-xl p-3.5 focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-medium" required>
          </div>
          <div>
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted mb-2">Harga Patungan (Rp)</label>
            <input v-model="form.price_per_person" type="number" class="w-full border theme-border rounded-xl p-3.5 focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-medium" required>
          </div>
        </div>

        <!-- Tombol Aksi (Gradasi Merah-Oranye Senada) -->
        <div class="pt-4">
          <button type="submit" class="w-full bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white font-bold py-4 rounded-xl shadow-md transition duration-200 text-sm active:scale-95">
            {{ isEditMode ? 'Simpan Perubahan Jadwal' : 'Simpan & Publikasikan Jadwal' }}
          </button>
        </div>
      </form>
    </div>

    <!-- Bagian Tabel Daftar Seluruh Jadwal -->
    <div class="theme-bg-surface rounded-2xl shadow-sm border theme-border overflow-hidden transition-colors duration-300">
      <div class="p-5 border-b theme-border theme-bg-card">
        <h3 class="font-bold theme-text-main text-base">Daftar Seluruh Jadwal</h3>
      </div>
      
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="theme-bg-card theme-text-muted text-xs uppercase tracking-wider">
            <th class="p-4 font-bold border-b theme-border">Nama Jadwal</th>
            <th class="p-4 font-bold border-b theme-border">Tanggal</th>
            <th class="p-4 font-bold border-b theme-border">Lokasi</th>
            <th class="p-4 font-bold border-b theme-border">Harga</th>
            <th class="p-4 font-bold border-b theme-border text-center">Aksi Koreksi</th>
          </tr>
        </thead>
        <tbody class="text-sm">
          <tr v-for="event in eventsData?.data" :key="event.id" class="border-b theme-border-subtle hover:bg-orange-500/5 transition">
            <td class="p-4 theme-text-main font-bold">{{ event.title }}</td>
            <td class="p-4 theme-text-muted">{{ new Date(event.match_date).toLocaleDateString('id-ID') }}</td>
            <td class="p-4 theme-text-muted">{{ event.location }}</td>
            <td class="p-4 theme-text-main font-semibold">Rp {{ (event.price_per_person || 0).toLocaleString('id-ID') }}</td>
            <td class="p-4 text-center space-x-2">
              <button @click="pilihEdit(event)" class="px-3 py-1.5 bg-amber-500/10 text-amber-500 hover:bg-amber-500/20 rounded-lg text-xs font-bold transition">
                Edit
              </button>
              <button @click="hapusJadwal(event.id)" class="px-3 py-1.5 bg-red-500/10 text-red-500 hover:bg-red-500/20 rounded-lg text-xs font-bold transition">
                Hapus
              </button>
            </td>
          </tr>
          <tr v-if="!eventsData?.data || eventsData.data.length === 0">
            <td colspan="5" class="p-8 text-center theme-text-muted text-sm">Belum ada jadwal yang terdaftar.</td>
          </tr>
        </tbody>
      </table>
    </div>

  </div>
</template>

<script setup>
import { ref } from 'vue'

const { $api } = useNuxtApp()
const { useAutoRefresh } = useRealtime()
const toast = useToast()

const { data: eventsData, refresh } = await useApiFetch('/events')

// Pasang Auto-Refresh Realtime
useAutoRefresh(['EVENT_UPDATED', 'REGISTRATION_UPDATED'], () => {
  refresh()
})

const isEditMode = ref(false)
const selectedEventId = ref(null)

const form = ref({
  title: '',
  location: '',
  match_date: '',
  quota_max: 20,
  price_per_person: 50000,
  payment_deadline: ''
})

const pilihEdit = (event) => {
  isEditMode.value = true
  selectedEventId.value = event.id
  
  const formatDateTime = (dateStr) => {
    if (!dateStr) return ''
    const d = new Date(dateStr)
    return d.toISOString().slice(0, 16)
  }

  form.value = {
    title: event.title,
    location: event.location,
    match_date: formatDateTime(event.match_date),
    payment_deadline: formatDateTime(event.payment_deadline),
    quota_max: event.quota_max,
    price_per_person: event.price_per_person
  }
  
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const resetForm = () => {
  isEditMode.value = false
  selectedEventId.value = null
  form.value = {
    title: '',
    location: '',
    match_date: '',
    quota_max: 20,
    price_per_person: 50000,
    payment_deadline: ''
  }
}

const submitForm = async () => {
  try {
    const payload = {
      ...form.value,
      match_date: new Date(form.value.match_date).toISOString(),
      payment_deadline: new Date(form.value.payment_deadline).toISOString(),
      status: "OPEN"
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
    title: 'Hapus Jadwal',
    message: 'Apakah Anda yakin ingin menghapus jadwal pertandingan ini? Data polling yang terkait akan dibersihkan.',
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
</script>