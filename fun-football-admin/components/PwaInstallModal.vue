<template>
  <div v-if="show" class="fixed inset-0 bg-black/80 backdrop-blur-md z-50 flex items-end sm:items-center justify-center p-0 sm:p-4 animate-fade-in">
    <div class="theme-bg-card-solid border theme-border rounded-t-[32px] sm:rounded-3xl p-6 max-w-md w-full shadow-2xl relative overflow-hidden theme-text-main animate-slide-up">
      
      <!-- Aksen Cahaya Atas -->
      <div class="absolute top-0 left-0 right-0 h-1.5 bg-gradient-to-r from-red-600 via-orange-500 to-amber-400"></div>

      <!-- Tombol Tutup -->
      <button @click="$emit('close')" class="absolute top-5 right-5 theme-text-muted hover:theme-text-main w-9 h-9 rounded-full theme-bg-surface border theme-border flex items-center justify-center transition active:scale-90">
        <Icon name="ph:x-bold" class="text-base" />
      </button>

      <!-- Header Logo & Info -->
      <div class="flex items-center gap-4 mt-1 mb-5">
        <div class="w-16 h-16 rounded-2xl bg-white p-2 shadow-xl border-2 border-orange-500/30 shrink-0 flex items-center justify-center">
          <img :src="'/logo-jmt.png'" alt="Logo JMT Sport" class="w-full h-full object-contain" />
        </div>
        <div>
          <div class="flex items-center gap-1.5">
            <h3 class="text-lg font-black theme-text-main tracking-wide">JMT Sport</h3>
            <span class="bg-gradient-to-r from-orange-500 to-red-600 text-white text-[9px] font-black px-2 py-0.5 rounded-full uppercase tracking-wider">OFFICIAL PWA</span>
          </div>
          <p class="text-xs theme-text-muted mt-0.5">Aplikasi Komunitas Fun Football & Mini Soccer</p>
        </div>
      </div>

      <!-- Tombol Utama Pasang / Install Sekarang -->
      <div class="mb-5">
        <button 
          type="button"
          @click="$emit('install')" 
          class="w-full bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black py-3.5 px-5 rounded-2xl shadow-xl shadow-orange-600/30 transition duration-200 active:scale-95 text-sm flex items-center justify-center gap-2">
          <Icon name="ph:download-simple-bold" class="text-xl" />
          <span>Pasang Aplikasi Sekarang</span>
        </button>
      </div>

      <!-- Tab Panduan (Android vs iPhone vs Desktop) -->
      <div class="flex theme-bg-surface p-1 rounded-xl mb-4 border theme-border gap-1">
        <button 
          @click="activeTab = 'android'" 
          :class="activeTab === 'android' ? 'bg-gradient-to-r from-orange-600 to-red-600 text-white font-bold shadow' : 'theme-text-muted hover:theme-text-main'"
          class="flex-1 py-1.5 rounded-lg text-[11px] transition flex items-center justify-center gap-1">
          <Icon name="ph:android-logo-bold" class="text-sm" /> Android
        </button>
        <button 
          @click="activeTab = 'ios'" 
          :class="activeTab === 'ios' ? 'bg-gradient-to-r from-orange-600 to-red-600 text-white font-bold shadow' : 'theme-text-muted hover:theme-text-main'"
          class="flex-1 py-1.5 rounded-lg text-[11px] transition flex items-center justify-center gap-1">
          <Icon name="ph:apple-logo-bold" class="text-sm" /> iPhone
        </button>
        <button 
          @click="activeTab = 'desktop'" 
          :class="activeTab === 'desktop' ? 'bg-gradient-to-r from-orange-600 to-red-600 text-white font-bold shadow' : 'theme-text-muted hover:theme-text-main'"
          class="flex-1 py-1.5 rounded-lg text-[11px] transition flex items-center justify-center gap-1">
          <Icon name="ph:desktop-bold" class="text-sm" /> PC / Laptop
        </button>
      </div>

      <!-- Konten Panduan Android -->
      <div v-if="activeTab === 'android'" class="theme-bg-surface border theme-border-subtle rounded-2xl p-4 text-xs space-y-3">
        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            1
          </div>
          <p class="theme-text-muted leading-relaxed">
            Tekan ikon <strong class="theme-text-main">Titik Tiga (<Icon name="ph:dots-three-vertical-bold" class="inline text-orange-500" />)</strong> di pojok kanan atas browser Google Chrome.
          </p>
        </div>

        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            2
          </div>
          <p class="theme-text-muted leading-relaxed">
            Pilih menu <strong class="text-orange-500">"Install Aplikasi"</strong> (atau <strong class="theme-text-main">"Tambahkan ke Layar Utama"</strong>).
          </p>
        </div>

        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            3
          </div>
          <p class="theme-text-muted leading-relaxed">
            Tekan <strong class="theme-text-main">"Install"</strong>. Ikon aplikasi JMT Sport akan langsung terpasang di menu HP Anda!
          </p>
        </div>
      </div>

      <!-- Konten Panduan iOS Safari -->
      <div v-if="activeTab === 'ios'" class="theme-bg-surface border theme-border-subtle rounded-2xl p-4 text-xs space-y-3">
        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            1
          </div>
          <p class="theme-text-muted leading-relaxed">
            Buka di browser <strong class="theme-text-main">Safari</strong>, lalu tekan tombol <strong class="text-orange-500">Bagikan (Share / <Icon name="ph:export-bold" class="inline" />)</strong> di bar bawah.
          </p>
        </div>

        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            2
          </div>
          <p class="theme-text-muted leading-relaxed">
            Gulir ke bawah dan pilih <strong class="text-orange-500">"Tambahkan ke Layar Utama" (Add to Home Screen)</strong>.
          </p>
        </div>

        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            3
          </div>
          <p class="theme-text-muted leading-relaxed">
            Tekan <strong class="theme-text-main">"Tambah" (Add)</strong> di pojok kanan atas.
          </p>
        </div>
      </div>

      <!-- Konten Panduan Desktop PC/Laptop (Edge / Chrome) -->
      <div v-if="activeTab === 'desktop'" class="theme-bg-surface border theme-border-subtle rounded-2xl p-4 text-xs space-y-3">
        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            1
          </div>
          <p class="theme-text-muted leading-relaxed">
            Perhatikan bagian <strong class="theme-text-main">sebelah kanan kolom URL (Address Bar)</strong> browser Anda di atas.
          </p>
        </div>

        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            2
          </div>
          <p class="theme-text-muted leading-relaxed">
            Klik ikon <strong class="text-orange-500">App Tersedia / Pasang (<Icon name="ph:desktop-tower-bold" class="inline text-orange-500" />)</strong> di samping ikon bintang/bookmark.
          </p>
        </div>

        <div class="flex items-start gap-3">
          <div class="w-6 h-6 rounded-full bg-orange-500/20 text-orange-500 flex items-center justify-center font-bold text-xs shrink-0 border border-orange-500/30">
            3
          </div>
          <p class="theme-text-muted leading-relaxed">
            Klik <strong class="theme-text-main">"Install"</strong> pada jendela pop-up untuk membuka aplikasi dalam jendela mandiri tanpa tab browser!
          </p>
        </div>
      </div>

      <!-- Fitur & Keuntungan -->
      <div class="mt-4 pt-4 border-t theme-border-subtle flex justify-between items-center text-[11px] theme-text-muted">
        <span class="flex items-center gap-1 text-emerald-500 font-bold">
          <Icon name="ph:check-circle-fill" class="text-sm" /> Realtime Sync
        </span>
        <span class="flex items-center gap-1 text-emerald-500 font-bold">
          <Icon name="ph:check-circle-fill" class="text-sm" /> Bebas Tab Browser
        </span>
        <span class="flex items-center gap-1 text-emerald-500 font-bold">
          <Icon name="ph:check-circle-fill" class="text-sm" /> Standalone PWA
        </span>
      </div>

      <button @click="$emit('close')" class="w-full mt-4 py-2.5 text-xs theme-text-muted hover:theme-text-main font-semibold transition">
        Tutup
      </button>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const props = defineProps({
  show: Boolean,
  isIOS: Boolean,
  hasPrompt: Boolean
})

defineEmits(['close', 'install'])

const activeTab = ref('android')

onMounted(() => {
  if (typeof window !== 'undefined') {
    const userAgent = window.navigator.userAgent.toLowerCase()
    if (/iphone|ipad|ipod/.test(userAgent)) {
      activeTab.value = 'ios'
    } else if (!/android|mobile/.test(userAgent)) {
      activeTab.value = 'desktop'
    } else {
      activeTab.value = 'android'
    }
  }
})
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
