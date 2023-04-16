package service

import (
	"context"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/service/dto"
)

type Dashboard struct {
	service.Service
}

// All 获取所有看板
func (d *Dashboard) All(ctx context.Context) (*dto.AllDashboardRsp, error) {
	return &dto.AllDashboardRsp{CampProgressions: &dto.CampaignProgression{
		Dungeon: &dto.Progression{
			Done:       20,
			Unfinished: 80,
			Total:      100,
		},
		Forrest: &dto.Progression{
			Done:       40,
			Unfinished: 60,
			Total:      100,
		},
		Desert: &dto.Progression{
			Done:       50,
			Unfinished: 50,
			Total:      100,
		},
		Mountain: &dto.Progression{
			Done:       80,
			Unfinished: 20,
			Total:      100,
		},
		Glacier: &dto.Progression{
			Done:       100,
			Unfinished: 0,
			Total:      100,
		},
	}}, nil
}
