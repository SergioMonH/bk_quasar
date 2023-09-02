package util

import "strings"

// func trackPhone(x1, y1, r1, x2, y2, r2, x3, y3, r3 float32) (float32, float32) {
// 	A := 2*x2 - 2*x1
// 	B := 2*y2 - 2*y1
// 	C := r1*r1 - r2*r2 - x1*x1 + x2*x2 - y1*y1 + y2*y2
// 	D := 2*x3 - 2*x2
// 	E := 2*y3 - 2*y2
// 	F := r2*r2 - r3*r3 - x2*x2 + x3*x3 - y2*y2 + y3*y3
// 	x := (C*E - F*B) / (E*A - B*D)
// 	y := (C*D - A*F) / (B*D - A*E)
// 	return x, y
// }

func CalculateCoefficients(x1, y1, r1, x2, y2, r2, x3, y3, r3 float32) (float32, float32, float32, float32, float32, float32) {
	A := 2 * (x2 - x1)
	B := 2 * (y2 - y1)
	C := r1*r1 - r2*r2 - x1*x1 + x2*x2 - y1*y1 + y2*y2
	D := 2 * (x3 - x2)
	E := 2 * (y3 - y2)
	F := r2*r2 - r3*r3 - x2*x2 + x3*x3 - y2*y2 + y3*y3
	return A, B, C, D, E, F
}

func Trilateration(x1, y1, r1, x2, y2, r2, x3, y3, r3 float32) (float32, float32) {
	A, B, C, D, E, F := CalculateCoefficients(x1, y1, r1, x2, y2, r2, x3, y3, r3)
	x := (C*E - F*B) / (E*A - B*D)
	y := (C*D - A*F) / (B*D - A*E)
	return x, y
}

func AppendMessage(previousMessage, newMessage []string) []string {
	for i := 0; i < len(previousMessage); i++ {
		if previousMessage[i] == "" {
			previousMessage[i] = newMessage[i]
		}
	}
	return previousMessage
}

func DiscoverMessage(messages ...[]string) string {
	if len(messages) < 1 || len(messages[0]) == 0 {
		return ""
	}

	result := make([]string, len(messages[0]))

	for i := 0; i < len(messages[0]); i++ {
		var selectedMessage string

		for j := 0; j < len(messages); j++ {
			if len(messages[j]) <= i || messages[j][i] == "" {
				continue
			}
			selectedMessage = messages[j][i]
			break
		}

		result[i] = selectedMessage
	}

	return strings.Join(result, " ")
}
