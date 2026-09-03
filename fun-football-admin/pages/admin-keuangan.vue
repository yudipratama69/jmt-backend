<template>
  <div class="space-y-8 w-full max-w-7xl mx-auto px-2 sm:px-4 pb-16 transition-colors duration-300">
    
    <!-- ================================================================= -->
    <!-- HEADER LAPORAN KEUANGAN                                           -->
    <!-- ================================================================= -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 border-b theme-border-subtle pb-6">
      <div>
        <div class="flex items-center gap-2 mb-1">
          <span class="w-8 h-8 rounded-xl bg-gradient-to-tr from-emerald-600 to-teal-500 text-white flex items-center justify-center shadow-md shadow-emerald-600/30 text-base">
            <Icon name="ph:chart-line-up-bold" />
          </span>
          <h1 class="text-2xl font-black theme-text-main">Laporan & Buku Kas Keuangan</h1>
          <span class="text-[10px] uppercase tracking-wider font-extrabold px-2.5 py-0.5 rounded-full bg-emerald-500/10 text-emerald-500 border border-emerald-500/30">
            JMT Finance Pro
          </span>
        </div>
        <p class="text-xs theme-text-muted">
          Transparansi pembukuan kas komunitas, arus masuk tiket & deposit, serta pencatatan pengeluaran operasional.
        </p>
      </div>

      <!-- Action Buttons (Print & Export CSV) -->
      <div class="flex items-center gap-2.5 flex-wrap print:hidden">
        <button 
          @click="cetakLaporan" 
          type="button"
          class="px-3.5 py-2 rounded-xl theme-bg-surface border theme-border hover:border-orange-500 text-xs font-bold theme-text-main transition shadow-sm flex items-center gap-1.5 active:scale-95">
          <Icon name="ph:printer-bold" class="text-base text-orange-500" />
          <span>Cetak Laporan</span>
        </button>

        <button 
          @click="eksporCSV" 
          type="button"
          class="px-3.5 py-2 rounded-xl theme-bg-surface border theme-border hover:border-emerald-500 text-xs font-bold theme-text-main transition shadow-sm flex items-center gap-1.5 active:scale-95">
          <Icon name="ph:file-csv-bold" class="text-base text-emerald-500" />
          <span>Ekspor CSV (Excel)</span>
        </button>

        <button 
          @click="activeTab = 'catat'" 
          type="button"
          class="px-4 py-2 rounded-xl bg-gradient-to-r from-red-600 to-orange-500 hover:from-red-700 hover:to-orange-700 text-white text-xs font-black shadow-md shadow-orange-600/20 transition flex items-center gap-1.5 active:scale-95">
          <Icon name="ph:plus-circle-bold" class="text-base" />
          <span>+ Catat Kas</span>
        </button>
      </div>
    </div>

    <!-- ================================================================= -->
    <!-- KOP SURAT RESMI (HANYA MUNCUL SAAT DI-PRINT / CETAK)               -->
    <!-- ================================================================= -->
    <div class="hidden print:block border-b-2 border-black pb-4 mb-6">
      <div class="text-center space-y-1">
        <h2 class="text-2xl font-black tracking-tight text-black uppercase">JMT SPORT FUN FOOTBALL COMMUNITY</h2>
        <p class="text-xs text-gray-700">Laporan Arus Kas, Pemasukan Tiket/Deposit & Pengeluaran Operasional</p>
        <p class="text-[10px] text-gray-500">Dicetak pada: {{ formatTanggalLengkap(new Date()) }} | Oleh: Pengurus / Admin Kas</p>
      </div>
    </div>

    <!-- ================================================================= -->
    <!-- 4 KARTU METRIK UTAMA LAPORAN KEUANGAN (EXECUTIVE SUMMARY)         -->
    <!-- ================================================================= -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      
      <!-- Kartu 1: Total Saldo Kas Berjalan -->
      <div class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm relative overflow-hidden flex flex-col justify-between group transition-colors duration-300">
        <div class="flex justify-between items-start">
          <div class="space-y-1">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-orange-500 bg-orange-500/10 px-2.5 py-0.5 rounded-full border border-orange-500/20">
              💰 Saldo Kas Berjalan
            </span>
            <p class="text-xs theme-text-muted mt-1 font-medium">Likuiditas Kas Riil</p>
          </div>
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-red-500 to-orange-500 text-white flex items-center justify-center shadow-md">
            <Icon name="ph:wallet-bold" class="text-xl" />
          </div>
        </div>
        <div class="mt-4">
          <p class="text-2xl font-black theme-text-main tracking-tight">
            Rp {{ (totalSaldo || 0).toLocaleString('id-ID') }}
          </p>
          <p class="text-[11px] text-emerald-500 font-bold mt-1 flex items-center gap-1">
            <Icon name="ph:shield-check-bold" />
            <span>Kas Positif (Tersedia)</span>
          </p>
        </div>
      </div>

      <!-- Kartu 2: Total Pemasukan Kas (Inflow) -->
      <div class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm relative overflow-hidden flex flex-col justify-between group transition-colors duration-300">
        <div class="flex justify-between items-start">
          <div class="space-y-1">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-emerald-500 bg-emerald-500/10 px-2.5 py-0.5 rounded-full border border-emerald-500/20">
              📥 Total Pemasukan
            </span>
            <p class="text-xs theme-text-muted mt-1 font-medium">Tiket, Deposit & Kas Masuk</p>
          </div>
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-emerald-600 to-teal-500 text-white flex items-center justify-center shadow-md">
            <Icon name="ph:arrow-down-left-bold" class="text-xl" />
          </div>
        </div>
        <div class="mt-4">
          <p class="text-2xl font-black text-emerald-500 tracking-tight">
            + Rp {{ (totalMasuk || 0).toLocaleString('id-ID') }}
          </p>
          <p class="text-[11px] theme-text-muted font-medium mt-1">
            {{ allTransactions.filter(t => t.type === 'IN').length }} Transaksi Masuk
          </p>
        </div>
      </div>

      <!-- Kartu 3: Total Pengeluaran (Outflow) -->
      <div class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm relative overflow-hidden flex flex-col justify-between group transition-colors duration-300">
        <div class="flex justify-between items-start">
          <div class="space-y-1">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-rose-500 bg-rose-500/10 px-2.5 py-0.5 rounded-full border border-rose-500/20">
              📤 Total Pengeluaran
            </span>
            <p class="text-xs theme-text-muted mt-1 font-medium">Sewa, Alat & Operasional</p>
          </div>
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-rose-600 to-pink-500 text-white flex items-center justify-center shadow-md">
            <Icon name="ph:arrow-up-right-bold" class="text-xl" />
          </div>
        </div>
        <div class="mt-4">
          <p class="text-2xl font-black text-rose-500 tracking-tight">
            - Rp {{ (totalKeluar || 0).toLocaleString('id-ID') }}
          </p>
          <p class="text-[11px] theme-text-muted font-medium mt-1">
            {{ manualTransactions.filter(t => t.type === 'OUT').length }} Transaksi Beban
          </p>
        </div>
      </div>

      <!-- Kartu 4: Net Cash Flow / Surplus Kas -->
      <div class="theme-bg-surface p-5 rounded-3xl border theme-border shadow-sm relative overflow-hidden flex flex-col justify-between group transition-colors duration-300">
        <div class="flex justify-between items-start">
          <div class="space-y-1">
            <span class="text-[10px] font-extrabold uppercase tracking-wider text-blue-500 bg-blue-500/10 px-2.5 py-0.5 rounded-full border border-blue-500/20">
              ⚖️ Net Cash Flow
            </span>
            <p class="text-xs theme-text-muted mt-1 font-medium">Arus Kas Bersih (Surplus)</p>
          </div>
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-blue-600 to-indigo-500 text-white flex items-center justify-center shadow-md">
            <Icon name="ph:scales-bold" class="text-xl" />
          </div>
        </div>
        <div class="mt-4">
          <p class="text-2xl font-black theme-text-main tracking-tight">
            Rp {{ (totalSaldo || 0).toLocaleString('id-ID') }}
          </p>
          <p class="text-[11px] text-emerald-500 font-bold mt-1">
            {{ totalMasuk > 0 ? Math.round(((totalSaldo) / (totalMasuk || 1)) * 100) : 0 }}% Efisiensi Arus Kas
          </p>
        </div>
      </div>

    </div>

    <!-- ================================================================= -->
    <!-- TAB NAVIGASI LAPORAN (BUKU KAS / LABA RUGI / CATAT TRANSAKSI)      -->
    <!-- ================================================================= -->
    <div class="flex border-b theme-border gap-2 overflow-x-auto pb-px print:hidden">
      <button 
        @click="activeTab = 'buku_kas'" 
        :class="activeTab === 'buku_kas' ? 'border-orange-500 text-orange-500 font-black' : 'border-transparent theme-text-muted hover:theme-text-main font-semibold'"
        class="pb-3 px-4 border-b-2 text-xs flex items-center gap-2 transition whitespace-nowrap">
        <Icon name="ph:book-bookmark-bold" class="text-base" />
        <span>Buku Kas & Ledger Transaksi</span>
        <span class="text-[10px] px-2 py-0.5 rounded-full bg-orange-500/10 text-orange-500 font-bold">
          {{ allTransactions.length }}
        </span>
      </button>

      <button 
        @click="activeTab = 'laba_rugi'" 
        :class="activeTab === 'laba_rugi' ? 'border-orange-500 text-orange-500 font-black' : 'border-transparent theme-text-muted hover:theme-text-main font-semibold'"
        class="pb-3 px-4 border-b-2 text-xs flex items-center gap-2 transition whitespace-nowrap">
        <Icon name="ph:receipt-bold" class="text-base" />
        <span>Laporan Laba Rugi Sederhana</span>
      </button>

      <button 
        @click="activeTab = 'catat'" 
        :class="activeTab === 'catat' ? 'border-orange-500 text-orange-500 font-black' : 'border-transparent theme-text-muted hover:theme-text-main font-semibold'"
        class="pb-3 px-4 border-b-2 text-xs flex items-center gap-2 transition whitespace-nowrap">
        <Icon name="ph:pencil-simple-line-bold" class="text-base" />
        <span>Catat Pemasukan / Pengeluaran Baru</span>
      </button>
    </div>

    <!-- ================================================================= -->
    <!-- TAB 1: BUKU KAS & LEDGER TRANSAKSI (RUNNING BALANCE)              -->
    <!-- ================================================================= -->
    <div v-show="activeTab === 'buku_kas'" class="space-y-6">
      
      <!-- Filter Bar & Search -->
      <div class="theme-bg-surface rounded-3xl p-5 border theme-border shadow-sm flex flex-col md:flex-row md:items-center justify-between gap-4 print:hidden">
        
        <!-- Search Input -->
        <div class="relative flex-1 max-w-md">
          <Icon name="ph:magnifying-glass-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 theme-text-muted text-xs" />
          <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="Cari transaksi, nama pemain, atau nota..." 
            class="w-full pl-9 pr-3.5 py-2.5 border theme-border rounded-2xl text-xs theme-text-main theme-bg-card focus:border-orange-500 outline-none transition font-medium" 
          />
        </div>

        <!-- Filter Dropdowns -->
        <div class="flex items-center gap-2.5 flex-wrap">
          
          <!-- Filter Tipe -->
          <select 
            v-model="filterType" 
            class="border theme-border rounded-2xl px-3 py-2 text-xs theme-bg-card theme-text-main font-semibold outline-none focus:border-orange-500">
            <option value="ALL">Semua Arus Kas</option>
            <option value="IN">🟢 Uang Masuk (+)</option>
            <option value="OUT">🔴 Uang Keluar (-)</option>
          </select>

          <!-- Filter Kategori -->
          <select 
            v-model="filterCategory" 
            class="border theme-border rounded-2xl px-3 py-2 text-xs theme-bg-card theme-text-main font-semibold outline-none focus:border-orange-500">
            <option value="ALL">Semua Kategori</option>
            <option value="Deposit">Deposit Member</option>
            <option value="Tiket">Tiket Pertandingan</option>
            <option value="Sewa Lapangan">Sewa Lapangan</option>
            <option value="Alat & Bola">Peralatan & Bola</option>
            <option value="Operasional">Operasional & Lainnya</option>
          </select>

          <!-- Reset Filter -->
          <button 
            v-if="searchQuery || filterType !== 'ALL' || filterCategory !== 'ALL'" 
            @click="resetFilter"
            type="button" 
            class="text-xs text-red-500 font-bold hover:underline px-2 py-1">
            Reset
          </button>
        </div>

      </div>

      <!-- Tabel Buku Kas Akuntansi (General Ledger) -->
      <div class="theme-bg-surface rounded-3xl shadow-sm border theme-border overflow-hidden transition-colors duration-300">
        
        <div class="p-5 border-b theme-border theme-bg-card flex justify-between items-center">
          <div>
            <h3 class="font-black theme-text-main text-base flex items-center gap-2">
              <Icon name="ph:rows-bold" class="text-orange-500 text-lg" />
              Buku Kas Umum (General Ledger)
            </h3>
            <p class="text-xs theme-text-muted mt-0.5">Daftar mutasi kas dengan perhitungan saldo kumulatif berjalan.</p>
          </div>
          <span class="text-xs theme-text-muted font-bold">
            Menampilkan {{ filteredLedger.length }} Baris
          </span>
        </div>

        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="theme-bg-card theme-text-muted text-[11px] uppercase tracking-wider font-extrabold">
                <th class="p-4 border-b theme-border">Tanggal & Jam</th>
                <th class="p-4 border-b theme-border">Keterangan / Uraian</th>
                <th class="p-4 border-b theme-border">Kategori</th>
                <th class="p-4 border-b theme-border text-right">Debit (Masuk)</th>
                <th class="p-4 border-b theme-border text-right">Kredit (Keluar)</th>
                <th class="p-4 border-b theme-border text-right">Saldo Kas</th>
                <th class="p-4 border-b theme-border text-center print:hidden">Aksi / Sumber</th>
              </tr>
            </thead>
            <tbody class="text-xs divide-y theme-border-subtle">
              <tr 
                v-for="tx in filteredLedger" 
                :key="tx.id" 
                class="hover:bg-orange-500/5 transition">
                
                <!-- Tanggal & Jam -->
                <td class="p-4 whitespace-nowrap">
                  <p class="font-bold theme-text-main">{{ formatTanggal(tx.timestamp) }}</p>
                  <p class="text-[10px] theme-text-muted">{{ formatJam(tx.timestamp) }} WIB</p>
                </td>

                <!-- Keterangan -->
                <td class="p-4">
                  <p class="font-black theme-text-main text-xs leading-snug">{{ tx.description }}</p>
                  <p v-if="tx.notes" class="text-[10px] theme-text-muted mt-0.5">{{ tx.notes }}</p>
                </td>

                <!-- Kategori -->
                <td class="p-4 whitespace-nowrap">
                  <span 
                    :class="getCategoryBadgeClass(tx.category)"
                    class="px-2.5 py-1 rounded-full text-[10px] font-extrabold border flex items-center gap-1 w-max">
                    <span class="w-1.5 h-1.5 rounded-full" :class="getCategoryDotClass(tx.category)"></span>
                    {{ tx.category }}
                  </span>
                </td>

                <!-- Debit (Uang Masuk) -->
                <td class="p-4 text-right font-black text-emerald-500 whitespace-nowrap text-sm">
                  {{ tx.type === 'IN' ? '+ Rp ' + (tx.amount || 0).toLocaleString('id-ID') : '-' }}
                </td>

                <!-- Kredit (Uang Keluar) -->
                <td class="p-4 text-right font-black text-rose-500 whitespace-nowrap text-sm">
                  {{ tx.type === 'OUT' ? '- Rp ' + (tx.amount || 0).toLocaleString('id-ID') : '-' }}
                </td>

                <!-- Saldo Kumulatif Berjalan -->
                <td class="p-4 text-right font-black theme-text-main whitespace-nowrap text-sm">
                  Rp {{ (tx.runningBalance || 0).toLocaleString('id-ID') }}
                </td>

                <!-- Aksi / Sumber -->
                <td class="p-4 text-center whitespace-nowrap print:hidden">
                  <span v-if="tx.isAuto" class="px-2.5 py-1 bg-gray-500/10 theme-text-muted rounded-full text-[10px] font-bold border theme-border">
                    ⚡ Auto-Sync
                  </span>
                  <button 
                    v-else 
                    @click="hapusTransaksi(tx.id)"
                    title="Hapus Catatan Kas"
                    class="px-3 py-1 bg-rose-500/10 text-rose-500 hover:bg-rose-500/20 rounded-xl text-xs font-bold transition active:scale-95">
                    Hapus
                  </button>
                </td>

              </tr>

              <!-- Jika Kosong -->
              <tr v-if="filteredLedger.length === 0">
                <td colspan="7" class="p-12 text-center theme-text-muted space-y-2">
                  <Icon name="ph:receipt-x-bold" class="text-4xl mx-auto opacity-40" />
                  <p class="font-bold theme-text-main text-sm">Belum Ada Transaksi Kas</p>
                  <p class="text-xs">Catatan transaksi masuk atau keluar akan ditampilkan di sini secara urut.</p>
                </td>
              </tr>
            </tbody>

            <!-- Total Footer Ledger -->
            <tfoot>
              <tr class="theme-bg-card font-black theme-text-main text-xs border-t-2 theme-border">
                <td colspan="3" class="p-4 uppercase tracking-wider">TOTAL MUTASI & SALDO AKHIR</td>
                <td class="p-4 text-right text-emerald-500 text-sm font-black">+ Rp {{ (totalMasuk || 0).toLocaleString('id-ID') }}</td>
                <td class="p-4 text-right text-rose-500 text-sm font-black">- Rp {{ (totalKeluar || 0).toLocaleString('id-ID') }}</td>
                <td class="p-4 text-right text-orange-500 text-base font-black">Rp {{ (totalSaldo || 0).toLocaleString('id-ID') }}</td>
                <td class="p-4 print:hidden"></td>
              </tr>
            </tfoot>

          </table>
        </div>

      </div>

    </div>

    <!-- ================================================================= -->
    <!-- TAB 2: LAPORAN LABA RUGI SEDERHANA (INCOME STATEMENT)             -->
    <!-- ================================================================= -->
    <div v-show="activeTab === 'laba_rugi'" class="space-y-6">
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        
        <!-- Kolom Kiri: Rincian Pemasukan (Revenue Breakdown) -->
        <div class="theme-bg-surface rounded-3xl p-6 border theme-border shadow-sm space-y-4">
          <div class="flex justify-between items-center border-b theme-border-subtle pb-3">
            <h3 class="font-black theme-text-main text-sm flex items-center gap-2">
              <span class="w-2.5 h-2.5 rounded-full bg-emerald-500"></span>
              Rincian Sumber Pemasukan
            </h3>
            <span class="text-xs font-black text-emerald-500">+ Rp {{ (totalMasuk || 0).toLocaleString('id-ID') }}</span>
          </div>

          <div class="space-y-3 text-xs">
            
            <div class="flex justify-between items-center p-3 rounded-2xl theme-bg-card border theme-border-subtle">
              <div class="space-y-0.5">
                <p class="font-bold theme-text-main">1. Pembayaran Tiket Main Langsung</p>
                <p class="text-[11px] theme-text-muted">{{ autoRegistrations.length }} Pendaftar Lunas via Transfer</p>
              </div>
              <p class="font-black text-emerald-500">Rp {{ (subtotalTiket || 0).toLocaleString('id-ID') }}</p>
            </div>

            <div class="flex justify-between items-center p-3 rounded-2xl theme-bg-card border theme-border-subtle">
              <div class="space-y-0.5">
                <p class="font-bold theme-text-main">2. Top Up Saldo Deposit Member</p>
                <p class="text-[11px] theme-text-muted">{{ autoTopups.length }} Transaksi Deposit Disetujui</p>
              </div>
              <p class="font-black text-teal-500">Rp {{ (subtotalTopup || 0).toLocaleString('id-ID') }}</p>
            </div>

            <div class="flex justify-between items-center p-3 rounded-2xl theme-bg-card border theme-border-subtle">
              <div class="space-y-0.5">
                <p class="font-bold theme-text-main">3. Pemasukan Kas Lainnya / Sponsor</p>
                <p class="text-[11px] theme-text-muted">{{ manualTransactions.filter(t => t.type === 'IN').length }} Catatan Manual</p>
              </div>
              <p class="font-black text-emerald-500">Rp {{ (subtotalManualMasuk || 0).toLocaleString('id-ID') }}</p>
            </div>

          </div>
        </div>

        <!-- Kolom Kanan: Rincian Pengeluaran (Expense Breakdown) -->
        <div class="theme-bg-surface rounded-3xl p-6 border theme-border shadow-sm space-y-4">
          <div class="flex justify-between items-center border-b theme-border-subtle pb-3">
            <h3 class="font-black theme-text-main text-sm flex items-center gap-2">
              <span class="w-2.5 h-2.5 rounded-full bg-rose-500"></span>
              Rincian Pos Pengeluaran
            </h3>
            <span class="text-xs font-black text-rose-500">- Rp {{ (totalKeluar || 0).toLocaleString('id-ID') }}</span>
          </div>

          <div class="space-y-3 text-xs">
            
            <div class="flex justify-between items-center p-3 rounded-2xl theme-bg-card border theme-border-subtle">
              <div class="space-y-0.5">
                <p class="font-bold theme-text-main">1. Sewa Lapangan Pertandingan</p>
                <p class="text-[11px] theme-text-muted">Biaya sewa lapangan mini soccer / futsal</p>
              </div>
              <p class="font-black text-rose-500">Rp {{ (subtotalSewa || 0).toLocaleString('id-ID') }}</p>
            </div>

            <div class="flex justify-between items-center p-3 rounded-2xl theme-bg-card border theme-border-subtle">
              <div class="space-y-0.5">
                <p class="font-bold theme-text-main">2. Bola, Rompi & Peralatan</p>
                <p class="text-[11px] theme-text-muted">Pembelian dan perawatan aset olahraga</p>
              </div>
              <p class="font-black text-rose-500">Rp {{ (subtotalAlat || 0).toLocaleString('id-ID') }}</p>
            </div>

            <div class="flex justify-between items-center p-3 rounded-2xl theme-bg-card border theme-border-subtle">
              <div class="space-y-0.5">
                <p class="font-bold theme-text-main">3. Operasional, Wasit, Medis & Konsumsi</p>
                <p class="text-[11px] theme-text-muted">Biaya fotografer, air minum, p3k, wasit</p>
              </div>
              <p class="font-black text-rose-500">Rp {{ (subtotalOperasional || 0).toLocaleString('id-ID') }}</p>
            </div>

          </div>
        </div>

      </div>

      <!-- Ringkasan Laba / Rugi Bersih Komunitas -->
      <div class="theme-bg-surface rounded-3xl p-6 sm:p-8 border theme-border shadow-sm flex flex-col sm:flex-row sm:items-center justify-between gap-6">
        <div class="space-y-1">
          <span class="text-xs uppercase tracking-wider font-extrabold text-orange-500">
            Laporan Kinerja Arus Kas Komunitas
          </span>
          <h4 class="text-lg font-black theme-text-main">
            Surplus Kas Bersih (Sisa Saldo Kas):
          </h4>
          <p class="text-xs theme-text-muted">
            Total Kas Masuk dikurangi Seluruh Beban Operasional Komunitas.
          </p>
        </div>

        <div class="text-left sm:text-right">
          <p class="text-3xl font-black text-orange-500 tracking-tight">
            Rp {{ (totalSaldo || 0).toLocaleString('id-ID') }}
          </p>
          <p class="text-xs text-emerald-500 font-bold mt-1">
            Status Keuangan: Sehat & Likuid
          </p>
        </div>
      </div>

    </div>

    <!-- ================================================================= -->
    <!-- TAB 3: FORM CATAT TRANSAKSI KAS BARU                              -->
    <!-- ================================================================= -->
    <div v-show="activeTab === 'catat'" class="theme-bg-surface rounded-3xl shadow-sm border theme-border p-6 sm:p-8 space-y-6 transition-colors duration-300 print:hidden">
      
      <div class="flex justify-between items-center border-b theme-border-subtle pb-4">
        <div>
          <h2 class="text-lg font-black theme-text-main flex items-center gap-2">
            <Icon name="ph:pencil-simple-bold" class="text-orange-500 text-xl" />
            Catat Mutasi Kas Baru
          </h2>
          <p class="text-xs theme-text-muted mt-0.5">Pilih kategori cepat atau masukkan rincian pengeluaran/pemasukan manual.</p>
        </div>
      </div>

      <!-- Kategori Cepat (Quick Expense/Income Presets) -->
      <div class="space-y-2">
        <label class="block text-[11px] font-bold uppercase tracking-wider theme-text-muted">
          ⚡ Pilih Keperluan Cepat (Satu Klik):
        </label>
        <div class="grid grid-cols-2 sm:grid-cols-4 gap-2.5">
          <button 
            v-for="preset in quickTxPresets" 
            :key="preset.name"
            type="button"
            @click="applyTxPreset(preset)"
            class="p-3 rounded-2xl border theme-border theme-bg-card hover:border-orange-500/60 text-left transition active:scale-95 group shadow-sm">
            <span class="text-2xl block mb-1">{{ preset.icon }}</span>
            <p class="text-xs font-bold theme-text-main group-hover:text-orange-500 transition leading-tight">{{ preset.name }}</p>
            <span :class="preset.type === 'IN' ? 'text-emerald-500' : 'text-rose-500'" class="text-[10px] font-bold block mt-0.5">
              {{ preset.type === 'IN' ? '🟢 Pemasukan' : '🔴 Pengeluaran' }}
            </span>
          </button>
        </div>
      </div>

      <!-- Form Inputs -->
      <form @submit.prevent="submitTransaksi" class="space-y-5">
        
        <!-- Baris 1: Jenis Transaksi & Kategori -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          
          <!-- Jenis Transaksi -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Jenis Arus Kas <span class="text-red-500">*</span>
            </label>
            <div class="grid grid-cols-2 gap-3">
              <button 
                type="button"
                @click="form.type = 'OUT'"
                :class="form.type === 'OUT' ? 'bg-rose-600 text-white font-black shadow-md' : 'theme-bg-card border theme-border theme-text-muted'"
                class="py-3 px-4 rounded-2xl text-xs flex items-center justify-center gap-2 transition active:scale-95">
                <span>🔴 Uang Keluar (Beban)</span>
              </button>

              <button 
                type="button"
                @click="form.type = 'IN'"
                :class="form.type === 'IN' ? 'bg-emerald-600 text-white font-black shadow-md' : 'theme-bg-card border theme-border theme-text-muted'"
                class="py-3 px-4 rounded-2xl text-xs flex items-center justify-center gap-2 transition active:scale-95">
                <span>🟢 Uang Masuk (Kas)</span>
              </button>
            </div>
          </div>

          <!-- Kategori Transaksi -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Kategori Pos Kas <span class="text-red-500">*</span>
            </label>
            <select 
              v-model="form.category" 
              class="w-full border theme-border rounded-2xl p-3 theme-bg-card theme-text-main text-xs font-bold outline-none focus:border-orange-500 transition">
              <option value="Sewa Lapangan">🏟️ Sewa Lapangan Pertandingan</option>
              <option value="Alat & Bola">⚽ Bola, Rompi & Peralatan</option>
              <option value="Operasional">🥤 Wasit, Medis & Konsumsi</option>
              <option value="Kas Masuk">💰 Pemasukan Kas / Sponsor Tambahan</option>
              <option value="Lainnya">🔖 Keperluan Lainnya</option>
            </select>
          </div>

        </div>

        <!-- Baris 2: Keterangan & Nominal -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          
          <!-- Keterangan -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Uraian / Keterangan Keperluan <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <Icon name="ph:notepad-bold" class="absolute left-3.5 top-1/2 -translate-y-1/2 text-orange-500 text-base" />
              <input 
                v-model="form.description" 
                type="text" 
                placeholder="Misal: Sewa Lapangan Cozy Infinity 2 Jam" 
                class="w-full pl-10 pr-4 py-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-xs theme-text-main theme-bg-card font-medium transition" 
                required 
              />
            </div>
          </div>

          <!-- Nominal -->
          <div class="space-y-2">
            <label class="block text-xs font-bold uppercase tracking-wider theme-text-muted">
              Nominal Transaksi (Rp) <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <span class="absolute left-3.5 top-1/2 -translate-y-1/2 text-orange-500 font-bold text-xs">Rp</span>
              <input 
                v-model.number="form.amount" 
                type="number" 
                step="5000" 
                placeholder="Contoh: 300000" 
                class="w-full pl-10 pr-4 py-3 border theme-border rounded-2xl focus:border-orange-500 outline-none text-xs theme-text-main theme-bg-card font-black transition" 
                required 
              />
            </div>

            <!-- Quick Nominal Pills -->
            <div class="flex flex-wrap gap-1.5 pt-0.5">
              <button 
                v-for="amt in [50000, 100000, 250000, 500000, 1000000]" 
                :key="amt"
                type="button"
                @click="form.amount = amt"
                class="text-[10px] px-2 py-0.5 rounded-lg border theme-border-subtle theme-bg-card theme-text-muted hover:theme-text-main transition">
                + Rp {{ (amt / 1000) }}rb
              </button>
            </div>
          </div>

        </div>

        <!-- Tombol Simpan -->
        <div class="pt-4">
          <button 
            type="submit" 
            class="w-full bg-gradient-to-r from-red-600 via-orange-600 to-amber-500 hover:from-red-700 hover:to-orange-700 text-white font-black py-4 px-6 rounded-2xl shadow-lg shadow-orange-600/30 transition active:scale-95 text-sm flex items-center justify-center gap-2">
            <Icon name="ph:check-circle-bold" class="text-lg" />
            <span>Simpan Transaksi Kas ke Buku Kas</span>
          </button>
        </div>

      </form>

    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const { data: regData, refresh: refreshReg } = await useApiFetch('/registrations')
