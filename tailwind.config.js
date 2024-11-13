/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ", "./node_modules/flowbite/**/*.js"],
  theme: {
    extend: {
      colors: require("tailwindcss/colors"),
    },
  },
  plugins: [
    require('flowbite/plugin'),
  ],
}
