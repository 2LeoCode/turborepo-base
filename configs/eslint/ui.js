const base = require("./base");
const globals = require("globals");
const ts = require("typescript-eslint");
const svelte = require("eslint-plugin-svelte");
svelte.configs["flat/recommended"];

/** @type {import("eslint").Linter.FlatConfig[]} */
module.exports = [
  ...base,
  ...svelte.configs["flat/recommended"],
  ...svelte.configs["flat/prettier"],
  {
    languageOptions: {
      parserOptions: {
        parser: ts.parser,
      },
      globals: {
        ...globals.browser,
      },
    },
  },
  {
    ignores: [".*.svelte"],
  },
];
