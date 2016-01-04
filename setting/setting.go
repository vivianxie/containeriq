package setting

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

var (
	conf config.ConfigContainer
)

var (
	//Global
	AppName    string
	Usage      string
	Version    string
	Author     string
	Email      string
	RunMode    string
	ListenMode string
	LogPath    string
	DBURI      string
	DBUser     string
	DBPasswd   string
	DBSchema   string
)

func SetConfig(path string) error {
	var err error

	conf, err = config.NewConfig("ini", path)

	if err != nil {
		return fmt.Errorf("Read %s error: %v", path, err.Error())
	}

	if appname := conf.String("appname"); appname != "" {
		AppName = appname
	} else if appname == "" {
		err = fmt.Errorf("AppName config value is null")
	}

	if usage := conf.String("usage"); usage != "" {
		Usage = usage
	} else if usage == "" {
		err = fmt.Errorf("Usage config value is null")
	}

	if version := conf.String("version"); version != "" {
		Version = version
	} else if version == "" {
		err = fmt.Errorf("Version config value is null")
	}

	if author := conf.String("author"); author != "" {
		Author = author
	} else if author == "" {
		err = fmt.Errorf("Author config value is null")
	}

	if email := conf.String("email"); email != "" {
		Email = email
	} else if email == "" {
		err = fmt.Errorf("Email config value is null")
	}

	if runmode := conf.String("runmode"); runmode != "" {
		RunMode = runmode
	} else if runmode == "" {
		err = fmt.Errorf("RunMode config value is null")
	}

	if listenmode := conf.String("listenmode"); listenmode != "" {
		ListenMode = listenmode
	} else if listenmode == "" {
		err = fmt.Errorf("ListenMode config value is null")
	}

	if logpath := conf.String("log::filepath"); logpath != "" {
		LogPath = logpath
	} else if logpath == "" {
		err = fmt.Errorf("LogPath config value is null")
	}

	if dburi := conf.String("db::uri"); dburi != "" {
		DBURI = dburi
	} else if dburi == "" {
		err = fmt.Errorf("DBURI config value is null")
	}

	if dbuser := conf.String("db::user"); dbuser != "" {
		DBUser = dbuser
	} else if dbuser == "" {
		err = fmt.Errorf("DBUser config value is null")
	}

	if dbpass := conf.String("db::passwd"); dbpass != "" {
		DBPasswd = dbpass
	} else if dbpass == "" {
		err = fmt.Errorf("DBPasswd config value is null")
	}

	if dbschema := conf.String("db::schema"); dbschema != "" {
		DBSchema = dbschema
	} else if dbschema == "" {
		err = fmt.Errorf("DBSchema config value is null")
	}
	return nil
}
