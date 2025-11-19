import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import tailwindcss from '@tailwindcss/vite'
import fs from 'fs'
import path from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  server: {
    https: {
      cert: fs.readFileSync(path.resolve(__dirname, '/root/.acme.sh/mcax.cn_ecc/fullchain.cer')),
      key: fs.readFileSync(path.resolve(__dirname, '/root/.acme.sh/mcax.cn_ecc/mcax.cn.key')),
    },
    host: true,
    allowedHosts: true,
  },
})
