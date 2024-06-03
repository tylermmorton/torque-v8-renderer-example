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
            preserveEntrySignatures: 'strict',
            input: {
                "login": './app/routes/login/login.client.ts',
            },
            output: {
                format: 'esm',
                exports: 'named',
                entryFileNames: '[name].js',
            },
        },
    },
})