const { data: eventsData, refresh: refreshEvents } = await useApiFetch('/events')
const { data: topupData, refresh: refreshTopup } = await useApiFetch('/approved-topups')

const { useAutoRefresh } = useRealtime()
const toast = useToast()

// Pasang Auto-Refresh Realtime
useAutoRefresh(['PAYMENT_UPDATED', 'TOPUP_UPDATED', 'REGISTRATION_UPDATED', 'EVENT_UPDATED'], () => {
  refreshReg()
  refreshEvents()
  refreshTopup()
})

const activeTab = ref('buku_kas')
const searchQuery = ref('')
const filterType = ref('ALL')
const filterCategory = ref('ALL')

const manualTransactions = ref([])

onMounted(() => {
  manualTransactions.value = JSON.parse(localStorage.getItem('jmt_manual_tx') || '[]')
})

const form = ref({
  type: 'OUT',
  category: 'Sewa Lapangan',
  description: '',
  amount: ''
})

const quickTxPresets = [
  { name: 'Sewa Lapangan Mini Soccer', icon: '🏟️', type: 'OUT', category: 'Sewa Lapangan', defaultAmt: 400000 },
  { name: 'Beli Bola & Rompi Baru', icon: '⚽', type: 'OUT', category: 'Alat & Bola', defaultAmt: 250000 },
  { name: 'Wasit & Fotografer Match', icon: '🧑‍⚖️', type: 'OUT', category: 'Operasional', defaultAmt: 150000 },
  { name: 'Kas Masuk / Sponsor', icon: '💰', type: 'IN', category: 'Kas Masuk', defaultAmt: 200000 }
]

