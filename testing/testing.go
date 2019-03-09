package testing

var isTestModeOn bool = false

func SetTestMode(action bool) {
	isTestModeOn = action
}

func IsTestModeOn() (isTestModeOnParam bool) {
	return isTestModeOn
}
