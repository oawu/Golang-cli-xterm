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

func Black(str string) *Str  { return New(str).Color(0) }
func Red(str string) *Str    { return New(str).Color(1) }
func Green(str string) *Str  { return New(str).Color(2) }
func Yellow(str string) *Str { return New(str).Color(3) }
func Blue(str string) *Str   { return New(str).Color(4) }
func Purple(str string) *Str { return New(str).Color(5) }
func Cyan(str string) *Str   { return New(str).Color(6) }
func Gray(str string) *Str   { return New(str).Color(7) }

func LightBlack(str string) *Str  { return New(str).Color(8) }
func LightRed(str string) *Str    { return New(str).Color(9) }
func LightGreen(str string) *Str  { return New(str).Color(10) }
func LightYellow(str string) *Str { return New(str).Color(11) }
func LightBlue(str string) *Str   { return New(str).Color(12) }
func LightPurple(str string) *Str { return New(str).Color(13) }
func LightCyan(str string) *Str   { return New(str).Color(14) }
func LightGray(str string) *Str   { return New(str).Color(15) }

func BackgroundBlack(str string) *Str  { return New(str).Background(0) }
func BackgroundRed(str string) *Str    { return New(str).Background(1) }
func BackgroundGreen(str string) *Str  { return New(str).Background(2) }
func BackgroundYellow(str string) *Str { return New(str).Background(3) }
func BackgroundBlue(str string) *Str   { return New(str).Background(4) }
func BackgroundPurple(str string) *Str { return New(str).Background(5) }
func BackgroundCyan(str string) *Str   { return New(str).Background(6) }
func BackgroundGray(str string) *Str   { return New(str).Background(7) }

func BgBlack(str string) *Str  { return New(str).Background(0) }
func BgRed(str string) *Str    { return New(str).Background(1) }
func BgGreen(str string) *Str  { return New(str).Background(2) }
func BgYellow(str string) *Str { return New(str).Background(3) }
func BgBlue(str string) *Str   { return New(str).Background(4) }
func BgPurple(str string) *Str { return New(str).Background(5) }
func BgCyan(str string) *Str   { return New(str).Background(6) }
func BgGray(str string) *Str   { return New(str).Background(7) }

func BackgroundLightBlack(str string) *Str  { return New(str).Background(8) }
func BackgroundLightRed(str string) *Str    { return New(str).Background(9) }
func BackgroundLightGreen(str string) *Str  { return New(str).Background(10) }
func BackgroundLightYellow(str string) *Str { return New(str).Background(11) }
func BackgroundLightBlue(str string) *Str   { return New(str).Background(12) }
func BackgroundLightPurple(str string) *Str { return New(str).Background(13) }
func BackgroundLightCyan(str string) *Str   { return New(str).Background(14) }
func BackgroundLightGray(str string) *Str   { return New(str).Background(15) }

func BgLightBlack(str string) *Str  { return New(str).Background(8) }
func BgLightRed(str string) *Str    { return New(str).Background(9) }
func BgLightGreen(str string) *Str  { return New(str).Background(10) }
func BgLightYellow(str string) *Str { return New(str).Background(11) }
func BgLightBlue(str string) *Str   { return New(str).Background(12) }
func BgLightPurple(str string) *Str { return New(str).Background(13) }
func BgLightCyan(str string) *Str   { return New(str).Background(14) }
func BgLightGray(str string) *Str   { return New(str).Background(15) }
