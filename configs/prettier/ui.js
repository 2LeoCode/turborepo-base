const base = require("./base");

/** @type {import("prettier").Config} */
module.exports = {
  ...base,
  overrides: [
    {
      files: "*.svelte",
      options: {
        plugins: ["prettier-plugin-svelte"],
        parser: "svelte",
      },
    },
  ],
};
