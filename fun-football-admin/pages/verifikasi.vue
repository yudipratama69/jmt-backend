<template>
  <div class="space-y-6 max-w-7xl mx-auto pb-12 transition-colors duration-300">
    
    <!-- Header Page -->
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-black theme-text-main flex items-center gap-2.5">
          <span class="w-9 h-9 rounded-xl bg-gradient-to-tr from-red-600 to-orange-500 text-white flex items-center justify-center shadow-md shadow-orange-600/30 text-lg">
            <Icon name="ph:shield-check-bold" />
          </span>
          Pusat Verifikasi
        </h1>
        <p class="text-xs theme-text-muted mt-1">
          Kelola dan setujui bukti transaksi pendaftaran pertandingan dan pengisian deposit saldo pemain.
        </p>
      </div>

      <!-- Quick Actions / Live Refresh Badge -->
      <div class="flex items-center gap-2">
        <button 
          @click="refreshSemua" 
          class="px-3 py-2 theme-bg-surface border theme-border rounded-xl text-xs font-bold theme-text-main hover:bg-orange-500/10 transition flex items-center gap-1.5 active:scale-95 shadow-sm">
          <Icon name="ph:arrows-clockwise-bold" class="text-sm text-orange-500" :class="{ 'animate-spin': isRefreshing }" />
          <span>Segarkan Data</span>
        </button>
      </div>
    </div>

    <!-- Quick Stats Cards (3 Kolom Ringkasan Cepat) -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      
      <!-- Card 1: Tiket Menunggu Verifikasi -->
      <div 
        @click="activeTab = 'tiket'"
        :class="activeTab === 'tiket' ? 'ring-2 ring-orange-500 shadow-orange-500/20' : ''"
        class="theme-bg-surface p-5 rounded-2xl border theme-border shadow-sm flex items-center justify-between cursor-pointer hover:border-orange-500/50 transition duration-200">
        <div class="space-y-1">
          <p class="text-[11px] font-bold uppercase tracking-wider text-amber-500 flex items-center gap-1.5">
            <span class="w-2 h-2 rounded-full bg-amber-500 animate-pulse"></span>
            Tiket Menunggu Verifikasi
          </p>
          <p class="text-2xl font-black theme-text-main">{{ pendingTicketsCount }} <span class="text-xs font-normal theme-text-muted">Resi</span></p>
          <p class="text-[10px] theme-text-muted">Pembayaran via transfer bank</p>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-amber-500/10 text-amber-500 border border-amber-500/20 flex items-center justify-center text-2xl shrink-0">
          <Icon name="ph:ticket-bold" />
        </div>
      </div>

      <!-- Card 2: Deposit Menunggu Verifikasi -->
      <div 
        @click="activeTab = 'deposit'"
        :class="activeTab === 'deposit' ? 'ring-2 ring-orange-500 shadow-orange-500/20' : ''"
        class="theme-bg-surface p-5 rounded-2xl border theme-border shadow-sm flex items-center justify-between cursor-pointer hover:border-orange-500/50 transition duration-200">
        <div class="space-y-1">
          <p class="text-[11px] font-bold uppercase tracking-wider text-blue-500 flex items-center gap-1.5">
            <span class="w-2 h-2 rounded-full bg-blue-500 animate-pulse"></span>
            Deposit Menunggu Verifikasi
          </p>
          <p class="text-2xl font-black theme-text-main">{{ pendingTopupsCount }} <span class="text-xs font-normal theme-text-muted">Permintaan</span></p>
          <p class="text-[10px] theme-text-muted">Top up saldo member</p>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-blue-500/10 text-blue-500 border border-blue-500/20 flex items-center justify-center text-2xl shrink-0">
          <Icon name="ph:wallet-bold" />
        </div>
      </div>

      <!-- Card 3: Total Transaksi Disetujui (Lunas) -->
      <div class="theme-bg-surface p-5 rounded-2xl border theme-border shadow-sm flex items-center justify-between">
        <div class="space-y-1">
          <p class="text-[11px] font-bold uppercase tracking-wider text-emerald-500 flex items-center gap-1.5">
            <Icon name="ph:check-circle-bold" class="text-sm" />
            Total Tiket Terdaftar
          </p>
          <p class="text-2xl font-black theme-text-main">{{ totalTicketsCount }} <span class="text-xs font-normal theme-text-muted">Pemain</span></p>
          <p class="text-[10px] theme-text-muted">{{ approvedTicketsCount }} pemain status Lunas</p>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-emerald-500/10 text-emerald-500 border border-emerald-500/20 flex items-center justify-center text-2xl shrink-0">
          <Icon name="ph:receipt-bold" />
        </div>
      </div>

    </div>

    <!-- Tab Selector Navigasi (Pill Switcher) -->
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-3 border-b theme-border-subtle pb-4">
      
      <!-- Nav Tabs -->
      <div class="flex items-center gap-2 p-1 theme-bg-card rounded-2xl border theme-border w-max">
        
        <!-- Tab Tiket -->
        <button 
          @click="changeTab('tiket')" 
          :class="activeTab === 'tiket' 
            ? 'bg-gradient-to-r from-red-600 to-orange-500 text-white shadow-md font-bold' 
            : 'theme-text-muted hover:theme-text-main font-semibold'"
          class="px-4 py-2 rounded-xl text-xs flex items-center gap-2 transition duration-200">
          <Icon name="ph:ticket-bold" class="text-sm" />
          <span>Verifikasi Tiket Main</span>
          <span 
            v-if="pendingTicketsCount > 0" 
            class="px-1.5 py-0.2 rounded-full text-[10px] font-black"
            :class="activeTab === 'tiket' ? 'bg-white text-orange-600' : 'bg-amber-500 text-white'">
            {{ pendingTicketsCount }}
          </span>
        </button>

        <!-- Tab Deposit -->
        <button 
          @click="changeTab('deposit')" 
          :class="activeTab === 'deposit' 
            ? 'bg-gradient-to-r from-red-600 to-orange-500 text-white shadow-md font-bold' 
            : 'theme-text-muted hover:theme-text-main font-semibold'"
          class="px-4 py-2 rounded-xl text-xs flex items-center gap-2 transition duration-200">
          <Icon name="ph:wallet-bold" class="text-sm" />
          <span>Verifikasi Deposit Saldo</span>
          <span 
            v-if="pendingTopupsCount > 0" 
            class="px-1.5 py-0.2 rounded-full text-[10px] font-black"
            :class="activeTab === 'deposit' ? 'bg-white text-orange-600' : 'bg-blue-500 text-white'">
            {{ pendingTopupsCount }}
          </span>
        </button>

      </div>

      <!-- Info Tab Aktif -->
      <span class="text-xs theme-text-muted">
        Menampilkan data <strong>{{ activeTab === 'tiket' ? 'Pembayaran Tiket Pertandingan' : 'Pengisian Saldo Deposit' }}</strong>
      </span>

    </div>

    <!-- ================================================================= -->
    <!-- BAGIAN 1: TAB VERIFIKASI TIKET PERTANDINGAN                       -->
    <!-- ================================================================= -->
    <div v-if="activeTab === 'tiket'" class="space-y-4 animate-fade-in">
      
      <!-- Filter Bar Tiket -->
      <div class="theme-bg-surface p-4 rounded-2xl border theme-border flex flex-col md:flex-row md:items-center justify-between gap-3">
        
        <!-- Search Input -->
        <div class="relative flex-1 max-w-sm">
          <Icon name="ph:magnifying-glass-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 theme-text-muted text-sm" />
          <input 
            v-model="searchTicket" 
            type="text" 
            placeholder="Cari nama pemain..." 
            class="w-full pl-9 pr-4 py-2 theme-bg-card border theme-border rounded-xl text-xs theme-text-main focus:border-orange-500 outline-none transition" 
          />
        </div>

        <!-- Filter Dropdown Status & Event -->
        <div class="flex flex-wrap items-center gap-2">
          
          <!-- Filter Status Bayar -->
          <select 
            v-model="filterTicketStatus" 
            class="px-3 py-2 theme-bg-card border theme-border rounded-xl text-xs font-semibold theme-text-main focus:border-orange-500 outline-none">
            <option value="ALL">Semua Status Bayar</option>
            <option value="VERIFYING">⏳ Menunggu Verifikasi</option>
            <option value="PAID">✓ Lunas</option>
            <option value="UNPAID">✕ Belum Bayar</option>
            <option value="REJECTED">✕ Ditolak</option>
          </select>

          <!-- Filter Sesi Jadwal -->
          <select 
            v-model="filterTicketEvent" 
            class="px-3 py-2 theme-bg-card border theme-border rounded-xl text-xs font-semibold theme-text-main focus:border-orange-500 outline-none max-w-[200px] truncate">
            <option value="ALL">Semua Sesi Jadwal</option>
            <option v-for="ev in eventsList" :key="ev.id" :value="ev.id">
              {{ ev.title }} ({{ formatShortDate(ev.match_date) }})
            </option>
          </select>

          <span class="text-xs theme-text-muted px-1">
            {{ filteredTickets.length }} Data
          </span>

        </div>

      </div>

      <!-- Tabel Tiket Pertandingan -->
      <div class="theme-bg-surface rounded-2xl shadow-sm border theme-border overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="theme-bg-card theme-text-muted text-[11px] uppercase tracking-wider font-bold">
                <th class="p-4 border-b theme-border">Nama Pemain</th>
                <th class="p-4 border-b theme-border">Sesi Pertandingan</th>
                <th class="p-4 border-b theme-border">Status Polling</th>
                <th class="p-4 border-b theme-border">Metode & Status</th>
                <th class="p-4 border-b theme-border text-center">Bukti Transfer</th>
                <th class="p-4 border-b theme-border text-center">Aksi Pengurus</th>
              </tr>
            </thead>
            <tbody class="text-xs divide-y theme-border-subtle">
              <tr 
                v-for="reg in filteredTickets" 
                :key="reg.id" 
                class="hover:bg-orange-500/5 transition">
                
                <!-- Nama Pemain -->
                <td class="p-4">
                  <div class="flex items-center gap-2.5">
                    <div class="w-8 h-8 rounded-full bg-gradient-to-br from-red-600 to-orange-500 text-white font-bold text-xs flex items-center justify-center shadow-sm shrink-0">
                      {{ reg.user_name ? reg.user_name.substring(0, 2).toUpperCase() : 'JM' }}
                    </div>
                    <div>
                      <p class="font-black theme-text-main text-sm leading-tight">{{ reg.user_name }}</p>
                      <p class="text-[10px] theme-text-muted mt-0.5">ID: #{{ reg.id.slice(-6).toUpperCase() }}</p>
                    </div>
                  </div>
                </td>

                <!-- Sesi Pertandingan -->
                <td class="p-4 theme-text-main font-medium">
                  <p class="font-bold leading-tight">{{ getEventTitle(reg.event_id) }}</p>
                  <p class="text-[10px] theme-text-muted mt-0.5">{{ getEventDate(reg.event_id) }}</p>
                </td>

                <!-- Status Polling -->
                <td class="p-4">
                  <span 
                    :class="reg.polling_status === 'JOIN' ? 'bg-emerald-500/10 text-emerald-500 border-emerald-500/30' : 'bg-amber-500/10 text-amber-500 border-amber-500/30'" 
                    class="px-2.5 py-1 border rounded-full text-[10px] font-black uppercase">
                    {{ reg.polling_status === 'JOIN' ? 'Masuk Kuota' : 'Waiting List' }}
                  </span>
                </td>

                <!-- Metode & Status Bayar -->
                <td class="p-4">
                  <div class="flex flex-col items-start gap-1">
                    <span 
                      v-if="reg.payment_status === 'VERIFYING'" 
                      class="px-2.5 py-0.5 bg-amber-500/15 text-amber-500 border border-amber-500/30 rounded-full text-[10px] font-black flex items-center gap-1 animate-pulse">
                      <Icon name="ph:hourglass-bold" class="text-xs" /> MENUNGGU VERIFIKASI
                    </span>
                    <span 
                      v-else-if="reg.payment_status === 'PAID'" 
                      class="px-2.5 py-0.5 bg-emerald-500/15 text-emerald-500 border border-emerald-500/30 rounded-full text-[10px] font-black flex items-center gap-1">
                      <Icon name="ph:check-bold" class="text-xs" /> LUNAS
                    </span>
                    <span 
                      v-else-if="reg.payment_status === 'REJECTED'" 
                      class="px-2.5 py-0.5 bg-rose-500/15 text-rose-500 border border-rose-500/30 rounded-full text-[10px] font-black">
                      DITOLAK
                    </span>
                    <span 
                      v-else 
                      class="px-2.5 py-0.5 theme-bg-card theme-text-muted border theme-border rounded-full text-[10px] font-bold">
                      BELUM BAYAR
                    </span>
                    <span class="text-[9px] font-mono theme-text-muted uppercase">
                      Via {{ reg.payment_method === 'deposit' ? 'Saldo Deposit' : 'Transfer Bank' }}
                    </span>
                  </div>
                </td>

                <!-- Bukti Transfer -->
                <td class="p-4 text-center">
                  <button 
                    v-if="reg.payment_proof_url" 
                    @click="bukaBukti(reg.payment_proof_url, reg.user_name, 'TIKET', reg)" 
                    class="px-3 py-1.5 bg-blue-500/10 text-blue-500 hover:bg-blue-500/20 rounded-xl text-xs font-bold transition inline-flex items-center gap-1 active:scale-95">
                    <Icon name="ph:image-bold" class="text-sm" /> Lihat Foto
                  </button>
                  <span v-else class="theme-text-muted text-xs">-</span>
                </td>

                <!-- Aksi Pengurus -->
                <td class="p-4 text-center">
                  <div v-if="reg.payment_status === 'VERIFYING'" class="flex items-center justify-center gap-1.5">
                    <button 
                      @click="prosesVerifikasiTiket(reg.id, 'APPROVE')" 
                      class="px-3 py-1.5 bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
                      Setujui
                    </button>
                    <button 
                      @click="prosesVerifikasiTiket(reg.id, 'REJECT')" 
                      class="px-3 py-1.5 bg-rose-600 hover:bg-rose-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
                      Tolak
                    </button>
                  </div>
                  <span v-else class="text-xs theme-text-muted">
                    {{ reg.payment_status === 'PAID' ? 'Selesai ✓' : '-' }}
                  </span>
                </td>

              </tr>

              <!-- State Kosong -->
              <tr v-if="filteredTickets.length === 0">
                <td colspan="6" class="p-10 text-center theme-text-muted space-y-2">
                  <Icon name="ph:ticket-bold" class="text-4xl mx-auto opacity-40" />
                  <p class="font-bold theme-text-main text-sm">Tidak Ada Data Pendaftaran Tiket</p>
                  <p class="text-xs">Tidak ditemukan tiket yang cocok dengan filter pencarian saat ini.</p>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

    </div>

    <!-- ================================================================= -->
    <!-- BAGIAN 2: TAB VERIFIKASI DEPOSIT SALDO PEMAIN                     -->
    <!-- ================================================================= -->
    <div v-else-if="activeTab === 'deposit'" class="space-y-4 animate-fade-in">
      
      <!-- Filter Bar Deposit -->
      <div class="theme-bg-surface p-4 rounded-2xl border theme-border flex flex-col md:flex-row md:items-center justify-between gap-3">
        
        <!-- Search Input -->
        <div class="relative flex-1 max-w-sm">
          <Icon name="ph:magnifying-glass-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 theme-text-muted text-sm" />
          <input 
            v-model="searchDeposit" 
            type="text" 
            placeholder="Cari nama pemain deposit..." 
            class="w-full pl-9 pr-4 py-2 theme-bg-card border theme-border rounded-xl text-xs theme-text-main focus:border-orange-500 outline-none transition" 
          />
        </div>

        <div class="flex items-center gap-2">
          <span class="text-xs font-bold text-orange-500 bg-orange-500/10 border border-orange-500/20 px-3 py-1 rounded-full">
            {{ filteredTopups.length }} Permintaan Menunggu
          </span>
        </div>

      </div>

      <!-- Tabel Permintaan Topup Deposit -->
      <div class="theme-bg-surface rounded-2xl shadow-sm border theme-border overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="theme-bg-card theme-text-muted text-[11px] uppercase tracking-wider font-bold">
                <th class="p-4 border-b theme-border">Waktu Permintaan</th>
                <th class="p-4 border-b theme-border">Nama Pemain</th>
                <th class="p-4 border-b theme-border">Nominal Top Up</th>
                <th class="p-4 border-b theme-border text-center">Bukti Transfer</th>
                <th class="p-4 border-b theme-border text-center">Aksi Pengurus</th>
              </tr>
            </thead>
            <tbody class="text-xs divide-y theme-border-subtle">
              <tr 
                v-for="item in filteredTopups" 
                :key="item._id" 
                class="hover:bg-orange-500/5 transition">
                
                <!-- Tanggal -->
                <td class="p-4 theme-text-muted font-medium">
                  {{ formatDateTime(item.created_at) }}
                </td>

                <!-- Nama Pemain -->
                <td class="p-4">
                  <div class="flex items-center gap-2.5">
                    <div class="w-8 h-8 rounded-full bg-gradient-to-tr from-amber-500 to-orange-500 text-white font-bold text-xs flex items-center justify-center shadow-sm shrink-0">
                      {{ item.user_name ? item.user_name.substring(0, 2).toUpperCase() : 'JM' }}
                    </div>
                    <div>
                      <p class="font-black theme-text-main text-sm leading-tight">{{ item.user_name || 'Tanpa Nama' }}</p>
                      <p class="text-[10px] theme-text-muted mt-0.5">Member ID: #{{ item._id.slice(-6).toUpperCase() }}</p>
                    </div>
                  </div>
                </td>

                <!-- Nominal Topup -->
                <td class="p-4">
                  <span class="font-black text-sm text-emerald-500">
                    Rp {{ (item.amount || 0).toLocaleString('id-ID') }}
                  </span>
                </td>

                <!-- Bukti Transfer -->
                <td class="p-4 text-center">
                  <button 
                    v-if="item.receipt" 
                    @click="bukaBukti(item.receipt, item.user_name, 'DEPOSIT', item)" 
                    class="px-3 py-1.5 bg-blue-500/10 text-blue-500 hover:bg-blue-500/20 rounded-xl text-xs font-bold transition inline-flex items-center gap-1 active:scale-95">
                    <Icon name="ph:image-bold" class="text-sm" /> Lihat Bukti
                  </button>
                  <span v-else class="theme-text-muted text-xs">-</span>
                </td>

                <!-- Tombol Aksi Pengurus -->
                <td class="p-4 text-center">
                  <div class="flex items-center justify-center gap-2">
                    <button 
                      @click="prosesTopup(item._id, 'APPROVE')" 
                      class="px-3.5 py-1.5 bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
                      Setujui
                    </button>
                    <button 
                      @click="prosesTopup(item._id, 'REJECT')" 
                      class="px-3.5 py-1.5 bg-rose-600 hover:bg-rose-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
                      Tolak
                    </button>
                  </div>
                </td>

              </tr>

              <!-- State Kosong Topup -->
              <tr v-if="filteredTopups.length === 0">
                <td colspan="5" class="p-10 text-center theme-text-muted space-y-2">
                  <Icon name="ph:wallet-bold" class="text-4xl mx-auto opacity-40" />
                  <p class="font-bold theme-text-main text-sm">Tidak Ada Permintaan Deposit Baru</p>
                  <p class="text-xs">Semua permintaan top up deposit saldo telah diverifikasi.</p>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

    </div>

    <!-- Modal Pratinjau Foto Bukti Pembayaran Universal -->
    <div v-if="showModalBukti" class="fixed inset-0 bg-black/80 backdrop-blur-sm z-50 flex items-center justify-center p-4 animate-fade-in">
      <div class="theme-bg-surface border theme-border p-5 rounded-3xl max-w-lg w-full relative shadow-2xl space-y-4">
        
        <div class="flex justify-between items-center border-b theme-border pb-3">
          <div>
            <h4 class="font-black theme-text-main text-sm flex items-center gap-1.5">
              <Icon name="ph:receipt-bold" class="text-orange-500" />
              Bukti Transfer Pembayaran
            </h4>
            <p class="text-[11px] theme-text-muted mt-0.5">
              Pemain: <strong>{{ modalUserName }}</strong> ({{ modalType === 'TIKET' ? 'Tiket Pertandingan' : 'Deposit Saldo' }})
            </p>
          </div>
          <div class="flex items-center gap-2">
            <a 
              v-if="selectedBuktiUrl"
              :href="useApiUrl(selectedBuktiUrl)" 
              target="_blank" 
              class="theme-text-muted hover:text-orange-500 text-xs px-2.5 py-1 rounded-xl theme-bg-card border theme-border flex items-center gap-1 font-bold transition"
              title="Buka di Tab Baru">
              <Icon name="ph:arrow-square-out-bold" />
              <span>Buka Tab Baru</span>
            </a>
            <button 
              @click="showModalBukti = false" 
              class="theme-text-muted hover:theme-text-main w-8 h-8 rounded-full theme-bg-card flex items-center justify-center transition">
              ✕
            </button>
          </div>
        </div>

        <!-- Gambar Bukti & Fallback -->
        <div class="max-h-[65vh] overflow-auto rounded-2xl border theme-border p-2 bg-black/10 flex flex-col items-center justify-center min-h-[220px]">
          <img 
            v-if="selectedBuktiUrl && !hasImageError"
            :src="useApiUrl(selectedBuktiUrl)" 
            @error="hasImageError = true"
            class="w-full h-auto max-h-[60vh] object-contain rounded-xl" 
            alt="Bukti Transfer" 
          />

          <!-- Fallback jika file fisik tidak ada di server -->
          <div v-else class="text-center p-6 space-y-3">
            <div class="w-12 h-12 rounded-2xl bg-amber-500/10 text-amber-500 flex items-center justify-center mx-auto text-2xl">
              <Icon name="ph:image-broken-bold" />
            </div>
            <div>
              <p class="font-bold theme-text-main text-xs">Foto Bukti Tidak Dapat Dimuat</p>
              <p class="text-[11px] theme-text-muted mt-0.5 max-w-xs mx-auto">
                File bukti transfer mungkin tersimpan di server host lain atau URL tidak dapat diakses langsung.
              </p>
            </div>
            <a 
              v-if="selectedBuktiUrl"
              :href="useApiUrl(selectedBuktiUrl)" 
              target="_blank" 
              class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-xl bg-orange-500 text-white text-xs font-bold shadow-sm transition hover:bg-orange-600">
              <Icon name="ph:link-bold" />
              <span>Akses Tautan Gambar Asli</span>
            </a>
          </div>
        </div>

        <!-- Tombol Aksi di dalam Modal -->
        <div class="flex justify-between items-center pt-2 border-t theme-border-subtle">
          <button 
            @click="showModalBukti = false" 
            class="px-4 py-2 theme-bg-card theme-text-main border theme-border text-xs font-bold rounded-xl hover:opacity-80 transition">
            Tutup
          </button>

          <!-- Aksi Cepat jika status masih pending -->
          <div v-if="modalItem" class="flex items-center gap-2">
            <button 
              v-if="modalType === 'TIKET' && modalItem.payment_status === 'VERIFYING'"
              @click="prosesVerifikasiTiket(modalItem.id, 'APPROVE'); showModalBukti = false" 
              class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
              Setujui Tiket
            </button>
            <button 
              v-if="modalType === 'DEPOSIT'"
              @click="prosesTopup(modalItem._id, 'APPROVE'); showModalBukti = false" 
              class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-black rounded-xl shadow transition active:scale-95">
              Setujui Deposit
            </button>
          </div>
        </div>

      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const { $api } = useNuxtApp()
