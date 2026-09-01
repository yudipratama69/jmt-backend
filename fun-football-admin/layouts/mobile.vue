<template>
  <div class="min-h-screen bg-gray-900 flex justify-center font-sans">
    <div class="w-full max-w-md bg-gray-50 min-h-screen relative shadow-2xl flex flex-col overflow-hidden">
      
      <!-- HEADER ATAS (Hanya tampil jika bukan di halaman login / register) -->
      <header v-if="!isAuthPage" class="bg-gradient-to-r from-red-600 via-orange-600 to-amber-600 text-white p-4 sticky top-0 z-40 shadow-md rounded-b-[24px]">
        <div class="flex justify-between items-center relative">
          
          <!-- Logo & Nama Klub -->
          <div class="flex items-center gap-3">
            <div class="w-11 h-11 bg-white rounded-full p-0.5 shadow-md flex items-center justify-center overflow-hidden border-2 border-white/40">
              <img src="/logo-jmt.png" class="w-full h-full object-cover" alt="Logo JMT" />
            </div>
            <div>
              <h1 class="text-lg font-extrabold tracking-wide leading-tight">JMT Sport</h1>
              <p class="text-[10px] text-orange-100 font-medium">Kuy Cari Keringat!</p>
            </div>
          </div>
          
          <!-- Tombol Profile / Menu -->
          <div class="relative">
            <button @click="isMenuOpen = !isMenuOpen" class="relative z-50 w-10 h-10 bg-white rounded-full overflow-hidden flex items-center justify-center text-orange-600 shadow-md hover:bg-orange-50 transition active:scale-95 border-2 border-white/80">
              <img v-if="userPhoto" :src="userPhoto" class="w-full h-full object-cover" />
              <!-- Berikan warna tegas (text-orange-600) dan ukuran yang pas agar terlihat jelas -->
              <Icon v-else name="ph:user-bold" class="text-xl text-orange-600" />
            </button>

            <!-- Dropdown Menu -->
            <div v-if="isMenuOpen" class="absolute right-0 mt-3 w-48 bg-white rounded-xl shadow-xl border border-gray-100 py-2 z-50 animate-fade-in-down overflow-hidden">
              <button @click="bukaPengaturan" class="w-full text-left px-4 py-3 text-sm text-gray-700 hover:bg-gray-50 flex items-center gap-3 transition">
                <Icon name="ph:gear-bold" class="text-lg text-gray-400" />
                Pengaturan Profile
              </button>
              
              <div class="border-t border-gray-100 my-1"></div>
              
              <button @click="logout" class="w-full text-left px-4 py-3 text-sm text-red-600 hover:bg-red-50 flex items-center gap-3 font-medium transition">
                <Icon name="ph:sign-out-bold" class="text-lg text-red-500" />
                Keluar (Logout)
              </button>
            </div>
          </div>

        </div>
      </header>

      <div v-if="isMenuOpen" @click="isMenuOpen = false" class="fixed inset-0 z-30"></div>

      <!-- Area Konten -->
      <main class="flex-1 overflow-y-auto pb-24 relative z-10">
        <slot />
      </main>

      <!-- Bottom Navigation Bar dengan SVG Langsung (Pasti Muncul) -->
      <nav class="absolute bottom-0 w-full bg-white border-t border-gray-100 grid grid-cols-4 items-center p-3 pb-5 shadow-[0_-10px_15px_-3px_rgba(0,0,0,0.05)] rounded-t-[24px] z-20">
        
  <!-- Menu Home -->
  <NuxtLink to="/player" class="flex flex-col items-center text-gray-400 hover:text-orange-600 transition-transform active:scale-95" exact-active-class="text-orange-600 font-bold">
    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 mb-1" fill="currentColor" viewBox="0 0 256 256">
      <path d="M218.12,103.87l-80-72a12,12,0,0,0-16.24,0l-80,72A12,12,0,0,0,36,121.31V216a12,12,0,0,0,12,12H96a12,12,0,0,0,12-12V160h32v56a12,12,0,0,0,12,12h48a12,12,0,0,0,12-12V121.31A12,12,0,0,0,218.12,103.87Z"></path>
    </svg>
    <span class="text-[10px]">Home</span>
  </NuxtLink>

  <!-- Menu Jadwalku -->
  <NuxtLink to="/player/jadwalku" class="flex flex-col items-center text-gray-400 hover:text-orange-600 transition-transform active:scale-95" exact-active-class="text-orange-600 font-bold">
    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 mb-1" fill="currentColor" viewBox="0 0 256 256">
      <path d="M216,48H176V32a12,12,0,0,0-24,0V48H104V32a12,12,0,0,0-24,0V48H40A16,16,0,0,0,24,64V208a16,16,0,0,0,16,16H216a16,16,0,0,0,16-16V64A16,16,0,0,0,216,48ZM204,204H52V80H204ZM92,112a12,12,0,0,1,12-12h48a12,12,0,0,1,0,24H104A12,12,0,0,1,92,112Zm0,40a12,12,0,0,1,12-12h48a12,12,0,0,1,0,24H104A12,12,0,0,1,92,152Z"></path>
    </svg>
    <span class="text-[10px]">Jadwalku</span>
  </NuxtLink>

  <!-- Menu Keuangan -->
  <NuxtLink to="/player/keuangan" class="flex flex-col items-center text-gray-400 hover:text-orange-600 transition-transform active:scale-95" exact-active-class="text-orange-600 font-bold">
    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 mb-1" fill="currentColor" viewBox="0 0 256 256">
      <path d="M224,72H40A16,16,0,0,0,24,88V184a16,16,0,0,0,16,16H224a16,16,0,0,0,16-16V88A16,16,0,0,0,224,72Zm-8,112H40V88H216Zm-24-40a12,12,0,1,1-12-12A12,12,0,0,1,192,144Z"></path>
    </svg>
    <span class="text-[10px]">Keuangan</span>
  </NuxtLink>

  <!-- Menu Pemain Baru -->
  <NuxtLink to="/player/pemain" class="flex flex-col items-center text-gray-400 hover:text-orange-600 transition-transform active:scale-95" exact-active-class="text-orange-600 font-bold">
    <Icon name="ph:users-bold" class="w-6 h-6 mb-1" />
    <span class="text-[10px]">Pemain</span>
  </NuxtLink>

