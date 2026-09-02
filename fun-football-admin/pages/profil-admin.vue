<template>
  <div class="p-4 md:p-8 max-w-4xl transition-colors duration-300">
    
    <!-- Header Halaman -->
    <div class="mb-8">
      <h1 class="text-3xl font-extrabold theme-text-main">Profil Pengurus</h1>
      <p class="text-sm theme-text-muted mt-1">Kelola informasi akun dan keamanan panel admin Anda.</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      
      <!-- Kartu Info Profil (Kiri) -->
      <div class="md:col-span-1 space-y-6">
        <div class="theme-bg-surface p-6 rounded-2xl shadow-sm border theme-border flex flex-col items-center text-center transition-colors duration-300">
          <div class="w-24 h-24 bg-gradient-to-br from-red-500 to-orange-500 rounded-full flex items-center justify-center text-white text-3xl font-black shadow-lg mb-4">
            {{ form.name ? form.name.substring(0, 1).toUpperCase() : 'A' }}
          </div>
          <h2 class="text-lg font-bold theme-text-main">{{ form.name }}</h2>
          <span class="px-3 py-1 bg-red-500/10 text-red-500 border border-red-500/20 text-xs font-bold rounded-full mt-2">Super Admin</span>
        </div>
      </div>

      <!-- Form Edit Profil & Password (Kanan) -->
      <div class="md:col-span-2 space-y-6">
        
        <!-- Form Pengaturan Akun -->
        <div class="theme-bg-surface p-6 rounded-2xl shadow-sm border theme-border transition-colors duration-300">
          <h3 class="text-lg font-bold theme-text-main mb-4 flex items-center gap-2 border-b theme-border-subtle pb-3">
            <Icon name="ph:user-circle-bold" class="text-orange-500 text-xl" />
            Pengaturan Akun
          </h3>
          
          <form @submit.prevent="simpanProfil" class="space-y-4">
            <div>
              <label class="block text-xs font-bold theme-text-muted uppercase tracking-wider mb-1.5">Nama Lengkap</label>
              <input v-model="form.name" type="text" class="w-full border theme-border rounded-xl p-3 theme-bg-card text-sm font-semibold theme-text-main outline-none focus:border-orange-500 transition" />
            </div>
            
            <div>
              <label class="block text-xs font-bold theme-text-muted uppercase tracking-wider mb-1.5">Email / Username</label>
              <input v-model="form.email" type="email" class="w-full border theme-border rounded-xl p-3 theme-bg-card text-sm font-semibold theme-text-main outline-none focus:border-orange-500 transition" />
            </div>

            <button type="submit" class="bg-gray-800 hover:bg-gray-900 dark:bg-orange-600 dark:hover:bg-orange-700 text-white font-bold py-2.5 px-6 rounded-xl shadow-md transition duration-200 text-sm active:scale-95">
              Simpan Perubahan
            </button>
          </form>
        </div>

        <!-- Form Ganti Password -->
        <div class="theme-bg-surface p-6 rounded-2xl shadow-sm border theme-border transition-colors duration-300">
          <h3 class="text-lg font-bold theme-text-main mb-4 flex items-center gap-2 border-b theme-border-subtle pb-3">
            <Icon name="ph:lock-key-bold" class="text-orange-500 text-xl" />
            Ganti Password
          </h3>
          
          <form @submit.prevent="gantiPassword" class="space-y-4">
            <div>
              <label class="block text-xs font-bold theme-text-muted uppercase tracking-wider mb-1.5">Password Baru</label>
              <input v-model="passForm.newPass" type="password" placeholder="Masukkan password baru" class="w-full border theme-border rounded-xl p-3 theme-bg-card text-sm theme-text-main outline-none focus:border-orange-500 transition placeholder-slate-400" />
            </div>
            
            <div>
              <label class="block text-xs font-bold theme-text-muted uppercase tracking-wider mb-1.5">Konfirmasi Password Baru</label>
              <input v-model="passForm.confirmPass" type="password" placeholder="Ulangi password baru" class="w-full border theme-border rounded-xl p-3 theme-bg-card text-sm theme-text-main outline-none focus:border-orange-500 transition placeholder-slate-400" />
            </div>

            <button type="submit" class="bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white font-bold py-2.5 px-6 rounded-xl shadow-md transition duration-200 text-sm active:scale-95">
              Update Password
            </button>
          </form>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const adminId = ref('')
const form = ref({
  name: 'Memuat...',
  email: 'Memuat...'
})

const passForm = ref({
  newPass: '',
  confirmPass: ''
})

const { $api } = useNuxtApp()
const toast = useToast()

// 1. Ambil data dari MongoDB saat halaman dibuka
const fetchProfile = async () => {
  try {
    const res = await $api('/admin-profile')
    adminId.value = res.id
    form.value.name = res.name
    form.value.email = res.email
  } catch (error) {
    console.error("Admin tidak ditemukan. Pastikan sudah menambahkan role: 'admin' di Compass.")
    form.value.name = 'Admin Tidak Ditemukan'
    form.value.email = '-'
  }
}

onMounted(() => {
  fetchProfile()
})

// 2. Fungsi Simpan Profil ke MongoDB
const simpanProfil = async () => {
  if (!adminId.value) return toast.error('ID Admin tidak ditemukan!', 'Error')

  try {
    await $api('/admin-profile', {
      method: 'PUT',
      body: {
        id: adminId.value,
        name: form.value.name,
        email: form.value.email
      }
    })
    toast.success(`Profil admin berhasil disimpan atas nama: ${form.value.name}`, 'Sukses Simpan')
  } catch (error) {
    toast.error('Terjadi kesalahan saat menyimpan profil.', 'Error')
  }
}

// 3. Fungsi Ganti Password ke MongoDB
const gantiPassword = async () => {
  if (!adminId.value) return toast.error('ID Admin tidak ditemukan!', 'Error')
  if (passForm.value.newPass !== passForm.value.confirmPass) {
    toast.warning('Password baru dan konfirmasi tidak cocok!', 'Perhatian')
    return
  }
  if (passForm.value.newPass.length < 6) {
    toast.warning('Password minimal harus 6 karakter!', 'Perhatian')
    return
  }

  try {
    await $api('/admin-password', {
      method: 'PUT',
      body: {
        id: adminId.value,
        password: passForm.value.newPass
      }
    })
    toast.success('Password admin berhasil diperbarui!', 'Sukses')
    passForm.value.newPass = ''
    passForm.value.confirmPass = ''
  } catch (error) {
    toast.error('Terjadi kesalahan saat memperbarui password.', 'Error')
  }
}
</script>