package utils

import "log"

func GlobalErrorWithBool(errMsg error) bool {
	if errMsg != nil {
		log.Printf("Error Exception occured : %s", errMsg)
		return false
	}
	return errMsg == nil
}

func GlobalQueryErrorWithBool(errMsg error) bool {
	if errMsg != nil {
		log.Printf("Query Exception occured : %s", errMsg)
		return false
	}
	return errMsg == nil
}
