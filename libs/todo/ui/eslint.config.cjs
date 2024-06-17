// eslint-disable no-undef

const ui = require("@repo/configs-eslint/ui");

/** @type {import("eslint").Linter.FlatConfig[]} */
module.exports = [
  ...ui,
  {
    ignores: ["*.cjs", ".svelte-kit/"],
  },
];
