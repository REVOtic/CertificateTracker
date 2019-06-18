// Node Server to render the ÐApp
var express = require('express');
var connect = require('connect');
var http = require('http');

var app = connect().use(express.static(__dirname));
http.createServer(app).listen(9080);
console.log("Access the website on port: 9080");