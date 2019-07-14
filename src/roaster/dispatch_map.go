package roaster

type Patcher interface {
	ParseParameters(map[string]string) error
    Run() error
}

var dispatch = map[string]interface{Patcher} {
	"text_change" : new(TextChange),
	"stats_change" : new(StatsChange),
    "menu_image_change" : new(MenuImageChange),
}
