import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tailwindcss from '@tailwindcss/vite'
import tsconfigPaths from 'vite-tsconfig-paths'
import { tanstackRouter } from '@tanstack/router-plugin/vite'

export default defineConfig({
  server: {
    host: true
  },
  preview: {
    host: true
  },
  plugins: [
    tanstackRouter({
      target: 'react',
      autoCodeSplitting: true
    }),
    react(),
    tsconfigPaths(),
    tailwindcss()
  ]
})