const applyTxPreset = (preset) => {
  form.value.type = preset.type
  form.value.category = preset.category
  form.value.description = preset.name
  form.value.amount = preset.defaultAmt
}

// 1. Data Auto Registrations (Pembayaran Tiket Lunas via Transfer Langsung)
const autoRegistrations = computed(() => {
  if (!regData.value?.data || !eventsData.value?.data) return []
  
  return regData.value.data
    .filter(r => r.payment_status === 'PAID' && r.payment_method !== 'deposit')
    .map(r => {
      const evt = eventsData.value.data.find(e => e.id === r.event_id)
      return {
        id: 'reg-' + r.id,
        timestamp: r.created_at || new Date().toISOString(),
        type: 'IN',
        category: 'Tiket',
        description: `Patungan Tiket: ${r.user_name || 'Member'} (${evt?.title || 'Jadwal'})`,
        amount: evt ? evt.price_per_person : 0,
        isAuto: true
      }
    })
})

// 2. Data Auto Topups (Deposit Terverifikasi)
const autoTopups = computed(() => {
  if (!topupData.value?.data) return []
  return topupData.value.data.map(t => ({
    id: 'topup-' + t._id,
    timestamp: t.created_at || new Date().toISOString(),
    type: 'IN',
    category: 'Deposit',
    description: `Top Up Saldo: ${t.user_name || 'Member'}`,
    amount: t.amount || 0,
    isAuto: true
  }))
})

