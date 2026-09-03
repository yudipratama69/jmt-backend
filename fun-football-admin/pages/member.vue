<template>
  <div class="space-y-8 w-full max-w-7xl mx-auto px-2 sm:px-4 pb-16 transition-colors duration-300">
    
    <!-- ================================================================= -->
    <!-- HEADER DAFTAR MEMBER & SQUAD                                      -->
    <!-- ================================================================= -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 border-b theme-border-subtle pb-6">
      <div>
        <div class="flex items-center gap-2 mb-1">
          <span class="w-8 h-8 rounded-xl bg-gradient-to-tr from-blue-600 to-indigo-500 text-white flex items-center justify-center shadow-md shadow-blue-600/30 text-base">
            <Icon name="ph:users-three-bold" />
          </span>
          <h1 class="text-2xl font-black theme-text-main">Daftar Member & Line-up Squad</h1>
          <span class="text-[10px] uppercase tracking-wider font-extrabold px-2.5 py-0.5 rounded-full bg-blue-500/10 text-blue-500 border border-blue-500/30">
            JMT Community
          </span>
        </div>
        <p class="text-xs theme-text-muted">
          Kelola database member komunitas, pantau saldo deposit pemain, dan tinjau line-up squad resmi per sesi pertandingan.
        </p>
      </div>

      <!-- Quick Counter & Tab Switcher -->
      <div class="flex items-center gap-2">
        <span class="px-3.5 py-2 theme-bg-surface border theme-border rounded-xl text-xs font-bold theme-text-main flex items-center gap-2 shadow-sm">
          <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
          <span>{{ members.length }} Member Terdaftar</span>
        </span>
      </div>
    </div>

    <!-- ================================================================= -->
    <!-- TAB SWITCHER: MEMBER DATABASE vs SQUAD LINE-UP                    -->
    <!-- ================================================================= -->
    <div class="flex border-b theme-border gap-2 overflow-x-auto pb-px">
      <button 
        @click="switchTab('member')" 
        :class="activeTab === 'member' ? 'border-orange-500 text-orange-500 font-black' : 'border-transparent theme-text-muted hover:theme-text-main font-semibold'"
        class="pb-3 px-4 border-b-2 text-xs flex items-center gap-2 transition whitespace-nowrap">
        <Icon name="ph:address-book-bold" class="text-base" />
        <span>Database Member & Saldo Deposit</span>
        <span class="text-[10px] px-2 py-0.5 rounded-full bg-orange-500/10 text-orange-500 font-bold">
          {{ members.length }}
        </span>
      </button>

      <button 
        @click="switchTab('squad')" 
        :class="activeTab === 'squad' ? 'border-orange-500 text-orange-500 font-black' : 'border-transparent theme-text-muted hover:theme-text-main font-semibold'"
        class="pb-3 px-4 border-b-2 text-xs flex items-center gap-2 transition whitespace-nowrap">
        <Icon name="ph:soccer-ball-bold" class="text-base" />
        <span>Line-up Squad per Jadwal Match</span>
        <span class="text-[10px] px-2 py-0.5 rounded-full bg-blue-500/10 text-blue-500 font-bold">
          {{ events.length }} Sesi
        </span>
      </button>
    </div>

    <!-- ================================================================= -->
    <!-- TAB 1: DATABASE MEMBER & SALDO DEPOSIT                            -->
    <!-- ================================================================= -->
    <div v-show="activeTab === 'member'" class="space-y-6">
      
      <!-- 4 Summary KPI Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
        
        <!-- Total Member -->
        <div class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm flex flex-col justify-between">
          <div class="flex justify-between items-start">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-blue-500 bg-blue-500/10 px-2.5 py-0.5 rounded-full border border-blue-500/20">
              👥 Member
            </span>
            <div class="w-9 h-9 rounded-2xl bg-gradient-to-tr from-blue-600 to-indigo-500 text-white flex items-center justify-center shadow-md">
              <Icon name="ph:users-bold" class="text-lg" />
            </div>
          </div>
          <div class="mt-4">
            <p class="text-2xl font-black theme-text-main tracking-tight">{{ members.length }} Orang</p>
            <p class="text-[11px] theme-text-muted mt-0.5">Total Akun Terdaftar</p>
          </div>
        </div>

        <!-- Total Saldo Deposit Komunitas -->
        <div class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm flex flex-col justify-between">
          <div class="flex justify-between items-start">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-emerald-500 bg-emerald-500/10 px-2.5 py-0.5 rounded-full border border-emerald-500/20">
              💳 Total Deposit
            </span>
            <div class="w-9 h-9 rounded-2xl bg-gradient-to-tr from-emerald-600 to-teal-500 text-white flex items-center justify-center shadow-md">
              <Icon name="ph:wallet-bold" class="text-lg" />
            </div>
          </div>
          <div class="mt-4">
            <p class="text-2xl font-black text-emerald-500 tracking-tight">
              Rp {{ totalDepositMembers.toLocaleString('id-ID') }}
            </p>
            <p class="text-[11px] theme-text-muted mt-0.5">Dana Deposit Tersimpan</p>
          </div>
        </div>

        <!-- Rata-rata Saldo per Member -->
        <div class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm flex flex-col justify-between">
          <div class="flex justify-between items-start">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-purple-500 bg-purple-500/10 px-2.5 py-0.5 rounded-full border border-purple-500/20">
              📊 Rata-rata
            </span>
            <div class="w-9 h-9 rounded-2xl bg-gradient-to-tr from-purple-600 to-pink-500 text-white flex items-center justify-center shadow-md">
              <Icon name="ph:chart-bar-bold" class="text-lg" />
            </div>
          </div>
          <div class="mt-4">
            <p class="text-2xl font-black theme-text-main tracking-tight">
              Rp {{ avgDeposit.toLocaleString('id-ID') }}
            </p>
            <p class="text-[11px] theme-text-muted mt-0.5">Rata-rata Saldo / Member</p>
          </div>
        </div>

        <!-- Member Bersaldo Aktif -->
        <div class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm flex flex-col justify-between">
          <div class="flex justify-between items-start">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-orange-500 bg-orange-500/10 px-2.5 py-0.5 rounded-full border border-orange-500/20">
              ⚡ Saldo Aktif
            </span>
            <div class="w-9 h-9 rounded-2xl bg-gradient-to-tr from-red-600 to-orange-500 text-white flex items-center justify-center shadow-md">
              <Icon name="ph:lightning-bold" class="text-lg" />
            </div>
          </div>
          <div class="mt-4">
            <p class="text-2xl font-black text-orange-500 tracking-tight">{{ membersWithDeposit.length }} Member</p>
            <p class="text-[11px] theme-text-muted mt-0.5">Memiliki Saldo > Rp 0</p>
          </div>
        </div>

      </div>

      <!-- Filter Bar & Search Member -->
      <div class="theme-bg-surface rounded-3xl p-5 border theme-border shadow-sm flex flex-col md:flex-row md:items-center justify-between gap-4">
        
        <!-- Search -->
        <div class="relative flex-1 max-w-md">
          <Icon name="ph:magnifying-glass-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 theme-text-muted text-xs" />
          <input 
            v-model="searchMember" 
            type="text" 
            placeholder="Cari nama member atau email..." 
            class="w-full pl-9 pr-3.5 py-2.5 border theme-border rounded-2xl text-xs theme-text-main theme-bg-card focus:border-orange-500 outline-none transition font-medium" 
          />
        </div>

        <!-- Filter Dropdowns -->
        <div class="flex items-center gap-2.5 flex-wrap">
          <select 
            v-model="filterDeposit" 
            class="border theme-border rounded-2xl px-3 py-2 text-xs theme-bg-card theme-text-main font-semibold outline-none focus:border-orange-500">
            <option value="ALL">Semua Member</option>
            <option value="HAS_DEPOSIT">🟢 Memiliki Saldo (> Rp 0)</option>
            <option value="ZERO_DEPOSIT">⚪ Saldo Kosong (Rp 0)</option>
          </select>

          <button 
            v-if="searchMember || filterDeposit !== 'ALL'" 
            @click="resetMemberFilter"
            type="button" 
            class="text-xs text-red-500 font-bold hover:underline px-2 py-1">
            Reset
          </button>
        </div>

      </div>

      <!-- Tabel Database Member -->
      <div class="theme-bg-surface rounded-3xl shadow-sm border theme-border overflow-hidden transition-colors duration-300">
        
        <div class="p-5 border-b theme-border theme-bg-card flex justify-between items-center">
          <div>
            <h3 class="font-black theme-text-main text-base flex items-center gap-2">
              <Icon name="ph:user-list-bold" class="text-orange-500 text-lg" />
              Daftar Member Komunitas
            </h3>
            <p class="text-xs theme-text-muted mt-0.5">Informasi akun, saldo deposit, serta total keikutsertaan main.</p>
          </div>
          <span class="text-xs theme-text-muted font-bold">
            {{ filteredMembers.length }} Member
          </span>
        </div>

        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="theme-bg-card theme-text-muted text-[11px] uppercase tracking-wider font-extrabold">
                <th class="p-4 border-b theme-border">Profil Member</th>
                <th class="p-4 border-b theme-border">Email / Kontak</th>
                <th class="p-4 border-b theme-border text-center">Partisipasi Match</th>
                <th class="p-4 border-b theme-border text-right">Saldo Deposit</th>
                <th class="p-4 border-b theme-border text-center">Aksi Pengurus</th>
              </tr>
            </thead>
            <tbody class="text-xs divide-y theme-border-subtle">
              <tr 
                v-for="user in filteredMembers" 
                :key="user.id" 
                class="hover:bg-orange-500/5 transition">
                
                <!-- Profil Member -->
                <td class="p-4">
                  <div class="flex items-center gap-3">
                    <div class="w-10 h-10 rounded-2xl overflow-hidden bg-gradient-to-tr from-orange-500 to-amber-500 text-white flex items-center justify-center font-black text-sm shrink-0 shadow-sm border theme-border relative">
                      <img 
                        v-if="user.profile_pic && !imageErrors[user.id]" 
                        :src="useApiUrl(user.profile_pic)" 
                        @error="imageErrors[user.id] = true"
                        class="w-full h-full object-cover" 
                        :alt="user.name" 
                      />
                      <span v-else>{{ getInitials(user.name) }}</span>
                    </div>
                    <div>
                      <p class="font-black theme-text-main text-sm">{{ user.name || 'Tanpa Nama' }}</p>
                      <p class="text-[10px] theme-text-muted">ID: {{ (user.id || '').slice(-6) }}</p>
                    </div>
                  </div>
                </td>

                <!-- Email / Kontak -->
                <td class="p-4">
                  <p class="theme-text-main font-medium">{{ user.email || '-' }}</p>
                  <p class="text-[10px] theme-text-muted">Terdaftar sejak {{ formatJoinDate(user.created_at) }}</p>
                </td>

                <!-- Partisipasi Match -->
                <td class="p-4 text-center">
                  <span class="px-3 py-1 rounded-full text-xs font-black bg-blue-500/10 text-blue-500 border border-blue-500/20">
                    {{ getUserMatchCount(user.id, user.name) }} Sesi Match
                  </span>
                </td>

                <!-- Saldo Deposit -->
                <td class="p-4 text-right whitespace-nowrap">
                  <span 
                    :class="(user.deposit || 0) > 0 ? 'text-emerald-500 bg-emerald-500/10 border-emerald-500/20' : 'theme-text-muted theme-bg-card border theme-border'"
                    class="px-3 py-1.5 rounded-2xl text-xs font-black border inline-block">
                    Rp {{ (user.deposit || 0).toLocaleString('id-ID') }}
                  </span>
                </td>

                <!-- Aksi Pengurus -->
                <td class="p-4 text-center whitespace-nowrap">
                  <div class="flex items-center justify-center gap-2">
                    <button 
                      @click="bukaModalSesuaikanDeposit(user)" 
                      class="px-3 py-1.5 bg-gradient-to-r from-emerald-600 to-teal-500 hover:from-emerald-700 hover:to-teal-600 text-white rounded-xl text-xs font-bold shadow-sm transition active:scale-95 flex items-center gap-1">
                      <Icon name="ph:wallet-bold" />
                      <span>Atur Saldo</span>
                    </button>
                    <button 
                      @click="bukaModalDetailMember(user)" 
                      class="px-3 py-1.5 theme-bg-card border theme-border hover:border-orange-500 theme-text-main rounded-xl text-xs font-bold transition">
                      Riwayat
                    </button>
                  </div>
                </td>

              </tr>

              <!-- Kosong -->
              <tr v-if="filteredMembers.length === 0">
                <td colspan="5" class="p-12 text-center theme-text-muted space-y-2">
                  <Icon name="ph:users-slash-bold" class="text-4xl mx-auto opacity-40" />
                  <p class="font-bold theme-text-main text-sm">Tidak Ada Member Ditemukan</p>
                  <p class="text-xs">Coba sesuaikan kata kunci pencarian atau filter saldo di atas.</p>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

      </div>

    </div>

    <!-- ================================================================= -->
    <!-- TAB 2: LINE-UP SQUAD PER JADWAL MATCH                             -->
    <!-- ================================================================= -->
    <div v-show="activeTab === 'squad'" class="space-y-6">
      
      <!-- Horizontal Session Pills Selector -->
      <div class="space-y-2">
        <div class="flex justify-between items-center">
          <label class="block text-xs font-extrabold uppercase tracking-wider theme-text-muted">
            ⚽ Pilih Sesi Jadwal Pertandingan:
          </label>
          <span class="text-xs font-bold text-orange-500">
            {{ events.length }} Sesi Terjadwal
          </span>
        </div>

        <div class="flex gap-2.5 overflow-x-auto pb-2 scrollbar-thin">
          <button 
            v-for="evt in events" 
            :key="evt.id"
            type="button"
            @click="selectedEventId = evt.id"
            :class="String(selectedEvent?.id) === String(evt.id) ? 'bg-gradient-to-r from-red-600 to-orange-600 text-white shadow-md shadow-orange-600/30 font-black' : 'theme-bg-surface theme-text-main border theme-border hover:border-orange-500/50'"
            class="px-4 py-3 rounded-2xl text-xs transition whitespace-nowrap shrink-0 flex items-center gap-2.5 active:scale-95">
            <span class="text-base">⚽</span>
            <div class="text-left">
              <p class="font-bold leading-tight">{{ evt.title }}</p>
              <p class="text-[10px] opacity-80 mt-0.5">
                {{ formatShortDate(evt.match_date) }} • {{ formatTime(evt.match_date) }} WIB
              </p>
            </div>
            <span 
              :class="String(selectedEvent?.id) === String(evt.id) ? 'bg-white/20 text-white' : 'bg-orange-500/10 text-orange-500 border border-orange-500/20'"
              class="px-2 py-0.5 rounded-full text-[10px] font-black ml-1">
              {{ getEventRegistrationCount(evt.id) }}/{{ evt.quota_max }}
            </span>
          </button>
        </div>
      </div>

      <!-- Match Banner Terpilih -->
      <div v-if="selectedEvent" class="theme-bg-surface rounded-3xl border theme-border p-6 shadow-sm space-y-4 transition-colors duration-300">
        
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
          <div class="space-y-1">
            <div class="flex items-center gap-2">
              <span class="px-2.5 py-0.5 rounded-full text-[10px] font-black uppercase tracking-wider bg-orange-500/10 text-orange-500 border border-orange-500/20">
                Sesi Terpilih
              </span>
              <span 
                v-if="selectedEvent.status === 'OPEN'" 
                class="px-2.5 py-0.5 bg-emerald-500/10 text-emerald-500 rounded-full text-[10px] font-black border border-emerald-500/30">
                OPEN
              </span>
              <span 
                v-else 
                class="px-2.5 py-0.5 theme-bg-card theme-text-muted rounded-full text-[10px] font-bold border theme-border">
                {{ selectedEvent.status }}
              </span>
            </div>
            <h2 class="text-xl font-black theme-text-main">{{ selectedEvent.title }}</h2>
            <div class="flex flex-wrap items-center gap-4 text-xs theme-text-muted pt-1">
              <span class="flex items-center gap-1 font-bold theme-text-main">
                <Icon name="ph:calendar-blank-bold" class="text-orange-500" />
                {{ formatFullDate(selectedEvent.match_date) }}
              </span>
              <span class="flex items-center gap-1 font-bold theme-text-main">
                <Icon name="ph:clock-bold" class="text-orange-500" />
                {{ formatTime(selectedEvent.match_date) }} WIB
              </span>
              <span class="flex items-center gap-1">
                <Icon name="ph:map-pin-bold" class="text-orange-500" />
                {{ selectedEvent.location }}
              </span>
              <span class="flex items-center gap-1 font-black text-emerald-500">
                <Icon name="ph:tag-bold" />
                Rp {{ (selectedEvent.price_per_person || 0).toLocaleString('id-ID') }} / orang
              </span>
            </div>
          </div>

          <!-- Tombol Copy WhatsApp Squad & Kelola -->
          <div class="flex items-center gap-2 flex-wrap">
            <button 
              @click="salinFormatWA" 
              type="button"
              class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl text-xs font-bold shadow-sm transition active:scale-95 flex items-center gap-1.5">
              <Icon name="ph:whatsapp-logo-bold" class="text-base" />
              <span>Salin Format WA</span>
            </button>
            <NuxtLink 
              to="/jadwal" 
              class="px-3.5 py-2 theme-bg-card border theme-border hover:border-orange-500 theme-text-main rounded-xl text-xs font-bold transition">
              Edit Jadwal
            </NuxtLink>
          </div>
        </div>

        <!-- Quota Progress Bar -->
        <div class="space-y-1.5 pt-2 border-t theme-border-subtle">
          <div class="flex justify-between items-center text-xs font-bold">
            <span class="theme-text-main">
              Keterisian Squad: {{ currentOfficialSquad.length }} / {{ selectedEvent.quota_max }} Pemain
            </span>
            <span class="text-emerald-500">
              {{ currentOfficialSquad.filter(p => p.payment_status === 'PAID').length }} Terkonfirmasi Lunas
            </span>
          </div>
          <div class="w-full theme-bg-card rounded-full h-2.5 overflow-hidden border theme-border">
            <div 
              class="h-full rounded-full bg-gradient-to-r from-amber-500 to-emerald-400 transition-all duration-300"
              :style="{ width: `${Math.min(100, Math.round((currentOfficialSquad.length / selectedEvent.quota_max) * 100))}%` }">
            </div>
          </div>
        </div>

      </div>

      <!-- Squad Line-up Cards (Official Quota) -->
      <div class="space-y-4">
        
        <div class="flex justify-between items-center">
          <h3 class="font-black theme-text-main text-base flex items-center gap-2">
            <span class="w-2.5 h-2.5 rounded-full bg-emerald-500"></span>
            Line-up Squad Resmi (Masuk Kuota)
          </h3>
          <span class="text-xs font-bold text-emerald-500">
            {{ currentOfficialSquad.length }} Pemain
          </span>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3.5">
          
          <div 
            v-for="(player, idx) in currentOfficialSquad" 
            :key="player.id"
            class="p-4 rounded-3xl theme-bg-surface border theme-border shadow-sm flex items-center justify-between gap-3 hover:border-orange-500/40 transition">
            
            <div class="flex items-center gap-3 overflow-hidden">
              <span class="w-8 h-8 rounded-xl bg-orange-500/10 text-orange-500 font-black text-xs flex items-center justify-center shrink-0 border border-orange-500/20">
                #{{ idx + 1 }}
              </span>

              <!-- Avatar Pemain di Squad -->
              <div class="w-9 h-9 rounded-2xl overflow-hidden bg-gradient-to-tr from-orange-500 to-amber-500 text-white flex items-center justify-center font-black text-xs shrink-0 shadow-sm border theme-border relative">
                <img 
                  v-if="getPlayerAvatar(player) && !imageErrors[player.user_id || player.id]" 
                  :src="useApiUrl(getPlayerAvatar(player))" 
                  @error="imageErrors[player.user_id || player.id] = true"
                  class="w-full h-full object-cover" 
                  :alt="player.user_name" 
                />
                <span v-else>{{ getInitials(player.user_name) }}</span>
              </div>

              <div class="truncate">
                <p class="font-black theme-text-main text-sm truncate">{{ player.user_name || 'Member' }}</p>
                <p class="text-[10px] theme-text-muted mt-0.5">{{ formatJam(player.registered_at) }} WIB</p>
              </div>
            </div>

            <!-- Status Pembayaran -->
            <div class="shrink-0 text-right">
              <span 
                v-if="player.payment_status === 'PAID'" 
                class="px-2.5 py-1 rounded-full text-[10px] font-black bg-emerald-500/10 text-emerald-500 border border-emerald-500/30">
                LUNAS
              </span>
              <span 
                v-else-if="player.payment_status === 'VERIFYING'" 
                class="px-2.5 py-1 rounded-full text-[10px] font-black bg-amber-500/10 text-amber-500 border border-amber-500/30">
                VERIFIKASI
              </span>
              <span 
                v-else 
                class="px-2.5 py-1 rounded-full text-[10px] font-black bg-red-500/10 text-red-500 border border-red-500/30">
                BELUM BAYAR
              </span>
            </div>

          </div>

          <!-- Jika Kosong -->
          <div v-if="currentOfficialSquad.length === 0" class="col-span-full p-10 text-center theme-bg-surface rounded-3xl border theme-border space-y-2">
            <Icon name="ph:users-three-bold" class="text-4xl text-orange-500 mx-auto opacity-50" />
            <p class="font-bold theme-text-main text-sm">Belum Ada Pemain Masuk Kuota</p>
            <p class="text-xs theme-text-muted">Pemain yang mendaftar di aplikasi akan langsung muncul di sini.</p>
          </div>

        </div>

      </div>

      <!-- Waiting List / Cadangan (Jika Ada) -->
      <div v-if="currentWaitingList.length > 0" class="space-y-4 pt-4 border-t theme-border-subtle">
        
        <div class="flex justify-between items-center">
          <h3 class="font-black theme-text-main text-base flex items-center gap-2">
            <span class="w-2.5 h-2.5 rounded-full bg-amber-500"></span>
            Waiting List / Cadangan (Di Luar Kuota)
          </h3>
          <span class="text-xs font-bold text-amber-500">
            {{ currentWaitingList.length }} Pemain
          </span>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3.5">
          <div 
            v-for="(wl, idx) in currentWaitingList" 
            :key="wl.id"
            class="p-4 rounded-3xl theme-bg-card border theme-border-subtle flex items-center justify-between gap-3 opacity-90">
            <div class="flex items-center gap-3 overflow-hidden">
              <span class="w-8 h-8 rounded-xl bg-gray-500/10 theme-text-muted font-bold text-xs flex items-center justify-center shrink-0 border theme-border">
                W{{ idx + 1 }}
              </span>
              <div class="truncate">
                <p class="font-bold theme-text-main text-xs truncate">{{ wl.user_name || 'Member' }}</p>
                <p class="text-[10px] theme-text-muted">Antrean Cadangan</p>
              </div>
            </div>
            <span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold bg-amber-500/10 text-amber-500 border border-amber-500/20">
              Menunggu Slot
            </span>
          </div>
        </div>

      </div>

    </div>

    <!-- ================================================================= -->
    <!-- MODAL SESUAIKAN SALDO DEPOSIT MEMBER                              -->
    <!-- ================================================================= -->
    <div v-if="showModalDeposit" class="fixed inset-0 bg-black/80 backdrop-blur-sm z-50 flex items-center justify-center p-4">
      <div class="theme-bg-surface border theme-border p-6 rounded-3xl relative max-w-md w-full shadow-2xl space-y-5 transition-colors duration-300">
        
        <div class="flex justify-between items-center border-b theme-border pb-3">
          <div>
            <h4 class="font-black theme-text-main text-base">Atur Saldo Deposit Member</h4>
            <p class="text-xs theme-text-muted mt-0.5">Pemain: <span class="font-bold text-orange-500">{{ selectedUser?.name }}</span></p>
          </div>
          <button @click="showModalDeposit = false" class="theme-text-muted hover:theme-text-main w-8 h-8 rounded-full theme-bg-card flex items-center justify-center transition">
            ✕
          </button>
        </div>

        <div class="p-3.5 rounded-2xl theme-bg-card border theme-border-subtle flex justify-between items-center">
          <div>
            <p class="text-[10px] uppercase font-bold theme-text-muted tracking-wider">Saldo Saat Ini</p>
            <p class="text-lg font-black text-emerald-500">Rp {{ (selectedUser?.deposit || 0).toLocaleString('id-ID') }}</p>
          </div>
          <Icon name="ph:wallet-bold" class="text-2xl text-emerald-500" />
        </div>

        <!-- Pilihan Aksi Deposit -->
        <div class="space-y-2">
          <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">Jenis Penyesuaian Saldo</label>
          <div class="grid grid-cols-3 gap-2">
            <button 
              type="button" 
              @click="depositAction = 'ADD'"
              :class="depositAction === 'ADD' ? 'bg-emerald-600 text-white font-black shadow-md' : 'theme-bg-card border theme-border theme-text-muted'"
              class="py-2.5 rounded-xl text-xs flex items-center justify-center gap-1 transition active:scale-95">
              <span>+ Tambah</span>
            </button>
            <button 
              type="button" 
              @click="depositAction = 'DEDUCT'"
              :class="depositAction === 'DEDUCT' ? 'bg-rose-600 text-white font-black shadow-md' : 'theme-bg-card border theme-border theme-text-muted'"
              class="py-2.5 rounded-xl text-xs flex items-center justify-center gap-1 transition active:scale-95">
              <span>- Potong</span>
            </button>
            <button 
              type="button" 
              @click="depositAction = 'SET'"
              :class="depositAction === 'SET' ? 'bg-blue-600 text-white font-black shadow-md' : 'theme-bg-card border theme-border theme-text-muted'"
              class="py-2.5 rounded-xl text-xs flex items-center justify-center gap-1 transition active:scale-95">
              <span>⚙️ Set Saldo</span>
            </button>
          </div>
        </div>

        <!-- Input Nominal -->
        <div class="space-y-2">
          <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">Nominal Saldo (Rp)</label>
          <div class="relative">
            <span class="absolute left-3.5 top-1/2 -translate-y-1/2 text-orange-500 font-bold text-xs">Rp</span>
            <input 
              v-model.number="depositAmount" 
              type="number" 
              step="5000" 
              placeholder="Contoh: 50000" 
              class="w-full pl-10 pr-4 py-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-sm theme-text-main theme-bg-card font-black transition" 
              required 
            />
          </div>

          <!-- Quick Nominal Pills -->
          <div class="flex flex-wrap gap-1.5 pt-0.5">
            <button 
              v-for="amt in [25000, 50000, 100000, 200000]" 
              :key="amt"
              type="button"
              @click="depositAmount = amt"
              class="text-[10px] px-2 py-0.5 rounded-lg border theme-border-subtle theme-bg-card theme-text-muted hover:theme-text-main transition">
              + Rp {{ (amt / 1000) }}rb
            </button>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="flex gap-2 pt-2">
          <button 
            type="button" 
            @click="showModalDeposit = false" 
            class="w-1/3 py-3 theme-bg-card border theme-border rounded-2xl text-xs font-bold theme-text-muted hover:theme-text-main transition">
            Batal
          </button>
          <button 
            type="button" 
            @click="simpanPenyesuaianDeposit" 
            class="w-2/3 py-3 bg-gradient-to-r from-red-600 to-orange-500 hover:from-red-700 hover:to-orange-700 text-white rounded-2xl text-xs font-black shadow-md transition active:scale-95">
            Konfirmasi Simpan
          </button>
        </div>

      </div>
    </div>

    <!-- ================================================================= -->
    <!-- MODAL DETAIL & RIWAYAT MEMBER                                     -->
    <!-- ================================================================= -->
    <div v-if="showModalDetail" class="fixed inset-0 bg-black/80 backdrop-blur-sm z-50 flex items-center justify-center p-4">
      <div class="theme-bg-surface border theme-border p-6 rounded-3xl relative max-w-lg w-full shadow-2xl space-y-4 transition-colors duration-300">
        
        <div class="flex justify-between items-center border-b theme-border pb-3">
          <div>
            <h4 class="font-black theme-text-main text-base">Riwayat & Info Member</h4>
            <p class="text-xs theme-text-muted mt-0.5">{{ selectedUser?.name }} ({{ selectedUser?.email }})</p>
          </div>
          <button @click="showModalDetail = false" class="theme-text-muted hover:theme-text-main w-8 h-8 rounded-full theme-bg-card flex items-center justify-center transition">
            ✕
          </button>
        </div>

        <div class="space-y-3 max-h-[60vh] overflow-y-auto pr-1">
          <div class="p-3.5 rounded-2xl theme-bg-card border theme-border-subtle flex justify-between items-center">
            <div>
              <p class="text-[10px] uppercase font-bold theme-text-muted">Total Saldo Deposit</p>
              <p class="text-xl font-black text-emerald-500">Rp {{ (selectedUser?.deposit || 0).toLocaleString('id-ID') }}</p>
            </div>
            <span class="px-3 py-1 rounded-full text-xs font-bold bg-blue-500/10 text-blue-500 border border-blue-500/20">
              {{ getUserMatchCount(selectedUser?.id, selectedUser?.name) }} Match Diikuti
            </span>
          </div>

          <h5 class="font-bold theme-text-main text-xs pt-2">Riwayat Pendaftaran Pertandingan:</h5>
          <div class="space-y-2">
            <div 
              v-for="reg in getUserRegistrations(selectedUser?.id, selectedUser?.name)" 
              :key="reg.id"
              class="p-3 rounded-2xl theme-bg-card border theme-border-subtle flex justify-between items-center text-xs">
              <div>
                <p class="font-bold theme-text-main">{{ getEventTitle(reg.event_id) }}</p>
                <p class="text-[10px] theme-text-muted">{{ formatShortDate(reg.registered_at) }}</p>
              </div>
              <span 
                :class="reg.payment_status === 'PAID' ? 'text-emerald-500 bg-emerald-500/10' : 'text-amber-500 bg-amber-500/10'"
                class="px-2.5 py-1 rounded-full text-[10px] font-black border border-current">
                {{ reg.payment_status }}
              </span>
            </div>

            <p v-if="getUserRegistrations(selectedUser?.id, selectedUser?.name).length === 0" class="text-xs theme-text-muted text-center py-4">
              Belum pernah mendaftar di pertandingan mana pun.
            </p>
          </div>
        </div>

        <div class="flex justify-end pt-2 border-t theme-border">
          <button @click="showModalDetail = false" class="px-4 py-2 theme-bg-card border theme-border text-xs font-bold rounded-xl theme-text-main">
            Tutup
          </button>
        </div>

      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'

