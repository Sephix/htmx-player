/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates//**/*.html"],
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