// Gabungkan Seluruh Transaksi & Hitung Running Balance
const allTransactions = computed(() => {
  const combined = [
    ...autoRegistrations.value,
    ...autoTopups.value,
    ...manualTransactions.value.map(m => ({
      ...m,
      timestamp: m.timestamp || new Date(Number(m.id) || Date.now()).toISOString(),
      category: m.category || (m.type === 'IN' ? 'Kas Masuk' : 'Operasional'),
      isAuto: false
    }))
  ]

  // Urutkan dari transaksi terlama ke terbaru untuk hitung saldo berjalan
  combined.sort((a, b) => new Date(a.timestamp) - new Date(b.timestamp))

  let running = 0
  const withBalance = combined.map(tx => {
    if (tx.type === 'IN') {
      running += Number(tx.amount || 0)
    } else {
      running -= Number(tx.amount || 0)
    }
    return {
      ...tx,
      runningBalance: running
    }
  })

  // Balikkan urutan agar yang terbaru muncul di paling atas pada tabel
  return withBalance.reverse()
})

// Filter Ledger
const filteredLedger = computed(() => {
  return allTransactions.value.filter(tx => {
    // Filter Search
    if (searchQuery.value) {
      const q = searchQuery.value.toLowerCase()
      const matchDesc = tx.description?.toLowerCase().includes(q)
      const matchCat = tx.category?.toLowerCase().includes(q)
      if (!matchDesc && !matchCat) return false
    }

    // Filter Type
    if (filterType.value !== 'ALL' && tx.type !== filterType.value) {
      return false
    }

    // Filter Category
    if (filterCategory.value !== 'ALL') {
      if (filterCategory.value === 'Deposit' && tx.category !== 'Deposit') return false
      if (filterCategory.value === 'Tiket' && tx.category !== 'Tiket') return false
      if (filterCategory.value === 'Sewa Lapangan' && tx.category !== 'Sewa Lapangan') return false
      if (filterCategory.value === 'Alat & Bola' && tx.category !== 'Alat & Bola') return false
      if (filterCategory.value === 'Operasional' && tx.category !== 'Operasional') return false
    }

    return true
  })
})

