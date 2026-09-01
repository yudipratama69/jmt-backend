<template>
  <div class="min-h-screen bg-gradient-to-br from-red-600 via-orange-600 to-amber-600 flex flex-col justify-center px-6 py-12 relative overflow-hidden font-sans">
    
    <!-- Efek Dekorasi Lingkaran Latar Belakang -->
    <div class="absolute -top-20 -left-20 w-64 h-64 rounded-full bg-white/15 blur-3xl pointer-events-none"></div>
    <div class="absolute -bottom-20 -right-20 w-64 h-64 rounded-full bg-black/15 blur-3xl pointer-events-none"></div>

    <div class="w-full max-w-sm mx-auto relative z-10">
      
      <!-- Kartu Login Putih Bersih dengan Shadow Lembut -->
      <div class="bg-white/95 backdrop-blur-md rounded-3xl p-8 shadow-2xl border border-white/40">
        
        <!-- Logo JMT Sport -->
        <div class="text-center mb-6">
          <div class="w-24 h-24 mx-auto flex items-center justify-center mb-3">
            <img src="/logo-jmt.png" class="w-full h-full object-contain drop-shadow-md" alt="Logo JMT" />
          </div>
          <h2 class="text-2xl font-extrabold text-gray-800 tracking-tight">Selamat Datang!</h2>
          <p class="text-xs text-gray-500 mt-1 font-medium">Masuk untuk mulai ikut main dan cari keringat</p>
        </div>

        <!-- Form Login -->
        <form @submit.prevent="handleLogin" class="space-y-4">
          <div>
            <label class="block text-xs font-bold uppercase tracking-wider text-gray-500 mb-1.5">Email</label>
            <input v-model="form.email" type="email" placeholder="nama@gmail.com" class="w-full border border-gray-200 rounded-xl p-3.5 bg-gray-50/70 text-sm text-gray-800 outline-none focus:ring-2 focus:ring-orange-500 transition" required />
          </div>

          <div>
            <label class="block text-xs font-bold uppercase tracking-wider text-gray-500 mb-1.5">Password</label>
            <input v-model="form.password" type="password" placeholder="••••••••" class="w-full border border-gray-200 rounded-xl p-3.5 bg-gray-50/70 text-sm text-gray-800 outline-none focus:ring-2 focus:ring-orange-500 transition" required />
          </div>

          <div class="pt-2">
            <button type="submit" class="w-full bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white font-bold py-4 rounded-xl shadow-lg shadow-orange-600/30 transition duration-200 text-sm active:scale-95">
              Masuk
            </button>
          </div>
        </form>

        <!-- Link Register -->
        <div class="mt-6 text-center">
          <p class="text-xs text-gray-500">
            Belum punya akun? 
            <NuxtLink to="/player/register" class="text-orange-600 font-bold hover:underline">Daftar di sini</NuxtLink>
          </p>
        </div>

      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

definePageMeta({ layout: 'mobile' })

const form = ref({
  email: '',
  password: ''
})

const handleLogin = async () => {
  try {
    const res = await $fetch('http://localhost:8080/login', {
      method: 'POST',
      body: form.value
    })
    
    localStorage.setItem('user_id', res.data.id)
    localStorage.setItem('user_name', res.data.name)
    
    alert('Login Berhasil! 🔥')
    navigateTo('/player')
  } catch (error) {
    alert(error.response?._data?.error || 'Email atau password salah.')
  }
}
</script>