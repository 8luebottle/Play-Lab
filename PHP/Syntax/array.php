<?php

# ARRAY
/*
 * There are 3 types of arrays
 *  1. Indexed arrays
 *  2. Associative arrays
 *  3. Multidimensional arrays
 */

## create an array
$numbs = [1, 2, 3]; // short syntax (version >= php 5.4)
var_dump($numbs);

$years = array(2021, 2022, 2023); // (version < php 5.4)
var_dump($years);

$empty = array(); // create an empty array
var_dump($empty);
/*
    array(0) {
    }
 */

$no_numbs = []; // create an empty array by short syntax
var_dump($no_numbs);
/*
    array(0) {
    }
 */

## length of array
### count($array_var);
echo count($years)."\n";    // 3
echo count($empty)."\n";    // 0
echo count($no_numbs)."\n"; // 0

## get element from an array
### $array_var[index];
echo $numbs[1]."\n"; // 2
// echo $numbs[3];   // index out of range. (PHP Warning:  Undefined array key 3 in path:line)
//echo $numbs[-1];   // unlike python there is no -1 syntax.

### use loop
for ($i = 0; $i < count($years); $i++) {
    echo $years[$i]."\n";
}
/*
    2021
    2022
    2023
 */

## push to an array
### array_push($array_var, $var);
$news = ['north', 'east', 'west', 'south'];
array_push($news, 'north-east');
var_dump($news);
/*
    array(5) {
      [0]=>
      string(5) "north"
      [1]=>
      string(4) "east"
      [2]=>
      string(4) "west"
      [3]=>
      string(5) "south"
      [4]=>
      string(10) "north-east"
    }
 */

## pop from an array
### array_pop($array_var);
array_pop($news);
var_dump($news);
/*
    array(4) {
      [0]=>
      string(5) "north"
      [1]=>
      string(4) "east"
      [2]=>
      string(4) "west"
      [3]=>
      string(5) "south"
    }
 */

## shift
### array_shift($array_var);
$code = ['dirty', 'code', 'smell'];
array_shift($code);
var_dump($code);
/*
    array(2) {
      [0]=>
      string(4) "code"
      [1]=>
      string(5) "smell"
    }
 */

## unshift
### array_unshift($array_var);
array_unshift($code, 'good');
var_dump($code);
/*
    array(3) {
      [0]=>
      string(4) "good"
      [1]=>
      string(4) "code"
      [2]=>
      string(5) "smell"
    }
 */

## add multiple elements
### $array_var = array_merge($array_var, $array_var2)
$languages = ['php', 'javascript'];
$prefer_languages = ['go', 'python'];
array_merge($languages, $prefer_languages);
var_dump($languages);
/*
    array(2) {
      [0]=>
      string(3) "php"
      [1]=>
      string(10) "javascript"
    }
 */

$languages = array_merge($languages, $prefer_languages);
var_dump($languages);
/*
    array(4) {
      [0]=>
      string(3) "php"
      [1]=>
      string(10) "javascript"
      [2]=>
      string(2) "go"
      [3]=>
      string(6) "python"
    }
 */

## add element to specific position
### array_splice($array_var, offset, length $var);
$alpha = ['a', 'c', 'd'];
array_splice($alpha, 1, 0,'b');
var_dump($alpha);
/*
    array(4) {
      [0]=>
      string(1) "a"
      [1]=>
      string(1) "b"
      [2]=>
      string(1) "c"
      [3]=>
      string(1) "d"
    }
 */