import defaultTheme from 'tailwindcss/defaultTheme';

export const fontFamily = {
  body: ['system-ui', 'sans-serif', ...defaultTheme.fontFamily.sans],
  heading: ['system-ui', 'sans-serif', ...defaultTheme.fontFamily.sans],
  monospace: ['Menlo', 'monospace', ...defaultTheme.fontFamily.mono],
};
