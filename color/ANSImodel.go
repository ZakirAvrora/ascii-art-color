package color

func IsColorANSI(color string) bool {
	if color == "black" || color == "blue" || color == "green" || color == "cyan" ||
		color == "red" || color == "purple" || color == "brown" || color == "gray" ||
		color == "darkgray" || color == "lightblue" || color == "lightgreen" || color == "lightcyan" ||
		color == "lightred" || color == "lightpurple" || color == "yellow" || color == "white" || color == "orange" {
		return true
	}
	return false
}

func ANSIColor(color string) string {
	switch color {
	case "black":
		return "\033[0;30m"
	case "blue":
		return "\033[0;34m"
	case "green":
		return "\033[0;32m"
	case "cyan":
		return "\033[0;36m"
	case "red":
		return "\033[0;31m"
	case "purple":
		return "\033[0;35m"
	case "brown":
		return "\033[0;33m"
	case "gray":
		return "\033[0;37m"
	case "darkgray":
		return "\033[1;30m"
	case "lightblue":
		return "\033[1;34m"
	case "lightgreen":
		return "\033[1;32m"
	case "lightcyan":
		return "\033[1;36m"
	case "lightred":
		return "\033[1;31m"
	case "lightpurple":
		return "\033[1;35m"
	case "yellow":
		return "\033[1;33m"
	case "white":
		return "\033[1;37m"
	case "orange":
		return "\033[38;2;255;165;0m"
	default:
		return ""
	}
}