const { $api } = useNuxtApp()
const route = useRoute()
const toast = useToast()
const { useAutoRefresh } = useRealtime()

const { data: usersData, refresh: refreshUsers } = await useApiFetch('/users')
const { data: eventsData, refresh: refreshEvents } = await useApiFetch('/events')
const { data: regData, refresh: refreshReg } = await useApiFetch('/registrations')

// Pasang Auto-Refresh Realtime
useAutoRefresh(['USER_UPDATED', 'TOPUP_UPDATED', 'REGISTRATION_UPDATED', 'EVENT_UPDATED', 'PAYMENT_UPDATED'], () => {
  refreshUsers()
  refreshEvents()
  refreshReg()
})

const activeTab = ref('member')
const selectedEventId = ref(null)
const imageErrors = ref({})

const members = computed(() => usersData.value?.data || [])
const events = computed(() => eventsData.value?.data || [])
const registrations = computed(() => regData.value?.data || [])

// Watcher untuk inisialisasi selectedEventId begitu events ter-load
watch(events, (newEvents) => {
  if (newEvents && newEvents.length > 0 && !selectedEventId.value) {
    if (route.query.event_id) {
      selectedEventId.value = route.query.event_id
    } else {
      selectedEventId.value = newEvents[0].id
    }
  }
}, { immediate: true })

