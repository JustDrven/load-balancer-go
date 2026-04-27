package manager

import (
	"fmt"

	"dev.justdrven/loadbalancer/data"
	"dev.justdrven/loadbalancer/engine"
	"github.com/go-ping/ping"
	"gorm.io/gorm"
)

var services []data.ManagedService

func SvcLoadServices(db gorm.DB) {

	var fetchServices []data.Service

	db.Find(&fetchServices).Where("type = ?", engine.GetApplication().ServiceType)
	for i := range fetchServices {
		dbService := fetchServices[i]

		if makeDiagnostic(dbService.Address) == data.Failed {
			continue
		}

		saveService(mapToManagedService(dbService))
		fmt.Printf("[SVC-MANAGER]  Service %d->%s\n", i+1, dbService.Address)
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
	addr := service.Address
	maxReferences := service.MaxReferences

	return &data.ManagedService{
		Address:       addr,
		MaxReferences: maxReferences,
		References:    0,
	}
}

func saveService(service *data.ManagedService) {
	services = append(services, *service)
}
