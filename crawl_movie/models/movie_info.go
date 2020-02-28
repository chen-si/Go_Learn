package models

import (
	"regexp"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var db orm.Ormer

type MovieInfo struct {
	Id                 int64
	MovieId            int64
	MovieName          string
	MoviePic           string
	MovieDirector      string
	WovieWriter        string
	MovieCountry       string
	MovieLanguage      string
	MovieMainCharacter string
	MovieType          string
	MovieOnTime        string
	MovieSpan          string
	MovieGrade         string
	Remark             string
	_create_time       string
}

func init() {
	orm.Debug = true
	//链接数据库
	orm.RegisterDataBase("default", "mysql", "root:1234w5asd@tcp(localhost:3306)/douban?charset=utf8", 30)
	//绑定表
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func Addmovie(movie_info *MovieInfo) (int64, error) {
	//插入
	return db.Insert(movie_info)
}

func GetMovieDirector(moviehtml string) string {
	if moviehtml == "" {
		return ""
	}
	//下面使用正则匹配

	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*)</a>`)
	result := reg.FindAllStringSubmatch(moviehtml, -1)
	return string(result[0][1])
}

func GetMovieName(moviehtml string) string {
	if moviehtml == "" {
		return ""
	}
	//下面使用正则匹配

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(moviehtml, -1)
	return string(result[0][1])
}

func GetMovieMainCharactor(moviehtml string) string {
	if moviehtml == "" {
		return ""
	}
	//下面使用正则匹配

	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(moviehtml, -1)

	mainCharactor := ""
	for _, v := range result {
		mainCharactor += v[1] + "//"
	}

	return mainCharactor
}
