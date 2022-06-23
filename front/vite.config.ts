import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  base: "./",
  server: {
    watch: { usePolling: true },
    port: 3000,
    proxy: { "^/api/.*": "http://localhost:3001" },
  },
})
