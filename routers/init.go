package routers

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/howeyc/fsnotify" //File system notifications for Go

	"github.com/astaxie/beego"
	"github.com/beego/compress"
	"github.com/beego/i18n"
)

var (
	CompressConfPath = "conf/compress.json"
)

func initLocales() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

func settingCompress() {
	setting, err := compress.LoadJsonConf(CompressConfPath, IsPro, "/")
	if err != nil {
		beego.Error(err)
		return
	}

	setting.RunCommand()

	if IsPro {
		setting.RunCompress(true, false, true)
	}

	beego.AddFuncMap("compress_js", setting.Js.CompressJs)
	beego.AddFuncMap("compress_css", setting.Css.CompressCss)
}

func dict(values ...interface{}) (map[string]interface{}, error) {

	//beego.Info("values len:", len(values))
	//for i,v := range values {
	//	beego.Info(i, fmt.Sprintf("%+v", v))
	//}

	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	//beego.Info(fmt.Sprintf("dict:%+v", dict))
	return dict, nil
}

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func initTemplates() {
	beego.AddFuncMap("dict", dict)
	beego.AddFuncMap("loadtimes", loadtimes)
}

func InitApp() {
	//在模版中注册函数
	initTemplates()
	// 选择语言
	initLocales()
	// 压缩js css
	settingCompress()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("Failed start app watcher: " + err.Error())
	}

	// ini json变更动态更改
	go func() {
		for {
			select {
			case event := <-watcher.Event:
				switch filepath.Ext(event.Name) {
				case ".ini":
					beego.Info("event:", event)

					if err := i18n.ReloadLangs(); err != nil {
						beego.Error("Conf Reload: ", err)
					}

					beego.Info("Config Reloaded")

				case ".json":
					if event.Name == CompressConfPath {
						settingCompress()
						beego.Info("Beego Compress Reloaded")
					}
				}
			}
		}
	}()

	if err := watcher.WatchFlags("conf", fsnotify.FSN_MODIFY); err != nil {
		beego.Error(err)
	}
}
