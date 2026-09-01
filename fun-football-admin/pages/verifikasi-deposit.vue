<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold text-gray-800 mb-2">Verifikasi Deposit</h1>
    <p class="text-gray-500 mb-6 text-sm">Daftar permintaan isi saldo deposit pemain yang menunggu persetujuan.</p>

    <!-- Tabel Daftar Top Up -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-100 text-gray-600 text-sm">
            <th class="p-4 font-medium border-b">Tanggal</th>
            <th class="p-4 font-medium border-b">Nama Pemain</th>
            <th class="p-4 font-medium border-b">Nominal</th>
            <th class="p-4 font-medium border-b text-center">Bukti Transfer</th>
            <th class="p-4 font-medium border-b text-center">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in pendingTopups?.data" :key="item._id" class="border-b hover:bg-gray-50 transition">
            <td class="p-4 text-gray-600">{{ new Date(item.created_at).toLocaleDateString('id-ID') }}</td>
            <td class="p-4 text-gray-800 font-medium">{{ item.user_name || 'Tanpa Nama' }}</td>
            <td class="p-4 font-bold text-orange-600">Rp {{ item.amount.toLocaleString('id-ID') }}</td>
            <td class="p-4 text-center">
              <button @click="lihatResi(item.receipt)" class="text-blue-500 hover:underline text-sm font-medium">
                Lihat Foto
              </button>
            </td>
            <td class="p-4 flex justify-center gap-2">
              <button @click="prosesTopup(item._id, 'APPROVE')" class="px-3 py-1 bg-green-500 text-white text-sm rounded hover:bg-green-600 transition">
                Approve
              </button>
              <button @click="prosesTopup(item._id, 'REJECT')" class="px-3 py-1 bg-red-500 text-white text-sm rounded hover:bg-red-600 transition">
                Reject
              </button>
            </td>
          </tr>

          <!-- Jika Data Kosong -->
          <tr v-if="!pendingTopups?.data || pendingTopups.data.length === 0">
            <td colspan="5" class="p-6 text-center text-gray-400">
              Tidak ada permintaan deposit baru.
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal Lihat Foto Bukti -->
    <div v-if="showModal" class="fixed inset-0 bg-black/70 z-50 flex items-center justify-center p-4">
      <div class="bg-white p-2 rounded-xl relative max-w-lg w-full">
        <button @click="showModal = false" class="absolute -top-10 right-0 text-white hover:text-gray-300 font-bold">
          Tutup (X)
        </button>
        <img :src="`http://localhost:8080${selectedReceipt}`" class="w-full h-auto rounded-lg" alt="Bukti Transfer" />
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref } from 'vue'

// Menggunakan useFetch persis seperti di verifikasi.vue
const { data: pendingTopups, refresh } = await useFetch('http://localhost:8080/pending-topups')

const showModal = ref(false)
const selectedReceipt = ref('')

const lihatResi = (url) => {
  selectedReceipt.value = url
  showModal.value = true
}

const prosesTopup = async (id, action) => {
  const confirmMsg = action === 'APPROVE' ? 'Setujui dan tambahkan saldo deposit ini?' : 'Tolak deposit ini?'
  if (!confirm(confirmMsg)) return

  try {
    await $fetch('http://localhost:8080/approve-topup', {
      method: 'POST',
      body: { topup_id: id, action: action }
    })
    alert(`Sukses! Permintaan deposit berhasil di-${action}`)
    
    // Muat ulang tabel otomatis setelah sukses (sama seperti verifikasi.vue)
    refresh() 
  } catch (error) {
    alert(error.response?._data?.error || 'Terjadi kesalahan saat memproses data')
  }
}
</script>