onMounted(() => {
  if (route.query.tab === 'squad') {
    activeTab.value = 'squad'
  }
  if (route.query.event_id) {
    selectedEventId.value = route.query.event_id
  } else if (events.value.length > 0 && !selectedEventId.value) {
    selectedEventId.value = events.value[0].id
  }
})

const switchTab = (tab) => {
  activeTab.value = tab
  if (tab === 'squad' && !selectedEventId.value && events.value.length > 0) {
    selectedEventId.value = events.value[0].id
  }
}

// Filter Member
const searchMember = ref('')
const filterDeposit = ref('ALL')

const filteredMembers = computed(() => {
  return members.value.filter(u => {
    // Search
    if (searchMember.value) {
      const q = searchMember.value.toLowerCase()
      const matchName = u.name && u.name.toLowerCase().includes(q)
      const matchEmail = u.email && u.email.toLowerCase().includes(q)
      if (!matchName && !matchEmail) return false
    }

    // Filter Deposit
    if (filterDeposit.value === 'HAS_DEPOSIT' && (u.deposit || 0) <= 0) return false
    if (filterDeposit.value === 'ZERO_DEPOSIT' && (u.deposit || 0) > 0) return false

    return true
  })
})

const resetMemberFilter = () => {
  searchMember.value = ''
  filterDeposit.value = 'ALL'
}

