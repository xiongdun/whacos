const path = require('path')

module.exports = {
  configureWebpack: {
    resolve: {
      alias: {
        api: path.resolve(__dirname, 'src/api/'),
        common: path.resolve(__dirname, 'src/common/'),
        components: path.resolve(__dirname, 'src/components/'),
        assets: path.resolve(__dirname, 'src/assets/'),
        views: path.resolve(__dirname, 'src/views/'),
        style: path.resolve(__dirname, 'src/style/'),
      }
    }
  },
  // devServer: {
  // 	proxy: {
  // 		'/api': {
  // 			target: 'https://c.y.qq.com',
  // 			changeOrigin: true,
  // 			pathRewrite: {
  // 				'^/api': ''
  // 			}
  // 		}
  // 	}
  // }
}
