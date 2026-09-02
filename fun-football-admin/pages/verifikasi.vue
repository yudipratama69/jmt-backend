<template>
  <div class="space-y-6 max-w-6xl mx-auto pb-10">
    <div>
      <h1 class="text-2xl font-extrabold text-gray-800">Verifikasi Pembayaran</h1>
      <p class="text-xs text-gray-400 mt-0.5">Periksa bukti transfer dan setujui slot pemain.</p>
    </div>
    
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <div class="p-5 border-b border-gray-100 bg-gray-50/50 flex justify-between items-center">
        <h3 class="font-bold text-gray-700 text-base">Daftar Pendaftar & Bukti Pembayaran</h3>
        <span class="text-xs text-orange-600 font-bold bg-orange-50 px-3 py-1 rounded-full border border-orange-100">
          {{ registrations?.data?.length || 0 }} Pendaftar
        </span>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-gray-50 text-gray-400 text-xs uppercase tracking-wider">
              <th class="p-4 font-bold border-b border-gray-100">Nama Pemain</th>
              <th class="p-4 font-bold border-b border-gray-100">Status Polling</th>
              <th class="p-4 font-bold border-b border-gray-100">Metode & Status</th>
              <th class="p-4 font-bold border-b border-gray-100 text-center">Bukti Transfer</th>
              <th class="p-4 font-bold border-b border-gray-100 text-center">Aksi Pengurus</th>
            </tr>
          </thead>
          <tbody class="text-sm">
            <tr v-for="reg in registrations?.data" :key="reg.id" class="border-b border-gray-50 hover:bg-gray-50/80 transition">
              <td class="p-4 text-gray-800 font-bold">{{ reg.user_name }}</td>
              <td class="p-4">
                <span :class="reg.polling_status === 'JOIN' ? 'bg-emerald-50 text-emerald-700 border-emerald-200' : 'bg-amber-50 text-amber-700 border-amber-200'" class="px-2.5 py-1 border rounded-full text-xs font-bold">
                  {{ reg.polling_status }}
                </span>
              </td>
              <td class="p-4">
                <span v-if="reg.payment_status === 'VERIFYING'" class="px-3 py-1 bg-amber-100 text-amber-800 rounded-full text-xs font-black animate-pulse">VERIFIKASI</span>
                <span v-else-if="reg.payment_status === 'PAID'" class="px-3 py-1 bg-emerald-100 text-emerald-800 rounded-full text-xs font-black">LUNAS</span>
                <span v-else class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs font-bold">{{ reg.payment_status }}</span>
              </td>
              <td class="p-4 text-center">
                <!-- Tombol Lihat Foto Bukti -->
                <button 
                  v-if="reg.payment_proof_url" 
                  @click="bukaBukti(reg.payment_proof_url)" 
                  class="px-3 py-1 bg-blue-50 text-blue-600 hover:bg-blue-100 rounded-xl text-xs font-bold transition inline-flex items-center gap-1">
                  <Icon name="ph:image-bold" class="text-sm" /> Lihat Foto
                </button>
                <span v-else class="text-gray-400 text-xs">-</span>
              </td>
              <td class="p-4 text-center space-x-2">
                <button 
                  v-if="reg.payment_status === 'VERIFYING'" 
                  @click="prosesVerifikasi(reg.id, 'APPROVE')" 
                  class="px-3.5 py-1.5 bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
                  Approve
                </button>
                <button 
                  v-if="reg.payment_status === 'VERIFYING'" 
                  @click="prosesVerifikasi(reg.id, 'REJECT')" 
                  class="px-3.5 py-1.5 bg-rose-600 hover:bg-rose-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
                  Reject
                </button>
                <span v-if="reg.payment_status !== 'VERIFYING'" class="text-xs text-gray-400">Selesai</span>
              </td>
            </tr>
            
            <tr v-if="!registrations?.data || registrations.data.length === 0">
              <td colspan="5" class="p-8 text-center text-gray-400 text-sm">Belum ada data pendaftaran yang masuk.</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal Pratinjau Foto Bukti Pembayaran -->
    <div v-if="showModalBukti" class="fixed inset-0 bg-black/80 backdrop-blur-sm z-50 flex items-center justify-center p-4">
      <div class="bg-white p-4 rounded-2xl max-w-lg w-full relative shadow-2xl space-y-3">
        <div class="flex justify-between items-center border-b pb-2">
          <h4 class="font-bold text-gray-800 text-sm">Bukti Transfer Pembayaran</h4>
          <button @click="showModalBukti = false" class="text-gray-400 hover:text-gray-800 w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center transition">
            ✕
          </button>
        </div>
        <div class="max-h-[70vh] overflow-auto rounded-xl border">
          <img :src="useApiUrl(selectedBuktiUrl)" class="w-full h-auto object-contain" alt="Bukti Transfer" />
        </div>
        <div class="flex justify-end pt-1">
          <button @click="showModalBukti = false" class="px-4 py-2 bg-gray-800 text-white text-xs font-bold rounded-xl hover:bg-gray-900 transition">
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

// Menarik data pendaftar dari backend
const { data: registrations, refresh } = await useApiFetch('/registrations')

// Pasang Auto-Refresh Realtime
useAutoRefresh(['PAYMENT_UPDATED', 'REGISTRATION_UPDATED'], () => {
  refresh()
})

const toast = useToast()
const showModalBukti = ref(false)
const selectedBuktiUrl = ref('')

const bukaBukti = (url) => {
  selectedBuktiUrl.value = url
  showModalBukti.value = true
}

// Fungsi untuk tombol Approve / Reject
const prosesVerifikasi = (id, action) => {
  const isApprove = action === 'APPROVE'
  toast.confirm({
    title: isApprove ? 'Setujui Pembayaran' : 'Tolak Pembayaran',
    message: isApprove ? 'Apakah Anda yakin ingin menyetujui pembayaran pemain ini? Status tiket pemain akan langsung LUNAS.' : 'Apakah Anda yakin ingin menolak bukti transfer ini?',
    confirmText: isApprove ? 'Ya, Setujui' : 'Tolak',
    cancelText: 'Batal',
    onConfirm: async () => {
      try {
        await $api('/verify-payment', {
          method: 'PUT',
          body: { registration_id: id, action: action }
        })
        toast.success(`Pembayaran berhasil di-${isApprove ? 'setujui (LUNAS)' : 'tolak'}!`, 'Verifikasi Berhasil')
        refresh() 
      } catch (error) {
        toast.error('Terjadi kesalahan saat memverifikasi pembayaran.', 'Error')
      }
    }
  })
}
</script>