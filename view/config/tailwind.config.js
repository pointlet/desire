const { colors, fontFamily } = require('./theme');
const plugin = require("tailwindcss/plugin");

const colorsToCssVariablesPlugin = plugin(({ addUtilities }) => {
  const toCssVariables = (colors, prefix = "color") => {
    return Object
      .entries(colors)
      .reduce(
        ((acc, [name, value]) => {
          if(typeof value === "object") {
            return { 
              ...acc, 
              ...toCssVariables(value, `${prefix}-${name}`) 
            }
          }

          acc[`--${prefix}-${name}`] = value
          return acc
        }),
        {}
      );
  }

  const cssVariables = toCssVariables(colors);

  addUtilities({
    ":root": cssVariables
  });
});

/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  content: ["view/**/*.html", "view/**/*.templ"],
  theme: {
    extend: {
      colors,
      fontFamily
    },
  },
  plugins: [
    colorsToCssVariablesPlugin
  ],
};
