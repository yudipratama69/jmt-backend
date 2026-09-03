<template>
  <div class="h-[100dvh] w-full theme-bg-page flex justify-center font-sans antialiased selection:bg-orange-500 selection:text-white transition-colors duration-300 overflow-hidden">
    <div class="w-full max-w-md theme-bg-surface theme-text-main h-[100dvh] flex flex-col relative overflow-hidden border-x theme-border transition-colors duration-300">
      
      <!-- HEADER ATAS (Fixed di atas kontainer, tidak bergeser saat scrolling) -->
      <header v-if="!isAuthPage" class="shrink-0 bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 text-white px-4 py-3 z-30 shadow-md shadow-orange-950/20 rounded-b-[24px]">
        <div class="flex justify-between items-center relative">
          
          <!-- Logo & Brand JMT Sport -->
          <NuxtLink to="/player" class="flex items-center gap-2.5 active:scale-95 transition">
            <div class="w-10 h-10 bg-white rounded-full p-0.5 shadow-md flex items-center justify-center overflow-hidden border-2 border-white/60 shrink-0">
              <img :src="'/logo-jmt.png'" class="w-full h-full object-contain" alt="Logo JMT" />
            </div>
            <div>
              <div class="flex items-center gap-1">
                <h1 class="text-sm font-black tracking-wide leading-tight text-white drop-shadow">JMT SPORT</h1>
                <span class="bg-black/30 backdrop-blur-sm text-[8px] font-black px-1.5 py-0.2 rounded text-amber-200 border border-white/20 uppercase">PRO</span>
              </div>
              <p class="text-[10px] text-orange-100 font-medium leading-none mt-0.5">Fun Football ⚽</p>
            </div>
          </NuxtLink>
          
          <!-- Menu Kanan: Theme Toggle, Saldo Chip & Profil Avatar -->
          <div class="flex items-center gap-1.5">
            <!-- Tombol Lonceng Notifikasi -->
            <button 
              @click="isDrawerOpen = true" 
              title="Pusat Notifikasi"
              class="w-8 h-8 rounded-full bg-black/25 hover:bg-black/35 backdrop-blur-md border border-white/20 flex items-center justify-center text-amber-200 active:scale-95 transition relative">
              <Icon name="ph:bell-bold" class="text-sm" />
              <span 
                v-if="unreadCount > 0" 
                class="absolute -top-1 -right-1 bg-red-500 text-white font-black text-[9px] w-4 h-4 rounded-full flex items-center justify-center border border-white animate-pulse">
                {{ unreadCount > 9 ? '9+' : unreadCount }}
              </span>
            </button>

            <!-- Theme Toggle Button -->
            <button 
              @click="toggleTheme" 
              :title="isDark ? 'Ganti ke Light Mode' : 'Ganti ke Dark Mode'"
              class="w-8 h-8 rounded-full bg-black/25 hover:bg-black/35 backdrop-blur-md border border-white/20 flex items-center justify-center text-amber-200 active:scale-95 transition">
              <Icon :name="isDark ? 'ph:sun-bold' : 'ph:moon-bold'" class="text-sm" />
            </button>

            <!-- Tombol Avatar Profile / Menu dengan Stacking Context yang Benar -->
            <div class="relative">
              <button 
                type="button"
                @click.stop="isMenuOpen = !isMenuOpen" 
                class="w-9 h-9 bg-white rounded-full overflow-hidden flex items-center justify-center text-orange-600 shadow-md hover:ring-2 hover:ring-white transition active:scale-95 border-2 border-white/80">
                <img v-if="userPhoto" :src="userPhoto" class="w-full h-full object-cover" />
                <Icon v-else name="ph:user-bold" class="text-lg text-orange-600" />
              </button>

              <!-- Backdrop Penutup Menu saat Klik Luar (Di bawah dropdown, z-40) -->
              <div 
                v-if="isMenuOpen" 
                @click="isMenuOpen = false" 
                class="fixed inset-0 z-40 cursor-default">
              </div>

              <!-- Dropdown Menu (Z-50 Berada di atas backdrop) -->
              <div 
                v-if="isMenuOpen" 
                class="absolute right-0 mt-3 w-56 theme-bg-card-solid theme-text-main rounded-2xl shadow-2xl border theme-border py-2 z-50 animate-fade-in-down overflow-hidden">
                <div class="px-4 py-2.5 border-b theme-border-subtle">
                  <p class="text-[10px] theme-text-muted uppercase tracking-wider font-bold">Akun Pemain</p>
                  <p class="text-sm font-bold theme-text-main truncate">{{ userName || 'Member' }}</p>
                </div>

                <button 
                  type="button"
                  @click="bukaPengaturan" 
                  class="w-full text-left px-4 py-3 text-xs theme-text-main hover:bg-orange-500/10 flex items-center gap-3 transition font-medium">
                  <Icon name="ph:gear-bold" class="text-base theme-text-muted" />
                  Pengaturan Profile
                </button>

                <button 
                  v-if="!isInstalled"
                  type="button"
                  @click="bukaInstallModal" 
                  class="w-full text-left px-4 py-3 text-xs text-orange-500 hover:bg-orange-500/10 flex items-center gap-3 font-bold transition">
                  <Icon name="ph:download-simple-bold" class="text-base text-orange-500" />
                  Pasang Aplikasi di HP
                </button>

                <button 
                  type="button"
                  @click="gantiTema" 
                  class="w-full text-left px-4 py-3 text-xs text-amber-400 hover:bg-amber-500/10 flex items-center gap-3 font-bold transition">
                  <Icon :name="isDark ? 'ph:sun-bold' : 'ph:moon-bold'" class="text-base text-amber-400" />
                  Mode: {{ isDark ? 'Terang (Light)' : 'Gelap (Dark)' }}
                </button>
                
                <div class="border-t theme-border-subtle my-1"></div>
                
                <button 
                  type="button"
                  @click="logout" 
                  class="w-full text-left px-4 py-3 text-xs text-red-400 hover:bg-red-500/10 flex items-center gap-3 font-semibold transition">
                  <Icon name="ph:sign-out-bold" class="text-base text-red-400" />
                  Keluar (Logout)
                </button>
              </div>
            </div>
          </div>

        </div>
      </header>

      <!-- Modal Panduan Install PWA -->
      <PwaInstallModal 
        :show="showPwaModal" 
        :is-i-o-s="isIOS" 
        :has-prompt="hasPrompt"
        @close="closeModal" 
        @install="installPwa" 
      />

      <!-- BANNER PASANG APLIKASI (Hanya jika belum diinstall) -->
      <div 
        v-if="showInstallBanner && !isAuthPage && !isInstalled" 
        class="bg-gradient-to-r from-orange-600 via-amber-600 to-orange-500 text-white px-3.5 py-2.5 flex items-center justify-between shadow-inner text-xs shrink-0 transition-all duration-300">
        <div class="flex items-center gap-2">
          <div class="w-7 h-7 rounded-lg bg-white p-0.5 flex items-center justify-center shrink-0">
            <img :src="'/logo-jmt.png'" class="w-full h-full object-contain rounded-md" alt="Icon" />
          </div>
          <div class="leading-tight">
            <p class="font-bold text-[11px]">Pasang Aplikasi JMT Sport</p>
            <p class="text-[9px] text-orange-100">Bebas buka browser & lebih cepat ⚡</p>
          </div>
        </div>
        <div class="flex items-center gap-1.5">
          <button 
            @click="triggerInstall" 
            class="px-2.5 py-1 bg-white text-orange-600 font-extrabold rounded-lg text-[10px] shadow active:scale-95 transition">
            Pasang
          </button>
          <button 
            @click="dismissInstall" 
            class="text-orange-200 hover:text-white p-1 text-xs">
            ✕
          </button>
        </div>
      </div>

      <!-- MAIN CONTENT AREA (Scroll strictly inside main) -->
      <main class="flex-1 overflow-y-auto overscroll-contain mobile-content-area transition-colors duration-300">
        <slot />
      </main>

      <!-- BOTTOM NAVIGATION DOCK (PINNED FIXED AT BOTTOM) -->
      <nav 
        v-if="!isAuthPage" 
        class="mobile-bottom-dock grid grid-cols-4 items-center text-center select-none shadow-2xl transition-colors duration-300">
        
        <!-- Menu Home -->
        <NuxtLink to="/player" class="flex flex-col items-center py-1 theme-text-muted hover:text-orange-500 transition-all duration-200 active:scale-95 group" exact-active-class="!text-orange-500 font-bold">
          <div class="p-1 rounded-xl group-[.router-link-exact-active]:bg-orange-500/10 group-[.router-link-exact-active]:text-orange-500 transition">
            <Icon name="ph:squares-four-bold" class="text-2xl mb-0.5" />
          </div>
          <span class="text-[10px] tracking-tight">Home</span>
        </NuxtLink>

        <!-- Menu Jadwalku -->
        <NuxtLink to="/player/jadwalku" class="flex flex-col items-center py-1 theme-text-muted hover:text-orange-500 transition-all duration-200 active:scale-95 group" exact-active-class="!text-orange-500 font-bold">
          <div class="p-1 rounded-xl group-[.router-link-exact-active]:bg-orange-500/10 group-[.router-link-exact-active]:text-orange-500 transition">
            <Icon name="ph:calendar-check-bold" class="text-2xl mb-0.5" />
          </div>
          <span class="text-[10px] tracking-tight">Jadwalku</span>
        </NuxtLink>

        <!-- Menu Keuangan -->
        <NuxtLink to="/player/keuangan" class="flex flex-col items-center py-1 theme-text-muted hover:text-orange-500 transition-all duration-200 active:scale-95 group" exact-active-class="!text-orange-500 font-bold">
          <div class="p-1 rounded-xl group-[.router-link-exact-active]:bg-orange-500/10 group-[.router-link-exact-active]:text-orange-500 transition">
            <Icon name="ph:wallet-bold" class="text-2xl mb-0.5" />
          </div>
          <span class="text-[10px] tracking-tight">Keuangan</span>
        </NuxtLink>

        <!-- Menu Pemain -->
        <NuxtLink to="/player/pemain" class="flex flex-col items-center py-1 theme-text-muted hover:text-orange-500 transition-all duration-200 active:scale-95 group" exact-active-class="!text-orange-500 font-bold">
          <div class="p-1 rounded-xl group-[.router-link-exact-active]:bg-orange-500/10 group-[.router-link-exact-active]:text-orange-500 transition">
            <Icon name="ph:users-three-bold" class="text-2xl mb-0.5" />
          </div>
          <span class="text-[10px] tracking-tight">Pemain</span>
        </NuxtLink>

      </nav>
    </div>

    <!-- Drawer Notifikasi Broadcast & Pengingat -->
    <NotificationDrawer />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useNotification } from '~/composables/useNotification'

