package error_judge

import "github.com/gogf/gf/text/gstr"

func IsDatabaseUnique(error error) bool {
	if gstr.Contains(error.Error(), "Error 1062: Duplicate entry") {
		return true
	}
	return false
}
