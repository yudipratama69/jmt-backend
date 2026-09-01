<template>
  <div class="p-5 space-y-6 pb-24">
    
    <!-- Header -->
    <div>
      <h2 class="text-xl font-bold text-gray-800">Daftar Pemain</h2>
      <p class="text-xs text-gray-500 mt-1">List pemain yang sudah polling (JOIN) dan Lunas.</p>
    </div>

    <!-- Daftar Pemain -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <div v-if="pending" class="p-6 text-center text-gray-400 text-sm">
        Memuat data pemain...
      </div>
      
      <div v-else-if="pemainLunas.length === 0" class="p-6 text-center text-gray-400 text-sm">
        Belum ada pemain yang polling dan lunas.
      </div>

      <ul v-else class="divide-y divide-gray-50">
        <li v-for="(pemain, index) in pemainLunas" :key="pemain.id" class="p-4 flex items-center justify-between hover:bg-gray-50/50 transition">
          <div class="flex items-center gap-3">
            <!-- Nomor Urut -->
            <span class="text-xs font-bold text-gray-400 w-5">{{ index + 1 }}</span>
            
            <!-- Avatar Inisial -->
            <div class="w-10 h-10 rounded-full bg-gradient-to-br from-orange-100 to-orange-200 flex items-center justify-center text-orange-700 font-bold text-sm">
              {{ pemain.user_name ? pemain.user_name.substring(0, 2).toUpperCase() : 'NN' }}
            </div>
            
            <!-- Nama Pemain -->
            <div>
              <p class="text-sm font-bold text-gray-800">{{ pemain.user_name }}</p>
              <p class="text-[10px] text-gray-400 font-medium mt-0.5">Sudah Polling</p>
            </div>
          </div>

          <!-- Status Lunas -->
          <div class="px-3 py-1 bg-green-100 text-green-700 rounded-full text-[10px] font-extrabold tracking-wider flex items-center gap-1">
            <Icon name="ph:check-circle-fill" class="text-xs" />
            LUNAS
          </div>
        </li>
      </ul>
    </div>
    
  </div>
</template>

<script setup>
import { computed } from 'vue'

// Sesuaikan layout yang dipakai untuk player
definePageMeta({ layout: 'mobile' })

// Tarik data semua pendaftaran
const { data: regData, pending } = await useFetch('http://localhost:8080/registrations')

// Filter hanya pemain yang Polling = JOIN dan Bayar = PAID
const pemainLunas = computed(() => {
  if (!regData.value?.data) return []
  
  return regData.value.data.filter(
    (reg) => reg.polling_status === 'JOIN' && reg.payment_status === 'PAID'
  )
})
</script>