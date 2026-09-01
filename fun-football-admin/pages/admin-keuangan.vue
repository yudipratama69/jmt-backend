<template>
  <div class="space-y-8 w-full px-8 pb-10">
    
    <!-- Judul Halaman -->
    <div>
      <h1 class="text-2xl font-extrabold text-gray-800">Dashboard Kas & Keuangan</h1>
      <p class="text-xs text-gray-400 mt-0.5">Kelola arus kas masuk dan uang keluar komunitas secara transparan.</p>
    </div>

    <!-- Kartu Statistik Keuangan dengan Latar Gradasi Penuh -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      
      <!-- Kartu 1: Total Saldo Kas (Gradasi Oranye ke Merah) -->
      <div style="background: linear-gradient(135deg, #ea580c, #dc2626);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40">
        <div class="absolute right-4 -bottom-4 text-white/20 pointer-events-none">
          <Icon name="ph:wallet-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            💰 Saldo Utama
          </div>
          <h3 class="text-orange-100 text-xs font-bold uppercase tracking-wider">Total Saldo Kas</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">Rp {{ totalSaldo.toLocaleString('id-ID') }}</p>
      </div>

      <!-- Kartu 2: Total Uang Masuk (Gradasi Hijau) -->
      <div style="background: linear-gradient(135deg, #059669, #047857);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40">
        <div class="absolute right-4 -bottom-4 text-white/20 pointer-events-none">
          <Icon name="ph:arrow-down-left-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            📥 Pemasukan
          </div>
          <h3 class="text-emerald-100 text-xs font-bold uppercase tracking-wider">Total Uang Masuk</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">+ Rp {{ totalMasuk.toLocaleString('id-ID') }}</p>
      </div>

      <!-- Kartu 3: Total Uang Keluar (Gradasi Merah Tua) -->
      <div style="background: linear-gradient(135deg, #e11d48, #9f1239);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40">
        <div class="absolute right-4 -bottom-4 text-white/20 pointer-events-none">
          <Icon name="ph:arrow-up-right-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            📤 Pengeluaran
          </div>
          <h3 class="text-rose-100 text-xs font-bold uppercase tracking-wider">Total Uang Keluar</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">- Rp {{ totalKeluar.toLocaleString('id-ID') }}</p>
      </div>

    </div>

    <!-- Bagian Catat Transaksi Kas Baru -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-6">
      <div class="flex items-center gap-2 mb-4 border-b border-gray-100 pb-3">
        <Icon name="ph:plus-circle-bold" class="text-xl text-orange-600" />
        <h3 class="font-bold text-gray-800 text-base">Catat Transaksi Kas Baru</h3>
      </div>
      
      <form @submit.prevent="submitTransaksi" class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          
          <div>
            <label class="block text-xs font-bold text-gray-500 uppercase tracking-wider mb-1.5">Jenis Transaksi</label>
            <select v-model="form.type" class="w-full border border-gray-200 rounded-xl p-3.5 bg-gray-50/50 text-sm font-semibold outline-none focus:ring-2 focus:ring-orange-500">
              <option value="IN">🟢 Uang Masuk</option>
              <option value="OUT">🔴 Uang Keluar</option>
            </select>
          </div>

          <div>
            <label class="block text-xs font-bold text-gray-500 uppercase tracking-wider mb-1.5">Keterangan / Keperluan</label>
            <input v-model="form.description" type="text" placeholder="Misal: Sewa Lapangan / Beli Bola" class="w-full border border-gray-200 rounded-xl p-3.5 bg-gray-50/50 text-sm outline-none focus:ring-2 focus:ring-orange-500" required />
          </div>

          <div>
            <label class="block text-xs font-bold text-gray-500 uppercase tracking-wider mb-1.5">Nominal (Rp)</label>
            <input v-model="form.amount" type="number" placeholder="Misal: 150000" class="w-full border border-gray-200 rounded-xl p-3.5 bg-gray-50/50 text-sm outline-none focus:ring-2 focus:ring-orange-500" required />
          </div>

        </div>

        <div class="pt-2">
          <button type="submit" class="w-full bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white font-bold py-3.5 rounded-xl shadow-md transition duration-200 text-sm active:scale-95 flex items-center justify-center gap-2">
            <Icon name="ph:check-bold" class="text-lg" />
            Simpan Transaksi Kas
          </button>
        </div>
      </form>
    </div>

    <!-- Tabel Riwayat Semua Transaksi Kas -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <div class="p-5 border-b border-gray-100 bg-gray-50/50 flex justify-between items-center">
        <h3 class="font-bold text-gray-700 text-base">Riwayat Arus Kas</h3>
        <span class="text-xs text-gray-400 font-medium">Terakumulasi secara otomatis</span>
      </div>
      
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-50 text-gray-400 text-xs uppercase tracking-wider">
            <th class="p-4 font-bold border-b border-gray-100">Keterangan</th>
            <th class="p-4 font-bold border-b border-gray-100">Kategori</th>
            <th class="p-4 font-bold border-b border-gray-100">Jumlah</th>
            <th class="p-4 font-bold border-b border-gray-100 text-center">Aksi</th>
          </tr>
        </thead>
        <tbody class="text-sm">
          
          <!-- 1. Riwayat Deposit Masuk -->
          <tr v-for="topup in autoTopups" :key="'topup-' + topup._id" class="border-b border-gray-50 hover:bg-gray-50/80 transition">
            <td class="p-4 text-gray-800 font-bold">Top Up Deposit: {{ topup.user_name }}</td>
            <td class="p-4">
              <span class="px-3 py-1 bg-teal-50 text-teal-700 border border-teal-100 rounded-full text-xs font-bold flex items-center gap-1 w-max">
                <span class="w-1.5 h-1.5 rounded-full bg-teal-500"></span> Uang Masuk (Deposit)
              </span>
            </td>
            <td class="p-4 text-teal-600 font-extrabold">+ Rp {{ topup.amount?.toLocaleString('id-ID') || 0 }}</td>
            <td class="p-4 text-center text-xs text-gray-400 font-medium">Sistem</td>
          </tr>

          <!-- 2. Riwayat Pendaftaran (Patungan Main) -->
          <tr v-for="item in autoRegistrations" :key="'reg-' + item.id" class="border-b border-gray-50 hover:bg-gray-50/80 transition">
            <td class="p-4 text-gray-800 font-bold">Patungan Member: {{ item.user_name }}</td>
            <td class="p-4">
              <span class="px-3 py-1 bg-green-50 text-green-700 border border-green-100 rounded-full text-xs font-bold flex items-center gap-1 w-max">
                <span class="w-1.5 h-1.5 rounded-full bg-green-500"></span> Uang Masuk (Jadwal)
              </span>
            </td>
            <td class="p-4 text-green-600 font-extrabold">+ Rp {{ item.amount?.toLocaleString('id-ID') || 0 }}</td>
            <td class="p-4 text-center text-xs text-gray-400 font-medium">Sistem</td>
          </tr>

          <!-- 3. Riwayat Transaksi Manual -->
          <tr v-for="tx in manualTransactions" :key="'tx-' + tx.id" class="border-b border-gray-50 hover:bg-gray-50/80 transition">
            <td class="p-4 text-gray-800 font-bold">{{ tx.description }}</td>
            <td class="p-4">
              <span v-if="tx.type === 'IN'" class="px-3 py-1 bg-green-50 text-green-700 border border-green-100 rounded-full text-xs font-bold flex items-center gap-1 w-max">
                <span class="w-1.5 h-1.5 rounded-full bg-green-500"></span> Uang Masuk
              </span>
              <span v-else class="px-3 py-1 bg-red-50 text-red-700 border border-red-100 rounded-full text-xs font-bold flex items-center gap-1 w-max">
                <span class="w-1.5 h-1.5 rounded-full bg-red-500"></span> Uang Keluar
              </span>
            </td>
            <td class="p-4 font-extrabold" :class="tx.type === 'IN' ? 'text-green-600' : 'text-red-600'">
              {{ tx.type === 'IN' ? '+' : '-' }} Rp {{ tx.amount.toLocaleString('id-ID') }}
            </td>
            <td class="p-4 text-center">
              <button @click="hapusTransaksi(tx.id)" class="px-3 py-1.5 bg-red-50 text-red-600 hover:bg-red-100 rounded-xl text-xs font-bold transition">
                Hapus
              </button>
            </td>
          </tr>

          <!-- 4. Jika Kosong Semua -->
          <tr v-if="autoRegistrations.length === 0 && manualTransactions.length === 0 && autoTopups.length === 0">
            <td colspan="4" class="p-8 text-center text-gray-400 text-sm">Belum ada catatan transaksi kas.</td>
          </tr>

        </tbody>
      </table>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const { data: regData } = await useFetch('http://localhost:8080/registrations')