const resetFilter = () => {
  searchQuery.value = ''
  filterType.value = 'ALL'
  filterCategory.value = 'ALL'
}

// Subtotals untuk Laporan Laba Rugi
const subtotalTiket = computed(() => autoRegistrations.value.reduce((acc, curr) => acc + (curr.amount || 0), 0))
const subtotalTopup = computed(() => autoTopups.value.reduce((acc, curr) => acc + (curr.amount || 0), 0))
const subtotalManualMasuk = computed(() => {
  return manualTransactions.value
    .filter(t => t.type === 'IN')
    .reduce((acc, curr) => acc + Number(curr.amount || 0), 0)
})

const subtotalSewa = computed(() => {
  return manualTransactions.value
    .filter(t => t.type === 'OUT' && (t.category === 'Sewa Lapangan' || t.description?.toLowerCase().includes('sewa') || t.description?.toLowerCase().includes('lapangan')))
    .reduce((acc, curr) => acc + Number(curr.amount || 0), 0)
})

const subtotalAlat = computed(() => {
  return manualTransactions.value
    .filter(t => t.type === 'OUT' && (t.category === 'Alat & Bola' || t.description?.toLowerCase().includes('bola') || t.description?.toLowerCase().includes('rompi')))
    .reduce((acc, curr) => acc + Number(curr.amount || 0), 0)
})

