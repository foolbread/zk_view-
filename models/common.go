package models

import (
	. "github.com/foolbread/zk_view/models/user"
	. "github.com/foolbread/zk_view/models/zookeeper"
)

func InitModels() {
	InitUser()
	InitZookeeper()
}

