<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-800 mb-6">Verifikasi Pembayaran</h1>
    
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-100 text-gray-600 text-sm">
            <th class="p-4 font-medium border-b">Nama Pemain</th>
            <th class="p-4 font-medium border-b">Status Polling</th>
            <th class="p-4 font-medium border-b">Status Bayar</th>
            <th class="p-4 font-medium border-b">Bukti Transfer</th>
            <th class="p-4 font-medium border-b">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="reg in registrations?.data" :key="reg.id" class="border-b hover:bg-gray-50 transition">
            <td class="p-4 text-gray-800 font-medium">{{ reg.user_name }}</td>
            <td class="p-4 text-gray-600">{{ reg.polling_status }}</td>
            <td class="p-4">
              <span v-if="reg.payment_status === 'VERIFYING'" class="px-3 py-1 bg-yellow-100 text-yellow-700 rounded-full text-xs font-bold">VERIFYING</span>
              <span v-else-if="reg.payment_status === 'PAID'" class="px-3 py-1 bg-green-100 text-green-700 rounded-full text-xs font-bold">PAID</span>
              <span v-else class="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-xs font-bold">{{ reg.payment_status }}</span>
            </td>
            <td class="p-4">
              <!-- Membuka gambar di tab baru -->
              <a v-if="reg.payment_proof_url" :href="'http://localhost:8080' + reg.payment_proof_url" target="_blank" class="text-blue-500 hover:underline text-sm font-medium">
                Lihat Bukti
              </a>
              <span v-else class="text-gray-400 text-sm">-</span>
            </td>
            <td class="p-4 space-x-2">
              <button v-if="reg.payment_status === 'VERIFYING'" @click="prosesVerifikasi(reg.id, 'APPROVE')" class="px-3 py-1 bg-green-500 text-white text-sm rounded hover:bg-green-600 transition">Approve</button>
              <button v-if="reg.payment_status === 'VERIFYING'" @click="prosesVerifikasi(reg.id, 'REJECT')" class="px-3 py-1 bg-red-500 text-white text-sm rounded hover:bg-red-600 transition">Reject</button>
            </td>
          </tr>
          
          <tr v-if="!registrations?.data || registrations.data.length === 0">
            <td colspan="5" class="p-6 text-center text-gray-400">Belum ada data pendaftar.</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
// Menarik data pendaftar dari backend
const { data: registrations, refresh } = await useFetch('http://localhost:8080/registrations')

// Fungsi untuk tombol Approve / Reject
const prosesVerifikasi = async (id, action) => {
  try {
    await $fetch('http://localhost:8080/verify-payment', {
      method: 'PUT',
      body: { registration_id: id, action: action }
    })
    alert(`Sukses! Pembayaran berhasil di-${action}`)
    
    // Muat ulang tabel otomatis setelah sukses
    refresh() 
  } catch (error) {
    alert('Terjadi kesalahan saat memverifikasi')
  }
}
</script>