const subtotalOperasional = computed(() => {
  return totalKeluar.value - subtotalSewa.value - subtotalAlat.value
})

const totalMasuk = computed(() => subtotalTiket.value + subtotalTopup.value + subtotalManualMasuk.value)
const totalKeluar = computed(() => {
  return manualTransactions.value
    .filter(t => t.type === 'OUT')
    .reduce((acc, curr) => acc + Number(curr.amount || 0), 0)
})
const totalSaldo = computed(() => totalMasuk.value - totalKeluar.value)

// Submit Transaksi Manual
const submitTransaksi = () => {
  const newTx = {
    id: Date.now().toString(),
    timestamp: new Date().toISOString(),
    type: form.value.type,
    category: form.value.category,
    description: form.value.description,
    amount: Number(form.value.amount)
  }
  
  manualTransactions.value.unshift(newTx)
  localStorage.setItem('jmt_manual_tx', JSON.stringify(manualTransactions.value))
  
  form.value.description = ''
  form.value.amount = ''
  activeTab.value = 'buku_kas'
  toast.success('Transaksi kas berhasil dicatat ke buku kas!', 'Catatan Disimpan')
}

// Hapus Transaksi Manual
const hapusTransaksi = (id) => {
  toast.confirm({
    title: 'Hapus Transaksi Kas',
    message: 'Apakah Anda yakin ingin menghapus catatan kas manual ini?',
    confirmText: 'Ya, Hapus',
    cancelText: 'Batal',
    onConfirm: () => {
      manualTransactions.value = manualTransactions.value.filter(t => t.id !== id)
      localStorage.setItem('jmt_manual_tx', JSON.stringify(manualTransactions.value))
      toast.success('Transaksi berhasil dihapus!', 'Terhapus')
    }
  })
}

