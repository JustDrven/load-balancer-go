package reference

import "dev.justdrven/loadbalancer/internal/service"

func Close(service *service.ManagedService) {
	if service.References <= 0 {
		return
	}

	service.References--
}

func Use(service *service.ManagedService) {
	if service.References >= service.MaxReferences {
		return
	}

	service.References++

}
