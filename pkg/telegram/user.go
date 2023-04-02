package telegram

type User struct {
	age           int
	gender        string
	height        int
	weight        int
	activityLevel string
}

func displayStart() string {
	message := `
	You can control me by sending these commands:

	/start - start a bot
	/membership - view membership prices
	/trainings - see training plans and prices
`
	return message
}
