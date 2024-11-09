/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./components/**/*.templ"],
  theme: {
    extend: {
      colors: require("tailwindcss/colors"),
    },
  },
  plugins: [],
}
