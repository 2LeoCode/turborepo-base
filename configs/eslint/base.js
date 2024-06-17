const path = require("path");

const { FlatCompat } = require("@eslint/eslintrc");
const js = require("@eslint/js");
const ts = require("typescript-eslint");
const prettier = require("eslint-config-prettier");

const project = path.resolve(process.cwd(), "tsconfig.json");
const compat = new FlatCompat({
  baseDirectory: __dirname,
});

/** @type {import("eslint").Linter.FlatConfig[]} */
module.exports = [
  js.configs.recommended,
  ...ts.configs.recommended,
  ...compat.extends("eslint-config-turbo"),
  prettier,
  {
    files: ["*.ts"],
    settings: {
      "import/resolver": {
        typescript: {
          project,
        },
      },
    },
    languageOptions: {
      parser: ts.parser,
    },
  },
  {
    ignores: [".*.js", ".*.ts", "node_modules/", "dist/"],
  },
];
