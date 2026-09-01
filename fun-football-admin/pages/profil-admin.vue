<template>
  <div class="p-8 max-w-4xl">
    
    <!-- Header Halaman -->
    <div class="mb-8">
      <h1 class="text-3xl font-extrabold text-gray-800">Profil Pengurus</h1>
      <p class="text-sm text-gray-500 mt-1">Kelola informasi akun dan keamanan panel admin Anda.</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      
      <!-- Kartu Info Profil (Kiri) -->
      <div class="md:col-span-1 space-y-6">
        <div class="bg-white p-6 rounded-2xl shadow-sm border border-gray-100 flex flex-col items-center text-center">
          <div class="w-24 h-24 bg-gradient-to-br from-red-500 to-orange-500 rounded-full flex items-center justify-center text-white text-3xl font-black shadow-lg mb-4">
            {{ form.name.substring(0, 1).toUpperCase() }}
          </div>
          <h2 class="text-lg font-bold text-gray-800">{{ form.name }}</h2>
          <span class="px-3 py-1 bg-red-50 text-red-600 text-xs font-bold rounded-full mt-2 border border-red-100">Super Admin</span>
        </div>
      </div>

      <!-- Form Edit Profil & Password (Kanan) -->
      <div class="md:col-span-2 space-y-6">
        
        <!-- Form Pengaturan Akun -->
        <div class="bg-white p-6 rounded-2xl shadow-sm border border-gray-100">
          <h3 class="text-lg font-bold text-gray-800 mb-4 flex items-center gap-2 border-b border-gray-50 pb-3">
            <Icon name="ph:user-circle-bold" class="text-orange-500 text-xl" />
            Pengaturan Akun
          </h3>
          
          <form @submit.prevent="simpanProfil" class="space-y-4">
            <div>
              <label class="block text-xs font-bold text-gray-500 uppercase tracking-wider mb-1.5">Nama Lengkap</label>
              <input v-model="form.name" type="text" class="w-full border border-gray-200 rounded-xl p-3 bg-gray-50/50 text-sm font-semibold text-gray-700 outline-none focus:ring-2 focus:ring-orange-500 transition" />
            </div>
            
            <div>
              <label class="block text-xs font-bold text-gray-500 uppercase tracking-wider mb-1.5">Email / Username</label>
              <input v-model="form.email" type="email" class="w-full border border-gray-200 rounded-xl p-3 bg-gray-50/50 text-sm font-semibold text-gray-700 outline-none focus:ring-2 focus:ring-orange-500 transition" />
            </div>

            <button type="submit" class="bg-gray-800 hover:bg-gray-900 text-white font-bold py-2.5 px-6 rounded-xl shadow-md transition duration-200 text-sm active:scale-95">
              Simpan Perubahan
            </button>
          </form>
        </div>

        <!-- Form Ganti Password -->
        <div class="bg-white p-6 rounded-2xl shadow-sm border border-gray-100">
          <h3 class="text-lg font-bold text-gray-800 mb-4 flex items-center gap-2 border-b border-gray-50 pb-3">
            <Icon name="ph:lock-key-bold" class="text-orange-500 text-xl" />
            Ganti Password
          </h3>
          
          <form @submit.prevent="gantiPassword" class="space-y-4">
            <div>
              <label class="block text-xs font-bold text-gray-500 uppercase tracking-wider mb-1.5">Password Baru</label>
              <input v-model="passForm.newPass" type="password" placeholder="Masukkan password baru" class="w-full border border-gray-200 rounded-xl p-3 bg-gray-50/50 text-sm outline-none focus:ring-2 focus:ring-orange-500 transition" />
            </div>
            
            <div>
              <label class="block text-xs font-bold text-gray-500 uppercase tracking-wider mb-1.5">Konfirmasi Password Baru</label>
              <input v-model="passForm.confirmPass" type="password" placeholder="Ulangi password baru" class="w-full border border-gray-200 rounded-xl p-3 bg-gray-50/50 text-sm outline-none focus:ring-2 focus:ring-orange-500 transition" />
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

// 1. Ambil data dari MongoDB saat halaman dibuka
const fetchProfile = async () => {
  try {
    const res = await $fetch('http://localhost:8080/admin-profile')
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
  if (!adminId.value) return alert('ID Admin tidak ditemukan!')

  try {
    await $fetch('http://localhost:8080/admin-profile', {
      method: 'PUT',
      body: {
        id: adminId.value,
        name: form.value.name,
        email: form.value.email
      }
    })
    alert(`Sukses! Profil berhasil disimpan atas nama: ${form.value.name}`)
  } catch (error) {
    alert('Terjadi kesalahan saat menyimpan profil.')
  }
}

// 3. Fungsi Ganti Password ke MongoDB
const gantiPassword = async () => {
  if (!adminId.value) return alert('ID Admin tidak ditemukan!')
  if (passForm.value.newPass !== passForm.value.confirmPass) {
    alert('Password baru dan konfirmasi tidak cocok!')
    return
  }
  if (passForm.value.newPass.length < 6) {
    alert('Password minimal 6 karakter!')
    return
  }

  try {
    await $fetch('http://localhost:8080/admin-password', {
      method: 'PUT',
      body: {
        id: adminId.value,
        password: passForm.value.newPass
      }
    })
    alert('Sukses! Password admin berhasil diperbarui!')
    passForm.value.newPass = ''
    passForm.value.confirmPass = ''
  } catch (error) {
    alert('Terjadi kesalahan saat memperbarui password.')
  }
}
</script>