// Cetak Laporan Formal
const cetakLaporan = () => {
  activeTab.value = 'buku_kas'
  setTimeout(() => {
    window.print()
  }, 200)
}

// Ekspor ke CSV Excel
const eksporCSV = () => {
  if (allTransactions.value.length === 0) {
    toast.error('Tidak ada data transaksi untuk diekspor.')
    return
  }

  let csvContent = 'data:text/csv;charset=utf-8,'
  csvContent += 'Tanggal,Jam,Uraian Transaksi,Kategori,Tipe,Debit (Masuk),Kredit (Keluar),Saldo Berjalan\n'

  allTransactions.value.forEach(tx => {
    const d = new Date(tx.timestamp)
    const dateStr = d.toLocaleDateString('id-ID')
    const timeStr = d.toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
    const desc = `"${(tx.description || '').replace(/"/g, '""')}"`
    const cat = `"${tx.category || ''}"`
    const debit = tx.type === 'IN' ? tx.amount : 0
    const kredit = tx.type === 'OUT' ? tx.amount : 0
    const balance = tx.runningBalance || 0

    csvContent += `${dateStr},${timeStr},${desc},${cat},${tx.type},${debit},${kredit},${balance}\n`
  })

  const encodedUri = encodeURI(csvContent)
  const link = document.createElement('a')
  link.setAttribute('href', encodedUri)
  link.setAttribute('download', `Laporan_Keuangan_JMT_${new Date().toISOString().slice(0, 10)}.csv`)
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  toast.success('Laporan kas berhasil diekspor ke file CSV/Excel!')
}

