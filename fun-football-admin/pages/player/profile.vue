<template>
  <div class="p-4 space-y-6">
    
    <!-- Tombol Kembali & Judul -->
    <div class="flex items-center gap-3">
      <button @click="navigateTo('/player')" class="p-2.5 theme-bg-surface rounded-full theme-text-muted hover:theme-text-main transition active:scale-95 border theme-border">
        <Icon name="ph:arrow-left-bold" class="text-xl" />
      </button>
      <div>
        <h2 class="text-xl font-black theme-text-main tracking-wide">Pengaturan Profil</h2>
        <p class="text-xs theme-text-muted">Atur avatar dan nama panggilan lapanganmu.</p>
      </div>
    </div>

    <!-- Profil Card Box -->
    <div class="theme-bg-card backdrop-blur-md rounded-3xl p-6 border theme-border shadow-xl text-center transition-colors duration-300">
      
      <!-- Area Foto Profil Bulat Sporty -->
      <div class="relative w-32 h-32 mx-auto mb-4">
        
        <!-- Preview Foto Bulat -->
        <img v-if="previewImage" :src="previewImage" class="w-full h-full object-cover rounded-full border-4 border-orange-500/40 shadow-xl" />
        <div v-else class="w-full h-full theme-bg-surface rounded-full flex items-center justify-center border-4 theme-border shadow-xl">
          <Icon name="ph:user-bold" class="text-6xl theme-text-muted" />
        </div>
        
        <!-- Tombol Kamera Bulat -->
        <label class="absolute bottom-1 right-1 bg-gradient-to-r from-red-600 to-orange-500 text-white w-10 h-10 flex items-center justify-center rounded-full cursor-pointer shadow-lg hover:brightness-110 transition active:scale-95 border-2 border-white">
          <Icon name="ph:camera-plus-bold" class="text-xl" />
          <input type="file" @change="handleFileChange" accept="image/*" style="display: none;" />
        </label>

      </div>

      <p class="text-xs theme-text-muted mb-6">Sentuh ikon kamera untuk mengganti foto profil</p>

      <!-- Form Edit Nama -->
      <form @submit.prevent="simpanProfile" class="space-y-4 text-left">
        <div>
          <label class="block text-xs font-black uppercase tracking-wider theme-text-muted mb-1.5">Nama Panggilan di Lapangan</label>
          <input 
            v-model="formName" 
            type="text" 
            class="w-full border theme-border rounded-2xl p-3.5 theme-bg-surface text-sm theme-text-main placeholder-slate-400 outline-none focus:border-orange-500 font-bold transition" 
            required />
        </div>

        <button 
          type="submit" 
          :disabled="isSaving"
          class="w-full bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black py-4 rounded-2xl shadow-xl shadow-orange-600/30 transition active:scale-95 text-sm flex items-center justify-center gap-2 disabled:opacity-50">
          <Icon v-if="isSaving" name="ph:spinner-gap-bold" class="text-lg animate-spin" />
          <Icon v-else name="ph:floppy-disk-back-bold" class="text-lg" />
          <span>{{ isSaving ? 'Menyimpan...' : 'Simpan Perubahan' }}</span>
        </button>
      </form>

    </div>

    <!-- Card Pasang PWA (Hanya jika belum diinstall) -->
    <div v-if="!isInstalled" class="theme-bg-card backdrop-blur-md rounded-3xl p-5 border theme-border shadow-xl flex items-center justify-between transition-colors duration-300">
      <div class="flex items-center gap-3.5">
        <div class="w-12 h-12 rounded-2xl bg-orange-500/20 border border-orange-500/30 text-orange-500 flex items-center justify-center shrink-0">
          <Icon name="ph:device-mobile-bold" class="text-2xl" />
        </div>
        <div>
          <h4 class="text-sm font-black theme-text-main">Aplikasi HP (PWA)</h4>
          <p class="text-[11px] theme-text-muted">Pasang di layar utama HP untuk akses cepat</p>
        </div>
      </div>
      <button 
        @click="triggerInstall" 
        class="px-4 py-2.5 bg-gradient-to-r from-orange-500 to-red-600 text-white text-xs font-black rounded-xl shadow-md transition active:scale-95 flex items-center gap-1.5 shrink-0">
        <Icon name="ph:download-simple-bold" class="text-base" />
        Pasang
      </button>
    </div>

    <!-- Modal Panduan Install PWA -->
    <PwaInstallModal 
      :show="showPwaModal" 
      :is-i-o-s="isIOS" 
      :has-prompt="hasPrompt"
      @close="closeModal" 
      @install="triggerInstall" 
    />

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

definePageMeta({ layout: 'mobile' })

const { $api } = useNuxtApp()
const { isInstalled, isIOS, hasPrompt, showModal: showPwaModal, triggerInstall, closeModal } = usePwaInstall()

const userId = ref('')
const formName = ref('')
const selectedFile = ref(null)
const previewImage = ref(null)
const isSaving = ref(false)

onMounted(async () => {
  userId.value = localStorage.getItem('user_id')
  if (!userId.value) return navigateTo('/player/login')

  try {
    const res = await $api(`/user?id=${userId.value}`)
    formName.value = res.data.name
    if (res.data.profile_pic) {
      previewImage.value = useApiUrl(res.data.profile_pic)
    }
  } catch (error) {
    console.error("Gagal memuat profil")
  }
})

const handleFileChange = (event) => {
  const file = event.target.files[0]
  if (file) {
    selectedFile.value = file
    previewImage.value = URL.createObjectURL(file)
  }
}

const toast = useToast()

const simpanProfile = async () => {
  isSaving.value = true
  try {
    const formData = new FormData()
    formData.append('user_id', userId.value)
    formData.append('name', formName.value)
    if (selectedFile.value) {
      formData.append('profile_pic', selectedFile.value)
    }

    const res = await $api('/update-profile', {
      method: 'POST',
      body: formData
    })

    if (res.new_name) localStorage.setItem('user_name', res.new_name)
    if (res.new_profile_pic) previewImage.value = useApiUrl(res.new_profile_pic)

    toast.success('Profil pemain Anda berhasil diperbarui!', 'Sukses Simpan')
    setTimeout(() => {
      navigateTo('/player')
    }, 800)
  } catch (error) {
    toast.error('Gagal memperbarui profil.', 'Error')
  } finally {
    isSaving.value = false
  }
}
</script>