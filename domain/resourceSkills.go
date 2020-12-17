package domain

type ResourceSkills struct {
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
	UserID       string `json:"userId"`
	TypeID       string `json:"typeId"`
	Active       bool   `json:"active"`
	Content      struct {
		JobRoles []struct {
			JRID      string `json:"JR_ID"`
			JobRole   string `json:"jobRole"`
			Primary   bool   `json:"primary"`
			SkillSets []struct {
				JRSSID   string `json:"JRSS_ID"`
				SkillSet string `json:"skillSet"`
				Primary  bool   `json:"primary,omitempty"`
			} `json:"skillSets"`
		} `json:"jobRoles"`
		PrimaryJobCategory struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"primaryJobCategory"`
		SecondaryJobCategory struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"secondaryJobCategory"`
		Certifications struct {
			Badges []struct {
				BadgeID          string `json:"badgeId"`
				BadgeName        string `json:"badgeName"`
				BadgeDescription string `json:"badgeDescription"`
				BadgeImageURL    string `json:"badgeImageUrl"`
				IssueID          string `json:"issueId"`
				IssueDate        string `json:"issueDate"`
				BadgeURL         string `json:"badgeUrl"`
				PublicURL        string `json:"publicUrl"`
			} `json:"badges"`
		} `json:"certifications"`
		UID string `json:"uid"`
	} `json:"content"`
	ID  string `json:"_id"`
	Rev string `json:"_rev"`
}
