// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package metrics

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/metrics/metricutils"
)

type ItemAction struct {
	actionutils.ParentAction
}

func (this *ItemAction) Init() {
	this.Nav("", "", "item")
}

func (this *ItemAction) RunGet(params struct {
	ItemId int64
}) {
	_, err := metricutils.InitItem(this.Parent(), params.ItemId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Show()
}
