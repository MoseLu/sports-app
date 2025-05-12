declare module 'vite-plugin-vue-svg' {
  import type { Plugin } from 'vite';
  function svg(options?: {
    defaultExport?: 'component' | 'url';
    svgoConfig?: Record<string, unknown>;
  }): Plugin;
  export default svg;
}