const { useAutoRefresh } = useRealtime()
const toast = useToast()

// State Tab Navigasi ('tiket' atau 'deposit')
const activeTab = ref('tiket')

onMounted(() => {
  if (route.query.tab === 'deposit') {
    activeTab.value = 'deposit'
  }
})

const changeTab = (tab) => {
  activeTab.value = tab
  router.replace({ query: { ...route.query, tab } })
}

// Tarik data pendaftaran tiket, permintaan topup deposit, dan events
const { data: registrations, refresh: refreshRegs } = await useApiFetch('/registrations')
const { data: pendingTopups, refresh: refreshTopups } = await useApiFetch('/pending-topups')
const { data: eventsData, refresh: refreshEvents } = await useApiFetch('/events')

const isRefreshing = ref(false)
const refreshSemua = async () => {
  isRefreshing.value = true
  try {
    await Promise.all([refreshRegs(), refreshTopups(), refreshEvents()])
    toast.success('Data verifikasi berhasil diperbarui!', 'Segar')
  } catch (err) {
    console.error(err)
  } finally {
    setTimeout(() => { isRefreshing.value = false }, 500)
  }
}

// Pasang Auto-Refresh Realtime WebSocket
useAutoRefresh(['PAYMENT_UPDATED', 'REGISTRATION_UPDATED', 'TOPUP_UPDATED', 'EVENT_UPDATED'], () => {
  refreshRegs()
  refreshTopups()
  refreshEvents()
})

