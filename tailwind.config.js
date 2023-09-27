/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates//**/*.html", "./assets/components/**/*.js"],
  theme: {
    extend: {
      transitionProperty: {
        bubble: "transform",
      },
    },
  },
  plugins: [],
  darkMode: "class",
};
