package server

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"

	pb "github.com/Liberxue/gateway_auth/protocol/proto"
)

var searchTags = map[string]string{
	"Allow me":                         "让我来",
	"Hold on":                          "请稍等",
	"Whatever":                         "随你的便!",
	"After you":                        "你先请。",
	"Nonsense":                         "胡说八道！",
	"It depends":                       "看情形。",
	"Who Cares":                        "谁管你呀！",
	"Calm down":                        "冷静一点。",
	"It's on me":                       "我来付。",
	"I mean it":                        "我是说真的",
	"It's a deal":                      "一言为定",
	"It won't work":                    "行不通",
	"What's up?":                       "近来过得如何？",
	"I'll get it":                      "我来接电话",
	"Let's face it":                    "面对现实吧",
	"That depends":                     "看情况再说",
	"It takes time":                    "这需要时间",
	"It depends on you":                "取决于你",
	"Out of the question":              "不可能的!",
	"I've done my best":                "我已尽力了",
	"It slipped mymind":                "我不留神忘了。",
	"It will cometo me":                "我会想起来的",
	"Don't be so modest":               "别这么谦虚",
	"Don't give up":                    "别放弃",
	"I apologize":                      "我很抱歉。",
	"I assure you":                     "我向你保证。",
	"Don't bother":                     "不用麻烦了。",
	"I can't afford it":                "我买不起。",
	"Don't get me wrong":               "别误会我。",
	"I bet you can":                    "我确信你能做到。",
	"I can manage":                     "我自己可以应付。",
	"It won't happen again":            "下不为例。",
	"Mind your own business":           "别多管闲事。",
	"It's nice meeting you":            "很高兴认识你。",
	"Still up":                         "还没睡呀？",
	"Who knows":                        "天晓得！",
	"It is urgent":                     "有急事。",
	"Easy does it":                     "慢慢来。",
	"God works":                        "上帝的安排。",
	"Take your time":                   "慢慢吃。 ",
	"Don't push me":                    "别逼我。 ",
	"Come on":                          "快点，振作起来！  ",
	"What is the fuss":                 "吵什么？",
	"Have a good of it":                "玩得很高兴。",
	"Don't let me down":                "别让我失望。",
	"Never heard of it":                "没听说过！",
	"No doubt about it":                "勿庸置疑！",
	"same as usual":                    "一如既往！",
	"walls have ears":                  "隔墙有耳！",
	"There you go again":               "你又来了！",
	"Time is running out":              "没有时间了！",
	"A thousand times no":              "绝对办不到！ ",
	"Don't mention it":                 "没关系，别客气。",
	"It is not a big deal":             "没什么了不起！",
	"Please lay them straight":         "请把他们放好。",
	"It doesn't make any differences":  "没关系。 ",
	"Please wash your hands with soap": "打打肥皂。 ",
	"No way":                           "不行!",
	"Cheer up":                         "振作点!",
	"Go for it":                        "加油！",
	"I don't have a clue":              "我不知道。",
	"You've gone too far":              "你太过分了",
	"cut the crap":                     "少废话，别兜圈子。",
	"Cut it out":                       "省省吧，少来这套。",
	" don't lie to me":                 "别骗我",
	"take your time":                   "慢慢来",
	"anything is possible":             "一切皆有可能",
	"as you wish":                      "如你所愿",
}

func (h *GateWayServer) SearchTag(ctx context.Context, r *pb.SearchTagRequest) (*pb.SearchTagResponse, error) {
	rand.Seed(time.Now().UnixNano())
	tags := make([]*pb.Tags, 0)
	for k, v := range searchTags {
		if len(tags) < 6 {
			tags = append(tags, &pb.Tags{
				TagKey:    k,
				TagValue:  v,
				TagImages: fmt.Sprintf("%d.png", rand.Intn(8)),
			})
		}
		continue
	}
	return &pb.SearchTagResponse{
		Code:    pb.ResponseCode_SUCCESSFUL,
		Message: pb.ResponseCode_SUCCESSFUL.String(),
		Tags:    tags,
	}, nil
}
