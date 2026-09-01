<template>
  <div class="space-y-8 w-full px-8 pb-10">
    
    <div>
      <h1 class="text-2xl font-extrabold text-gray-800">Ringkasan Dashboard</h1>
      <p class="text-xs text-gray-400 mt-0.5">Statistik utama aktivitas dan partisipasi JMT Sport.</p>
    </div>
    
    <!-- Kartu Statistik dengan Warna Gradasi Penuh & Simbol Menarik -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      
      <!-- Kartu 1: Jadwal Terdaftar (Gradasi Merah - Oranye) -->
      <div style="background: linear-gradient(135deg, #ea580c, #dc2626);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40">
        <div class="absolute right-4 -bottom-4 text-white/15 pointer-events-none">
          <Icon name="ph:calendar-blank-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            📅 Jadwal
          </div>
          <h3 class="text-orange-100 text-xs font-bold uppercase tracking-wider">Jadwal Terdaftar</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">{{ stats?.total_events || 0 }} Agenda</p>
      </div>

      <!-- Kartu 2: Total Pemain (Gradasi Biru - Indigo) -->
      <div style="background: linear-gradient(135deg, #2563eb, #1d4ed8);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40">
        <div class="absolute right-4 -bottom-4 text-white/15 pointer-events-none">
          <Icon name="ph:users-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            👥 Partisipasi
          </div>
          <h3 class="text-blue-100 text-xs font-bold uppercase tracking-wider">Total Pemain (Ikut)</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">{{ stats?.total_players || 0 }} Orang</p>
      </div>

      <!-- Kartu 3: Menunggu Verifikasi (Gradasi Amber - Kuning Emas) -->
      <div style="background: linear-gradient(135deg, #d97706, #b45309);" class="p-6 rounded-2xl shadow-xl text-white relative overflow-hidden flex flex-col justify-between h-40">
        <div class="absolute right-4 -bottom-4 text-white/15 pointer-events-none">
          <Icon name="ph:clock-countdown-bold" class="text-9xl transform -rotate-12" />
        </div>
        <div class="relative z-10">
          <div class="inline-block bg-black/20 text-white text-[10px] font-extrabold px-3 py-1 rounded-full uppercase tracking-wider mb-2">
            ⏳ Verifikasi
          </div>
          <h3 class="text-amber-100 text-xs font-bold uppercase tracking-wider">Menunggu Verifikasi</h3>
        </div>
        <p class="text-3xl font-black tracking-tight relative z-10 text-white">{{ stats?.pending_verification || 0 }} Resi</p>
      </div>

    </div>

    <!-- Tabel Daftar Jadwal Main -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <div class="p-5 border-b border-gray-100 flex justify-between items-center bg-gray-50/50">
        <h2 class="font-bold text-gray-700 text-base">Daftar Jadwal Main</h2>
      </div>
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-50 text-gray-400 text-xs uppercase tracking-wider">
            <th class="p-4 font-bold border-b border-gray-100">Nama Jadwal</th>
            <th class="p-4 font-bold border-b border-gray-100">Tanggal</th>
            <th class="p-4 font-bold border-b border-gray-100">Lokasi</th>
            <th class="p-4 font-bold border-b border-gray-100">Harga</th>
            <th class="p-4 font-bold border-b border-gray-100">Status</th>
          </tr>
        </thead>
        <tbody class="text-sm">
          <tr v-for="event in eventsData?.data" :key="event.id" class="border-b border-gray-50 hover:bg-gray-50/80 transition">
            <td class="p-4 text-gray-800 font-bold">{{ event.title }}</td>
            <td class="p-4 text-gray-600">{{ new Date(event.match_date).toLocaleDateString('id-ID') }}</td>
            <td class="p-4 text-gray-600">{{ event.location }}</td>
            <td class="p-4 text-gray-600 font-semibold">Rp {{ event.price_per_person.toLocaleString('id-ID') }}</td>
            <td class="p-4">
              <span class="px-3 py-1 bg-orange-100 text-orange-700 rounded-full text-xs font-extrabold">
                {{ event.status }}
              </span>
            </td>
          </tr>
          <tr v-if="!eventsData?.data || eventsData.data.length === 0">
            <td colspan="5" class="p-8 text-center text-gray-400 text-sm">Belum ada jadwal yang dibuat.</td>
          </tr>
        </tbody>
      </table>
    </div>

  </div>
</template>

<script setup>
const { data: eventsData } = await useFetch('http://localhost:8080/events')
const { data: stats } = await useFetch('http://localhost:8080/dashboard-stats')
</script>