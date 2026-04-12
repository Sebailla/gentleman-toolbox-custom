package catalog

import "github.com/gentleman-programming/gentle-ai/internal/model"

type Skill struct {
	ID       model.SkillID
	Name     string
	Category string
	Priority string
}

var mvpSkills = []Skill{
	// SDD skills
	{ID: model.SkillSDDInit, Name: "sdd-init", Category: "sdd", Priority: "p0"},

	{ID: model.SkillSDDApply, Name: "sdd-apply", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDVerify, Name: "sdd-verify", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDExplore, Name: "sdd-explore", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDPropose, Name: "sdd-propose", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDSpec, Name: "sdd-spec", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDDesign, Name: "sdd-design", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDTasks, Name: "sdd-tasks", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDArchive, Name: "sdd-archive", Category: "sdd", Priority: "p0"},
	// Foundation skills
	{ID: model.SkillGoTesting, Name: "go-testing", Category: "testing", Priority: "p0"},
	{ID: model.SkillCreator, Name: "skill-creator", Category: "workflow", Priority: "p0"},
	{ID: model.SkillDistiller, Name: "gentleman-distiller", Category: "maintenance", Priority: "p0"},
	{ID: model.SkillContextGuardian, Name: "context-guardian", Category: "context", Priority: "p0"},
	{ID: model.SkillGoReviewer, Name: "go-reviewer", Category: "quality", Priority: "p0"},
	{ID: model.SkillTSReviewer, Name: "typescript-reviewer", Category: "quality", Priority: "p0"},
	{ID: model.SkillDriver, Name: "gentleman-driver", Category: "autonomy", Priority: "p0"},
	{ID: model.SkillLearn, Name: "instinct-learning", Category: "intelligence", Priority: "p0"},
	{ID: model.SkillSentinel, Name: "gentleman-sentinel", Category: "maintenance", Priority: "p0"},
	{ID: model.SkillStoryteller, Name: "gentleman-storyteller", Category: "intelligence", Priority: "p0"},
	{ID: model.SkillCoach, Name: "gentleman-coach", Category: "intelligence", Priority: "p0"},
	{ID: model.SkillBriefing, Name: "gentleman-briefing", Category: "maintenance", Priority: "p0"},
	{ID: model.SkillPlan, Name: "gentleman-plan", Category: "maintenance", Priority: "p0"},
	{ID: model.SkillJudge, Name: "judgment-day", Category: "quality", Priority: "p0"},
	{ID: model.SkillBlueprint, Name: "gentleman-blueprint", Category: "maintenance", Priority: "p0"},
	{ID: model.SkillRefactor, Name: "instinct-refactor", Category: "maintenance", Priority: "p0"},
	{ID: model.SkillPRFix, Name: "pr-fix", Category: "automation", Priority: "p0"},
	{ID: model.SkillConsole, Name: "gentleman-console", Category: "intelligence", Priority: "p0"},
	{ID: model.SkillSync, Name: "fleet-sync", Category: "maintenance", Priority: "p0"},
}

func MVPSkills() []Skill {
	skills := make([]Skill, len(mvpSkills))
	copy(skills, mvpSkills)
	return skills
}
