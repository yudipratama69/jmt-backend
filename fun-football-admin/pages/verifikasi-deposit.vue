<template>
  <div class="space-y-6 max-w-6xl mx-auto pb-10 transition-colors duration-300">
    <div>
      <h1 class="text-2xl font-extrabold theme-text-main">Verifikasi Deposit Saldo</h1>
      <p class="text-xs theme-text-muted mt-0.5">Daftar permintaan isi saldo deposit pemain yang menunggu persetujuan.</p>
    </div>

    <!-- Tabel Daftar Top Up -->
    <div class="theme-bg-surface rounded-2xl shadow-sm border theme-border overflow-hidden transition-colors duration-300">
      <div class="p-5 border-b theme-border bg-gray-50/50 theme-bg-card flex justify-between items-center">
        <h3 class="font-bold theme-text-main text-base">Permintaan Top Up Saldo</h3>
        <span class="text-xs text-orange-500 font-bold bg-orange-500/10 px-3 py-1 rounded-full border border-orange-500/20">
          {{ pendingTopups?.data?.length || 0 }} Antrean
        </span>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="theme-bg-card theme-text-muted text-xs uppercase tracking-wider">
              <th class="p-4 font-bold border-b theme-border">Tanggal</th>
              <th class="p-4 font-bold border-b theme-border">Nama Pemain</th>
              <th class="p-4 font-bold border-b theme-border">Nominal</th>
              <th class="p-4 font-bold border-b theme-border text-center">Bukti Transfer</th>
              <th class="p-4 font-bold border-b theme-border text-center">Aksi Pengurus</th>
            </tr>
          </thead>
          <tbody class="text-sm">
            <tr v-for="item in pendingTopups?.data" :key="item._id" class="border-b theme-border-subtle hover:bg-orange-500/5 transition">
              <td class="p-4 theme-text-muted">{{ new Date(item.created_at).toLocaleDateString('id-ID') }}</td>
              <td class="p-4 theme-text-main font-bold">{{ item.user_name || 'Tanpa Nama' }}</td>
              <td class="p-4 font-black text-amber-500">Rp {{ (item.amount || 0).toLocaleString('id-ID') }}</td>
              <td class="p-4 text-center">
                <button @click="lihatResi(item.receipt)" class="px-3 py-1 bg-blue-500/10 text-blue-500 hover:bg-blue-500/20 rounded-xl text-xs font-bold transition inline-flex items-center gap-1">
                  <Icon name="ph:image-bold" class="text-sm" /> Lihat Foto
                </button>
              </td>
              <td class="p-4 flex justify-center gap-2">
                <button @click="prosesTopup(item._id, 'APPROVE')" class="px-3.5 py-1.5 bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
                  Approve
                </button>
                <button @click="prosesTopup(item._id, 'REJECT')" class="px-3.5 py-1.5 bg-rose-600 hover:bg-rose-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
                  Reject
                </button>
              </td>
            </tr>

            <!-- Jika Data Kosong -->
            <tr v-if="!pendingTopups?.data || pendingTopups.data.length === 0">
              <td colspan="5" class="p-8 text-center theme-text-muted text-sm">
                Tidak ada permintaan deposit baru yang menunggu.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal Lihat Foto Bukti -->
    <div v-if="showModal" class="fixed inset-0 bg-black/80 backdrop-blur-sm z-50 flex items-center justify-center p-4">
      <div class="theme-bg-surface border theme-border p-4 rounded-2xl relative max-w-lg w-full shadow-2xl space-y-3">
        <div class="flex justify-between items-center border-b theme-border pb-2">
          <h4 class="font-bold theme-text-main text-sm">Bukti Transfer Deposit</h4>
          <button @click="showModal = false" class="theme-text-muted hover:theme-text-main w-8 h-8 rounded-full theme-bg-card flex items-center justify-center transition">
            ✕
          </button>
        </div>
        <div class="max-h-[70vh] overflow-auto rounded-xl border theme-border">
          <img :src="useApiUrl(selectedReceipt)" class="w-full h-auto object-contain" alt="Bukti Transfer" />
        </div>
        <div class="flex justify-end pt-1">
          <button @click="showModal = false" class="px-4 py-2 theme-bg-card theme-text-main border theme-border text-xs font-bold rounded-xl hover:opacity-80 transition">
            Tutup
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref } from 'vue'

const { $api } = useNuxtApp()
const { useAutoRefresh } = useRealtime()

const { data: pendingTopups, refresh } = await useApiFetch('/pending-topups')

// Pasang Auto-Refresh Realtime
useAutoRefresh(['TOPUP_UPDATED'], () => {
  refresh()
})

const showModal = ref(false)
const selectedReceipt = ref('')

const lihatResi = (url) => {
  selectedReceipt.value = url
  showModal.value = true
}

const toast = useToast()

const prosesTopup = (id, action) => {
  const isApprove = action === 'APPROVE'
  toast.confirm({
    title: isApprove ? 'Setujui Deposit' : 'Tolak Deposit',
    message: isApprove ? 'Apakah Anda yakin ingin menyetujui dan menambahkan saldo deposit pemain ini?' : 'Apakah Anda yakin ingin menolak permintaan deposit ini?',
    confirmText: isApprove ? 'Ya, Setujui' : 'Tolak',
    cancelText: 'Batal',
    onConfirm: async () => {
      try {
        await $api('/approve-topup', {
          method: 'POST',
          body: { topup_id: id, action: action }
        })
        toast.success(`Permintaan deposit berhasil di-${isApprove ? 'setujui dan saldo ditambahkan' : 'tolak'}!`, 'Sukses')
        refresh() 
      } catch (error) {
        toast.error(error.response?._data?.error || 'Terjadi kesalahan saat memproses data', 'Error')
      }
    }
  })
}
</script>