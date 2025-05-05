import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  // optimizeDeps can likely be removed completely if no other complex deps exist
  optimizeDeps: {
    include: [
      // Removed deck.gl/luma.gl entries
    ],
    force: true // Keep force: true for now, can remove later if stable
  }
})
