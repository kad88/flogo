package slack_activity



// Activity that sends a message to Slack

// Should take for input the Slack token, channel target, message
// Returns output message sent successfully or not




import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/TIBCOSoftware/flogo-cli/tools/activity"
)

//log is the default package logger
var log = logger.GetLogger("KD_logger")

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

	name := context.GetInput("name").(string)
	salutation := context.GetInput("salutation").(string)

	//use the log object to log the greetings
	log.Debugf("The Flogo engine says [%s] to [%s]", salutation, name)

	// Set the result as part of the context
	context.setOutput("result", "The Flogo engine says "+salutation+" to "+name)
////////////////////////////////////////////////////////////////////////////////////////////////////////////

	api := slack.New("xoxp-2227445904-4843514457-260501703687-6de177cb22ae24b837b9357f5c96822b")


	//Gets user details
	/*user, err := api.GetUserInfo("U04QTF4DF")  //user: kaddour
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
 */

	//Sends message to user
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "TestP",
		Text:    "testT",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage("U04QTF4DF", "Some text", params)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	return true, nil
}

