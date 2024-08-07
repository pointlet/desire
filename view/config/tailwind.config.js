import { colors, fontFamily } from "./theme/index";
import { createColorsToCssVariablesPlugin } from "./plugins/colorsToCssVariables";

/** @type {import('tailwindcss').Config} */
export default {
  mode: "jit",
  content: ["view/**/*.html", "view/**/*.templ"],
  theme: {
    extend: {
      colors,
      fontFamily
    },
  },
  plugins: [
    createColorsToCssVariablesPlugin(colors)
  ],
}