// Data references
const eventsList = computed(() => eventsData.value?.data || [])
const ticketsList = computed(() => registrations.value?.data || [])
const topupsList = computed(() => pendingTopups.value?.data || [])

// Counter Stat
const pendingTicketsCount = computed(() => {
  return ticketsList.value.filter((r) => r.payment_status === 'VERIFYING').length
})

const totalTicketsCount = computed(() => ticketsList.value.length)

const approvedTicketsCount = computed(() => {
  return ticketsList.value.filter((r) => r.payment_status === 'PAID').length
})

const pendingTopupsCount = computed(() => topupsList.value.length)

// Filter & Search Tiket
const searchTicket = ref('')
const filterTicketStatus = ref('ALL')
const filterTicketEvent = ref('ALL')

const filteredTickets = computed(() => {
  return ticketsList.value.filter((reg) => {
    // Search nama
    const matchName = !searchTicket.value || 
      (reg.user_name && reg.user_name.toLowerCase().includes(searchTicket.value.toLowerCase()))

    // Filter status
    const matchStatus = filterTicketStatus.value === 'ALL' || reg.payment_status === filterTicketStatus.value

    // Filter event
    const matchEvent = filterTicketEvent.value === 'ALL' || String(reg.event_id) === String(filterTicketEvent.value)

    return matchName && matchStatus && matchEvent
  })
})