// Format Helpers
const formatTanggal = (dateStr) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

const formatTanggalLengkap = (d) => {
  return d.toLocaleDateString('id-ID', { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' })
}

const formatJam = (dateStr) => {
  if (!dateStr) return '--:--'
  const d = new Date(dateStr)
  return d.toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
}

const getCategoryBadgeClass = (category) => {
  switch (category) {
    case 'Deposit':
      return 'bg-teal-500/10 text-teal-500 border-teal-500/20'
    case 'Tiket':
      return 'bg-emerald-500/10 text-emerald-500 border-emerald-500/20'
    case 'Sewa Lapangan':
      return 'bg-rose-500/10 text-rose-500 border-rose-500/20'
    case 'Alat & Bola':
      return 'bg-amber-500/10 text-amber-500 border-amber-500/20'
    case 'Operasional':
      return 'bg-purple-500/10 text-purple-500 border-purple-500/20'
    default:
      return 'bg-blue-500/10 text-blue-500 border-blue-500/20'
  }
}

const getCategoryDotClass = (category) => {
  switch (category) {
    case 'Deposit':
      return 'bg-teal-500'
    case 'Tiket':
      return 'bg-emerald-500'
    case 'Sewa Lapangan':
      return 'bg-rose-500'
    case 'Alat & Bola':
      return 'bg-amber-500'
    case 'Operasional':
      return 'bg-purple-500'
    default:
      return 'bg-blue-500'
  }
}
</script>