// Summary Metrics Member
const totalDepositMembers = computed(() => {
  return members.value.reduce((acc, curr) => acc + (curr.deposit || 0), 0)
})

const membersWithDeposit = computed(() => {
  return members.value.filter(u => (u.deposit || 0) > 0)
})

const avgDeposit = computed(() => {
  if (members.value.length === 0) return 0
  return Math.round(totalDepositMembers.value / members.value.length)
})

// Match / Registrations Helpers
const isPlayerJoining = (pollingStatus) => {
  if (!pollingStatus) return true
  const status = pollingStatus.toUpperCase()
  return status === 'JOIN' || status === 'IN' || status === 'YES' || status === 'IKUT'
}

const isPlayerWaitingList = (pollingStatus) => {
  if (!pollingStatus) return false
  const status = pollingStatus.toUpperCase()
  return status === 'WAITING_LIST' || status === 'WAITING' || status === 'CADANGAN'
}

const getUserRegistrations = (userId, userName) => {
  return registrations.value.filter(r => {
    const matchId = userId && (String(r.user_id) === String(userId))
    const matchName = userName && r.user_name && (r.user_name.toLowerCase() === userName.toLowerCase())
    return (matchId || matchName) && isPlayerJoining(r.polling_status)
  })
}

const getUserMatchCount = (userId, userName) => {
  return getUserRegistrations(userId, userName).length
}

