package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineID int) error {
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1e6
	node, err = sf.NewNode(int64(machineID))
	return err
}

func GenID() int64 {
	return node.Generate().Int64()
}
