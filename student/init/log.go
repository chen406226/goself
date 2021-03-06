package init

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/op/go-logging"
	"io"
	"os"
	"strings"
	"student/config"
	"student/global"
	"student/utils"
	"time"
)

const (
	logDir		= "log"
	logSoftLink	= "latest_log"
	module		= "LOG_MODULE"
)

var (
	defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
)

func init() {
	fmt.Errorf("logger prefix not found")
	c := global.GL_CONFIG.Log
	fmt.Errorf("logger preffdssfdix not found")
	if c.Prefix == "" {
		fmt.Errorf("logger prefix not found")
	}
	logger := logging.MustGetLogger(module)
	var backends []logging.Backend
	backends = registerStdout(c,backends)
	backends = registerFile(c,backends)
	logging.SetBackend(backends...)
	global.GL_LOG = logger
}

func registerFile(c config.Log, backends []logging.Backend) []logging.Backend {
	if c.File != "" {
		if ok, _ := utils.PathExists(logDir); !ok {
			// directory not exist
			fmt.Println("create log directory")
			_ = os.Mkdir(logDir, os.ModePerm)
		}
		fileWriter, err := rotatelogs.New(
			logDir+string(os.PathSeparator)+"%Y-%m-%d-%H-%M.log",
			// generate soft link, point to latest log file
			rotatelogs.WithLinkName(logSoftLink),
			// maximum time to save log files
			rotatelogs.WithMaxAge(7*24*time.Hour),
			// time period of log file switching
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Println(err)
			return backends
		}
		level, err := logging.LogLevel(c.File)
		if err != nil {
			fmt.Println(err)
		}
		backends = append(backends, createBackend(fileWriter, c, level))
	}

	return backends
}

func registerStdout(c config.Log, backends []logging.Backend)[]logging.Backend  {
	if c.Stdout != "" {
		level,err := logging.LogLevel(c.Stdout)
		if err != nil {
			fmt.Println("log_init_err:37",err)
		}
		backends = append(backends,createBackend(os.Stdout,c,level))
	}
	return backends
}


func createBackend(w io.Writer, c config.Log, level logging.Level) logging.Backend {
	backend := logging.NewLogBackend(w, c.Prefix, 0)
	stdoutWriter := false
	if w == os.Stdout {
		stdoutWriter = true
	}
	format := getLogFormatter(c, stdoutWriter)
	backendLeveled := logging.AddModuleLevel(logging.NewBackendFormatter(backend, format))
	backendLeveled.SetLevel(level, module)
	return backendLeveled
}

func getLogFormatter(c config.Log, stdoutWriter bool) logging.Formatter {
	pattern := defaultFormatter
	if !stdoutWriter {
		// Color is only required for console output
		// Other writers don't need %{color} tag
		pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
		pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	}
	if !c.LogFile {
		// Remove %{logfile} tag
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	return logging.MustStringFormatter(pattern)
}
