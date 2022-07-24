package texture

import (
	"fmt"
	"strings"
)

type Car []string
type Train []Car

const ()

var (
	hankyu_window_top    = "\u001b[38;5;81m\u001b[48;5;88m▄▄\u001b[00m"
	hankyu_window_middle = "\u001b[38;5;88m\u001b[48;5;81m▄▄\u001b[00m"
	// hankyu_window_middle    = "\u001b[38;5;81m██\u001b[00m"
	hankyu_window_buttom       = "\u001b[38;5;88m██\u001b[00m"
	hankyu_window_small_top    = "\u001b[38;5;81m\u001b[48;5;88m▄\u001b[00m"
	hankyu_window_small_middle = "\u001b[38;5;88m\u001b[48;5;81m▄\u001b[00m"
	// hankyu_window_small_middle = "\u001b[38;5;81m█\u001b[00m"
	hankyu_window_small_buttom = "\u001b[38;5;88m█\u001b[00m"
	// hankyu_door_pixel   = "\u001b[38;5;52m█\u001b[00m"
	hankyu_door_top = "\u001b[38;5;52m█\u001b[00m" +
		"\u001b[38;5;81m\u001b[48;5;52m▄\u001b[00m" +
		"\u001b[38;5;52m█\u001b[00m" +
		"\u001b[38;5;81m\u001b[48;5;52m▄\u001b[00m" +
		"\u001b[38;5;52m█\u001b[00m"
	hankyu_door_middle = "\u001b[38;5;52m█\u001b[00m" +
		"\u001b[38;5;52m\u001b[48;5;81m▄\u001b[00m" +
		// "\u001b[38;5;81m█\u001b[00m" +
		"\u001b[38;5;52m█\u001b[00m" +
		"\u001b[38;5;52m\u001b[48;5;81m▄\u001b[00m" +
		// "\u001b[38;5;81m█\u001b[00m" +
		"\u001b[38;5;52m█\u001b[00m"
	hankyu_door_buttom       = "\u001b[38;5;52m█████\u001b[00m"
	hankyu_door_small_top    = "\u001b[38;5;52m█\u001b[00m" + "\u001b[38;5;81m\u001b[48;5;52m▄\u001b[00m" + "\u001b[38;5;52m█\u001b[00m"
	hankyu_door_small_middle = "\u001b[38;5;52m█\u001b[00m" + "\u001b[38;5;52m\u001b[48;5;81m▄\u001b[00m" + "\u001b[38;5;52m█\u001b[00m"
	hankyu_door_small_buttom = "\u001b[38;5;52m███\u001b[00m"
	hankyu_body              = "\u001b[38;5;88m█\u001b[00m"
	hankyu_aircon            = "\u001b[38;5;251m\u001b[48;5;242m▄▄▄▄▄▄▄\u001b[00m"
	Hankyu_front             = Car{
		"   ◇                                             ◇   ",
		"\u001b[38;5;251m▄▄▄\u001b[00m" + "\u001b[38;5;251m\u001b[48;5;242m▄▄\u001b[00m" + "\u001b[38;5;251m▄▄▄\u001b[00m" + hankyu_aircon +
			"\u001b[38;5;251m▄▄▄▄▄▄▄▄▄\u001b[00m" + hankyu_aircon +
			"\u001b[38;5;251m▄▄▄▄▄▄▄▄▄\u001b[00m" + hankyu_aircon +
			"\u001b[38;5;251m▄▄▄\u001b[00m" + "\u001b[38;5;251m\u001b[48;5;242m▄▄\u001b[00m" + "\u001b[38;5;251m▄▄▄\u001b[00m",
		hankyu_body + hankyu_door_small_top + hankyu_body + hankyu_window_small_top + hankyu_body + hankyu_body +
			strings.Repeat(hankyu_door_top+hankyu_body+strings.Repeat(hankyu_body+hankyu_window_top, 3)+hankyu_body+hankyu_body, 2) +
			hankyu_door_top + hankyu_body + hankyu_body + strings.Repeat(hankyu_window_top+hankyu_body, 2),
		hankyu_body + hankyu_door_small_middle + hankyu_body + hankyu_window_small_middle + hankyu_body + hankyu_body +
			strings.Repeat(hankyu_door_middle+hankyu_body+strings.Repeat(hankyu_body+hankyu_window_middle, 3)+hankyu_body+hankyu_body, 2) +
			hankyu_door_middle + hankyu_body + hankyu_body + strings.Repeat(hankyu_window_middle+hankyu_body, 2),
		hankyu_body + hankyu_door_small_buttom + hankyu_body + hankyu_window_small_buttom + hankyu_body + hankyu_body +
			strings.Repeat(hankyu_door_buttom+hankyu_body+strings.Repeat(hankyu_body+hankyu_window_buttom, 3)+hankyu_body+hankyu_body, 2) +
			hankyu_door_buttom + hankyu_body + hankyu_body + strings.Repeat(hankyu_window_buttom+hankyu_body, 2),
		"\u001b[38;5;242m/   o▀▀▀▀o   ▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀▀▀▀   o▀▀▀▀o    \u001b[00m",
	}
	Hankyu_rear = Car{
		Hankyu_front[0],
		Hankyu_front[1],

		hankyu_body + strings.Repeat(hankyu_window_top+hankyu_body, 2) + hankyu_body +
			strings.Repeat(hankyu_door_top+hankyu_body+strings.Repeat(hankyu_body+hankyu_window_top, 3)+hankyu_body+hankyu_body, 2) +
			hankyu_door_top + hankyu_body + hankyu_body + hankyu_window_small_top + hankyu_body + hankyu_door_small_top + hankyu_body,
		hankyu_body + strings.Repeat(hankyu_window_middle+hankyu_body, 2) + hankyu_body +
			strings.Repeat(hankyu_door_middle+hankyu_body+strings.Repeat(hankyu_body+hankyu_window_middle, 3)+hankyu_body+hankyu_body, 2) +
			hankyu_door_middle + hankyu_body + hankyu_body + hankyu_window_small_middle + hankyu_body + hankyu_door_small_middle + hankyu_body,
		hankyu_body + strings.Repeat(hankyu_window_buttom+hankyu_body, 2) + hankyu_body +
			strings.Repeat(hankyu_door_buttom+hankyu_body+strings.Repeat(hankyu_body+hankyu_window_buttom, 3)+hankyu_body+hankyu_body, 2) +
			hankyu_door_buttom + hankyu_body + hankyu_body + hankyu_window_small_buttom + hankyu_body + hankyu_door_small_buttom + hankyu_body,
		"\u001b[38;5;242m    o▀▀▀▀o   ▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀▀▀▀   o▀▀▀▀o   \\\u001b[00m",
	}
	Hankyu = Train{
		Hankyu_front, Hankyu_rear,
	}
)

func (train *Train) Render() {
	for i := 0; i < 6; i++ {
		fmt.Println((*train)[0][i] + "   " + (*train)[1][i])
	}
}
