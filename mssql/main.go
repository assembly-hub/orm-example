package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/assembly-hub/mssql"
	"github.com/assembly-hub/orm"
	"github.com/assembly-hub/orm/dbtype"

	"github.com/assembly-hub/orm-example/mssql/dao"
)

var dbRef = orm.NewReference(dbtype.SQLServer)

// 定义数据表
func initRef() {
	// 表定义：建议放在表结构文件的init函数中
	dbRef.AddTableDef("table1", dao.Table1{})
	dbRef.AddTableDef("table2", dao.Table2{})
	dbRef.AddTableDef("table3", dao.Table3{})
	dbRef.BuildRefs()
}

// 获取数据库连接
func getDB() *sql.DB {
	cli := mssql.NewClient(&mssql.Config{
		Username: "sa",
		Password: "1aA123456",
		Server:   "localhost",
		Port:     1443,
		DBName:   "example",
	})
	db, err := cli.Connect()
	if err != nil {
		panic(err)
	}
	return db
}

func insertData(tb *orm.ORM) {
	dt := map[string]interface{}{
		"name": "test",
		"json": []string{"1", "2", "3"}, // 框架会进行处理
	}
	// 参数可以是 map 或 table1 struct，当id == nil or id == "" or id == 0 将忽略id字段，否则将插入
	// tb.InsertMany()
	// tb.InsertManySameClos()
	_, err := tb.InsertOne(dt)
	if err != nil {
		fmt.Println(err)
	}
}

func upsertData(tb *orm.ORM) {
	dt := map[string]interface{}{
		"name": "test",
		"json": []string{"1", "2", "3"}, // 框架会进行处理
	}
	// 参数可以是 map 或 table1 struct，当id == nil or id == "" or id == 0 将忽略id字段，否则将插入
	// tb.UpsertMany()
	// tb.UpsertManySameClos()
	_, err := tb.UpsertOne(dt)
	if err != nil {
		fmt.Println(err)
	}
}

func replaceData(tb *orm.ORM) {
	dt := map[string]interface{}{
		"name": "test",
		"json": []string{"1", "2", "3"}, // 框架会进行处理
	}
	// 参数可以是 map 或 table1 struct，当id == nil or id == "" or id == 0 将忽略id字段，否则将插入
	// tb.ReplaceMany()
	// tb.ReplaceManySameClos()
	_, err := tb.ReplaceOne(dt)
	if err != nil {
		fmt.Println(err)
	}
}

func saveData(tb *orm.ORM) {
	dt := map[string]interface{}{
		"name": "test",
		"json": []string{"1", "2", "3"}, // 框架会进行处理
	}
	// 参数可以是 map 或 table1 struct，当id == nil or id == "" or id == 0 将忽略id字段，否则将插入
	_, err := tb.SaveMany([]interface{}{dt}, true)
	if err != nil {
		fmt.Println(err)
	}
}

func updateByWhere(tb *orm.ORM) {
	upData := map[string]interface{}{
		"name": "test",
		"json": []string{"1", "2", "3"},
	}
	upWhere := orm.Where{
		"id__gt": 1,
	}
	// 自定义更新条件
	_, err := tb.UpdateByWhere(upData, upWhere)
	if err != nil {
		fmt.Println(err)
	}
}

func pageData(tb *orm.ORM) {
	// 排序字段，默认升序，- 降序
	tb.Order("id", "-ref") // id asc, ref desc

	// tb1.Limit(n) limit n
	// tb1.OverLimit(n, m) limit n,m
	// tb1.Page(no, size) limit size*(no-1),size
	// tb1.Distinct(b)
	var data []dao.Table1
	pd, err := tb.PageData(&data, false, 1, 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pd)
}

// 查询数据
func query(tb *orm.ORM) {
	// 条件查询
	/**
	每种算子均支持 强制忽略大小写查询和强制区分大小写查询，用法为：key__ignore_eq，key__bin_eq 或简写 key__i_eq，key__b_eq
	where 条件 支持子查询
	= : "key": "val" or "key__eq": "val"
	< : "key__lt": 1
	<= : "key__lte": 1
	> : "key__gt": 1
	>= : "key__gte": 1
	!= : "key__ne": 1
	in : "key__in": [1]
	not in : "key__nin": [1]
	date : "key__date": "2022-01-01" or time.Time(oracle必须)
	between : "key__between": [1, 2]

	以下不支持子查询
	is null : "key__null": true
	is not null : "key__null": false
	$or : map[string]interface{} or []map[string]interface{}
	$and : map[string]interface{} or []map[string]interface{}
	and_like : 参数为数组，针对数组每个元素分别取 like，之后条件之间取 and
			"key__startswith": "123"
			"key__endswith": "123"
			"key__contains": "123" or ["123", "123"]
			"key__customlike": "__st" or ["%test", "%test%", "test%"] // 自定义查询语句
	or_like : 参数为数组，针对数组每个元素分别取 like，之后条件之间取 or
			"key__orstartswith": "123" or ["123", "123"]
			"key__orendswith": "123" or ["123", "123"]
			"key__orcontains": "123" or ["123", "123"]
	        "key__orcustomlike": "__st" or ["%test", "%test%", "test%"] // 自定义查询语句

	原始数据，#修饰的字段为原始字段，不做处理，其他的字段会根据tag计算
	~ 为条件取反，必须在最前面，可用在所有算子前面，如果与#连用，#应在~后面，如：~#test
	*/
	tb.Query("id__gt", 0, "name__startswith", "test")
	// 等价
	tb.Wheres(orm.Where{
		"id__gt":           0,      // table1的id > 1
		"name__startswith": "test", // table1的name like 'test%'
		// tb2：table1的tag字段（可以理解为table2的别名），tb3：table2的tag字段（可以理解为table3的别名）
		// 含义：select * from table1 left join table2 left join table3 where table3.id > 1
		// "tb2.tb3.id__gt": 1,
	})

	exist, err := tb.Exist()
	if err != nil {
		fmt.Println(exist, err.Error())
	}

	// 查询数据
	var data dao.Table1
	// 第一个参数：接收数据的容器，可以是 struct 或 map 或 []map 或 []struct
	// 		参数非slice时，查询会只关心第一条
	// 第二个参数：false，子表数据查出之后会以嵌套的方式写入到第一个参数
	/** 例子
	{
		"id": 1,
		"name": "test",
		"json": ["test1", "test2", "test3"],
		"ref": 1,
		"tb2": {
			"id": 1,
			"name": "test",
			"age": 1,
			"ref": 0,
			"tb3": null
		}
	}
	*/
	// 		参数：true，会根据 SelectColLinkStr 指定的链接串，默认是 _ ，将tag数据组装，写到第一个参数中
	/** 例子
	{
		"id": 1,
		"name": "test",
		"json": ["test1", "test2", "test3"],
		"ref": 1,
		"tb2_id": 1,
		"tb2_name": "test",
		"tb2_age": 1,
		"tb2_ref": 0
	}
	*/
	err = tb.Select("*3").ToData(&data, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

// 编译需要加上环境变量 CGO_ENABLED=1
func main() {
	// 初始化表定义
	initRef()

	// 初始化数据库链接
	db := getDB()

	// 上下文，可控制超时
	ctx := context.Background()

	tb := orm.NewORM(ctx, "table1", db, dbRef)
	insertData(tb)
	query(tb)

	// 事务demo
	// 返回nil事务执行成功
	err := orm.TransSession(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		// 用 tx 生成新的orm对象即可，使用与普通的orm对象没有区别
		sessTb := orm.NewORMWithTx(ctx, "table1", tx, dbRef)
		updateByWhere(sessTb)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
