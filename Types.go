package main

type JSON map[string]interface{}

// Uses the HSB colorspace
type HSColor struct {
	H, S float64
}

// Uses the CIE colorspace
type XYColor struct {
	X, Y float64
}

var Amber = HSColor{0.125, 1.0}
var AntiqueWhite = HSColor{0.944444, 0.14}
var Aqua = HSColor{0.5, 1.0}
var Ash = HSColor{0.375, 0.6}
var Azure = HSColor{0.5833333, 1.0}
var Black = HSColor{0.0, 0.0}
var Blue = HSColor{0.66666667, 1.0}

/*
var Brick = HSColor{
var Bronze = HSColor{
var Brown = HSColor{
var Crimson = HSColor{
var DarkGold = HSColor{
var DeepPink = HSColor{
var ForestGreen = HSColor{
var Fuschia = HSColor{
var Gold = HSColor{
var Green = HSColor{
var Indigo = HSColor{
var Khaki = HSColor{
var LightGold = HSColor{
var LightPink = HSColor{
var Lime = HSColor{
var Magenta = HSColor{
var Maroon = HSColor{
var NavyBlue = HSColor{
var Orange = HSColor{
var OrangeRed = HSColor{
var Periwinkle = HSColor{
var Pink = HSColor{
var Red = HSColor{
var RedBrown = HSColor{
var SaddleBrown = HSColor{
var SkyBlue = HSColor{
var Turquoise = HSColor{
var Violet = HSColor{
var Wheat = HSColor{
var White = HSColor{
var Yellow = HSColor{
var YellowGreen = HSColor{
*/
