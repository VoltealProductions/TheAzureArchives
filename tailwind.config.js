/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/*.html",
    "./public/js/*.js"
  ],
  theme: {
    extend: {
      backgroundImage: {
        'header-image': "url('/public/img/library.jpg')",
        'backdrop-image': "url('/public/img/tww-bg-2.jpeg')",
        'footer-image': "url('/public/img/entrance.jpg')",
      },
      textColor: {
        skin: {
          base: 'var(--color-base)',
          dark: 'var(--color-dark)',
          muted: 'var(--color-muted)',
          theme: 'var(--color-theme)',
          accent: 'var(--color-accent)',
          'accent-light': 'var(--color-accent-light)',
          parchment: 'var(--color-parchment)',
        },
      },
      backgroundColor: {
        skin: {
          base: 'var(--color-base)',
          dark: 'var(--color-dark)',
          muted: 'var(--color-muted)',
          theme: 'var(--color-theme)',
          'theme-dark': 'var(--color-theme-dark)',
          accent: 'var(--color-accent)',
          'accent-light': 'var(--color-accent-light)',
          parchment: 'var(--color-parchment)',
        },
      },
      borderColor: {
        skin: {
          base: 'var(--color-base)',
          dark: 'var(--color-dark)',
          muted: 'var(--color-muted)',
          theme: 'var(--color-theme)',
          accent: 'var(--color-accent)',
          'accent-light': 'var(--color-accent-light)',
          parchment: 'var(--color-parchment)',
        },
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}
