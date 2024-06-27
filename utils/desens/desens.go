package desens

import "strings"

func DesensitizationEmail(email string) string {

	elist := strings.Split(email, "@")
	if len(elist) != 2 {
		return ""
	}

	return elist[0][:1] + "******@" + elist[1]
}