const { data: eventsData } = await useFetch('http://localhost:8080/events')
const { data: topupData } = await useFetch('http://localhost:8080/approved-topups')

const manualTransactions = ref([])

// Ambil data dari localStorage setelah halaman dimuat di browser (Mencegah error SSR)
onMounted(() => {
  manualTransactions.value = JSON.parse(localStorage.getItem('jmt_manual_tx') || '[]')
})

const form = ref({
  type: 'IN',
  description: '',
  amount: ''
})

const autoRegistrations = computed(() => {
  if (!regData.value?.data || !eventsData.value?.data) return []
  
  return regData.value.data
    // PERBAIKAN: Hanya hitung yang statusnya PAID DAN metodenya BUKAN deposit
    .filter(r => r.payment_status === 'PAID' && r.payment_method !== 'deposit')
    .map(r => {
      const evt = eventsData.value.data.find(e => e.id === r.event_id)
      return {
        ...r,
        amount: evt ? evt.price_per_person : 0
      }
    })
})

const autoTopups = computed(() => {
  return topupData.value?.data || []
})

const totalMasuk = computed(() => {
  const fromReg = autoRegistrations.value.reduce((acc, curr) => acc + (curr.amount || 0), 0)
  const fromTopup = autoTopups.value.reduce((acc, curr) => acc + (curr.amount || 0), 0)
  const fromManualIN = manualTransactions.value
    .filter(t => t.type === 'IN')
    .reduce((acc, curr) => acc + Number(curr.amount), 0)
    
  return fromReg + fromTopup + fromManualIN
})

const totalKeluar = computed(() => {
  return manualTransactions.value
    .filter(t => t.type === 'OUT')
    .reduce((acc, curr) => acc + Number(curr.amount), 0)
})

const totalSaldo = computed(() => {
  return totalMasuk.value - totalKeluar.value
})

const submitTransaksi = () => {
  const newTx = {
    id: Date.now().toString(),
    type: form.value.type,
    description: form.value.description,
    amount: Number(form.value.amount)
  }
  
  manualTransactions.value.unshift(newTx)
  localStorage.setItem('jmt_manual_tx', JSON.stringify(manualTransactions.value))
  
  form.value.description = ''
  form.value.amount = ''
  alert('Transaksi kas berhasil dicatat!')
}

const hapusTransaksi = (id) => {
  if (confirm('Hapus catatan transaksi ini?')) {
    manualTransactions.value = manualTransactions.value.filter(t => t.id !== id)
    localStorage.setItem('jmt_manual_tx', JSON.stringify(manualTransactions.value))
  }
}
</script>