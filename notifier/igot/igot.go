package igot

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
	Content string          `yaml:"content" json:"content"`
	Title   *string         `yaml:"title,omitempty" json:"title,omitempty"`
	Url     *string         `yaml:"url,omitempty" json:"url,omitempty"`
	Detail  *map[string]any `yaml:"detail,omitempty" json:"detail,omitempty"`
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
	utils.FormatAnyWithMap(&n.Webhook, &formatMap)
	json := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(json)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok && resp.Json().Get("ret").Int() == 0 {
		return nil
	}
	return fmt.Errorf("[iGot] [%v] %s", resp.StatusCode, resp.Content)
}
