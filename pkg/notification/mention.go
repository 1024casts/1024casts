package notification

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"
	"github.com/lexkong/log"
)

type Mention struct {
	users         []*model.UserModel
	userNames     []string
	originContent string
	parsedContent string
}

func NewMention() *Mention {
	return &Mention{}
}

// parse @test1 @test2... to user link: /users/test1, /users/test2
func (m *Mention) Parse(content string) string {
	m.originContent = content
	m.getMentionedUsername()

	// multi get users by username
	if len(m.userNames) > 0 {
		userRepo := repository.NewUserRepo()
		users, err := userRepo.GetUserByUserNames(m.userNames)
		if err != nil {
			log.Warnf("[mention] get user names err: %v", err)
			return m.originContent
		}
		m.users = users
	}

	m.replace()

	return m.parsedContent
}

func (m *Mention) replace() {
	m.parsedContent = m.originContent

	for _, user := range m.users {
		search := "@" + user.Username
		replace := fmt.Sprintf("[%s](/users/%s)", user.Username, user.Username)
		log.Infof("[notification] replace, search: %s", search)
		log.Infof("[notification] replace, replace: %s", replace)
		m.parsedContent = strings.ReplaceAll(m.parsedContent, search, replace)
	}
}

// get @test1, @test2... from originContent
func (m *Mention) getMentionedUsername() {
	reg := regexp.MustCompile(`(\S*)\@([^\r\n\s]*)`)
	regRet := reg.FindAll([]byte(m.originContent), 100)

	for key, userName := range regRet {
		log.Infof("[mention] reg find key: %d, ret: %+v", key, string(userName))
		m.userNames = append(m.userNames, strings.ReplaceAll(string(userName), "@", ""))
	}
}
