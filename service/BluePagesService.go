package service

import (
	"encoding/json"
	"ibm-resource-finder-api/domain"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	url   = goDotEnvVariable("BP_LDAP_URL")
	upUrl = goDotEnvVariable("BP_INTRA_URL")
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetResourceData(email string) domain.Resource {
	// Request to IBM webservice
	resp, err := http.Get(url + "mail=" + email + ".list/byjson")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	// Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Convert JSON to Go Struct
	var resource domain.Resource
	json.Unmarshal(body, &resource)
	// Return Domain Struct
	return resource
}

func GetManagerResourcesData(dn string) domain.Resource {
	// Request to IBM webservice
	resp, err := http.Get(url + "manager=" + dn + ".list/byjson")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	// Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Convert JSON to Go Struct
	var resource domain.Resource
	json.Unmarshal(body, &resource)
	// Return Domain Struct
	return resource
}

func GetResourceDetailsByUid(uid string) domain.ResourceSkills {
	// Request to IBM webservice
	resp, err := http.Get(upUrl + uid + "/profile_extended")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	// Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Convert JSON to Go Struct
	var resource domain.ResourceSkills
	json.Unmarshal(body, &resource)
	// Return Domain Struct
	return resource
}

func GetResourceByManagerDnWithDetails(dn string) domain.Resource {
	var newManagerResource domain.Resource
	managerResource := GetManagerResourcesData(dn)
	for _, entry := range managerResource.Search.Entry {
		var skills domain.Skills
		for _, attribute := range entry.Attribute {
			if attribute.Name == "uid" {
				var detail domain.ResourceSkills
				detail = GetResourceDetailsByUid(attribute.Value[0])
				skills.Name = "badges"
				for _, badges := range detail.Content.Certifications.Badges {
					skills.Value = append(skills.Value, string(badges.BadgeName))
				}
			}
		}
		entry.Attribute = append(entry.Attribute, skills)
		newManagerResource.Search.Entry = append(newManagerResource.Search.Entry, entry)
	}
	managerResource.Search.Entry = newManagerResource.Search.Entry
	return managerResource
}
