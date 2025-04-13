import { sentryVitePlugin } from "@sentry/vite-plugin";
import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from 'path'
import tailwindcss from "@tailwindcss/vite";

export default defineConfig(({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd(), "SENTRY") };

  return {
    build: {
      outDir: '../public',
      emptyOutDir: true,
      sourcemap: true,
    },
    plugins: [
      tailwindcss(),
      react(),
      sentryVitePlugin({
        org: "",
        project: "",
        authToken: process.env.SENTRY_AUTH_TOKEN,
        telemetry: false,
      })
    ],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
    server: {
      port: 3000,
      proxy: {
        '/api': {
          target: 'http://localhost:8000',
          changeOrigin: true,
        }
      }
    }
  }
})
