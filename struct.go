package ytUpload

type createVideoBody struct {
	ChannelID  string `json:"channelId"`
	ResourceID struct {
		ScottyResourceID struct {
			ID string `json:"id"`
		} `json:"scottyResourceId"`
	} `json:"resourceId"`
	FrontendUploadID string `json:"frontendUploadId"`
	InitialMetadata  struct {
		Title struct {
			NewTitle string `json:"newTitle"`
		} `json:"title"`
		Description struct {
			NewDescription string `json:"newDescription"`
			ShouldSegment  bool   `json:"shouldSegment"`
		} `json:"description"`
		Privacy struct {
			NewPrivacy string `json:"newPrivacy"`
		} `json:"privacy"`
		DraftState struct {
			IsDraft bool `json:"isDraft"`
		} `json:"draftState"`
	} `json:"initialMetadata"`
	Context struct {
		Client struct {
			ClientName       int    `json:"clientName"`
			ClientVersion    string `json:"clientVersion"`
			Hl               string `json:"hl"`
			Gl               string `json:"gl"`
			ExperimentsToken string `json:"experimentsToken"`
			UtcOffsetMinutes int    `json:"utcOffsetMinutes"`
		} `json:"client"`
		Request struct {
			ReturnLogEntry          bool          `json:"returnLogEntry"`
			InternalExperimentFlags []interface{} `json:"internalExperimentFlags"`
			SessionInfo             struct {
				Token string `json:"token"`
			} `json:"sessionInfo"`
		} `json:"request"`
		User struct {
			OnBehalfOfUser    string `json:"onBehalfOfUser"`
			DelegationContext struct {
				RoleType struct {
					ChannelRoleType string `json:"channelRoleType"`
				} `json:"roleType"`
				ExternalChannelID string `json:"externalChannelId"`
			} `json:"delegationContext"`
			SerializedDelegationContext string `json:"serializedDelegationContext"`
		} `json:"user"`
		ClientScreenNonce string `json:"clientScreenNonce"`
	} `json:"context"`
	DelegationContext struct {
		RoleType struct {
			ChannelRoleType string `json:"channelRoleType"`
		} `json:"roleType"`
		ExternalChannelID string `json:"externalChannelId"`
	} `json:"delegationContext"`
}
