module.exports = {
  mount: {
    public: { url: "/", static: true },
    src: { url: "/dist" },
  },
  plugins: ["@snowpack/plugin-typescript"],
  packageOptions: {
    /* ... */
  },
  devOptions: {
    port: 4090,
  },
  buildOptions: {},
  optimize: {
    bundle: true,
    splitting: true,
    treeshake: true,
    minify: true,
    target: "es6",
  },
};
