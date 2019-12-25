package rssreader

// errorList describes type for list of errors
type errorList []error

func (el *errorList) Error() string {
	result := ""
	for _, err := range *el {
		if len(result) > 0 {
			result += "\n"
		}
		result += err.Error()
	}

	return result
}

// ErrorOrNil method checks and returns object itself if there are any errors in list. Otherwise returns nil.
func (el *errorList) ErrorOrNil() error {
	for _, err := range *el {
		if err != nil {
			return el
		}
	}

	return nil
}
