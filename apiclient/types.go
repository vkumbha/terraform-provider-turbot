package apiclient

// Validation response - returned by API validation call
type ValidationResponse struct {
	Schema struct {
		QueryType struct {
			Name string
		}
	}
}

// is the validation response successful?
func (response *ValidationResponse) isValid() bool {
	return response.Schema.QueryType.Name == "Query"
}

// ApiResponse: used to unmarshall API error responses
type ApiResponse struct {
	Errors []Error
}
type Error struct {
	Message string
}

// PolicySettingResponse: must be consistent with fields defined in readPolicySettingQuery
type PolicySettingResponse struct {
	PolicySetting PolicySetting
}

type PolicySetting struct {
	Value              interface{}
	ValueSource        string
	Precedence         string
	Template           string
	TemplateInput      string
	Input              string
	Note               string
	ValidFromTimestamp string
	ValidToTimestamp   string
	Turbot             TurbotMetadata
}
type TurbotMetadata struct {
	Id       string
	ParentId string
	Akas     []string
}

// PolicyValueResponse: must be consistent with fields defined in readPolicyValueQuery
type PolicyValueResponse struct {
	PolicyValue PolicyValue
}
type PolicyValue struct {
	Value      interface{}
	Precedence string
	State      string
	Reason     string
	Details    string
	Setting    PolicySetting
	Turbot     TurbotMetadata
}

// InstallModResponse: must be consistent with with fields defined in installModMutation
type InstallModResponse struct {
	Mod struct {
		Turbot TurbotMetadata
	}
}

type ReadModResponse struct {
	Mod struct {
		Object  interface{}
		Uri     string
		Parent  string
		Version string
	}
}

type UninstallModResponse struct {
	ModUninstall struct {
		Success bool
	}
}

type Mod struct {
	Org     string
	Mod     string
	Version string
	Parent  string
}

type CreateResourceResponse struct {
	Resource struct {
		Turbot TurbotMetadata
	}
}

type ReadFolderResponse struct {
	Resource Folder
}

type Folder struct {
	Turbot      TurbotMetadata
	Title       string
	Description string
	Parent      string
}

// note: the Resource property is just an interface{} - this is mapped manually into a Resource object,
// rather than unmarshalled. This is to allow for dynamic data types, while always having the Turbot property
type ReadResourceResponse struct {
	Resource interface{}
}

type Resource struct {
	Turbot TurbotMetadata
	Data   map[string]interface{}
}
