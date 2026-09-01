<template>
  <div class="p-5 space-y-6">
    
    <!-- Tombol Kembali & Judul -->
    <div class="flex items-center gap-3 mb-6">
      <button @click="navigateTo('/player')" class="p-2 bg-gray-100 rounded-full text-gray-600 hover:bg-gray-200 transition active:scale-95">
        <Icon name="ph:arrow-left-bold" class="text-xl" />
      </button>
      <h2 class="text-xl font-bold text-gray-800">Pengaturan Profile</h2>
    </div>

    <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100 text-center">
      
      <!-- AREA FOTO PROFILE (YANG DIPERBAIKI) -->
      <div class="relative w-28 h-28 mx-auto mb-4">
        
        <!-- Preview Foto Bulat -->
        <img v-if="previewImage" :src="previewImage" class="w-full h-full object-cover rounded-full border-4 border-emerald-50 shadow-md" />
        <div v-else class="w-full h-full bg-emerald-100 rounded-full flex items-center justify-center border-4 border-emerald-50 shadow-md">
          <Icon name="ph:user-bold" class="text-5xl text-emerald-500" />
        </div>
        
        <!-- Tombol Kamera (Dibuat bulat sempurna dan input disembunyikan) -->
        <label class="absolute bottom-0 right-0 bg-emerald-600 text-white w-10 h-10 flex items-center justify-center rounded-full cursor-pointer shadow-lg hover:bg-emerald-700 transition active:scale-95">
          <Icon name="ph:camera-plus-bold" class="text-xl" />
          <!-- style="display: none;" akan memaksa teks 'Choose File' hilang -->
          <input type="file" @change="handleFileChange" accept="image/*" style="display: none;" />
        </label>

      </div>

      <p class="text-xs text-gray-400 mb-6">Klik ikon kamera untuk mengubah foto</p>

      <!-- Form Edit Nama -->
      <form @submit.prevent="simpanProfile" class="space-y-4 text-left">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Nama Panggilan</label>
          <input v-model="formName" type="text" class="w-full border border-gray-300 rounded-xl p-3 focus:ring-2 focus:ring-emerald-500 outline-none" required />
        </div>

        <button type="submit" class="w-full bg-gray-900 hover:bg-emerald-600 text-white font-bold py-3.5 rounded-xl active:scale-95 transition mt-4 flex items-center justify-center gap-2">
          <Icon name="ph:floppy-disk-back-bold" class="text-lg" />
          Simpan Perubahan
        </button>
      </form>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

definePageMeta({ layout: 'mobile' })

const userId = ref('')
const formName = ref('')
const selectedFile = ref(null)
const previewImage = ref(null)

// Tarik data user saat ini saat halaman dimuat
onMounted(async () => {
  userId.value = localStorage.getItem('user_id')
  if (!userId.value) return navigateTo('/player/login')

  try {
    const res = await $fetch(`http://localhost:8080/user?id=${userId.value}`)
    formName.value = res.data.name
    if (res.data.profile_pic) {
      previewImage.value = `http://localhost:8080${res.data.profile_pic}`
    }
  } catch (error) {
    console.error("Gagal memuat profil")
  }
})

// Fungsi untuk menangani pilihan gambar dan membuat preview
const handleFileChange = (event) => {
  const file = event.target.files[0]
  if (file) {
    selectedFile.value = file
    previewImage.value = URL.createObjectURL(file) // Buat URL lokal untuk preview cepat
  }
}

// Fungsi Simpan Profile ke Backend
const simpanProfile = async () => {
  try {
    const formData = new FormData()
    formData.append('user_id', userId.value)
    formData.append('name', formName.value)
    if (selectedFile.value) {
      formData.append('profile_pic', selectedFile.value)
    }

    const res = await $fetch('http://localhost:8080/update-profile', {
      method: 'PUT',
      body: formData
    })

    // Update nama di Local Storage agar sapaan di beranda ikut berubah
    localStorage.setItem('user_name', res.new_name)
    alert('Mantap! Profil berhasil diperbarui.')
    navigateTo('/player') // Kembali ke beranda
    
  } catch (error) {
    alert('Gagal memperbarui profil.')
  }
}
</script>