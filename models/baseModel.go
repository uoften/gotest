package base_model
/**
主要是用于链接数据库
分配链接Db
*/
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
// 数据库Db已连接
var Db *gorm.DB

func init(){
	db, err := gorm.Open("mysql", "jiangsheng:**123456a***_@tcp(rm-d5e97cc4d4o0698c4xo**.mysql.rds.aliyuncs.com:3306)/bc?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Panic("bc数据库链接失败",err)
		return
	}
	// 如果设置禁用表名复数形式属性为 true，`User` 的表名将是 `user`
	db.SingularTable(true)
	Db = db
	log.Println("---数据库初始化链接成功---")
	//defer db.Close()  // 后面做连接池
}
/**
分页封装
*/
func Paginate(page int,pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize=10
		}
		offset :=( page - 1 ) * pageSize
		return db.Where(map[string]interface{}{}).Offset(offset).Limit(pageSize)
	}
}
/**
排序判断
*/
func Order(sort string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		// 判断字符串第一个字符是 + -
		mark := sort[0:1]
		orderStr := sort[1:]
		switch {
		case mark == "-":
			orderStr =orderStr+` ASC`
		default:
			orderStr =orderStr+` DESC`
		}
		return db.Order(orderStr)
	}
}
/**
where查询
*/
func Where(wheres map[string]map[string]string) func(db *gorm.DB) *gorm.DB {
	var andWhere, orWhere = map[string]interface{}{}, map[string]interface{}{}
	var likeWhere string
	for key,value := range wheres {
		//组装where
		if key=="AND" {
			for key2,value2 := range value {
				andWhere[key2]=value2
			}
		}
		if key=="OR" {
			for key2,value2 := range value {
				orWhere[key2]=value2
			}
		}
		if key=="LIKE" {
			for key2,value2 := range value {
				likeWhere +=key2+` LIKE "`+value2+`%" OR `
			}
			// 去除尾部OR
			likeWhere = likeWhere[:len(likeWhere)-3]
		}
	}
	return func (db *gorm.DB) *gorm.DB {
		return db.Where(andWhere).Where(likeWhere).Or(orWhere)
	}
}