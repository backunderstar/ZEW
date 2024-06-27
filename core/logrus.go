package core

import (
	"bytes"
	"fmt"
	"github.com/backunderstar/zew/global"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

// 定义日志级别颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// LogFormatter 是用于格式化日志的结构体
type LogFormatter struct{}

// Format 格式化日志条目
// 它根据日志级别选择不同的颜色，并根据是否包含调用者信息来格式化输出
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据日志级别选择颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	// 使用entry提供的缓冲区或创建新的缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 获取日志前缀和时间戳
	logPrefix := global.Config.Logger.Prefix
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	// 如果包含调用者信息，格式化并添加到缓冲区
	if entry.HasCaller() {
		//	 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		//	 自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s \n", logPrefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		// 不包含调用者信息时的格式化
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s \n", logPrefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// InitLogger 初始化并返回一个配置好的logrus.Logger实例
// 它根据全局配置设置输出、格式化器、日志级别等
func InitLogger() *logrus.Logger {
	// 创建新的logger实例
	mLog := logrus.New() // 实例化
	// 设置日志输出为标准输出
	mLog.SetOutput(os.Stdout) // 设置输出类型
	// 根据配置决定是否显示调用者信息
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	// 设置自定义的格式化器
	mLog.SetFormatter(&LogFormatter{}) // 设置自定义的formatter
	// 根据配置设置日志级别
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) // 设置最低日志级别
	// 初始化默认logger
	InitDefaultLogger()
	// todo: 设置输出到按日期命名的文件
	logFilePath := getLogFilePath()
	file, _ := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	// mLog.Out = file
	// 设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	/* file, _ := os.OpenFile("checkemstools.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) */
	writers := []io.Writer{
		file,
		os.Stdout}
	//  同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	mLog.SetOutput(fileAndStdoutWriter)
	return mLog
}

// InitDefaultLogger 配置全局logrus实例
// 它的作用与InitLogger相似，但不返回logger实例
func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) // 设置最低日志级别
}

// getLogFilePath 返回日志文件的路径
// 它根据当前日期生成文件名，并确保日志目录存在
func getLogFilePath() string {
	// 指定日志文件存放目录
	logDir := "logs"

	// 创建或确认目录存在
	if err := os.MkdirAll(logDir, 0o755); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	// 生成以当前日期命名的日志文件名
	t := time.Now()
	fileName := fmt.Sprintf("%s_%s.log", t.Format("2006-01-02"), "app")

	// 返回完整文件路径
	return filepath.Join(logDir, fileName)
}
