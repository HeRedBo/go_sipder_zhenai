package persist

import (
	"GoSpider/ConcurrenceTask/zhenai/model"
	"testing"
)

func TestItemSaver(t *testing.T) {
	cityProfile := model.CityProfile{
		Name:      "雪梨",
		Photo:     "https://photo.zastatic.com/images/photo/25073/100290572/2483971534708801550.jpg",
		Gender:    "女士",
		Place:     "广东广州",
		Age:       "56",
		Height:    "165",
		Income:    "8001-12000元",
		Marriage:  "离异",
		Education: "大专",
		Introduce: "一个暖男家务一起做家庭观念强来一场结婚为目的的恋爱本人有点点内向，社恐，但认识久了你发现这货就是一个 2货哈哈还有最重要的一点，你觉得原生家庭很重要吗？我希望你是一个小话痨，话比我多，我是一个不错的聆听者。100度的水也会有0度的时候，双向奔赴的爱才是最想要的，让我携手同行。",
	}

	Save(cityProfile)


}