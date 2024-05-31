import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    build: {
        manifest: 'manifest.json',
        outDir: './app/.dist/client',
        emptyOutDir: true,
        rollupOptions:{
            input: {
                "login": './app/routes/login/login.client.ts',
            },
            output: {
                format: 'es',
                entryFileNames: '[name].js',
                inlineDynamicImports: true,
            },
        },
    },
})