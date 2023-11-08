/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}",
  "./components/**/*.{html,js,jsx}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('tailwindcss-animated')
  ],
}