const route = useRoute()
const isMenuOpen = ref(false)
const userPhoto = ref(null)
const userName = ref('')
const userDeposit = ref(0)

const { $api } = useNuxtApp()
const { isInstallable, isInstalled, isIOS, hasPrompt, showModal: showPwaModal, triggerInstall, closeModal } = usePwaInstall()
const { useAutoRefresh, on } = useRealtime()
const { isDrawerOpen, unreadCount, fetchNotifications, handleIncomingNotification } = useNotification()

const showInstallBanner = ref(false)

// Deteksi apakah sedang berada di halaman login atau register player
const isAuthPage = computed(() => {
  return route.path.includes('/player/login') || route.path.includes('/player/register')
})

// Fungsi untuk mengambil data profil user terbaru
const fetchUserData = async () => {
  const userId = localStorage.getItem('user_id')
  if (userId) {
    try {
      const res = await $api(`/user?id=${userId}`)
      if (res.data) {
        userName.value = res.data.name || ''
        userDeposit.value = res.data.deposit || 0
        if (res.data.profile_pic) {
          userPhoto.value = useApiUrl(res.data.profile_pic)
        }
      }
    } catch (error) {
      console.error("Gagal memuat profil header", error)
    }
  }
}

// Pasang Auto-Refresh Realtime di layout
useAutoRefresh(['TOPUP_UPDATED', 'PAYMENT_UPDATED', 'USER_UPDATED'], () => {
  fetchUserData()
})

