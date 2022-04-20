const production = !process.env.ROLLUP_WATCH

module.exports = {
  // https://tailwindcss.com/docs/content-configuration
  content: ['./src/**/*.{js,svelte}'],
  theme: {
    extend: {},
    fontFamily: {
      sans: ['Inter', 'sans-serif'],
    }
  },
  plugins: [],
};
