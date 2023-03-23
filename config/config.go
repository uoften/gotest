package config

import (
	"errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"gotest/common/utils"
	"os"
	"path/filepath"
)

type Global struct {
	SpiderAsync    bool `yaml:"spider_async"`
	SpiderAllowUrlRevisit bool `yaml:"spider_allow_url_revisit"`
	SpiderLimitRuleDomainGlob    string `yaml:"spider_limit_rule_domain_glob"`
	SpiderLimitRuleParallelism  int `yaml:"spider_limit_rule_parallelism"`
	SpiderLimitRuleRandomDelay int `yaml:"spider_limit_rule_random_delay"`
}

type config struct {
	appRoot      string
	confDir      string
}

var GlobalConfig *Global

func init() {
	conf := &config{}
	appRoot, err := os.Getwd()
	utils.ErrFatal(err)
	confDir := filepath.Join(appRoot, "config")

	conf.appRoot = appRoot
	conf.confDir = confDir

	if FilePathNotExist(confDir) {
		currentPath, err := filepath.Abs("")
		utils.ErrFatal(err)
		utils.Logger.Fatal("无法找到配置目录，在项目目录" + appRoot + "下创建confs文件夹，配置文件可参考" + currentPath + "下的confs_template，并配置相关参数，再启动")
	}
	utils.Logger.Warn("提前配置好所需用的配置文件", zap.String("配置文件所在目录", confDir))
	var confFileNameList []string
	walkFn := func(path string, fileInfo os.FileInfo, err error) error {
		if path == confDir {
			return nil
		}
		confFileNameList = append(confFileNameList, filepath.Base(path))
		if fileInfo.Name() == "config.yaml" {
			err = conf.initConf()
			if err != nil {
				utils.ErrFatal(err)
			}
			utils.Logger.Info("初始化redis配置文件成功")
		}
		return nil
	}
	err = filepath.Walk(confDir, walkFn)
	utils.ErrFatal(err)
	utils.Logger.Info("配置文件", zap.Any("已创建", confFileNameList))
}

func (c *config) initConf() error {
	GlobalConfig = new(Global)
	return ParseYamlToConfStruct(filepath.Join(c.confDir, "config.yaml"), GlobalConfig)
}

func FilePathNotExist(p string) bool {
	if _, err := os.Stat(p); err != nil {
		return true
	}
	return false
}

func ParseYamlToConfStruct(fp string, v interface{}) error {
	_, err := os.Stat(fp)
	if err != nil {
		return errors.New("文件不存在, path:" + fp)
	}

	fd, err := os.Open(fp)
	if err != nil {
		return errors.New("文件打开失败, path:" + fp)
	}
	defer fd.Close()
	err = yaml.NewDecoder(fd).Decode(v)
	if err != nil {
		return errors.New("yaml解析错误, path:" + fp + ", 错误：" + err.Error())
	}
	return nil
}
