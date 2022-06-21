import react from "@vitejs/plugin-react"
import { viteSingleFile } from "vite-plugin-singlefile"
import { defineConfig } from "vite"

export default defineConfig({
    plugins: [react(), viteSingleFile()],
    build: {
        target: "esnext",
        assetsInlineLimit: 100000000,
        chunkSizeWarningLimit: 100000000,
        cssCodeSplit: false,
        brotliSize: false,
        rollupOptions: {
            inlineDynamicImports: true,
        },
    },
    server: {
        watch: { usePolling: true },
        port: 3000,
        proxy: { "^/api/.*": "http://localhost:3001" },
    },
})