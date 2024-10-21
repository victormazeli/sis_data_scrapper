package transformdata

// Structs matching the input JSON structure
type BundleLevel struct {
	ID             string        `json:"id"`
	Credits        int           `json:"credits"`
	ModuleGroups   []ModuleGroup `json:"moduleGroups"`
	ProgramLevelID string        `json:"programLevelId"`
}

type ModuleGroup struct {
	Credits               int            `json:"credits"`
	ID                    string         `json:"id"`
	ModuleOptionGroupType string         `json:"moduleOptionGroupType"`
	ModuleOptions         []ModuleOption `json:"moduleOptions"`
	Name                  string         `json:"name"`
}

type ModuleOption struct {
	CourseID string `json:"courseId"`
	ID       string `json:"id"`
}

type Bundle struct {
	ID           string        `json:"id"`
	Name         string        `json:"bundleCode"`
	BundleLevels []BundleLevel `json:"bundleLevels"`
	ProgramID    string        `json:"programId"`
}

// Output Structs matching the new schema
type OutputBundle struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	ProgrammeID  string              `json:"programme_id"`
	BundleLevels []OutputBundleLevel `json:"bundle_levels"`
}

type OutputBundleLevel struct {
	ID               string              `json:"id"`
	ProgrammeLevelID string              `json:"programme_level_id"`
	CreditRequired   int                 `json:"credit_required"`
	ModuleGroups     []OutputModuleGroup `json:"module_groups"`
}

type OutputModuleGroup struct {
	ID              string               `json:"id"`
	GroupName       string               `json:"group_name"`
	GroupType       string               `json:"group_type"`
	NumberOfCourses int                  `json:"number_of_courses"`
	ModuleOptions   []OutputModuleOption `json:"module_options"`
}

type OutputModuleOption struct {
	ID       string `json:"id"`
	CourseID string `json:"course_id"`
}

// Transform a Bundle into the new format
func TransformBundle(bundle Bundle) OutputBundle {

	var bundleLevels []OutputBundleLevel
	for _, level := range bundle.BundleLevels {
		bundleLevels = append(bundleLevels, TransformBundleLevel(level))
	}

	return OutputBundle{
		ID:           bundle.ID,
		Name:         bundle.Name,
		ProgrammeID:  bundle.ProgramID,
		BundleLevels: bundleLevels,
	}
}

// Transform a BundleLevel into the new format
func TransformBundleLevel(level BundleLevel) OutputBundleLevel {

	var moduleGroups []OutputModuleGroup
	for _, group := range level.ModuleGroups {
		moduleGroups = append(moduleGroups, TransformModuleGroup(group))
	}

	return OutputBundleLevel{
		ID:               level.ID,
		ProgrammeLevelID: level.ProgramLevelID,
		CreditRequired:   level.Credits,
		ModuleGroups:     moduleGroups,
	}
}

// Transform a ModuleGroup into the new format
func TransformModuleGroup(group ModuleGroup) OutputModuleGroup {

	var moduleOptions []OutputModuleOption
	for _, option := range group.ModuleOptions {
		moduleOptions = append(moduleOptions, TransformModuleOption(option))
	}

	return OutputModuleGroup{
		ID:              group.ID,
		GroupName:       group.Name,
		GroupType:       group.ModuleOptionGroupType,
		NumberOfCourses: len(group.ModuleOptions),
		ModuleOptions:   moduleOptions,
	}
}

// Transform a ModuleOption into the new format
func TransformModuleOption(option ModuleOption) OutputModuleOption {

	return OutputModuleOption{
		ID:       option.ID,
		CourseID: option.CourseID,
	}
}