</nav>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const isMenuOpen = ref(false)
const userPhoto = ref(null)

// Deteksi apakah sedang berada di halaman login atau register player
const isAuthPage = computed(() => {
  return route.path.includes('/player/login') || route.path.includes('/player/register')
})

// Fungsi untuk mengambil data profil user terbaru
const fetchUserData = async () => {
  const userId = localStorage.getItem('user_id')
  if (userId) {
    try {
      const res = await $fetch(`http://localhost:8080/user?id=${userId}`)
      // Pastikan merujuk ke field profile_pic yang dikembalikan oleh backend
      if (res.data && res.data.profile_pic) {
        // Jika path foto di database sudah lengkap dengan http, gunakan langsung. 
        // Jika berupa relative path (misal: /uploads/xxx.jpg), gabungkan dengan domain backend.
        userPhoto.value = res.data.profile_pic.startsWith('http') 
          ? res.data.profile_pic 
          : `http://localhost:8080${res.data.profile_pic}`
      }
    } catch (error) {
      console.error("Gagal memuat foto profil header", error)
    }
  }
}

onMounted(() => {
  fetchUserData()
})

const bukaPengaturan = () => {
  isMenuOpen.value = false 
  navigateTo('/player/profile')
}

const logout = () => {
  isMenuOpen.value = false 
  if (confirm('Apakah Anda yakin ingin keluar dari akun?')) {
    localStorage.removeItem('user_id')
    localStorage.removeItem('user_name')
    navigateTo('/player/login')
  }
}
</script>

<style scoped>
.animate-fade-in-down {
  animation: fadeInDown 0.2s ease-out forwards;
}
@keyframes fadeInDown {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>