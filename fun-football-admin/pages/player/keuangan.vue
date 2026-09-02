<template>
  <div class="p-4 space-y-6">
    
    <!-- Header -->
    <div>
      <h2 class="text-xl font-black theme-text-main tracking-wide flex items-center gap-2">
        <Icon name="ph:wallet-bold" class="text-orange-500 text-2xl" /> Saldo & Deposit
      </h2>
      <p class="text-xs theme-text-muted mt-0.5">Deposit saldo untuk bayar instan tanpa upload resi berulang.</p>
    </div>

    <!-- Sporty VIP Card (Gradasi Metallic Flame) -->
    <div class="bg-gradient-to-tr from-slate-950 via-slate-900 to-orange-950 rounded-3xl p-6 text-white shadow-2xl border border-orange-500/30 relative overflow-hidden">
      <!-- Glow effect -->
      <div class="absolute -right-10 -top-10 w-40 h-40 bg-orange-500/20 rounded-full blur-3xl pointer-events-none"></div>
      <Icon name="ph:credit-card-bold" class="absolute -right-4 -bottom-4 text-9xl text-white/5 transform -rotate-12 pointer-events-none" />

      <div class="relative z-10 flex flex-col justify-between h-36">
        <div class="flex justify-between items-start">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-amber-500 via-orange-500 to-amber-600 text-white flex items-center justify-center shadow-lg border border-amber-300/40 shrink-0">
              <Icon name="ph:wallet-bold" class="text-2xl" />
            </div>
            <div>
              <p class="text-[10px] text-slate-300 uppercase font-black tracking-wider">Dompet Saldo JMT</p>
              <p class="text-xs font-bold text-white">{{ userData?.name || 'Member' }}</p>
            </div>
          </div>
          <span class="bg-emerald-500/20 border border-emerald-500/40 text-emerald-400 text-[10px] font-black px-2.5 py-0.5 rounded-full uppercase">
            ACTIVE
          </span>
        </div>

        <div>
          <p class="text-[10px] text-orange-300 font-black uppercase tracking-wider mb-0.5">Total Saldo Tersedia</p>
          <h3 class="text-3xl font-black text-white tracking-tight">
            Rp {{ userData?.deposit?.toLocaleString('id-ID') || '0' }}
          </h3>
        </div>
      </div>
    </div>

    <!-- Form Top Up Saldo -->
    <div class="theme-bg-card backdrop-blur-md rounded-3xl p-5 border theme-border shadow-xl space-y-4 transition-colors duration-300">
      <h4 class="font-black theme-text-main text-base flex items-center gap-2">
        <Icon name="ph:plus-circle-bold" class="text-orange-500 text-xl" /> Isi Saldo Deposit
      </h4>

      <!-- Quick Preset Nominal Buttons -->
      <div>
        <label class="block text-[11px] font-black uppercase tracking-wider theme-text-muted mb-2">Pilih Nominal Cepat</label>
        <div class="grid grid-cols-4 gap-2">
          <button 
            v-for="preset in [25000, 50000, 100000, 150000]" 
            :key="preset"
            type="button"
            @click="form.amount = preset"
            :class="form.amount == preset ? 'bg-orange-500 text-white border-orange-400 shadow-lg shadow-orange-600/30' : 'theme-bg-surface theme-text-main theme-border hover:border-orange-500/50'"
            class="py-2.5 px-2 rounded-xl text-xs font-black border transition active:scale-95 text-center">
            {{ preset / 1000 }}K
          </button>
        </div>
      </div>

      <form @submit.prevent="submitTopUp" class="space-y-4">
        <div>
          <label class="block text-[11px] font-black uppercase tracking-wider theme-text-muted mb-1.5">Atau Masukkan Nominal Sendiri (Rp)</label>
          <input 
            v-model="form.amount" 
            type="number" 
            min="10000" 
            placeholder="Contoh: 50000" 
            class="w-full border theme-border rounded-2xl p-3.5 theme-bg-surface text-sm theme-text-main placeholder-slate-400 outline-none focus:border-orange-500 transition font-bold" 
            required />
        </div>

        <div>
          <label class="block text-[11px] font-black uppercase tracking-wider theme-text-muted mb-1.5">Upload Bukti Transfer</label>
          <input 
            type="file" 
            @change="handleFileChange" 
            accept="image/*" 
            class="w-full text-xs theme-text-muted file:mr-3 file:py-2.5 file:px-4 file:rounded-xl file:border-0 file:text-xs file:font-black file:bg-orange-500 file:text-white hover:file:bg-orange-600 cursor-pointer" 
            required />
        </div>

        <button 
          type="submit" 
          :disabled="isLoading" 
          class="w-full bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black py-4 rounded-2xl shadow-xl shadow-orange-600/30 transition-all active:scale-95 disabled:opacity-50 text-sm flex items-center justify-center gap-2">
          <Icon v-if="isLoading" name="ph:spinner-gap-bold" class="text-lg animate-spin" />
          <span>{{ isLoading ? 'Mengirim Bukti...' : 'Kirim Permintaan Isi Saldo' }}</span>
        </button>
      </form>
    </div>

    <!-- Info Rekening Transfer Resmi -->
    <div class="theme-bg-card border theme-border rounded-3xl p-5 space-y-3 transition-colors duration-300">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-2 theme-text-main font-bold text-sm">
          <Icon name="ph:bank-bold" class="text-orange-500 text-xl" />
          <span>Rekening Tujuan Transfer</span>
        </div>
        <span class="text-[10px] theme-bg-surface theme-text-main font-bold px-2.5 py-0.5 rounded-full border theme-border">BNI</span>
      </div>

      <div class="theme-bg-surface p-3.5 rounded-2xl border theme-border flex justify-between items-center transition-colors duration-300">
        <div>
          <p class="text-[10px] theme-text-muted font-bold uppercase tracking-wider">Nomor Rekening</p>
          <p class="text-base font-black text-amber-500 font-mono tracking-wider mt-0.5">1148532156</p>
          <p class="text-xs theme-text-muted mt-0.5">a/n IPAN SOPANDI</p>
        </div>
        
        <button 
          type="button"
          @click="salinRekening" 
          class="px-3 py-1.5 theme-bg-card hover:opacity-90 text-orange-500 text-xs font-black rounded-xl border theme-border transition active:scale-95 flex items-center gap-1">
          <Icon :name="copied ? 'ph:check-bold' : 'ph:copy-bold'" class="text-sm" />
          <span>{{ copied ? 'Tersalin!' : 'Salin' }}</span>
        </button>
      </div>

      <p class="text-[11px] theme-text-muted leading-relaxed">
        💡 Setelah transfer berhasil, unggah struk/screenshot pembayaran melalui form di atas. Admin akan segera menyetujui deposit Anda.
      </p>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