const getEventTitle = (eventId) => {
  const evt = events.value.find(e => String(e.id) === String(eventId))
  return evt ? evt.title : 'Jadwal Pertandingan'
}

const getEventRegistrationCount = (eventId) => {
  return registrations.value.filter(r => String(r.event_id) === String(eventId) && isPlayerJoining(r.polling_status)).length
}

// Sesi Terpilih (Squad Tab)
const selectedEvent = computed(() => {
  if (!events.value || events.value.length === 0) return null
  if (selectedEventId.value) {
    const found = events.value.find(e => String(e.id) === String(selectedEventId.value))
    if (found) return found
  }
  return events.value[0]
})

// Line-up Sesi Terpilih
const currentOfficialSquad = computed(() => {
  if (!selectedEvent.value) return []
  return registrations.value
    .filter(r => String(r.event_id) === String(selectedEvent.value.id) && isPlayerJoining(r.polling_status))
    .sort((a, b) => new Date(a.registered_at) - new Date(b.registered_at))
})

const currentWaitingList = computed(() => {
  if (!selectedEvent.value) return []
  return registrations.value
    .filter(r => String(r.event_id) === String(selectedEvent.value.id) && isPlayerWaitingList(r.polling_status))
    .sort((a, b) => new Date(a.registered_at) - new Date(b.registered_at))
})

