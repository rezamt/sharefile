package main

import (
	"encoding/json"
	"net/http"
)

type Info struct {
	Name string `json:"name"`
	ID   string `json:"ver"`
}

type Employee struct {
	IsAdministrator         bool     `json:"IsAdministrator"`
	CanCreateFolders        bool     `json:"CanCreateFolders"`
	CanUseFileBox           bool     `json:"CanUseFileBox"`
	CanManageUsers          bool     `json:"CanManageUsers"`
	IsVirtualClient         bool     `json:"IsVirtualClient"`
	DiskSpace               int      `json:"DiskSpace"`
	Bandwidth               int      `json:"Bandwidth"`
	TotalSharedFiles        int      `json:"TotalSharedFiles"`
	Contacted               int      `json:"Contacted"`
	FullName                string   `json:"FullName"`
	ReferredBy              string   `json:"ReferredBy"`
	DateCreated             string   `json:"DateCreated"`
	FullNameShort           string   `json:"FullNameShort"`
	IsDeleted               bool     `json:"IsDeleted"`
	AffiliatedPartnerUserId string   `json:"AffiliatedPartnerUserId"`
	IsBillingContact        bool     `json:"IsBillingContact"`
	Username                string   `json:"Username"`
	Domain                  string   `json:"Domain"`
	FirstName               string   `json:"FirstName"`
	LastName                string   `json:"LastName"`
	Company                 string   `json:"Company"`
	IsConfirmed             bool     `json:"IsConfirmed"`
	IsDisabled              bool     `json:"IsDisabled"`
	LastAnyLogin            string   `json:"LastAnyLogin"`
	CreatedDate             string   `json:"CreateData"`
	Name                    string   `json:"Name"`
	Email                   string   `json:"Email"`
	Id                      string   `json:"Id"`
	Roles                   []string `json:"Roles"`
	Emails                  []string `json:"Emails"`
}

func Version(w http.ResponseWriter, r *http.Request) {
	Elog.Info(1, "Call Version")

	w.Header().Add("Content-Type", "application/json")

	var version = Info{"Sharefile Service Proxy", "v1.0"}
	json.NewEncoder(w).Encode(version)
}

func ListEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	Elog.Info(1, "Call ListEmployees")

	employees := []Employee{
		{
			IsAdministrator:         false,
			CanCreateFolders:        false,
			CanUseFileBox:           false,
			CanManageUsers:          false,
			IsVirtualClient:         true,
			DiskSpace:               -1,
			Bandwidth:               -1,
			TotalSharedFiles:        0,
			Contacted:               0,
			FullName:                "the emp1",
			ReferredBy:              "admin@domain.com",
			DateCreated:             "2019-08-13T00:43:18.803Z",
			FullNameShort:           "the emp1",
			IsDeleted:               false,
			AffiliatedPartnerUserId: "",
			IsBillingContact:        false,
			Username:                "emp1",
			Domain:                  "example.com",
			FirstName:               "the",
			LastName:                "emp1",
			Company:                 "examplecom",
			IsConfirmed:             true,
			IsDisabled:              false,
			LastAnyLogin:            "2019-17-13T00:43:18.803Z",
			Email:                   "theemp1@example.com",
			Id:                      "77192f38-c1a1-11e9-9cb5-2a2ae2dbcce4",
			Roles: []string{
				"Employee",
				"MasterAdmin",
				"SuperUser",
				"CanChangePassword",
			},
			Emails: []string{
				"emp1@nodomain.com",
			},
		},
		{
			IsAdministrator:         false,
			CanCreateFolders:        false,
			CanUseFileBox:           false,
			CanManageUsers:          false,
			IsVirtualClient:         true,
			DiskSpace:               -1,
			Bandwidth:               -1,
			TotalSharedFiles:        0,
			Contacted:               0,
			FullName:                "the emp2",
			ReferredBy:              "admin@domain.com",
			DateCreated:             "2019-08-13T00:43:18.803Z",
			FullNameShort:           "the emp1",
			IsDeleted:               false,
			AffiliatedPartnerUserId: "",
			IsBillingContact:        false,
			Username:                "emp2",
			Domain:                  "example.com",
			FirstName:               "the",
			LastName:                "emp2",
			Company:                 "examplecom",
			IsConfirmed:             true,
			IsDisabled:              false,
			LastAnyLogin:            "2019-17-13T00:43:18.803Z",
			Email:                   "theemp1@example.com",
			Id:                      "77192f38-c1a1-11e9-9cb5-2a2ae2dbcce4",
			Roles: []string{
				"Employee",
				"MasterAdmin",
				"SuperUser",
				"CanChangePassword",
				"CreateNetworkShareConnectors",
				"CreateSharePointConnectors",
				"CanSendDocumentsForSignature",
				"CanViewSignatureDocuments",
				"CanManageSignatureTemplates",
			},
			Emails: []string{
				"emp2@nodomain.com",
			},
		},
		{
			IsAdministrator:         true,
			CanCreateFolders:        true,
			CanUseFileBox:           true,
			CanManageUsers:          true,
			IsVirtualClient:         true,
			DiskSpace:               -1,
			Bandwidth:               -1,
			TotalSharedFiles:        0,
			Contacted:               0,
			FullName:                "the Mr Admin",
			ReferredBy:              "admin@domain.com",
			DateCreated:             "2019-08-13T00:43:18.803Z",
			FullNameShort:           "the emp1",
			IsDeleted:               false,
			AffiliatedPartnerUserId: "",
			IsBillingContact:        false,
			Username:                "Admin",
			Domain:                  "example.com",
			FirstName:               "the Mr.",
			LastName:                "Admininstrator",
			Company:                 "examplecom",
			IsConfirmed:             true,
			IsDisabled:              false,
			LastAnyLogin:            "2019-17-13T00:43:18.803Z",
			Email:                   "admin@example.com",
			Id:                      "77192f38-c1a1-11e9-9cb5-2a2ae2dbcce4",
			Roles: []string{
				"Employee",
				"MasterAdmin",
				"SuperUser",
				"CanChangePassword",
				"CanManageMySettings",
				"CanUseFileBox",
				"CanManageUsers",
				"CanCreateFolders",
				"CanUseDropBox",
				"CanSelectFolderZone",
				"CreateNetworkShareConnectors",
				"CreateSharePointConnectors",
				"CanSendDocumentsForSignature",
				"CanViewSignatureDocuments",
				"CanManageSignatureTemplates",
				"AdminAccountPolicies",
				"AdminBilling",
				"AdminBranding",
				"AdminChangeFolderPay",
				"AdminChangePlan",
				"AdminFileBoxAccess",
				"AdminManageEmployees",
				"AdminRemoteUploadForms",
				"AdminReporting",
				"AdminSharedDistGroups",
				"AdminSharedAddressBook",
				"AdminViewVdrAnalytics",
				"AdminViewReceipts",
				"AdminDelegate",
				"AdminManageDropBox",
				"AdminManageFolderTemplates",
				"AdminEmailMessages",
				"AdminSSO",
				"AdminSuperGroup",
				"AdminZones",
				"AdminCreateSharedGroups",
				"AdminArchivedSearch",
				"AdminConnectors",
				"AdminEmailArchiver",
				"AdminCanAdministerCustomerAccount",
				"AdminCustomWorkflow",
			},
			Emails: []string{
				"admin@example.com",
			},
		},
	}

	json.NewEncoder(w).Encode(employees)
}

func ListRoles(w http.ResponseWriter, r *http.Request) {
	Elog.Info(1, "Call ListRoles")
	roles := [] string{
		"Employee",
		"MasterAdmin",
		"SuperUser",
		"CanChangePassword",
		"CanManageMySettings",
		"CanUseFileBox",
		"CanManageUsers",
		"CanCreateFolders",
		"CanUseDropBox",
		"CanSelectFolderZone",
		"CreateNetworkShareConnectors",
		"CreateSharePointConnectors",
		"CanSendDocumentsForSignature",
		"CanViewSignatureDocuments",
		"CanManageSignatureTemplates",
		"AdminAccountPolicies",
		"AdminBilling",
		"AdminBranding",
		"AdminChangeFolderPay",
		"AdminChangePlan",
		"AdminFileBoxAccess",
		"AdminManageEmployees",
		"AdminRemoteUploadForms",
		"AdminReporting",
		"AdminSharedDistGroups",
		"AdminSharedAddressBook",
		"AdminViewVdrAnalytics",
		"AdminViewReceipts",
		"AdminDelegate",
		"AdminManageDropBox",
		"AdminManageFolderTemplates",
		"AdminEmailMessages",
		"AdminSSO",
		"AdminSuperGroup",
		"AdminZones",
		"AdminCreateSharedGroups",
		"AdminArchivedSearch",
		"AdminConnectors",
		"AdminEmailArchiver",
		"AdminCanAdministerCustomerAccount",
		"AdminCustomWorkflow",
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(roles)
}
