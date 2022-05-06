package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000 //初始化一个开始的时间
	node, err = sf.NewNode(machineID)  //指定一个机器的id
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2020-07-01", 1); err != nil {
//		fmt.Printf("init snowflake failed,err%v\n", err)
//		return
//	}
//	id := GenID()
//	fmt.Println(id)
//}
