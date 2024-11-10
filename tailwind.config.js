/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ"],
  theme: {
    extend: {
      colors: require("tailwindcss/colors"),
    },
  },
  plugins: [],
}
