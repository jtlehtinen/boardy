import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';
import css from 'rollup-plugin-css-only';
import sveltePreprocess from 'svelte-preprocess';
import autoprefixer from 'autoprefixer';
import tailwindcss from 'tailwindcss';

const production = !process.env.ROLLUP_WATCH;

function serve() {
  let server;

  function toExit() {
    if (server) server.kill(0);
  }

  return {
    writeBundle() {
      if (server) return;
      server = require('child_process').spawn('npm', ['run', 'start', '--', '--dev'], {
        stdio: ['ignore', 'inherit', 'inherit'],
        shell: true
      });

      process.on('SIGTERM', toExit);
      process.on('exit', toExit);
    }
  };
}

export default {
  input: 'src/main.js',
  output: {
    sourcemap: true,
    format: 'iife',
    name: 'app',
    file: 'public/build/bundle.js'
  },
  plugins: [
    // configuration for rollup-plugin-svelte:
    svelte({
      // https://svelte.dev/docs#compile-time-svelte-preprocess
      preprocess: sveltePreprocess({
        sourceMap: !production,
        postcss: {
          plugins: [
            autoprefixer(),
            tailwindcss(),
          ]
        }
      }
      ),

      // https://svelte.dev/docs#compile-time-svelte-compile
      compilerOptions: {
        dev: !production // Inject runtime checks and debug info.
      }
    }),
    css({ output: 'bundle.css' }),

    resolve({
      browser: true,
      dedupe: ['svelte']
    }),

    commonjs(),

    !production && serve(), // Call `npm run start` after bundle generated.
    !production && livereload('public'), // Watch 'public' dir and refresh browser on change.
    production && terser() // Minify
  ],
  watch: {
    clearScreen: false // Don't clear the terminal
  }
};
