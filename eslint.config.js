import globals from 'globals';
import js from '@eslint/js';
import eslintConfigPrettier from 'eslint-config-prettier';
import prettierPlugin from 'eslint-plugin-prettier';

/** @type { import("eslint").Linter.Config[] } */
export default [
  // Config files (node)
  {
    files: ['view/config/**/*.js'],
    rules: {
      ...js.configs.recommended.rules,
    },
    languageOptions: {
      globals: globals.node,
    },
    plugins: {
      format: prettierPlugin,
    },
  },
  // Browser logic files (browser)
  /*
  {
    files: [], // TODO
    rules: {
      ...js.configs.recommended.rules,
    },
    languageOptions: { 
      globals: globals.browser
    },
    plugins: {
      format: prettierPlugin,
    },
  },
  */
  eslintConfigPrettier,
];
