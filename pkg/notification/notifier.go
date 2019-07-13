package notification

import (
	"github.com/1024casts/1024casts/model"
)

type Notifier struct {
	notifiedUsers []*model.UserModel
}

func NewNotifier() *Notifier {
	return &Notifier{}
}

// parse @test1 @test2... to user link: /users/test1, /users/test2
func (n *Notifier) NewReplyNotify(fromUser model.UserModel) {

	// notify the author
	// Notification::batchNotify('new_reply', $fromUser, $this->removeDuplication([$topic->user]), $topic, $reply);

	// notify attented users
	//Notification::batchNotify('attention', $fromUser, $topic->attentedUsers(), $topic, $reply);

	// notify mentioned users
	// Notification::batchNotify('at', $fromUser, $this->removeDuplication($mentionParser->users), $topic, $reply);
}
