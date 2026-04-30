package manager

import (
	"fmt"

	"dev.justdrven/loadbalancer/data"
	"github.com/go-ping/ping"
	"gorm.io/gorm"
)

var services []*data.ManagedService

func SvcGetService() []*data.ManagedService {
	return services
}

func SvcLoadServices(db gorm.DB, applicationType string) {

	var fetchServices []data.Service

	db.Find(&fetchServices).Where("type = ?", applicationType)
	for i := range fetchServices {
		dbService := fetchServices[i]

		if makeDiagnostic(dbService.Address) == data.Failed {
			continue
		}

		saveService(mapToManagedService(dbService))
		fmt.Printf("[SVC-MANAGER]  Service #%d : %s\n", i+1, dbService.Address)
	}

	fmt.Printf("[SVC-MANAGER] Loaded %d services\n", len(services))

}
func makeDiagnostic(addr string) data.ServiceStatus {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return data.Failed
	}

	pinger.Count = 1
	err = pinger.Run()
	if err != nil {
		return data.Failed
	}

	return data.Success
}

func mapToManagedService(service data.Service) *data.ManagedService {
	addr := createAddressService(service)
	maxReferences := service.MaxReferences

	return &data.ManagedService{
		Address:       addr,
		MaxReferences: maxReferences,
		References:    0,
	}
}

func createAddressService(service data.Service) string {
	addr := service.Address
	schema := extractSchema(service)

	return fmt.Sprintf("%s://%s", schema, addr)
}

func extractSchema(service data.Service) string {
	var schema string
	if service.SSL {
		schema = "https"
	} else {
		schema = "http"
	}

	return schema
}

func saveService(service *data.ManagedService) {
	services = append(services, service)
}

func RefGetBestService() *data.ManagedService {
	var bestService *data.ManagedService
	services := SvcGetService()

	for i := range services {
		service := services[i]

		if !isBest(service) {
			continue
		}

		bestService = service
		break

	}

	return bestService

}

func RefClose(service *data.ManagedService) {
	if service.References <= 0 {
		return
	}

	service.References--
}

func RefUse(service *data.ManagedService) {
	if service.References >= service.MaxReferences {
		return
	}

	service.References++

}

func isBest(service *data.ManagedService) bool {

	if service.References < service.MaxReferences {
		return true
	}

	return false

}
