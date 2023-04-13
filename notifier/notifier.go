package notifier

import (
	"github.com/cvelab/notify/notifier/bark"
	"github.com/cvelab/notify/notifier/chanify"
	"github.com/cvelab/notify/notifier/dingtalk"
	"github.com/cvelab/notify/notifier/discord"
	"github.com/cvelab/notify/notifier/feishu"
	"github.com/cvelab/notify/notifier/ftqq"
	"github.com/cvelab/notify/notifier/gitter"
	"github.com/cvelab/notify/notifier/googlechat"
	"github.com/cvelab/notify/notifier/igot"
	"github.com/cvelab/notify/notifier/mailgun"
	"github.com/cvelab/notify/notifier/pushbullet"
	"github.com/cvelab/notify/notifier/pushdeer"
	"github.com/cvelab/notify/notifier/pushover"
	"github.com/cvelab/notify/notifier/qpush"
	"github.com/cvelab/notify/notifier/rocketchat"
	"github.com/cvelab/notify/notifier/showdoc"
	"github.com/cvelab/notify/notifier/slack"
	"github.com/cvelab/notify/notifier/telegram"
	"github.com/cvelab/notify/notifier/webhook"
	"github.com/cvelab/notify/notifier/xz"
	"github.com/cvelab/notify/notifier/zulip"
)

type NotifiesPackage struct {
	Bark       []*bark.Option       `yaml:"bark,omitempty"`
	Chanify    []*chanify.Option    `yaml:"chanify,omitempty"`
	Dingtalk   []*dingtalk.Option   `yaml:"dingtalk,omitempty"`
	Discord    []*discord.Option    `yaml:"discord,omitempty"`
	FeiShu     []*feishu.Option     `yaml:"feishu,omitempty"`
	FTQQ       []*ftqq.Option       `yaml:"ftqq,omitempty"`
	Gitter     []*gitter.Option     `yaml:"gitter,omitempty"`
	GoogleChat []*googlechat.Option `yaml:"googlechat,omitempty"`
	IGot       []*igot.Option       `yaml:"igot,omitempty"`
	Mailgun    []*mailgun.Option    `yaml:"mailgun,omitempty"`
	PushBullet []*pushbullet.Option `yaml:"pushbullet,omitempty"`
	PushDeer   []*pushdeer.Option   `yaml:"pushdeer,omitempty"`
	PushOver   []*pushover.Option   `yaml:"pushover,omitempty"`
	QPush      []*qpush.Option      `yaml:"qpush,omitempty"`
	RocketChat []*rocketchat.Option `yaml:"rocketchat,omitempty"`
	Slack      []*slack.Option      `yaml:"slack,omitempty"`
	ShowDoc    []*showdoc.Option    `yaml:"showdoc,omitempty"`
	Telegram   []*telegram.Option   `yaml:"telegram,omitempty"`
	XZ         []*xz.Option         `yaml:"xz,omitempty"`
	Webhook    []*webhook.Option    `yaml:"webhook,omitempty"`
	ZuLip      []*zulip.Option      `yaml:"zulip,omitempty"`

	// 计划中
	// Sendgraid  []*sendgraid.Option  `yaml:"sendgraid,omitempty"`
	// Mattermost []*mattermost.Option `yaml:"mattermost,omitempty" //https://api.mattermost.com/
}
