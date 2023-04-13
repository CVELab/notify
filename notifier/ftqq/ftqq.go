package ftqq

import (
	"fmt"
	"github.com/cvelab/notify/types"
	"github.com/cvelab/notify/utils"
	"github.com/cvelab/requests"
	"github.com/cvelab/requests/ext"
)

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Title       string  `yaml:"title" dict:"title"`
	Description *string `yaml:"desp,omitempty" dict:"desp,omitempty"`
	Channel     *string `yaml:"channel,omitempty" dict:"channel,omitempty"`
	OpenId      *string `yaml:"openId,omitempty" dict:"openId,omitempty"`
}

type notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *notifier {
	noticer := &notifier{}
	noticer.Option = opt
	return noticer
}

func (n *notifier) format(messages []string) (string, ext.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	data := utils.StructToDict(n.MessageParams)
	return n.Webhook, ext.Data(data)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok {
		return nil
	}
	return fmt.Errorf("[FTQQ] [%v] %s", resp.StatusCode, resp.Content)
}
