'use strict';

var del = require('del');
var gulp = require('gulp');
var csso = require('gulp-csso');
var sass = require('gulp-sass');
var babel = require('gulp-babel');
var uglify = require('gulp-uglify');
var htmlmin = require('gulp-htmlmin');

gulp.task('styles', function () {
    return gulp.src('./resources/assets/scss/*.scss')
        .pipe(sass({
            outputStyle: 'nested',
            precision: 10,
            includePaths: ['.'],
            onError: console.error.bind(console, 'Sass error:')
        }))
        .pipe(csso())
        .pipe(gulp.dest('./public/assets/css'))
  });

gulp.task('scripts', function() {
    return gulp.src('./resources/assets/js/*.js')
        .pipe(babel({
            presets: ['@babel/env']
        }))
        .pipe(uglify())
        .pipe(gulp.dest('./public/assets/js'))
});

gulp.task('views', function() {
    return gulp.src(['./resources/views/*.html'])
        .pipe(htmlmin({
            collapseWhitespace: true,
            removeComments: true
        }))
        .pipe(gulp.dest('./public/views'));
});

gulp.task('images', function() {
    return gulp.src(['./resources/assets/images/*'])
        .pipe(gulp.dest('./public/assets/images'));
});

gulp.task('clean', () => del(['public']));

gulp.task('default', gulp.series(
        'clean',
        'styles',
        'scripts',
        'views',
        'images'
    )
);
