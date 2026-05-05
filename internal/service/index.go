package service

import (
	"fmt"

	"dev.justdrven/loadbalancer/pkg"
	"gorm.io/gorm"
)

var services []*ManagedService

func GetService() []*ManagedService {
	return services
}

func LoadServices(db *gorm.DB, applicationType string) {

	var fetchServices []Service

	db.Find(&fetchServices).Where("type = ?", applicationType)
	for i := range fetchServices {
		dbService := fetchServices[i]

		if pkg.MakeDiagnostic(dbService.Address) == pkg.Failed {
			continue
		}

		saveService(mapToManagedService(dbService))
		fmt.Printf("[SVC-MANAGER]  Service #%d : %s\n", i+1, dbService.Address)
	}

	fmt.Printf("[SVC-MANAGER] Loaded %d services\n", len(services))

}

func mapToManagedService(service Service) *ManagedService {
	addr := createAddressService(service)
	maxReferences := service.MaxReferences

	return &ManagedService{
		Address:       addr,
		MaxReferences: maxReferences,
		References:    0,
	}
}

func createAddressService(service Service) string {
	addr := service.Address
	schema := extractSchema(service)

	return fmt.Sprintf("%s://%s", schema, addr)
}

func extractSchema(service Service) string {
	var schema string
	if service.SSL {
		schema = "https"
	} else {
		schema = "http"
	}

	return schema
}

func saveService(service *ManagedService) {
	services = append(services, service)
}

func RefGetBestService() *ManagedService {
	var bestService *ManagedService
	services := GetService()

	for i := range services {
		service := services[i]

		if !IsBest(service.References, service.MaxReferences) {
			continue
		}

		bestService = service
		break

	}

	return bestService

}

func IsBest(references int, maxReferences int) bool {
	if references < maxReferences {
		return true
	}

	return false

}
