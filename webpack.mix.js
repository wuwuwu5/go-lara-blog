let mix = require('laravel-mix');


mix
  .setPublicPath('public/')
  .ts('views/assets/js/app.js', 'public/js')
  .sass('views/assets/sass/app.scss', 'public/css').version();
