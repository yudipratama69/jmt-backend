<template>
  <div class="min-h-screen bg-gradient-to-br from-red-600 via-orange-600 to-amber-600 flex flex-col justify-center px-6 py-12 relative overflow-hidden font-sans">
    
    <!-- Efek Dekorasi Lingkaran Latar Belakang -->
    <div class="absolute -top-20 -left-20 w-64 h-64 rounded-full bg-white/15 blur-3xl pointer-events-none"></div>
    <div class="absolute -bottom-20 -right-20 w-64 h-64 rounded-full bg-black/15 blur-3xl pointer-events-none"></div>

    <div class="w-full max-w-sm mx-auto relative z-10">
      
      <!-- Kartu Register Putih Bersih -->
      <div class="bg-white/95 backdrop-blur-md rounded-3xl p-8 shadow-2xl border border-white/40">
        
        <!-- Logo JMT Sport -->
        <div class="text-center mb-6">
          <div class="w-20 h-20 mx-auto flex items-center justify-center mb-2">
            <img src="/logo-jmt.png" class="w-full h-full object-contain drop-shadow-md" alt="Logo JMT" />
          </div>
          <h2 class="text-2xl font-extrabold text-gray-800 tracking-tight">Buat Akun Baru</h2>
          <p class="text-xs text-gray-500 mt-1 font-medium">Daftar untuk mulai ikut main dan daftar jadwal</p>
        </div>

        <!-- Form Register -->
        <form @submit.prevent="handleRegister" class="space-y-4">
          <div>
            <label class="block text-xs font-bold uppercase tracking-wider text-gray-500 mb-1.5">Nama Lengkap</label>
            <input v-model="form.name" type="text" placeholder="Nama Anda" class="w-full border border-gray-200 rounded-xl p-3.5 bg-gray-50/70 text-sm text-gray-800 outline-none focus:ring-2 focus:ring-orange-500 transition" required />
          </div>

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
              Daftar Sekarang
            </button>
          </div>
        </form>

        <!-- Link ke Login -->
        <div class="mt-6 text-center">
          <p class="text-xs text-gray-500">
            Sudah punya akun? 
            <NuxtLink to="/player/login" class="text-orange-600 font-bold hover:underline">Masuk di sini</NuxtLink>
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
  name: '',
  email: '',
  password: ''
})

const handleRegister = async () => {
  try {
    await $fetch('http://localhost:8080/register-user', {
      method: 'POST',
      body: form.value
    })
    
    alert('Registrasi Berhasil! Silakan masuk.')
    navigateTo('/player/login')
  } catch (error) {
    alert(error.response?._data?.error || 'Gagal mendaftarkan akun.')
  }
}
</script>