definePageMeta({ layout: 'mobile' })

const { $api } = useNuxtApp()
const { useAutoRefresh } = useRealtime()

const userId = ref('')
const userData = ref({})
const isLoading = ref(false)
const copied = ref(false)

const form = ref({
  amount: 50000,
  receipt: null
})

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

// Pasang Auto-Refresh Realtime
useAutoRefresh(['TOPUP_UPDATED', 'PAYMENT_UPDATED'], () => {
  if (userId.value) {
    ambilDataUser(userId.value)
  }
})

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

const toast = useToast()

const salinRekening = async () => {
  const success = await copyToClipboard('1148532156')
  if (success) {
    copied.value = true
    toast.success('Nomor rekening BNI berhasil disalin!', 'Tersalin')
    setTimeout(() => {
      copied.value = false
    }, 2000)
  } else {
    toast.info('Nomor Rekening: 1148532156 (BNI a/n IPAN SOPANDI)', 'Salin Manual')
  }
}

const submitTopUp = async () => {
  if (!form.value.receipt) {
    return toast.warning('Harap unggah bukti transfer pembayaran!', 'Bukti Kosong')
  }

  isLoading.value = true
  try {
    const formData = new FormData()
    formData.append('user_id', userId.value)
    formData.append('amount', form.value.amount)
    formData.append('receipt', form.value.receipt)

    await $api('/request-topup', {
      method: 'POST',
      body: formData
    })

    toast.success('Permintaan isi saldo berhasil dikirim! Menunggu verifikasi admin.', 'Top Up Diproses')
    form.value.receipt = null
    if (userId.value) ambilDataUser(userId.value)
  } catch (error) {
    toast.error(error.response?._data?.error || 'Gagal mengirim permintaan deposit.', 'Top Up Gagal')
  } finally {
    isLoading.value = false
  }
}
</script>