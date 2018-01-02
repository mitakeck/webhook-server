package main

import "time"

type activity []struct {
	ID      int `json:"id"`
	Project struct {
		ID                                int         `json:"id"`
		ProjectKey                        string      `json:"projectKey"`
		Name                              string      `json:"name"`
		ChartEnabled                      bool        `json:"chartEnabled"`
		SubtaskingEnabled                 bool        `json:"subtaskingEnabled"`
		ProjectLeaderCanEditProjectLeader bool        `json:"projectLeaderCanEditProjectLeader"`
		TextFormattingRule                interface{} `json:"textFormattingRule"`
		Archived                          bool        `json:"archived"`
		DisplayOrder                      int         `json:"displayOrder"`
	} `json:"project"`
	Type    int `json:"type"`
	Content struct {
		ID          int    `json:"id"`
		KeyID       int    `json:"key_id"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Comment     struct {
			ID      int    `json:"id"`
			Content string `json:"content"`
		} `json:"comment"`
		Changes []struct {
			Field    string `json:"field"`
			NewValue string `json:"new_value"`
			OldValue string `json:"old_value"`
			Type     string `json:"type"`
		} `json:"changes"`
	} `json:"content"`
	Notifications []struct {
		ID          int  `json:"id"`
		AlreadyRead bool `json:"alreadyRead"`
		Reason      int  `json:"reason"`
		User        struct {
			ID          int    `json:"id"`
			UserID      string `json:"userId"`
			Name        string `json:"name"`
			RoleType    int    `json:"roleType"`
			Lang        string `json:"lang"`
			MailAddress string `json:"mailAddress"`
		} `json:"user"`
		ResourceAlreadyRead bool `json:"resourceAlreadyRead"`
	} `json:"notifications"`
	CreatedUser struct {
		ID          int    `json:"id"`
		UserID      string `json:"userId"`
		Name        string `json:"name"`
		RoleType    int    `json:"roleType"`
		Lang        string `json:"lang"`
		MailAddress string `json:"mailAddress"`
	} `json:"createdUser"`
	Created time.Time `json:"created"`
}
