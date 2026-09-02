<template>
  <div>
    <!-- Backdrop Overlay -->
    <div 
      v-if="isDrawerOpen" 
      @click="isDrawerOpen = false" 
      class="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 transition-opacity animate-fade-in">
    </div>

    <!-- Drawer Panel -->
    <div 
      :class="isDrawerOpen ? 'translate-x-0' : 'translate-x-full'"
      class="fixed top-0 right-0 bottom-0 w-full max-w-sm sm:max-w-md theme-bg-card-solid border-l theme-border z-50 shadow-2xl transition-transform duration-300 ease-out flex flex-col">
      
      <!-- Top Header -->
      <div class="p-4 border-b theme-border flex items-center justify-between theme-bg-surface shrink-0">
        <div class="flex items-center gap-2.5">
          <div class="w-9 h-9 rounded-xl bg-orange-500/10 border border-orange-500/30 text-orange-500 flex items-center justify-center font-black">
            <Icon name="ph:bell-ringing-bold" class="text-xl" />
          </div>
          <div>
            <h3 class="font-black text-sm theme-text-main tracking-tight flex items-center gap-1.5">
              Pusat Notifikasi
              <span v-if="unreadCount > 0" class="bg-red-500 text-white text-[10px] font-black px-2 py-0.2 rounded-full animate-pulse">
                {{ unreadCount }} Baru
              </span>
            </h3>
            <p class="text-[11px] theme-text-muted">Pengumuman & Pengingat Realtime</p>
          </div>
        </div>

        <button 
          @click="isDrawerOpen = false" 
          class="w-8 h-8 rounded-full theme-bg-surface border theme-border flex items-center justify-center theme-text-muted hover:theme-text-main transition active:scale-95">
          <Icon name="ph:x-bold" class="text-sm" />
        </button>
      </div>

      <!-- Banner Izin Push Notification Jika Belum Diaktifkan -->
      <div v-if="!hasPermission" class="p-3.5 mx-4 mt-4 rounded-2xl bg-gradient-to-r from-orange-500/15 via-amber-500/15 to-orange-500/15 border border-orange-500/30 shrink-0 space-y-2">
        <div class="flex items-start gap-2.5">
          <Icon name="ph:bell-simple-z-bold" class="text-orange-500 text-lg shrink-0 mt-0.5" />
          <div class="text-xs">
            <p class="font-black theme-text-main">Aktifkan Notifikasi Pop-up</p>
            <p class="theme-text-muted text-[11px] mt-0.5">Dapatkan peringatan jadwal main & info deposit langsung di bar atas HP Anda.</p>
          </div>
        </div>
        <button 
          @click="aktifkanIzinNotif"
          class="w-full py-2 bg-gradient-to-r from-orange-500 to-amber-500 hover:from-orange-600 hover:to-amber-600 text-white font-black text-xs rounded-xl shadow-md transition active:scale-95 flex items-center justify-center gap-1.5">
          <Icon name="ph:check-circle-bold" class="text-sm" /> Izinkan Notifikasi Pop-up
        </button>
      </div>

      <!-- Action Bar (Tandai Semua Dibaca) -->
      <div v-if="notifications.length > 0" class="px-4 py-2 flex justify-between items-center text-[11px] shrink-0 border-b theme-border-subtle mt-2">
        <span class="theme-text-muted font-semibold">{{ notifications.length }} Total Pemberitahuan</span>
        <button 
          v-if="unreadCount > 0"
          @click="markAllAsRead" 
          class="text-orange-500 hover:text-orange-600 font-black flex items-center gap-1 active:scale-95 transition">
          <Icon name="ph:checks-bold" class="text-sm" /> Tandai Semua Dibaca
        </button>
      </div>

      <!-- List Notifikasi Scrollable -->
      <div class="flex-1 overflow-y-auto p-4 space-y-3">
        <!-- Empty State -->
        <div v-if="notifications.length === 0" class="py-16 text-center space-y-3">
          <div class="w-16 h-16 rounded-full theme-bg-surface border theme-border mx-auto flex items-center justify-center text-orange-500/50">
            <Icon name="ph:bell-slash-bold" class="text-3xl" />
          </div>
          <div>
            <h4 class="font-bold text-sm theme-text-main">Belum Ada Notifikasi</h4>
            <p class="text-xs theme-text-muted mt-1 max-w-[220px] mx-auto">Pengumuman jadwal baru, info deposit, dan broadcast admin akan muncul di sini.</p>
          </div>
        </div>

        <!-- Looping Cards Notifikasi -->
        <div 
          v-for="notif in notifications" 
          :key="notif.id"
          @click="markAsRead(notif.id)"
          :class="notif.isRead ? 'opacity-80' : 'border-orange-500/40 shadow-lg ring-1 ring-orange-500/20'"
          class="theme-bg-surface border theme-border rounded-2xl p-4 transition-all duration-200 relative group cursor-pointer hover:border-orange-500/50">
          
          <!-- Indikator Belum Dibaca -->
          <span 
            v-if="!notif.isRead" 
            class="absolute top-3 right-3 w-2 h-2 rounded-full bg-orange-500 animate-ping"></span>
          <span 
            v-if="!notif.isRead" 
            class="absolute top-3 right-3 w-2 h-2 rounded-full bg-orange-500"></span>

          <!-- Header Card: Badge Kategori & Waktu -->
          <div class="flex items-center gap-2 mb-2 pr-4">
            <span 
              :class="getTypeBadgeClass(notif.type)"
              class="text-[10px] font-black px-2.5 py-0.5 rounded-full uppercase tracking-wider flex items-center gap-1">
              <Icon :name="getTypeIcon(notif.type)" class="text-xs" />
              {{ notif.type || 'INFO' }}
            </span>
            <span class="text-[10px] theme-text-muted font-mono">
              {{ formatWaktu(notif.created_at) }}
            </span>
          </div>

          <!-- Judul & Pesan -->
          <h4 class="font-black text-xs theme-text-main mb-1 leading-snug">
            {{ notif.title }}
          </h4>
          <p class="text-xs theme-text-muted leading-relaxed whitespace-pre-line">
            {{ notif.message }}
          </p>

          <!-- Footer Sender -->
          <div class="mt-2.5 pt-2 border-t theme-border-subtle flex items-center justify-between text-[10px] theme-text-dim">
            <span class="flex items-center gap-1 font-semibold">
              <Icon name="ph:user-circle-bold" class="text-xs text-orange-500" />
              {{ notif.sender || 'Admin JMT Sport' }}
            </span>
            <button 
              @click.stop="removeLocalNotification(notif.id)" 
              class="theme-text-dim hover:text-red-500 transition p-1" 
              title="Hapus">
              <Icon name="ph:trash-bold" class="text-xs" />
            </button>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useNotification } from '~/composables/useNotification'

