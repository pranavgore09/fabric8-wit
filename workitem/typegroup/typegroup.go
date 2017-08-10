package typegroup

import (
	uuid "github.com/satori/go.uuid"

	"github.com/fabric8-services/fabric8-wit/workitem"
)

// WorkItemTypeGroup represents the node in the group of work item types
type WorkItemTypeGroup struct {
	Level                  []int
	Group                  string
	Name                   string
	WorkItemTypeCollection []uuid.UUID
}

// Use following group constants while defining static groups
const (
	GroupPortfolio    = "portfolio"
	GroupRequirements = "requirements"
	GroupExecution    = "execution"
)

// Portfolio0 defines first level of typegroup Portfolio
var Portfolio0 = WorkItemTypeGroup{
	Group: GroupPortfolio,
	Level: []int{0, 0},
	Name:  "Portfolio",
	WorkItemTypeCollection: []uuid.UUID{
		workitem.SystemScenario,
		workitem.SystemFundamental,
		workitem.SystemPapercuts,
	},
}

// Portfolio1 defines second level of typegroup Portfolio
var Portfolio1 = WorkItemTypeGroup{
	Group: GroupPortfolio,
	Level: []int{0, 1},
	Name:  "Portfolio",
	WorkItemTypeCollection: []uuid.UUID{
		workitem.SystemExperience,
		workitem.SystemValueProposition,
	},
}

// Requirements0 defines first level of typegroup Requirements
// This group has less priority than Portfolio
var Requirements0 = WorkItemTypeGroup{
	Group: GroupRequirements,
	Level: []int{1, 0},
	Name:  "Requirements",
	WorkItemTypeCollection: []uuid.UUID{
		workitem.SystemFeature,
		workitem.SystemBug,
	},
}

// Execution0 defines first level of typegroup Execution
// This group has less priority than Requirements
var Execution0 = WorkItemTypeGroup{
	Group: GroupExecution,
	Level: []int{2, 0},
	Name:  "Execution",
	WorkItemTypeCollection: []uuid.UUID{
		workitem.SystemTask,
	},
}
