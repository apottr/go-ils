const path = require('path')

module.exports = {
    entry: "./src/main.js",
    output: {
        path: path.resolve(__dirname,'dist'),
        filename: "bundle"
    },
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                exclude: /(node_modules)/,
                use: "babel-loader"
            },
            {
                test: /\.css?$/,
                exclude: /(node_modules)/,
                use: ["style-loader","css-loader"]
            }
        ]
    }
}