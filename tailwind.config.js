/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./template/**/*.{html,js,templ}"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms")],
};
