<template>
  <div class="p-4 space-y-5">
    
    <!-- Hero Stadium Banner (Sporty Neon Gradient) -->
    <div class="bg-gradient-to-br from-red-600 via-orange-600 to-amber-500 rounded-3xl p-5 text-white shadow-xl shadow-orange-950/40 relative overflow-hidden border border-white/20">
      <!-- Glow Stadium Effect -->
      <div class="absolute -right-8 -bottom-8 w-36 h-36 bg-amber-300/20 rounded-full blur-2xl pointer-events-none"></div>
      <Icon name="ph:soccer-ball-bold" class="absolute -right-4 -bottom-4 text-9xl text-white/10 transform -rotate-12 pointer-events-none" />

      <div class="relative z-10">
        <div class="inline-flex items-center gap-1.5 bg-black/30 backdrop-blur-md px-3 py-1 rounded-full text-[10px] font-black uppercase tracking-wider text-amber-200 border border-white/20 mb-2.5">
          <span class="w-2 h-2 rounded-full bg-green-400 animate-pulse"></span>
          Live Match Ready
        </div>
        <h2 class="text-xl font-black text-white tracking-wide leading-tight">Halo, {{ userName || 'Pemain' }}! 🔥</h2>
        <p class="text-xs text-orange-100 mt-1 font-medium leading-relaxed">
          Amankan slot pertandingan mingguan & cek kuota pemain secara realtime!
        </p>
      </div>
    </div>

    <!-- Quick Action Cards (Saldo & PWA Download) -->
    <div class="grid grid-cols-1 gap-3">
      
      <!-- Card Saldo Deposit Pemain -->
      <div class="theme-bg-card backdrop-blur-md rounded-2xl p-4 border theme-border flex justify-between items-center shadow-lg transition-colors duration-300">
        <div class="flex items-center gap-3.5">
          <div class="w-11 h-11 rounded-2xl bg-gradient-to-tr from-amber-500 via-orange-500 to-amber-600 text-white flex items-center justify-center shadow-md shadow-orange-500/20 shrink-0">
            <Icon name="ph:wallet-bold" class="text-2xl" />
          </div>
          <div>
            <p class="text-[10px] theme-text-muted font-bold uppercase tracking-wider">Saldo Deposit Aktif</p>
            <p class="text-lg font-black theme-text-main tracking-tight">Rp {{ userData?.deposit?.toLocaleString('id-ID') || '0' }}</p>
          </div>
        </div>
        <NuxtLink to="/player/keuangan" class="px-4 py-2 bg-orange-500 hover:bg-orange-600 text-white rounded-xl text-xs font-black shadow-md shadow-orange-600/30 transition active:scale-95 flex items-center gap-1">
          <Icon name="ph:plus-bold" class="text-sm" /> Top Up
        </NuxtLink>
      </div>

      <!-- Card Tombol Download / Pasang Aplikasi PWA -->
      <div v-if="!isInstalled" class="theme-bg-card rounded-2xl p-4 shadow-xl flex items-center justify-between border border-orange-500/30 relative overflow-hidden transition-colors duration-300">
        <div class="flex items-center gap-3">
          <div class="w-11 h-11 rounded-2xl bg-white p-1.5 shrink-0 flex items-center justify-center shadow-md border border-orange-400/50">
            <img :src="'/logo-jmt.png'" alt="App Logo" class="w-full h-full object-contain" />
          </div>
          <div>
            <div class="flex items-center gap-1.5">
              <span class="text-xs font-black theme-text-main tracking-wide">Aplikasi JMT Sport</span>
              <span class="text-[9px] bg-gradient-to-r from-orange-500 to-red-600 text-white font-black px-1.5 py-0.2 rounded-full uppercase">PWA</span>
            </div>
            <p class="text-[11px] theme-text-muted">Pasang di Layar Utama HP</p>
          </div>
        </div>
        <button 
          @click="triggerInstall" 
          class="bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white text-xs font-black py-2.5 px-4 rounded-xl shadow-lg shadow-orange-600/30 transition active:scale-95 flex items-center gap-1.5">
          <Icon name="ph:download-simple-bold" class="text-base" />
          Install
        </button>
      </div>

    </div>

    <!-- Modal Panduan Install PWA -->
    <PwaInstallModal 
      :show="showPwaModal" 
      :is-i-o-s="isIOS" 
      :has-prompt="hasPrompt"
      @close="closeModal" 
      @install="triggerInstall" 
    />

    <!-- Header Section Jadwal Main -->
    <div class="flex justify-between items-center pt-2">
      <div>
        <h3 class="theme-text-main font-black text-lg tracking-wide flex items-center gap-2">
          <Icon name="ph:calendar-star-bold" class="text-orange-500 text-xl" /> Jadwal Pertandingan
        </h3>
        <p class="text-xs theme-text-muted">Pilih jadwal dan amankan posisimu di lapangan!</p>
      </div>
      <span class="theme-bg-card theme-text-main text-[11px] font-bold px-2.5 py-1 rounded-full border theme-border">
        {{ eventsData?.data?.length || 0 }} Agenda
      </span>
    </div>

    <!-- Looping Match Card Jadwal -->
    <div class="space-y-4">
      <div 
        v-for="event in eventsData?.data" 
        :key="event.id" 
        class="theme-bg-card backdrop-blur-md rounded-3xl p-5 border theme-border shadow-xl relative overflow-hidden transition hover:border-orange-500/50 duration-300">
        
        <!-- Header Card: Lokasi & Status Chip -->
        <div class="flex justify-between items-start mb-3 gap-2">
          <div class="flex items-center gap-1.5 text-xs text-orange-500 font-bold">
            <Icon name="ph:map-pin-bold" class="text-base shrink-0" />
            <span class="truncate">{{ event.location || 'Lapangan JMT' }}</span>
          </div>

          <!-- Status Pertandingan -->
          <div>
            <span v-if="event.status === 'OPEN'" class="bg-emerald-500/10 text-emerald-500 border border-emerald-500/30 text-[10px] font-black px-2.5 py-1 rounded-full flex items-center gap-1 uppercase">
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-500 animate-pulse"></span> BUKA
            </span>
            <span v-else-if="event.status === 'FULL'" class="bg-red-500/10 text-red-500 border border-red-500/30 text-[10px] font-black px-2.5 py-1 rounded-full uppercase">
              PENUH
            </span>
            <span v-else class="theme-bg-surface theme-text-muted text-[10px] font-black px-2.5 py-1 rounded-full uppercase">
              {{ event.status }}
            </span>
          </div>
        </div>

        <!-- Judul Pertandingan -->
        <h4 class="font-black theme-text-main text-base leading-snug">{{ event.title }}</h4>

        <!-- Waktu & Tanggal -->
        <div class="mt-3 grid grid-cols-2 gap-2 theme-bg-surface p-3 rounded-2xl border theme-border-subtle text-xs transition-colors duration-300">
          <div class="flex items-center gap-2 theme-text-main">
            <Icon name="ph:calendar-blank-bold" class="text-orange-500 text-base shrink-0" />
            <span class="font-semibold">{{ new Date(event.match_date).toLocaleDateString('id-ID', { weekday: 'short', day: 'numeric', month: 'short' }) }}</span>
          </div>
          <div class="flex items-center gap-2 theme-text-main">
            <Icon name="ph:clock-bold" class="text-orange-500 text-base shrink-0" />
            <span class="font-semibold">{{ new Date(event.match_date).toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' }) }} WIB</span>
          </div>
        </div>

        <!-- Live Slot Realtime Progress Bar -->
        <div class="mt-4 space-y-1.5">
          <div class="flex justify-between items-center text-xs">
            <span class="theme-text-muted font-bold flex items-center gap-1.5">
              <Icon name="ph:users-three-bold" class="text-sm theme-text-main" />
              <span>{{ event.registered_count || 0 }} / {{ event.quota_max }} Pemain</span>
            </span>
            <span :class="(event.slots_left || 0) <= 3 ? 'text-red-500 font-black' : 'text-emerald-500 font-bold'">
              {{ (event.slots_left || 0) > 0 ? `Sisa ${event.slots_left} Slot` : 'Slot Penuh' }}
            </span>
          </div>
          <!-- Bar Track -->
          <div class="w-full theme-bg-surface rounded-full h-2.5 overflow-hidden p-0.5 border theme-border">
            <div 
              class="h-full rounded-full transition-all duration-500"
              :class="((event.registered_count || 0) / event.quota_max) >= 0.8 ? 'bg-gradient-to-r from-orange-500 to-red-500' : 'bg-gradient-to-r from-amber-500 to-emerald-400'"
              :style="{ width: `${Math.min(100, Math.round(((event.registered_count || 0) / event.quota_max) * 100))}%` }">
            </div>
          </div>
        </div>

        <!-- Footer Card: Harga & Action Button -->
        <div class="mt-4 pt-4 border-t theme-border-subtle flex items-center justify-between gap-3">
          <div>
            <p class="text-[10px] theme-text-muted font-bold uppercase tracking-wider mb-0.5">HTM / Patungan</p>
            <p class="text-sm font-black text-amber-500">Rp {{ event.price_per_person?.toLocaleString('id-ID') }}</p>
          </div>

          <!-- Dynamic Status & Action Button -->
          <div>
            <!-- Case 1: Sudah Daftar & LUNAS -->
            <div v-if="getStatusReg(event.id)?.paymentStatus === 'PAID'" class="bg-emerald-500/20 text-emerald-500 border border-emerald-500/40 text-xs font-black px-4 py-2.5 rounded-xl flex items-center gap-1.5 shadow-sm">
              <Icon name="ph:check-circle-fill" class="text-sm" /> SUDAH LUNAS
            </div>

            <!-- Case 2: Sudah Daftar & Sedang Verifikasi -->
            <div v-else-if="getStatusReg(event.id)?.paymentStatus === 'VERIFYING'" class="bg-amber-500/20 text-amber-500 border border-amber-500/40 text-xs font-bold px-3.5 py-2 rounded-xl flex items-center gap-1.5">
              <Icon name="ph:hourglass-bold" class="text-sm animate-spin" /> PROSES VERIFIKASI
            </div>

            <!-- Case 3: Sudah Daftar tapi Belum Bayar -->
            <button 
              v-else-if="getStatusReg(event.id)?.paymentStatus === 'UNPAID'" 
              @click="bukaBayarUlang(getStatusReg(event.id))" 
              class="bg-gradient-to-r from-amber-500 to-orange-600 hover:from-amber-600 hover:to-orange-700 text-white text-xs font-black py-2.5 px-4 rounded-xl shadow-lg shadow-orange-600/30 active:scale-95 transition flex items-center gap-1.5 animate-pulse">
              <Icon name="ph:lightning-bold" class="text-sm" /> Bayar Sekarang
            </button>

            <!-- Case 4: Belum Daftar (Bisa Ikut Main) -->
            <button 
              v-else 
              @click="ikutMain(event.id)" 
              class="bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white text-xs font-black py-2.5 px-5 rounded-xl shadow-lg shadow-orange-600/30 transition active:scale-95 flex items-center gap-1.5">
              <Icon name="ph:sneaker-move-bold" class="text-base" /> Amankan Slot
            </button>
          </div>

        </div>

      </div>

      <!-- Jika Kosong -->
      <div v-if="!eventsData?.data || eventsData.data.length === 0" class="theme-bg-card rounded-3xl p-8 text-center border theme-border">
        <Icon name="ph:calendar-x-bold" class="text-5xl theme-text-muted mx-auto mb-2" />
        <p class="text-sm font-bold theme-text-muted">Belum ada agenda pertandingan.</p>
        <p class="text-xs theme-text-dim mt-1">Pengurus akan segera merilis jadwal baru!</p>
      </div>
    </div>

    <!-- POP-UP (MODAL) PEMBAYARAN INTERAKTIF -->
    <div v-if="showModal" class="fixed inset-0 bg-black/80 backdrop-blur-md z-50 flex items-end sm:items-center justify-center p-0 sm:p-4 animate-fade-in">
      <div class="theme-bg-card-solid border theme-border theme-text-main w-full max-w-md rounded-t-[32px] sm:rounded-3xl p-6 relative animate-slide-up shadow-2xl">
        
        <!-- Aksen Atas -->
        <div class="absolute top-0 left-0 right-0 h-1.5 bg-gradient-to-r from-red-600 via-orange-500 to-amber-400"></div>

        <button @click="tutupModal" class="absolute top-5 right-5 theme-text-muted hover:theme-text-main w-8 h-8 rounded-full theme-bg-surface border theme-border flex items-center justify-center transition">
          <Icon name="ph:x-bold" class="text-base" />
        </button>

        <div v-if="step === 1">
          <div class="flex items-center gap-2 mb-1">
            <span class="w-2.5 h-2.5 rounded-full bg-emerald-500"></span>
            <h3 class="text-lg font-black theme-text-main">Slot Berhasil Diamankan! 🚀</h3>
          </div>
          <p class="text-xs theme-text-muted mb-5">
            Pemain: <strong class="text-orange-500">{{ userName }}</strong>. Pilih metode pembayaran Anda:
          </p>
          
          <!-- Pilihan Metode Pembayaran -->
          <div class="mb-5 space-y-3">
            <div class="grid grid-cols-2 gap-3">
              <!-- Opsi Potong Deposit (Rekomendasi) -->
              <label 
                :class="paymentMethod === 'deposit' ? 'border-orange-500 bg-orange-500/10 text-orange-500' : 'theme-border theme-bg-surface theme-text-muted'" 
                class="border-2 rounded-2xl p-3.5 flex flex-col items-center justify-center gap-1.5 cursor-pointer transition active:scale-95 text-center">
                <input type="radio" v-model="paymentMethod" value="deposit" class="hidden" />
                <Icon name="ph:wallet-bold" class="text-2xl" />
                <span class="text-xs font-black">Saldo Deposit</span>
                <span class="text-[10px] theme-text-muted">Otomatis & Cepat</span>
              </label>

              <!-- Opsi Transfer Manual -->
              <label 
                :class="paymentMethod === 'transfer' ? 'border-orange-500 bg-orange-500/10 text-orange-500' : 'theme-border theme-bg-surface theme-text-muted'" 
                class="border-2 rounded-2xl p-3.5 flex flex-col items-center justify-center gap-1.5 cursor-pointer transition active:scale-95 text-center">
                <input type="radio" v-model="paymentMethod" value="transfer" class="hidden" />
                <Icon name="ph:bank-bold" class="text-2xl" />
                <span class="text-xs font-black">Transfer Bank</span>
                <span class="text-[10px] theme-text-muted">Upload Bukti</span>
              </label>
            </div>
          </div>

          <!-- JIKA PILIH DEPOSIT -->
          <div v-if="paymentMethod === 'deposit'" class="mb-5 p-4 theme-bg-surface border theme-border rounded-2xl space-y-2">
            <div class="flex justify-between text-xs">
              <span class="theme-text-muted">Saldo Deposit Anda:</span>
              <span class="font-bold theme-text-main">Rp {{ userData?.deposit?.toLocaleString('id-ID') || '0' }}</span>
            </div>
            <p class="text-[11px] theme-text-muted leading-relaxed">
              Saldo Anda akan dipotong otomatis & status tiket <strong class="text-emerald-500">LANGSUNG LUNAS</strong> tanpa perlu verifikasi manual!
            </p>
          </div>

          <!-- JIKA PILIH TRANSFER -->
          <div v-if="paymentMethod === 'transfer'" class="mb-5 space-y-3">
            <div class="theme-bg-surface border theme-border rounded-2xl p-3.5 text-xs space-y-1 theme-text-main">
              <p class="font-bold theme-text-main flex items-center gap-1">
                <Icon name="ph:info-bold" class="text-orange-500 text-sm" /> Rekening Pembayaran:
              </p>
              <p class="text-[11px] theme-text-main">Bank BNI: <strong class="text-amber-500">1148532156</strong></p>
              <p class="text-[11px] theme-text-muted">a/n IPAN SOPANDI</p>
            </div>

            <div>
              <label class="block text-xs font-bold theme-text-muted mb-1.5 uppercase tracking-wider">Upload Foto Resi / Bukti Transfer</label>
              <input type="file" @change="handleFileChange" accept="image/*" class="w-full text-xs theme-text-muted file:mr-3 file:py-2.5 file:px-4 file:rounded-xl file:border-0 file:text-xs file:font-black file:bg-orange-500 file:text-white hover:file:bg-orange-600 cursor-pointer" />
            </div>
          </div>
          
          <button 
            @click="submitUpload" 
            :disabled="isSubmitting || (paymentMethod === 'transfer' && !receiptFile)" 
            class="w-full bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black py-4 rounded-2xl disabled:opacity-50 transition active:scale-95 shadow-xl shadow-orange-600/30 text-sm flex items-center justify-center gap-2">
            <Icon v-if="isSubmitting" name="ph:spinner-gap-bold" class="text-lg animate-spin" />
            <span>{{ isSubmitting ? 'Memproses...' : (paymentMethod === 'transfer' ? 'Kirim Bukti Pembayaran' : 'Konfirmasi Bayar Pakai Saldo') }}</span>
          </button>
        </div>

        <!-- Step 2: Konfirmasi Sukses -->
        <div v-if="step === 2" class="text-center py-6">
          <div class="w-16 h-16 bg-emerald-500/20 text-emerald-500 rounded-full flex items-center justify-center mx-auto mb-3 border border-emerald-500/40">
            <Icon name="ph:check-bold" class="text-3xl" />
          </div>
          <h3 class="text-lg font-black theme-text-main mb-1">Pembayaran Berhasil Dicatat!</h3>
          <p class="text-xs theme-text-muted mb-6 max-w-xs mx-auto leading-relaxed">
            {{ paymentMethod === 'transfer' ? 'Bukti transfer berhasil dikirim. Admin akan segera memverifikasi tiketmu.' : 'Saldo deposit berhasil dipotong. Status tiket Anda sekarang sudah LUNAS!' }}
          </p>
          <button @click="tutupModal" class="w-full theme-bg-surface hover:opacity-90 theme-text-main font-bold py-3.5 rounded-2xl border theme-border active:scale-95 transition text-sm">
            Kembali ke Beranda
          </button>
        </div>

      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

definePageMeta({ layout: 'mobile' })

const { $api } = useNuxtApp()
const { isInstalled, isIOS, hasPrompt, showModal: showPwaModal, triggerInstall, closeModal } = usePwaInstall()
const { useAutoRefresh } = useRealtime()

const { data: eventsData, refresh: refreshEvents } = await useApiFetch('/events')
const riwayatDaftar = ref([])

const userId = ref('')
const userName = ref('')
const userData = ref({})
const showModal = ref(false)
const step = ref(1)
const registrationId = ref('')
const receiptFile = ref(null)
const isSubmitting = ref(false)

const paymentMethod = ref('deposit')

// Ambil data profil & saldo
const ambilDataUser = async (uid) => {
  const targetId = uid || userId.value
  if (!targetId) return
  try {
    const res = await $api(`/user?id=${targetId}`)
    userData.value = res.data || {}
  } catch (error) {
    console.error("Gagal mengambil data user")
  }
}

const ambilRiwayatDaftar = async (uid) => {
  const targetId = uid || userId.value
  if (!targetId) return
  try {
    const res = await $api(`/my-registrations?user_id=${targetId}`)
    riwayatDaftar.value = res.data || []
  } catch (error) {
    console.error("Gagal mengambil riwayat")
  }
}

// Pasang Auto-Refresh Realtime
useAutoRefresh(['EVENT_UPDATED', 'REGISTRATION_UPDATED', 'PAYMENT_UPDATED', 'TOPUP_UPDATED', 'USER_UPDATED'], () => {
  refreshEvents()
  if (userId.value) {
    ambilRiwayatDaftar(userId.value)
    ambilDataUser(userId.value)
  }
})

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

const getStatusReg = (idJadwal) => {
  const found = riwayatDaftar.value.find(reg => reg.event_id === idJadwal)
  if (!found) return null
  return {
    isRegistered: true,
    registrationId: found.id,
    pollingStatus: found.polling_status,
    paymentStatus: found.payment_status
  }
}

const tutupModal = () => {
  showModal.value = false
  refreshEvents()
  if (userId.value) {
    ambilRiwayatDaftar(userId.value)
    ambilDataUser(userId.value)
  }
}

const handleFileChange = (e) => {
  receiptFile.value = e.target.files[0]
}

const bukaBayarUlang = (statusObj) => {
  registrationId.value = statusObj.registrationId
  step.value = 1
  receiptFile.value = null
  paymentMethod.value = 'deposit'
  showModal.value = true
}

const toast = useToast()

const ikutMain = async (idJadwal) => {
  try {
    const res = await $api('/register', {
      method: 'POST',
      body: { event_id: idJadwal, user_id: userId.value }
    })
    registrationId.value = res.data.id
    step.value = 1 
    receiptFile.value = null
    paymentMethod.value = 'deposit'
    showModal.value = true
  } catch (error) {
    toast.error(error.response?._data?.error || 'Gagal mendaftar. Kuota mungkin penuh.', 'Pendaftaran Gagal')
  }
}

const submitUpload = async () => {
  isSubmitting.value = true
  try {
    if (paymentMethod.value === 'transfer') {
      const formData = new FormData()
      formData.append('registration_id', registrationId.value)
      formData.append('receipt', receiptFile.value)

      await $api('/upload-proof', {
        method: 'POST',
        body: formData
      })
      toast.success('Bukti transfer berhasil diunggah! Menunggu verifikasi admin.', 'Pembayaran Terkirim')
    } else {
      await $api('/pay-deposit', {
        method: 'POST',
        body: { 
          registration_id: registrationId.value, 
          user_id: userId.value 
        }
      })
      if (userId.value) ambilDataUser(userId.value)
      toast.success('Pembayaran pakai saldo deposit berhasil! Status tiket LUNAS.', 'Pembayaran Sukses')
    }
    
    step.value = 2 
  } catch (error) {
    toast.error(error.response?._data?.error || 'Gagal memproses pembayaran.', 'Pembayaran Gagal')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.2s ease-out forwards;
}
.animate-slide-up {
  animation: slideUp 0.3s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes slideUp {
  from { transform: translateY(100%); opacity: 0; }
  to { transform: translateY(0); opacity: 1; }
}
</style>