const getPlayerAvatar = (player) => {
  const member = members.value.find(m => 
    (player.user_id && String(m.id) === String(player.user_id)) || 
    (player.user_name && m.name?.toLowerCase() === player.user_name?.toLowerCase())
  )
  return member?.profile_pic || ''
}

const getInitials = (name) => {
  if (!name) return 'JM'
  const words = name.trim().split(' ')
  if (words.length >= 2) {
    return (words[0][0] + words[1][0]).toUpperCase()
  }
  return name.slice(0, 2).toUpperCase()
}

// Modal Penyesuaian Saldo Deposit
const showModalDeposit = ref(false)
const showModalDetail = ref(false)
const selectedUser = ref(null)
const depositAction = ref('ADD')
const depositAmount = ref(50000)

const bukaModalSesuaikanDeposit = (user) => {
  selectedUser.value = user
  depositAction.value = 'ADD'
  depositAmount.value = 50000
  showModalDeposit.value = true
}

const bukaModalDetailMember = (user) => {
  selectedUser.value = user
  showModalDetail.value = true
}

const simpanPenyesuaianDeposit = async () => {
  if (!selectedUser.value?.id) return
  if (!depositAmount.value || depositAmount.value <= 0) {
    toast.error('Masukkan nominal saldo yang valid!')
    return
  }

  try {
    await $api('/admin/adjust-deposit', {
      method: 'POST',
      body: {
        user_id: selectedUser.value.id,
        amount: Number(depositAmount.value),
        action: depositAction.value
      }
    })

    toast.success(`Saldo deposit ${selectedUser.value.name} berhasil disesuaikan!`, 'Sukses')
    showModalDeposit.value = false
    refreshUsers()
  } catch (error) {
    toast.error(error.response?._data?.error || 'Gagal menyesuaikan saldo deposit', 'Error')
  }
}

