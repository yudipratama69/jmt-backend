<template>
  <div class="p-5 space-y-5">
    
    <!-- Banner Sambutan (Gradasi Merah - Oranye Menyala) -->
    <div class="bg-gradient-to-r from-red-600 via-orange-600 to-amber-600 rounded-2xl p-5 text-white shadow-lg relative overflow-hidden">
      <div class="relative z-10">
        <h2 class="text-lg font-bold mb-1">Halo, {{ userName }}! 🔥</h2>
        <p class="text-xs text-orange-100">Siap cari keringat minggu ini? Langsung amankan slotmu!</p>
      </div>
      <Icon name="ph:sneaker-bold" class="absolute -right-4 -bottom-4 text-8xl text-white opacity-20 transform -rotate-12" />
    </div>

    <!-- Info Saldo Deposit (BARU) -->
    <div class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100 flex justify-between items-center">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 rounded-full bg-orange-100 text-orange-600 flex items-center justify-center">
          <Icon name="ph:wallet-bold" class="text-xl" />
        </div>
        <div>
          <p class="text-[10px] text-gray-400 font-bold uppercase tracking-wider">Saldo Deposit</p>
          <p class="text-lg font-black text-gray-800">Rp {{ userData?.deposit?.toLocaleString('id-ID') || '0' }}</p>
        </div>
      </div>
    </div>

    <div class="flex justify-between items-end">
      <h3 class="text-gray-800 font-bold text-lg">Jadwal Terdekat</h3>
    </div>

    <!-- Looping Card Jadwal -->
    <div v-for="event in eventsData?.data" :key="event.id" class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100 relative">
      
      <!-- Badge Status (Nuansa Merah-Oranye) -->
      <div v-if="cekStatusDaftar(event.id)" class="absolute top-4 right-4 bg-gray-100 text-gray-500 text-[10px] font-extrabold px-3 py-1 rounded-full flex items-center gap-1">
        <Icon name="ph:check-circle-fill" class="text-xs" /> SUDAH DAFTAR
      </div>
      <div v-else class="absolute top-4 right-4 bg-orange-100 text-orange-700 text-[10px] font-extrabold px-2 py-1 rounded-full">
        {{ event.status }}
      </div>

      <h4 class="font-bold text-gray-800 text-base w-3/4">{{ event.title }}</h4>
      <div class="mt-3 space-y-2">
        <div class="flex items-center text-xs text-gray-500 font-medium">
          <Icon name="ph:calendar-blank-bold" class="text-orange-600 mr-2 text-base" />
          {{ new Date(event.match_date).toLocaleDateString('id-ID', { weekday: 'long', day: 'numeric', month: 'long' }) }}
        </div>
        <div class="flex items-center text-xs text-gray-500 font-medium">
          <Icon name="ph:clock-bold" class="text-orange-600 mr-2 text-base" />
          {{ new Date(event.match_date).toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' }) }} WIB
        </div>
      </div>
      <div class="mt-4 pt-4 border-t border-gray-100 flex items-center justify-between">
        <div>
          <p class="text-[10px] text-gray-400 font-semibold mb-0.5">HTM / Patungan</p>
          <p class="text-sm font-extrabold text-gray-800">Rp {{ event.price_per_person.toLocaleString('id-ID') }}</p>
        </div>
        
        <button v-if="cekStatusDaftar(event.id)" disabled class="bg-gray-200 text-gray-400 text-xs font-bold py-2.5 px-6 rounded-xl shadow-none cursor-not-allowed">
          Terkunci
        </button>
        <!-- Tombol Ikut Main dengan Warna Merah-Oranye -->
        <button v-else @click="ikutMain(event.id)" class="bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white text-xs font-bold py-2.5 px-6 rounded-xl shadow-md transition-all active:scale-95">
          Ikut Main
        </button>
      </div>
    </div>

    <!-- POP-UP (MODAL) PEMBAYARAN -->
    <div v-if="showModal" class="fixed inset-0 bg-black/60 z-50 flex items-end justify-center sm:items-center">
      <div class="bg-white w-full max-w-md rounded-t-3xl sm:rounded-3xl p-6 relative animate-fade-in-up">
        
        <button @click="tutupModal" class="absolute top-4 right-4 text-gray-400 hover:text-gray-600">
          <Icon name="ph:x-circle-fill" class="text-3xl" />
        </button>

        <div v-if="step === 1">
          <h3 class="text-lg font-bold text-gray-800 mb-2">Slot Diamankan! 🚀</h3>
          <p class="text-xs text-gray-500 mb-5">Atas nama <span class="font-bold text-orange-600">{{ userName }}</span>. Silakan pilih metode pembayaran.</p>
          
          <!-- Pilihan Metode Pembayaran (BARU) -->
          <div class="mb-5 space-y-3">
            <div class="grid grid-cols-2 gap-3">
              <!-- Opsi Transfer Manual -->
              <label :class="paymentMethod === 'transfer' ? 'border-orange-500 bg-orange-50 text-orange-700' : 'border-gray-200 text-gray-500'" class="border-2 rounded-xl p-3 flex items-center justify-center gap-2 cursor-pointer transition">
                <input type="radio" v-model="paymentMethod" value="transfer" class="hidden" />
                <Icon name="ph:bank-bold" class="text-lg" />
                <span class="text-xs font-bold">Transfer</span>
              </label>

              <!-- Opsi Potong Deposit -->
              <label :class="paymentMethod === 'deposit' ? 'border-orange-500 bg-orange-50 text-orange-700' : 'border-gray-200 text-gray-500'" class="border-2 rounded-xl p-3 flex items-center justify-center gap-2 cursor-pointer transition">
                <input type="radio" v-model="paymentMethod" value="deposit" class="hidden" />
                <Icon name="ph:wallet-bold" class="text-lg" />
                <span class="text-xs font-bold">Deposit</span>
              </label>
            </div>
          </div>

          <!-- JIKA PILIH TRANSFER: Munculkan Input File -->
          <div v-if="paymentMethod === 'transfer'" class="mb-4">
            <input type="file" @change="handleFileChange" accept="image/*" class="w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-orange-50 file:text-orange-700 hover:file:bg-orange-100" />
          </div>

          <!-- JIKA PILIH DEPOSIT: Munculkan Pemberitahuan -->
          <div v-if="paymentMethod === 'deposit'" class="mb-4 p-3 bg-amber-50 border border-amber-200 rounded-xl flex items-start gap-3">
            <Icon name="ph:info-bold" class="text-amber-600 text-2xl shrink-0" />
            <div>
              <p class="text-sm font-bold text-amber-800">Tanpa Bukti Transfer!</p>
              <p class="text-xs text-amber-700 mt-1">Saldo deposit Anda (Rp {{ userData?.deposit?.toLocaleString('id-ID') || '0' }}) akan dipotong secara otomatis.</p>
            </div>
          </div>
          
          <button @click="submitUpload" :disabled="paymentMethod === 'transfer' && !receiptFile" class="w-full bg-gradient-to-r from-red-600 to-orange-600 text-white font-bold py-3 rounded-xl disabled:bg-gray-300 transition active:scale-95">
            {{ paymentMethod === 'transfer' ? 'Kirim Bukti Pembayaran' : 'Bayar Pakai Deposit' }}
          </button>
        </div>

        <div v-if="step === 2" class="text-center py-4">
          <Icon name="ph:check-circle-fill" class="text-6xl text-orange-600 mx-auto mb-3" />
          <h3 class="text-lg font-bold text-gray-800 mb-1">Berhasil!</h3>
          <p class="text-sm text-gray-500 mb-5">
            {{ paymentMethod === 'transfer' ? 'Admin akan memverifikasi bukti pembayaranmu.' : 'Pembayaran dengan deposit berhasil diproses.' }}
          </p>
          <button @click="tutupModal" class="w-full bg-orange-100 text-orange-700 font-bold py-3 rounded-xl active:scale-95">
            Tutup
          </button>
        </div>

      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

definePageMeta({ layout: 'mobile' })

const { data: eventsData, refresh: refreshEvents } = await useFetch('http://localhost:8080/events')
const riwayatDaftar = ref([])

const userId = ref('')
const userName = ref('')
const userData = ref({}) // Menyimpan data user dari backend (termasuk saldo deposit)
const showModal = ref(false)
const step = ref(1)
const registrationId = ref('')
const receiptFile = ref(null)

const paymentMethod = ref('transfer') // Default metode pembayaran

// Fungsi ambil profil pengguna (untuk cek saldo)
const ambilDataUser = async (uid) => {
  try {
    const res = await $fetch(`http://localhost:8080/user?id=${uid}`)
    userData.value = res.data || {}
  } catch (error) {
    console.error("Gagal mengambil data user")
  }
}

const ambilRiwayatDaftar = async (uid) => {
  try {
    const res = await $fetch(`http://localhost:8080/my-registrations?user_id=${uid}`)
    riwayatDaftar.value = res.data || []
  } catch (error) {
    console.error("Gagal mengambil riwayat")
  }
}

onMounted(() => {
  const storedUserId = localStorage.getItem('user_id')
  if (!storedUserId) {
    navigateTo('/player/login')
  } else {
    userId.value = storedUserId
    userName.value = localStorage.getItem('user_name')
    ambilRiwayatDaftar(storedUserId)
    ambilDataUser(storedUserId)
  }
})

const cekStatusDaftar = (idJadwal) => {
  return riwayatDaftar.value.some(reg => reg.event_id === idJadwal)
}

const tutupModal = () => {
  showModal.value = false
  refreshEvents()
  ambilRiwayatDaftar(userId.value)
}

const handleFileChange = (e) => {
  receiptFile.value = e.target.files[0]
}

const ikutMain = async (idJadwal) => {
  try {
    const res = await $fetch('http://localhost:8080/register', {
      method: 'POST',
      body: { event_id: idJadwal, user_id: userId.value }
    })
    registrationId.value = res.data.id
    step.value = 1 
    receiptFile.value = null
    paymentMethod.value = 'transfer' // Reset ke transfer saat modal dibuka
    showModal.value = true
  } catch (error) {
    alert(error.response?._data?.error || 'Gagal mendaftar. Kuota mungkin penuh.')
  }
}

const submitUpload = async () => {
  try {
    if (paymentMethod.value === 'transfer') {
      // PROSES 1: JIKA MENGGUNAKAN TRANSFER (UPLOAD RESI)
      const formData = new FormData()
      formData.append('registration_id', registrationId.value)
      formData.append('receipt', receiptFile.value)

      await $fetch('http://localhost:8080/upload-proof', {
        method: 'POST',
        body: formData
      })
    } else {
      // PROSES 2: JIKA MENGGUNAKAN DEPOSIT
      await $fetch('http://localhost:8080/pay-deposit', {
        method: 'POST',
        body: { 
          registration_id: registrationId.value, 
          user_id: userId.value 
        }
      })
      // Perbarui saldo deposit secara instan setelah pembayaran berhasil
      ambilDataUser(userId.value) 
    }
    
    step.value = 2 
  } catch (error) {
    alert(error.response?._data?.error || 'Gagal memproses pembayaran.')
  }
}
</script>

<style scoped>
.animate-fade-in-up {
  animation: fadeInUp 0.3s ease-out forwards;
}
@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>