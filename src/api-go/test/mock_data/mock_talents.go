package mock_data

import (
	"fmt"
	wh "github.com/jmilosze/wfrp-hammergen-go/internal/domain/warhammer"
)

var talent0 = wh.Wh{
	Id:      "500000000000000000000000",
	OwnerId: user1.Id,
	Object: &wh.Talent{
		Name:        "talent 0",
		Description: fmt.Sprintf("owned by %s", user1.Username),
		Tests:       "tests",
		MaxRank:     3,
		Attribute:   wh.AttBS,
		Attribute2:  wh.AttWS,
		IsGroup:     false,
		Modifiers: &wh.Modifiers{
			Size:     1,
			Movement: 1,
			Attributes: &wh.Attributes{
				WS:  1,
				BS:  2,
				S:   3,
				T:   4,
				I:   5,
				Ag:  6,
				Dex: 7,
				Int: 8,
				WP:  9,
				Fel: 10,
			},
		},
		Group:  []string{talent1.Id},
		Shared: false,
		Source: map[wh.Source]string{
			wh.SourceArchivesOfTheEmpireVolI: "d",
			wh.SourceSeaOfClaws:              "e",
		},
	},
}

var talent1 = wh.Wh{
	Id:      "500000000000000000000001",
	OwnerId: user1.Id,
	Object: &wh.Talent{
		Name:        "talent 1",
		Description: fmt.Sprintf("owned by %s", user1.Username),
	},
}

func NewMockTalents() []*wh.Wh {
	return []*wh.Wh{&talent0, &talent1}
}
