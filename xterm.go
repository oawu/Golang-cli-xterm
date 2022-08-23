/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2021
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package xterm

import (
	"fmt"
)

func light(num uint8, lights ...bool) uint8 {
	if len(lights) > 0 && lights[0] {
		return num + 8
	} else {
		return num
	}
}

type Str struct {
	val   string
	codes []string
}

func (str *Str) code(code string) *Str {
	if str == nil {
		return str
	}
	str.codes = append(str.codes, code)
	return str
}

func (str *Str) Bold() *Str                 { return str.code("\x1b[1m") }
func (str *Str) Dim() *Str                  { return str.code("\x1b[2m") }
func (str *Str) Italic() *Str               { return str.code("\x1b[3m") }
func (str *Str) Underline() *Str            { return str.code("\x1b[4m") }
func (str *Str) Blink() *Str                { return str.code("\x1b[5m") }
func (str *Str) Inverted() *Str             { return str.code("\x1b[7m") }
func (str *Str) Hidden() *Str               { return str.code("\x1b[8m") }
func (str *Str) Color(code uint8) *Str      { return str.code(fmt.Sprintf("\x1b[38;5;%dm", code)) }
func (str *Str) Background(code uint8) *Str { return str.code(fmt.Sprintf("\x1b[48;5;%dm", code)) }
func (str *Str) Bg(code uint8) *Str         { return str.Background(code) }

func (str *Str) Black(lights ...bool) *Str  { return str.Color(light(0, lights...)) }
func (str *Str) Red(lights ...bool) *Str    { return str.Color(light(1, lights...)) }
func (str *Str) Green(lights ...bool) *Str  { return str.Color(light(2, lights...)) }
func (str *Str) Yellow(lights ...bool) *Str { return str.Color(light(3, lights...)) }
func (str *Str) Blue(lights ...bool) *Str   { return str.Color(light(4, lights...)) }
func (str *Str) Purple(lights ...bool) *Str { return str.Color(light(5, lights...)) }
func (str *Str) Cyan(lights ...bool) *Str   { return str.Color(light(6, lights...)) }
func (str *Str) Gray(lights ...bool) *Str   { return str.Color(light(7, lights...)) }

func (str *Str) LightBlack() *Str  { return str.Color(light(0, true)) }
func (str *Str) LightRed() *Str    { return str.Color(light(1, true)) }
func (str *Str) LightGreen() *Str  { return str.Color(light(2, true)) }
func (str *Str) LightYellow() *Str { return str.Color(light(3, true)) }
func (str *Str) LightBlue() *Str   { return str.Color(light(4, true)) }
func (str *Str) LightPurple() *Str { return str.Color(light(5, true)) }
func (str *Str) LightCyan() *Str   { return str.Color(light(6, true)) }
func (str *Str) LightGray() *Str   { return str.Color(light(7, true)) }

func (str *Str) BackgroundBlack(lights ...bool) *Str  { return str.Background(light(0, lights...)) }
func (str *Str) BackgroundRed(lights ...bool) *Str    { return str.Background(light(1, lights...)) }
func (str *Str) BackgroundGreen(lights ...bool) *Str  { return str.Background(light(2, lights...)) }
func (str *Str) BackgroundYellow(lights ...bool) *Str { return str.Background(light(3, lights...)) }
func (str *Str) BackgroundBlue(lights ...bool) *Str   { return str.Background(light(4, lights...)) }
func (str *Str) BackgroundPurple(lights ...bool) *Str { return str.Background(light(5, lights...)) }
func (str *Str) BackgroundCyan(lights ...bool) *Str   { return str.Background(light(6, lights...)) }
func (str *Str) BackgroundGray(lights ...bool) *Str   { return str.Background(light(7, lights...)) }

func (str *Str) BackgroundLightBlack() *Str  { return str.Background(light(0, true)) }
func (str *Str) BackgroundLightRed() *Str    { return str.Background(light(1, true)) }
func (str *Str) BackgroundLightGreen() *Str  { return str.Background(light(2, true)) }
func (str *Str) BackgroundLightYellow() *Str { return str.Background(light(3, true)) }
func (str *Str) BackgroundLightBlue() *Str   { return str.Background(light(4, true)) }
func (str *Str) BackgroundLightPurple() *Str { return str.Background(light(5, true)) }
func (str *Str) BackgroundLightCyan() *Str   { return str.Background(light(6, true)) }
func (str *Str) BackgroundLightGray() *Str   { return str.Background(light(7, true)) }

func (str *Str) BgBlack(lights ...bool) *Str  { return str.Background(light(0, lights...)) }
func (str *Str) BgRed(lights ...bool) *Str    { return str.Background(light(1, lights...)) }
func (str *Str) BgGreen(lights ...bool) *Str  { return str.Background(light(2, lights...)) }
func (str *Str) BgYellow(lights ...bool) *Str { return str.Background(light(3, lights...)) }
func (str *Str) BgBlue(lights ...bool) *Str   { return str.Background(light(4, lights...)) }
func (str *Str) BgPurple(lights ...bool) *Str { return str.Background(light(5, lights...)) }
func (str *Str) BgCyan(lights ...bool) *Str   { return str.Background(light(6, lights...)) }
func (str *Str) BgGray(lights ...bool) *Str   { return str.Background(light(7, lights...)) }

