const baseColors = {
  white: '#ffffff',
  gray: '#e4e4e7',
  dim: '#333333',
  dark: '#202020',
  black: '#000000',

  blue: '#0000ff',
  blueLight: '#5555ff',
  blueDark: '#000055',
};

export const colors = {
  // ...baseColors,
  dark: baseColors.dark,
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
