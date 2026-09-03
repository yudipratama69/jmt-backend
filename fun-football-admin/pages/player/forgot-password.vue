<template>
  <div class="py-6 px-4 flex flex-col justify-center relative font-sans transition-colors duration-300">
    
    <!-- Stadium Lights & Glow Effects -->
    <div class="absolute -top-20 -left-20 w-60 h-60 rounded-full bg-orange-600/15 blur-[80px] pointer-events-none"></div>
    <div class="absolute -bottom-20 -right-20 w-60 h-60 rounded-full bg-red-600/15 blur-[80px] pointer-events-none"></div>

    <div class="w-full max-w-sm mx-auto relative z-10 space-y-5">
      
      <!-- Brand Header -->
      <div class="text-center pt-2">
        <div class="w-20 h-20 mx-auto mb-2 bg-white rounded-3xl p-1.5 shadow-xl border-2 border-orange-500/40 flex items-center justify-center">
          <img :src="'/logo-jmt.png'" class="w-full h-full object-contain" alt="Logo JMT Sport" />
        </div>
        <h1 class="text-2xl font-black theme-text-main tracking-wide">JMT SPORT</h1>
        <p class="text-xs theme-text-muted mt-0.5">Komunitas Fun Football & Mini Soccer</p>
      </div>

      <!-- Card Reset Password -->
      <div class="theme-bg-card backdrop-blur-xl rounded-3xl p-6 border theme-border shadow-2xl space-y-5 transition-colors duration-300">
        
        <div>
          <div class="flex items-center gap-2 mb-1">
            <span class="w-7 h-7 rounded-xl bg-orange-500/10 text-orange-500 flex items-center justify-center font-bold text-sm">
              🔑
            </span>
            <h2 class="text-lg font-black theme-text-main">Setel Ulang Password</h2>
          </div>
          <p class="text-xs theme-text-muted">Masukkan email akun pemainmu dan tentukan password baru.</p>
        </div>

        <!-- Form Reset Password -->
        <form @submit.prevent="handleResetPassword" class="space-y-4">
          
          <!-- Email Input -->
          <div>
            <label class="block text-xs font-black uppercase tracking-wider theme-text-muted mb-1.5">Email Terdaftar</label>
            <div class="relative">
              <Icon name="ph:envelope-simple-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 theme-text-muted text-base" />
              <input 
                v-model="form.email" 
                type="email" 
                placeholder="nama@gmail.com" 
                class="w-full border theme-border rounded-2xl pl-10 pr-3.5 py-3.5 theme-bg-surface text-sm theme-text-main placeholder-slate-400 outline-none focus:border-orange-500 font-bold transition" 
                required 
              />
            </div>
          </div>

          <!-- Password Baru Input -->
          <div>
            <label class="block text-xs font-black uppercase tracking-wider theme-text-muted mb-1.5">Password Baru</label>
            <div class="relative">
              <Icon name="ph:lock-key-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 theme-text-muted text-base" />
              <input 
                v-model="form.new_password" 
                :type="showPassword ? 'text' : 'password'" 
                placeholder="Minimal 6 karakter" 
                minlength="6"
                class="w-full border theme-border rounded-2xl pl-10 pr-10 py-3.5 theme-bg-surface text-sm theme-text-main placeholder-slate-400 outline-none focus:border-orange-500 font-bold transition" 
                required 
              />
              <button 
                type="button" 
                @click="showPassword = !showPassword" 
                class="absolute right-3.5 top-1/2 -translate-y-1/2 theme-text-muted hover:theme-text-main text-base">
                <Icon :name="showPassword ? 'ph:eye-slash-bold' : 'ph:eye-bold'" />
              </button>
            </div>
          </div>

          <!-- Konfirmasi Password Baru Input -->
          <div>
            <label class="block text-xs font-black uppercase tracking-wider theme-text-muted mb-1.5">Konfirmasi Password Baru</label>
            <div class="relative">
              <Icon name="ph:lock-key-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 theme-text-muted text-base" />
              <input 
                v-model="form.confirm_password" 
                :type="showPassword ? 'text' : 'password'" 
                placeholder="Ulangi password baru" 
                minlength="6"
                class="w-full border theme-border rounded-2xl pl-10 pr-10 py-3.5 theme-bg-surface text-sm theme-text-main placeholder-slate-400 outline-none focus:border-orange-500 font-bold transition" 
                required 
              />
            </div>
          </div>

          <!-- Tombol Submit -->
          <div class="pt-2">
            <button 
              type="submit" 
              :disabled="isLoading"
              class="w-full bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black py-4 rounded-2xl shadow-xl shadow-orange-600/30 transition duration-200 text-sm active:scale-95 disabled:opacity-50 flex items-center justify-center gap-2">
              <Icon v-if="isLoading" name="ph:spinner-gap-bold" class="text-lg animate-spin" />
              <span>{{ isLoading ? 'Menyimpan Password...' : 'Simpan Password Baru' }}</span>
            </button>
          </div>

        </form>

        <!-- Link Back to Login -->
        <div class="pt-2 text-center border-t theme-border-subtle">
          <p class="text-xs theme-text-muted">
            Sudah ingat passwordmu? 
            <NuxtLink to="/player/login" class="text-orange-500 font-black hover:underline ml-1">Masuk di sini</NuxtLink>
          </p>
        </div>

      </div>

      <!-- Quick Back Button -->
      <div class="text-center">
        <NuxtLink to="/player/login" class="text-[11px] theme-text-muted hover:theme-text-main font-medium flex items-center justify-center gap-1">
          <Icon name="ph:arrow-left-bold" />
          <span>Kembali ke Halaman Masuk</span>
        </NuxtLink>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

definePageMeta({ layout: 'mobile' })

const { $api } = useNuxtApp()
const router = useRouter()
const toast = useToast()

const form = ref({
  email: '',
  new_password: '',
  confirm_password: ''
})

const showPassword = ref(false)
const isLoading = ref(false)

const handleResetPassword = async () => {
  if (form.value.new_password.length < 6) {
    toast.error('Password baru minimal 6 karakter!')
    return
  }

  if (form.value.new_password !== form.value.confirm_password) {
    toast.error('Konfirmasi password tidak cocok dengan password baru!')
    return
  }

  isLoading.value = true
  try {
    const res = await $api('/auth/reset-password', {
      method: 'POST',
      body: {
        email: form.value.email,
        new_password: form.value.new_password,
        confirm_password: form.value.confirm_password
      }
    })

    toast.success(res.message || 'Password berhasil direset! Silakan login dengan password baru.', 'Sukses')
    router.push('/player/login')
  } catch (error) {
    toast.error(error.response?._data?.error || 'Gagal mereset password. Pastikan email Anda sudah terdaftar.', 'Error')
  } finally {
    isLoading.value = false
  }
}
</script>
