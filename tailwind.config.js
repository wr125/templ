/** @type {import('tailwindcss').Config} */
const withMT = require("@material-tailwind/html/utils/withMT")
module.exports = withMT({
    content: [
      "./templates/**/*.{html,js,templ,go}",
      "./templates/common/**/*.{html,js,templ,go}",
    ],
    theme: {
      screens: {
        sm: "640px",
        // rest of the breakpoints
        md: "720px",
        lg: "960px",
        "lg-max": { max: "960px" },
        xl: "1140px",
        "2xl": "1320px",
      },
      boxShadow: {
        sm: "0 2px 4px 0 rgb(0 0 0 / 0.05)",
        // rest of the box shadow values
      },
      extend: {
       /* gridColumn: {
          'span-16': 'span 16 / span 16',
        },*/
        colors: {
          sky: {
            50: "#f0f9ff",
            100: "#e0f2fe",
            200: "#bae6fd",
            300: "#7dd3fc",
            400: "#38bdf8",
            500: "#0ea5e9",
            600: "#0284c7",
            700: "#0369a1",
            800: "#075985",
            900: "#0c4a6e",
          },
        },
      },
      fontFamily: {
        sans: ["Quicksand"],
      },
    },
    plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
  });