<template>
  <div>
    <!-- Toast Notifications Container (Floating Top-Right / Center Top on Mobile) -->
    <div class="fixed top-4 left-1/2 -translate-x-1/2 sm:translate-x-0 sm:left-auto sm:right-5 z-[9999] flex flex-col gap-2.5 max-w-sm w-[92vw] sm:w-80 pointer-events-none">
      <transition-group name="toast-slide">
        <div 
          v-for="toast in toasts" 
          :key="toast.id" 
          class="pointer-events-auto rounded-2xl p-4 shadow-2xl backdrop-blur-xl border flex items-start gap-3 relative overflow-hidden transition-all duration-300"
          :class="{
            'bg-emerald-950/90 text-emerald-100 border-emerald-500/40 shadow-emerald-950/50': toast.type === 'success',
            'bg-red-950/90 text-red-100 border-red-500/40 shadow-red-950/50': toast.type === 'error',
            'bg-amber-950/90 text-amber-100 border-amber-500/40 shadow-amber-950/50': toast.type === 'warning',
            'bg-slate-900/90 text-slate-100 border-slate-700 shadow-slate-950/50': toast.type === 'info'
          }">
          
          <!-- Accent Line -->
          <div 
            class="absolute top-0 left-0 bottom-0 w-1.5"
            :class="{
              'bg-emerald-400': toast.type === 'success',
              'bg-red-500': toast.type === 'error',
              'bg-amber-400': toast.type === 'warning',
              'bg-orange-500': toast.type === 'info'
            }">
          </div>

          <!-- Icon Status -->
          <div class="shrink-0 mt-0.5 ml-1">
            <Icon v-if="toast.type === 'success'" name="ph:check-circle-fill" class="text-xl text-emerald-400" />
            <Icon v-else-if="toast.type === 'error'" name="ph:warning-circle-fill" class="text-xl text-red-400" />
            <Icon v-else-if="toast.type === 'warning'" name="ph:warning-bold" class="text-xl text-amber-400" />
            <Icon v-else name="ph:info-fill" class="text-xl text-orange-400" />
          </div>

          <!-- Message Content -->
          <div class="flex-1 pr-4">
            <h5 v-if="toast.title" class="text-xs font-black tracking-wide leading-tight mb-0.5">
              {{ toast.title }}
            </h5>
            <p class="text-xs font-medium opacity-90 leading-snug">
              {{ toast.message }}
            </p>
          </div>

          <!-- Close Button -->
          <button @click="remove(toast.id)" class="shrink-0 opacity-60 hover:opacity-100 text-sm p-0.5 transition">
            ✕
          </button>
        </div>
      </transition-group>
    </div>

    <!-- Confirm Dialog Modal (Sporty Replacement for window.confirm) -->
    <div v-if="confirmModal.show" class="fixed inset-0 bg-black/80 backdrop-blur-md z-[10000] flex items-center justify-center p-4 animate-fade-in">
      <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 max-w-sm w-full shadow-2xl text-white relative animate-scale-up overflow-hidden">
        
        <!-- Top Glow -->
        <div class="absolute top-0 left-0 right-0 h-1.5 bg-gradient-to-r from-red-600 via-orange-500 to-amber-400"></div>

        <!-- Header Icon -->
        <div class="w-14 h-14 rounded-2xl bg-orange-500/10 border border-orange-500/30 text-orange-400 flex items-center justify-center mx-auto mb-3 shadow-lg">
          <Icon name="ph:question-bold" class="text-3xl" />
        </div>

        <h4 class="text-base font-black text-center text-white mb-1">
          {{ confirmModal.options?.title || 'Konfirmasi Aksi' }}
        </h4>
        <p class="text-xs text-slate-300 text-center mb-6 leading-relaxed">
          {{ confirmModal.options?.message }}
        </p>

        <!-- Action Buttons -->
        <div class="grid grid-cols-2 gap-3">
          <button 
            @click="handleCancel" 
            class="py-3 px-4 bg-slate-800 hover:bg-slate-700 text-slate-300 font-bold rounded-2xl text-xs transition active:scale-95 border border-slate-700">
            {{ confirmModal.options?.cancelText || 'Batal' }}
          </button>
          <button 
            @click="handleConfirm" 
            class="py-3 px-4 bg-gradient-to-r from-red-600 to-orange-600 hover:from-red-700 hover:to-orange-700 text-white font-black rounded-2xl text-xs shadow-lg shadow-orange-600/30 transition active:scale-95">
            {{ confirmModal.options?.confirmText || 'Ya, Lanjutkan' }}
          </button>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
const { toasts, confirmModal, remove, handleConfirm, handleCancel } = useToast()
</script>

<style scoped>
.toast-slide-enter-active,
.toast-slide-leave-active {
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}
.toast-slide-enter-from {
  opacity: 0;
  transform: translateY(-20px) scale(0.95);
}
.toast-slide-leave-to {
  opacity: 0;
  transform: translateY(-10px) scale(0.9);
}

.animate-fade-in {
  animation: fadeIn 0.2s ease-out forwards;
}
.animate-scale-up {
  animation: scaleUp 0.25s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes scaleUp {
  from { opacity: 0; transform: scale(0.92); }
  to { opacity: 1; transform: scale(1); }
}
</style>