onMounted(() => {
  fetchUserData()
  fetchNotifications()
  if (!localStorage.getItem('pwa_dismissed') && !isInstalled.value) {
    showInstallBanner.value = true
  }

  // Tangkap notifikasi broadcast realtime
  on('BROADCAST_NOTIFICATION', (payload) => {
    if (payload) {
      handleIncomingNotification(payload)
    }
  })
})

const installPwa = () => {
  triggerInstall()
}

const dismissInstall = () => {
  showInstallBanner.value = false
  localStorage.setItem('pwa_dismissed', 'true')
}

const bukaInstallModal = () => {
  isMenuOpen.value = false
  triggerInstall()
}

const bukaPengaturan = () => {
  isMenuOpen.value = false 
  navigateTo('/player/profile')
}

const { isDark, toggleTheme } = useTheme()
const toast = useToast()

const gantiTema = () => {
  toggleTheme()
  isMenuOpen.value = false
}

const logout = () => {
  isMenuOpen.value = false 
  toast.confirm({
    title: 'Keluar Akun',
    message: 'Apakah Anda yakin ingin keluar dari akun pemain Anda?',
    confirmText: 'Ya, Keluar',
    cancelText: 'Batal',
    onConfirm: () => {
      localStorage.removeItem('user_id')
      localStorage.removeItem('user_name')
      toast.info('Anda telah keluar dari akun')
      navigateTo('/player/login')
    }
  })
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