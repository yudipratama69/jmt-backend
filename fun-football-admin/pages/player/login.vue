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

      <!-- Card Login Sporty -->
      <div class="theme-bg-card backdrop-blur-xl rounded-3xl p-6 border theme-border shadow-2xl space-y-5 transition-colors duration-300">
        
        <div>
          <h2 class="text-lg font-black theme-text-main">Masuk ke Akun</h2>
          <p class="text-xs theme-text-muted mt-0.5">Masukkan email & password akun pemainmu.</p>
        </div>

        <!-- Form Login -->
        <form @submit.prevent="handleLogin" class="space-y-4">
          <div>
            <label class="block text-xs font-black uppercase tracking-wider theme-text-muted mb-1.5">Email</label>
            <input 
              v-model="form.email" 
              type="email" 
              placeholder="nama@gmail.com" 
              class="w-full border theme-border rounded-2xl p-3.5 theme-bg-surface text-sm theme-text-main placeholder-slate-400 outline-none focus:border-orange-500 font-bold transition" 
              required />
          </div>

          <div>
            <label class="block text-xs font-black uppercase tracking-wider theme-text-muted mb-1.5">Password</label>
            <input 
              v-model="form.password" 
              type="password" 
              placeholder="••••••••" 
              class="w-full border theme-border rounded-2xl p-3.5 theme-bg-surface text-sm theme-text-main placeholder-slate-400 outline-none focus:border-orange-500 font-bold transition" 
              required />
          </div>

          <div class="pt-2">
            <button 
              type="submit" 
              :disabled="isLoading"
              class="w-full bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black py-4 rounded-2xl shadow-xl shadow-orange-600/30 transition duration-200 text-sm active:scale-95 disabled:opacity-50 flex items-center justify-center gap-2">
              <Icon v-if="isLoading" name="ph:spinner-gap-bold" class="text-lg animate-spin" />
              <span>{{ isLoading ? 'Memverifikasi...' : 'Masuk Sekarang' }}</span>
            </button>
          </div>
        </form>

        <!-- Link Register -->
        <div class="pt-2 text-center border-t theme-border-subtle">
          <p class="text-xs theme-text-muted">
            Belum punya akun pemain? 
            <NuxtLink to="/player/register" class="text-orange-500 font-black hover:underline ml-1">Daftar di sini</NuxtLink>
          </p>
        </div>

      </div>

      <!-- Quick Link to Admin Login -->
      <div class="text-center">
        <NuxtLink to="/admin-login" class="text-[11px] theme-text-muted hover:theme-text-main font-medium">
          Masuk sebagai Pengurus / Admin →
        </NuxtLink>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

definePageMeta({ layout: 'mobile' })

const { $api } = useNuxtApp()
const toast = useToast()

const form = ref({
  email: '',
  password: ''
})
const isLoading = ref(false)

const handleLogin = async () => {
  isLoading.value = true
  try {
    const res = await $api('/login', {
      method: 'POST',
      body: form.value
    })
    
    localStorage.setItem('user_id', res.data.id)
    localStorage.setItem('user_name', res.data.name)
    
    toast.success(`Selamat datang kembali, ${res.data.name}! 🔥`, 'Login Sukses')
    navigateTo('/player')
  } catch (error) {
    toast.error(error.response?._data?.error || 'Email atau password salah.', 'Gagal Masuk')
  } finally {
    isLoading.value = false
  }
}
</script>