func (str *Str) BgLightBlack() *Str  { return str.Background(light(0, true)) }
func (str *Str) BgLightRed() *Str    { return str.Background(light(1, true)) }
func (str *Str) BgLightGreen() *Str  { return str.Background(light(2, true)) }
func (str *Str) BgLightYellow() *Str { return str.Background(light(3, true)) }
func (str *Str) BgLightBlue() *Str   { return str.Background(light(4, true)) }
func (str *Str) BgLightPurple() *Str { return str.Background(light(5, true)) }
func (str *Str) BgLightCyan() *Str   { return str.Background(light(6, true)) }
func (str *Str) BgLightGray() *Str   { return str.Background(light(7, true)) }

func (str *Str) String() string {
	tmp := str.val
	for _, code := range str.codes {
		tmp = fmt.Sprintf("%s%s%s", code, tmp, "\x1b[0m")
	}
	return tmp
}
func New(str string) *Str {
	return &Str{val: str}
}

func Bold(str string) *Str      { return New(str).Bold() }
func Dim(str string) *Str       { return New(str).Dim() }
func Italic(str string) *Str    { return New(str).Italic() }
func Underline(str string) *Str { return New(str).Underline() }
func Blink(str string) *Str     { return New(str).Blink() }
func Inverted(str string) *Str  { return New(str).Inverted() }
func Hidden(str string) *Str    { return New(str).Hidden() }

func Black(str string, lights ...bool) *Str  { return New(str).Black(lights...) }
func Red(str string, lights ...bool) *Str    { return New(str).Red(lights...) }
func Green(str string, lights ...bool) *Str  { return New(str).Green(lights...) }
func Yellow(str string, lights ...bool) *Str { return New(str).Yellow(lights...) }
func Blue(str string, lights ...bool) *Str   { return New(str).Blue(lights...) }
func Purple(str string, lights ...bool) *Str { return New(str).Purple(lights...) }
func Cyan(str string, lights ...bool) *Str   { return New(str).Cyan(lights...) }
func Gray(str string, lights ...bool) *Str   { return New(str).Gray(lights...) }

func BackgroundBlack(str string, lights ...bool) *Str  { return New(str).BackgroundBlack(lights...) }
func BackgroundRed(str string, lights ...bool) *Str    { return New(str).BackgroundRed(lights...) }
func BackgroundGreen(str string, lights ...bool) *Str  { return New(str).BackgroundGreen(lights...) }
func BackgroundYellow(str string, lights ...bool) *Str { return New(str).BackgroundYellow(lights...) }
func BackgroundBlue(str string, lights ...bool) *Str   { return New(str).BackgroundBlue(lights...) }
func BackgroundPurple(str string, lights ...bool) *Str { return New(str).BackgroundPurple(lights...) }
func BackgroundCyan(str string, lights ...bool) *Str   { return New(str).BackgroundCyan(lights...) }
func BackgroundGray(str string, lights ...bool) *Str   { return New(str).BackgroundGray(lights...) }

func BgBlack(str string, lights ...bool) *Str  { return New(str).BackgroundBlack(lights...) }
func BgRed(str string, lights ...bool) *Str    { return New(str).BackgroundRed(lights...) }
func BgGreen(str string, lights ...bool) *Str  { return New(str).BackgroundGreen(lights...) }
func BgYellow(str string, lights ...bool) *Str { return New(str).BackgroundYellow(lights...) }
func BgBlue(str string, lights ...bool) *Str   { return New(str).BackgroundBlue(lights...) }
func BgPurple(str string, lights ...bool) *Str { return New(str).BackgroundPurple(lights...) }
func BgCyan(str string, lights ...bool) *Str   { return New(str).BackgroundCyan(lights...) }
func BgGray(str string, lights ...bool) *Str   { return New(str).BackgroundGray(lights...) }

func LightBlack(str string) *Str  { return New(str).LightBlack() }
func LightRed(str string) *Str    { return New(str).LightRed() }
func LightGreen(str string) *Str  { return New(str).LightGreen() }
func LightYellow(str string) *Str { return New(str).LightYellow() }
func LightBlue(str string) *Str   { return New(str).LightBlue() }
func LightPurple(str string) *Str { return New(str).LightPurple() }
func LightCyan(str string) *Str   { return New(str).LightCyan() }
func LightGray(str string) *Str   { return New(str).LightGray() }

func BackgroundLightBlack(str string) *Str  { return New(str).BackgroundLightBlack() }
func BackgroundLightRed(str string) *Str    { return New(str).BackgroundLightRed() }
func BackgroundLightGreen(str string) *Str  { return New(str).BackgroundLightGreen() }
func BackgroundLightYellow(str string) *Str { return New(str).BackgroundLightYellow() }
func BackgroundLightBlue(str string) *Str   { return New(str).BackgroundLightBlue() }
func BackgroundLightPurple(str string) *Str { return New(str).BackgroundLightPurple() }
func BackgroundLightCyan(str string) *Str   { return New(str).BackgroundLightCyan() }
func BackgroundLightGray(str string) *Str   { return New(str).BackgroundLightGray() }

func BgLightBlack(str string) *Str  { return New(str).BgLightBlack() }
func BgLightRed(str string) *Str    { return New(str).BgLightRed() }
func BgLightGreen(str string) *Str  { return New(str).BgLightGreen() }
func BgLightYellow(str string) *Str { return New(str).BgLightYellow() }
func BgLightBlue(str string) *Str   { return New(str).BgLightBlue() }
func BgLightPurple(str string) *Str { return New(str).BgLightPurple() }
func BgLightCyan(str string) *Str   { return New(str).BgLightCyan() }
func BgLightGray(str string) *Str   { return New(str).BgLightGray() }