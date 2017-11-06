package slack_activity

// Activity that sends a message to Slack

// Should take for input the Slack token, channel target, message
// Returns output message sent successfully or not

import (

/* "github.com/TIBCOSoftware/flogo-lib/logger"
 "github.com/TIBCOSoftware/flogo-lib/core/activity"
 "github.com/nlopes/slack" */

	
)

//log is the default package logger
//var log = logger.GetLogger("KD_logger")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	/*token := context.GetInput("token").(string)
	channel := context.GetInput("channelID").(string)
	message := context.GetInput("message").(string)

	//use the log object to log the greetings
	log.Debugf("token used: [%s]", token)
	log.Debugf("channelId used: [%s]", channel)
	log.Debugf("message sent: [%s]", message)

	// Set the result as part of the context

	//context.setOutput("result", "The Flogo engine says "+salutation+" to "+name)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//api := slack.New("xoxp-2227445904-4843514457-260501703687-6de177cb22ae24b837b9357f5c96822b")
	api:= slack.New(token)
	//Gets user details
	/*user, err := api.GetUserInfo("U04QTF4DF")  //user: kaddour
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	} ///
	log.Debugf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
	//fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)


	//Sends message to user
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "TestP",
		Text:    "testT",
		// Uncomment the following part to send a field too
		///
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		///
	}
	params.Attachments = []slack.Attachment{attachment}
	//channelID, timestamp, err := api.PostMessage("U04QTF4DF", "Some text", params)
	channelID, timestamp, err := api.PostMessage(channel, message, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	//fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	log.Debugf("Message successfully sent to channel %s at %s", channelID, timestamp)  */
	return true, nil
}

