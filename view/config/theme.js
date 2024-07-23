const defaultTheme = require("tailwindcss/defaultTheme");

const baseColors = {
  white: "#ffffff",
  gray: "#e4e4e7",
  dim: "#333333",
  black: "#000000",

  blue: "#0000ff",
  blueLight: "#5555ff",
  blueDark: "#000055",
};

export const colors = {
  // ...baseColors,
  background: {
    regular: baseColors.gray,
    light: baseColors.white,
    dark: baseColors.black,
  },
  accent: {
    regular: baseColors.blue,
    light: baseColors.blueLight,
    dark: baseColors.blueDark,
  },
  foreground: {
    regular: baseColors.dim,
    light: baseColors.white,
    dark: baseColors.black,
  },
};

export const fontFamily = {
  body: ["system-ui", "sans-serif", ...defaultTheme.fontFamily.sans],
  heading: ["system-ui", "sans-serif", ...defaultTheme.fontFamily.sans],
  monospace: ["Menlo", "monospace", ...defaultTheme.fontFamily.mono],
};

