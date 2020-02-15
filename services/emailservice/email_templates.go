package emailservice

const (
	// resetBaseURL   = "http://localhost:3000/api/update_password" // Change this for production
	welcomeSubject = "Welcome!"
	resetSubject   = "Instructions for resetting your password."
)

const welcomeText = `
	Hi there!

	Welcome! We really hope you enjoy using our application!

	Best,
	Yuichi
`

const welcomeHTML = `
	Hi there!<br/>
	<br/>
	Welcome to <a href="https://www.example.com">Example</a>! We really hope you enjoy using our application!<br/>
	<br/>
	Best,<br/>
	Yuichi
`

const resetTextTmpl = `
	Hi there!

	It appears that you have requested a password reset. If this was you, please follow the link below to update your password:

	%s

	If you are asked for a token, please use the following value:

	%s

	If you didn't request a password reset you can safely ignore this email and your account will not be changed.

	Best,
	Support
`

const resetHTMLTmpl = `
	Hi there!<br/>
	<br/>
	It appears that you have requested a password reset. If this was you, please follow the link below to update your password:<br/>
	<br/>
	<a href="%s">%s</a><br/>
	<br/>
	If you are asked for a token, please use the following value:<br/>
	<br/>
	%s<br/>
	<br/>
	If you didn't request a password reset you can safely ignore this email and your account will not be changed.<br/>
	<br/>
	Best,<br/>
	Support<br/>
`
