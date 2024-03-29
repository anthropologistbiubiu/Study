package model

type TimePorint struct {
	StartTime int64 `bson:"startTime"` //开始时间
	EndTime   int64 `bson:"endTime"`   //结束时间
}
type LogRecord struct {
	JobName string     `bson:"jobName"` //任务名
	Command string     `bson:"command"` //shell命令
	Err     string     `bson:"err"`     //脚本错误
	Content string     `bson:"content"` //脚本输出
	Tp      TimePorint //执行时间
}

// 查询实体
type FindByJobName struct {
	JobName string `bson:"jobName"` //任务名
}

// 更新实体
type UpdateByJobName struct {
	Command string `bson:"command"` //shell命令
	Content string `bson:"content"` //脚本输出
}

type TimePorintFilter struct {
	StartTime interface{} `bson:"tp.startTime,omitempty"` //开始时间
	EndTime   interface{} `bson:"tp.endTime,omitempty"`   //结束时间
}

type LogRecordFilter struct {
	ID      interface{} `bson:"_id,omitempty"`
	JobName interface{} `bson:"jobName,omitempty" json:"jobName"` //任务名
	Command interface{} `bson:"command,omitempty" `               //shell命令
	Err     interface{} `bson:"err,omitempty"`                    //脚本错误
	Content interface{} `bson:"content,omitempty"`                //脚本输出
	Tp      interface{} `bson:"tp,omitempty"`                     //执行时间
}

// 小于
type Lt struct {
	Lt int64 `bson:"$lt"` //小于
}

// 分组
type Group struct {
	Group interface{} `bson:"$group"` //小于
}

// 求和
type Sum struct {
	Sum interface{} `bson:"$sum"` //小于
}
