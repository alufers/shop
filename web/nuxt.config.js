module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: "shop-web",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "Shop project" },
      {
        rel: "stylesheet",
        href:
          "https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Material+Icons"
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: "#3B8070" },
  plugins: ["~/plugins/vuetify.js"],
  css: [
    "vuetify/dist/vuetify.min.css",
    "@mdi/font/css/materialdesignicons.css"
  ],
  /*
  ** Build configuration
  */
  build: {
    vendor: ["~/plugins/vuetify.js"],
    /*
    ** Run ESLint on save
    */
    extend(config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: "pre",
          test: /\.(js|vue)$/,
          loader: "eslint-loader",
          exclude: /(node_modules)/
        });
      }
    }
  },
  router: {
    extendRoutes(routes, resolve) {
      routes.push({
        name: "category-with-parent",
        path: "/admin/categories/:parentCategory",
        component: resolve(__dirname, "pages/admin/categories/index.vue")
      });
    }
  }
};