// Filter & Search Deposit
const searchDeposit = ref('')
const filteredTopups = computed(() => {
  return topupsList.value.filter((item) => {
    if (!searchDeposit.value) return true
    return item.user_name && item.user_name.toLowerCase().includes(searchDeposit.value.toLowerCase())
  })
})

// Modal Bukti Transfer
const showModalBukti = ref(false)
const selectedBuktiUrl = ref('')
const modalUserName = ref('')
const modalType = ref('TIKET')
const modalItem = ref(null)
const hasImageError = ref(false)

const bukaBukti = (url, userName, type, item) => {
  selectedBuktiUrl.value = url
  modalUserName.value = userName || 'Pemain'
  modalType.value = type
  modalItem.value = item
  hasImageError.value = false
  showModalBukti.value = true
}

// Aksi Verifikasi Tiket
const prosesVerifikasiTiket = (id, action) => {
  const isApprove = action === 'APPROVE'
  toast.confirm({
    title: isApprove ? 'Setujui Pembayaran Tiket' : 'Tolak Bukti Transfer',
    message: isApprove 
      ? 'Apakah Anda yakin ingin menyetujui pembayaran pemain ini? Status tiket akan langsung LUNAS dan slot resmi terisi.' 
      : 'Apakah Anda yakin ingin menolak bukti transfer ini?',
    confirmText: isApprove ? 'Ya, Setujui' : 'Tolak',
    cancelText: 'Batal',
    onConfirm: async () => {
      try {
        await $api('/verify-payment', {
          method: 'PUT',
          body: { registration_id: id, action: action }
        })
        toast.success(`Pembayaran tiket berhasil di-${isApprove ? 'setujui (LUNAS)' : 'tolak'}!`, 'Sukses')
        refreshRegs()
      } catch (error) {
        toast.error('Terjadi kesalahan saat memverifikasi pembayaran.', 'Error')
      }
    }
  })
}

