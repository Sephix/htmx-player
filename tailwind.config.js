/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal//**/*.templ"],
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
