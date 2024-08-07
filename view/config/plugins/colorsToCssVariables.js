import plugin from "tailwindcss/plugin";

export const createColorsToCssVariablesPlugin = colors => plugin(({ addUtilities }) => {
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