// Salin Format WA Squad
const salinFormatWA = () => {
  if (!selectedEvent.value) return

  let text = `⚽ *OFFICIAL SQUAD LINE-UP - JMT SPORT* ⚽\n`
  text += `🏟️ *Match:* ${selectedEvent.value.title}\n`
  text += `📍 *Lokasi:* ${selectedEvent.value.location}\n`
  text += `📅 *Hari/Tgl:* ${formatFullDate(selectedEvent.value.match_date)}\n`
  text += `⏰ *Jam:* ${formatTime(selectedEvent.value.match_date)} WIB\n`
  text += `💰 *Tarif:* Rp ${(selectedEvent.value.price_per_person || 0).toLocaleString('id-ID')} / orang\n`
  text += `------------------------------------\n`
  text += `👥 *DAFTAR SQUAD RESMI (${currentOfficialSquad.value.length}/${selectedEvent.value.quota_max}):*\n`

  currentOfficialSquad.value.forEach((p, idx) => {
    const statusIcon = p.payment_status === 'PAID' ? '✅ Lunas' : '⏳ Menunggu Bayar'
    text += `${idx + 1}. ${p.user_name || 'Member'} (${statusIcon})\n`
  })

  if (currentWaitingList.value.length > 0) {
    text += `\n📋 *WAITING LIST / CADANGAN:*\n`
    currentWaitingList.value.forEach((wl, idx) => {
      text += `W${idx + 1}. ${wl.user_name || 'Member'}\n`
    })
  }

  text += `\n_Generated by JMT Sport Admin Command Center_`

  navigator.clipboard.writeText(text)
  toast.success('Format list squad WhatsApp berhasil disalin ke clipboard!')
}

// Format Date Helpers
const formatFullDate = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' })
}

const formatShortDate = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

const formatJam = (dateStr) => {
  if (!dateStr) return '--:--'
  const d = new Date(dateStr)
  return d.toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
}

const formatTime = (dateStr) => {
  if (!dateStr) return '--:--'
  const d = new Date(dateStr)
  return d.toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
}

const formatJoinDate = (dateStr) => {
  if (!dateStr || dateStr.startsWith('0001')) return '2026'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { month: 'short', year: 'numeric' })
}
</script>
