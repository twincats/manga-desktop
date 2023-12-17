// vite.config.ts
import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "file:///D:/Web/desktop/mangav4/frontend/node_modules/.pnpm/vite@4.2.1_@types+node@18.15.11_less@4.1.3/node_modules/vite/dist/node/index.js";
import vue from "file:///D:/Web/desktop/mangav4/frontend/node_modules/.pnpm/@vitejs+plugin-vue@4.1.0_vite@4.2.1_vue@3.2.47/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import { ArcoResolver } from "file:///D:/Web/desktop/mangav4/frontend/node_modules/.pnpm/unplugin-vue-components@0.24.1_vue@3.2.47/node_modules/unplugin-vue-components/dist/resolvers.mjs";
import Components from "file:///D:/Web/desktop/mangav4/frontend/node_modules/.pnpm/unplugin-vue-components@0.24.1_vue@3.2.47/node_modules/unplugin-vue-components/dist/vite.mjs";
import AutoImport from "file:///D:/Web/desktop/mangav4/frontend/node_modules/.pnpm/unplugin-auto-import@0.15.2_@vueuse+core@10.0.2/node_modules/unplugin-auto-import/dist/vite.js";
import Icons from "file:///D:/Web/desktop/mangav4/frontend/node_modules/.pnpm/unplugin-icons@0.16.1/node_modules/unplugin-icons/dist/vite.mjs";
import IconsResolver from "file:///D:/Web/desktop/mangav4/frontend/node_modules/.pnpm/unplugin-icons@0.16.1/node_modules/unplugin-icons/dist/resolver.mjs";
import Unocss from "file:///D:/Web/desktop/mangav4/frontend/node_modules/.pnpm/unocss@0.51.4_postcss@8.4.23_vite@4.2.1/node_modules/unocss/dist/vite.mjs";
var __vite_injected_original_import_meta_url = "file:///D:/Web/desktop/mangav4/frontend/vite.config.ts";
var vite_config_default = defineConfig({
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", __vite_injected_original_import_meta_url)),
      "@wails": fileURLToPath(new URL("./wailsjs", __vite_injected_original_import_meta_url))
    }
  },
  plugins: [
    vue(),
    Unocss(),
    AutoImport({
      imports: ["vue", "vue-router", "@vueuse/core"],
      resolvers: [ArcoResolver()]
    }),
    Components({
      resolvers: [ArcoResolver({ sideEffect: true }), IconsResolver()]
    }),
    Icons({ autoInstall: true })
  ]
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJEOlxcXFxXZWJcXFxcZGVza3RvcFxcXFxtYW5nYXY0XFxcXGZyb250ZW5kXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJEOlxcXFxXZWJcXFxcZGVza3RvcFxcXFxtYW5nYXY0XFxcXGZyb250ZW5kXFxcXHZpdGUuY29uZmlnLnRzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9EOi9XZWIvZGVza3RvcC9tYW5nYXY0L2Zyb250ZW5kL3ZpdGUuY29uZmlnLnRzXCI7aW1wb3J0IHsgZmlsZVVSTFRvUGF0aCwgVVJMIH0gZnJvbSAnbm9kZTp1cmwnXG5cbmltcG9ydCB7IGRlZmluZUNvbmZpZyB9IGZyb20gJ3ZpdGUnXG5pbXBvcnQgdnVlIGZyb20gJ0B2aXRlanMvcGx1Z2luLXZ1ZSdcblxuaW1wb3J0IHsgQXJjb1Jlc29sdmVyIH0gZnJvbSAndW5wbHVnaW4tdnVlLWNvbXBvbmVudHMvcmVzb2x2ZXJzJ1xuaW1wb3J0IENvbXBvbmVudHMgZnJvbSAndW5wbHVnaW4tdnVlLWNvbXBvbmVudHMvdml0ZSdcbmltcG9ydCBBdXRvSW1wb3J0IGZyb20gJ3VucGx1Z2luLWF1dG8taW1wb3J0L3ZpdGUnXG5pbXBvcnQgSWNvbnMgZnJvbSAndW5wbHVnaW4taWNvbnMvdml0ZSdcbmltcG9ydCBJY29uc1Jlc29sdmVyIGZyb20gJ3VucGx1Z2luLWljb25zL3Jlc29sdmVyJ1xuaW1wb3J0IFVub2NzcyBmcm9tICd1bm9jc3Mvdml0ZSdcblxuLy8gaHR0cHM6Ly92aXRlanMuZGV2L2NvbmZpZy9cbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZyh7XG4gIHJlc29sdmU6IHtcbiAgICBhbGlhczoge1xuICAgICAgJ0AnOiBmaWxlVVJMVG9QYXRoKG5ldyBVUkwoJy4vc3JjJywgaW1wb3J0Lm1ldGEudXJsKSksXG4gICAgICAnQHdhaWxzJzogZmlsZVVSTFRvUGF0aChuZXcgVVJMKCcuL3dhaWxzanMnLCBpbXBvcnQubWV0YS51cmwpKSxcbiAgICB9LFxuICB9LFxuICBwbHVnaW5zOiBbXG4gICAgdnVlKCksXG4gICAgVW5vY3NzKCksXG4gICAgQXV0b0ltcG9ydCh7XG4gICAgICBpbXBvcnRzOiBbJ3Z1ZScsICd2dWUtcm91dGVyJywgJ0B2dWV1c2UvY29yZSddLFxuICAgICAgcmVzb2x2ZXJzOiBbQXJjb1Jlc29sdmVyKCldLFxuICAgIH0pLFxuICAgIENvbXBvbmVudHMoe1xuICAgICAgcmVzb2x2ZXJzOiBbQXJjb1Jlc29sdmVyKHsgc2lkZUVmZmVjdDogdHJ1ZSB9KSwgSWNvbnNSZXNvbHZlcigpXSxcbiAgICB9KSxcbiAgICBJY29ucyh7IGF1dG9JbnN0YWxsOiB0cnVlIH0pLFxuICBdLFxufSlcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBeVIsU0FBUyxlQUFlLFdBQVc7QUFFNVQsU0FBUyxvQkFBb0I7QUFDN0IsT0FBTyxTQUFTO0FBRWhCLFNBQVMsb0JBQW9CO0FBQzdCLE9BQU8sZ0JBQWdCO0FBQ3ZCLE9BQU8sZ0JBQWdCO0FBQ3ZCLE9BQU8sV0FBVztBQUNsQixPQUFPLG1CQUFtQjtBQUMxQixPQUFPLFlBQVk7QUFWNEosSUFBTSwyQ0FBMkM7QUFhaE8sSUFBTyxzQkFBUSxhQUFhO0FBQUEsRUFDMUIsU0FBUztBQUFBLElBQ1AsT0FBTztBQUFBLE1BQ0wsS0FBSyxjQUFjLElBQUksSUFBSSxTQUFTLHdDQUFlLENBQUM7QUFBQSxNQUNwRCxVQUFVLGNBQWMsSUFBSSxJQUFJLGFBQWEsd0NBQWUsQ0FBQztBQUFBLElBQy9EO0FBQUEsRUFDRjtBQUFBLEVBQ0EsU0FBUztBQUFBLElBQ1AsSUFBSTtBQUFBLElBQ0osT0FBTztBQUFBLElBQ1AsV0FBVztBQUFBLE1BQ1QsU0FBUyxDQUFDLE9BQU8sY0FBYyxjQUFjO0FBQUEsTUFDN0MsV0FBVyxDQUFDLGFBQWEsQ0FBQztBQUFBLElBQzVCLENBQUM7QUFBQSxJQUNELFdBQVc7QUFBQSxNQUNULFdBQVcsQ0FBQyxhQUFhLEVBQUUsWUFBWSxLQUFLLENBQUMsR0FBRyxjQUFjLENBQUM7QUFBQSxJQUNqRSxDQUFDO0FBQUEsSUFDRCxNQUFNLEVBQUUsYUFBYSxLQUFLLENBQUM7QUFBQSxFQUM3QjtBQUNGLENBQUM7IiwKICAibmFtZXMiOiBbXQp9Cg==