const { 
  notifications, 
  isDrawerOpen, 
  hasPermission, 
  unreadCount, 
  checkPermission, 
  requestPermission, 
  markAsRead, 
  markAllAsRead, 
  removeLocalNotification 
} = useNotification()

onMounted(() => {
  checkPermission()
})

const aktifkanIzinNotif = async () => {
  await requestPermission()
}

const getTypeBadgeClass = (type) => {
  switch (type) {
    case 'JADWAL':
      return 'bg-emerald-500/15 text-emerald-500 border border-emerald-500/30'
    case 'URGENT':
      return 'bg-rose-500/15 text-rose-500 border border-rose-500/30'
    case 'PROMO':
      return 'bg-amber-500/15 text-amber-500 border border-amber-500/30'
    case 'INFO':
    default:
      return 'bg-blue-500/15 text-blue-500 border border-blue-500/30'
  }
}

const getTypeIcon = (type) => {
  switch (type) {
    case 'JADWAL':
      return 'ph:soccer-ball-bold'
    case 'URGENT':
      return 'ph:warning-circle-bold'
    case 'PROMO':
      return 'ph:gift-bold'
    case 'INFO':
    default:
      return 'ph:info-bold'
  }
}

const formatWaktu = (isoDate) => {
  if (!isoDate) return ''
  try {
    const d = new Date(isoDate)
    const now = new Date()
    const diffSec = Math.floor((now.getTime() - d.getTime()) / 1000)

    if (diffSec < 60) return 'Baru saja'
    if (diffSec < 3600) return `${Math.floor(diffSec / 60)} mnt lalu`
    if (diffSec < 86400) return `${Math.floor(diffSec / 3600)} jam lalu`
    if (diffSec < 172800) return 'Kemarin'
    
    return d.toLocaleDateString('id-ID', {
      day: 'numeric',
      month: 'short',
      hour: '2-digit',
      minute: '2-digit'
    })
  } catch (e) {
    return isoDate
  }
}
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.2s ease-out forwards;
}
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>
