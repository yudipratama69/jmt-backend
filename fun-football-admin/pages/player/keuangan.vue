<template>
  <div class="p-5 space-y-6 pb-24">
    
    <!-- Header -->
    <div>
      <h2 class="text-xl font-bold text-gray-800">Keuangan</h2>
      <p class="text-xs text-gray-500 mt-1">Kelola saldo deposit untuk kemudahan bayar jadwal.</p>
    </div>

    <!-- Kartu Saldo -->
    <div class="bg-gradient-to-br from-orange-500 to-amber-500 rounded-2xl p-6 text-white shadow-lg shadow-orange-200 relative overflow-hidden">
      <Icon name="ph:wallet-bold" class="absolute -right-4 -bottom-4 text-8xl text-white opacity-20 transform -rotate-12" />
      <div class="relative z-10">
        <p class="text-xs font-medium text-orange-100 uppercase tracking-wider mb-1">Total Saldo Aktif</p>
        <h3 class="text-3xl font-black">Rp {{ userData?.deposit?.toLocaleString('id-ID') || '0' }}</h3>
      </div>
    </div>

    <!-- Form Top Up -->
    <div class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100">
      <h4 class="font-bold text-gray-800 mb-4 flex items-center gap-2">
        <Icon name="ph:money-bold" class="text-orange-600 text-xl" /> Isi Saldo Deposit
      </h4>

      <form @submit.prevent="submitTopUp" class="space-y-4">
        <div>
          <label class="block text-xs font-bold uppercase tracking-wider text-gray-500 mb-1.5">Nominal Top Up (Rp)</label>
          <input v-model="form.amount" type="number" min="10000" placeholder="Contoh: 50000" class="w-full border border-gray-200 rounded-xl p-3 bg-gray-50/70 text-sm text-gray-800 outline-none focus:ring-2 focus:ring-orange-500 transition" required />
        </div>

        <div>
          <label class="block text-xs font-bold uppercase tracking-wider text-gray-500 mb-1.5">Bukti Transfer</label>
          <input type="file" @change="handleFileChange" accept="image/*" class="w-full text-sm text-gray-500 file:mr-4 file:py-2.5 file:px-4 file:rounded-full file:border-0 file:text-xs file:font-bold file:bg-orange-50 file:text-orange-700 hover:file:bg-orange-100" required />
        </div>

        <button type="submit" :disabled="isLoading" class="w-full bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white font-bold py-3.5 rounded-xl shadow-md transition-all active:scale-95 disabled:opacity-70">
          {{ isLoading ? 'Mengirim...' : 'Kirim Permintaan Deposit' }}
        </button>
      </form>
    </div>

    <!-- Info Rekening -->
    <div class="bg-blue-50 border border-blue-100 rounded-2xl p-4 flex gap-3">
      <Icon name="ph:info-bold" class="text-blue-600 text-2xl shrink-0" />
      <div>
        <p class="text-sm font-bold text-blue-800">Info Transfer</p>
        <p class="text-xs text-blue-700 mt-1">Silakan transfer ke rekening <strong>BNI 1148532156 a/n IPAN SOPANDI</strong> sebelum mengunggah bukti pembayaran.</p>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

definePageMeta({ layout: 'mobile' })

const userId = ref('')
const userData = ref({})
const isLoading = ref(false)

const form = ref({
  amount: '',
  receipt: null
})

const ambilDataUser = async (uid) => {
  try {
    const res = await $fetch(`http://localhost:8080/user?id=${uid}`)
    userData.value = res.data || {}
  } catch (error) {
    console.error("Gagal mengambil data user")
  }
}

onMounted(() => {
  const storedUserId = localStorage.getItem('user_id')
  if (!storedUserId) {
    navigateTo('/player/login')
  } else {
    userId.value = storedUserId
    ambilDataUser(storedUserId)
  }
})

const handleFileChange = (e) => {
  form.value.receipt = e.target.files[0]
}

const submitTopUp = async () => {
  if (!form.value.receipt) {
    return alert('Harap unggah bukti transfer!')
  }

  isLoading.value = true
  try {
    const formData = new FormData()
    formData.append('user_id', userId.value)
    formData.append('amount', form.value.amount)
    formData.append('receipt', form.value.receipt)

    // Memanggil API backend untuk mencatat permintaan deposit
    await $fetch('http://localhost:8080/request-topup', {
      method: 'POST',
      body: formData
    })

    alert('Permintaan isi saldo berhasil dikirim! Menunggu verifikasi admin.')
    form.value.amount = ''
    form.value.receipt = null
  } catch (error) {
    alert(error.response?._data?.error || 'Gagal mengirim permintaan deposit.')
  } finally {
    isLoading.value = false
  }
}
</script>