// Aksi Verifikasi Deposit
const prosesTopup = (id, action) => {
  const isApprove = action === 'APPROVE'
  toast.confirm({
    title: isApprove ? 'Setujui Isi Deposit' : 'Tolak Deposit',
    message: isApprove 
      ? 'Apakah Anda yakin ingin menyetujui dan menambahkan saldo deposit ke akun pemain ini?' 
      : 'Apakah Anda yakin ingin menolak permintaan deposit ini?',
    confirmText: isApprove ? 'Ya, Tambah Saldo' : 'Tolak',
    cancelText: 'Batal',
    onConfirm: async () => {
      try {
        await $api('/approve-topup', {
          method: 'POST',
          body: { topup_id: id, action: action }
        })
        toast.success(`Permintaan deposit berhasil di-${isApprove ? 'setujui dan saldo ditambahkan' : 'tolak'}!`, 'Sukses')
        refreshTopups()
      } catch (error) {
        toast.error(error.response?._data?.error || 'Terjadi kesalahan saat memproses data', 'Error')
      }
    }
  })
}

// Helpers Info Event & Tanggal
const getEventTitle = (eventId) => {
  if (!eventId) return 'JMT Fun Football'
  const ev = eventsList.value.find((e) => String(e.id) === String(eventId))
  return ev ? ev.title : 'JMT Fun Football'
}

const getEventDate = (eventId) => {
  if (!eventId) return '-'
  const ev = eventsList.value.find((e) => String(e.id) === String(eventId))
  return ev ? formatShortDate(ev.match_date) : '-'
}

const formatShortDate = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { weekday: 'short', day: 'numeric', month: 'short' })
}

const formatDateTime = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' }) + ' ' + d.toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
}
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.2s ease-out forwards;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(4px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>