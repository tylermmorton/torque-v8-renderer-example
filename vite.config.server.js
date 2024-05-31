import { defineConfig } from 'vite'
import Vue from '@vitejs/plugin-vue'

export default defineConfig({
    plugins: [
        Vue({
            reactivityTransform: true,
        }),
    ],
    build: {
        manifest: 'manifest.json',
        outDir: './app/.dist/server',
        emptyOutDir: true,
        rollupOptions: {
            input: {
                "login": "./app/routes/login/login.server.ts",
            },
            output: {
                format: "cjs",
                entryFileNames: '[name].js',
                inlineDynamicImports: true,
            },
        },
    },
    ssr: {
        noExternal: /./,
    },
    resolve: {
        // necessary because vue.ssrUtils is only exported on cjs modules
        alias: [
            {
                find: '@vue/runtime-dom',
                replacement: '@vue/runtime-dom/dist/runtime-dom.cjs.js',
            },
            {
                find: '@vue/runtime-core',
                replacement: '@vue/runtime-core/dist/runtime-core.cjs.js',
            },
        ],
    },
})