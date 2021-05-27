module.exports = {
  publicPath: process.env.BASE_URL,
  devServer: {
    proxy: 'http://localhost:8000',
  },
  transpileDependencies: [
    'vuetify',
  ],
}
