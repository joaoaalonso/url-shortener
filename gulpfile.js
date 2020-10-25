'use strict';

var fs = require('fs');
var del = require('del');
var gulp = require('gulp');
var csso = require('gulp-csso');
var sass = require('gulp-sass');
var hash = require('gulp-hash');
var babel = require('gulp-babel');
var uglify = require('gulp-uglify');
var htmlmin = require('gulp-htmlmin');
var replace = require('gulp-replace');

gulp.task('clean', () => del(['public']));

gulp.task('styles', function () {
    return gulp.src('./resources/assets/scss/*.scss')
        .pipe(sass({
            outputStyle: 'nested',
            precision: 10,
            includePaths: ['.'],
            onError: console.error.bind(console, 'Sass error:')
        }))
        .pipe(csso())
        .pipe(hash())
        .pipe(gulp.dest('./public/assets/css'))
        .pipe(hash.manifest('assets.json'))
        .pipe(gulp.dest('./public/assets'));
  });

gulp.task('scripts', function() {
    return gulp.src('./resources/assets/js/*.js')
        .pipe(babel({
            presets: ['@babel/env']
        }))
        .pipe(uglify())
        .pipe(hash())
        .pipe(gulp.dest('./public/assets/js'))
        .pipe(hash.manifest('assets.json'))
        .pipe(gulp.dest('./public/assets'));
});

gulp.task('images', function() {
    return gulp.src(['./resources/assets/images/*'])
        .pipe(hash())
        .pipe(gulp.dest('./public/assets/images'))
        .pipe(hash.manifest('assets.json'))
        .pipe(gulp.dest('./public/assets'));
});

gulp.task('views', function() {
    var stream = gulp.src(['./resources/views/*.html'])
        .pipe(htmlmin({
            collapseWhitespace: true,
            removeComments: true
        }))
    
    var assetsPath = './public/assets/assets.json'
    if(fs.existsSync(assetsPath)) {
        var assetsMap = JSON.parse(fs.readFileSync(assetsPath))
        Object.entries(assetsMap).forEach(function(asset) {
            stream.pipe(replace(asset[0], asset[1]))
        })
    }

    stream.pipe(gulp.dest('./public/views'));

    return stream;
});

gulp.task('default', gulp.series(
        'clean',
        'styles',
        'scripts',
        'images',
        'views'
    )
);
