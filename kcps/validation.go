package kcps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func validateList(v interface{}, k string, ss []string) (errors []error) {
	found := false
	for _, z := range ss {
		if v.(string) == z {
			found = true
		}
	}
	if !found {
		errors = append(errors, fmt.Errorf("%s is invalid. It must be one of [%s]", k, strings.Join(ss, ",")))
	}
	return
}

func validateZone() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		existZones := []string{
			"jp2-east01",
			"jp2-east02",
			"jp2-east03",
			"jp2-east04",
			"jp2-west01",
		}
		return nil, validateList(v, k, existZones)
	}

}

func validateZoneId() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		existZoneIds := []string{
			"48889c46-dd04-4a50-af06-8aad7ea46f61",
			"9b295a0a-3374-4cb0-a144-88fbc1642305",
			"593697b6-c123-4025-b412-ef83822733e5",
			"c71ac6b8-498e-43e6-89ad-46944a67bce0",
			"805c4b2c-1b62-4cc8-a3b3-eadd4076ecee",
		}
		return nil, validateList(v, k, existZoneIds)
	}
}

func validateDiskOfferingId() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		//FAST_STORAGE and MIDDLE_STORAGE ID
		existDiskOfferingIds := []string{
			"bc1b5c0c-fcb3-4a7b-b8de-2c9d6952e0a5",
			"f361c03c-ed19-4228-823c-745a5569aa62",
		}
		return nil, validateList(v, k, existDiskOfferingIds)
	}
}

func validateDistributionGroup() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		existGroups := []string{
			"GROUP1",
			"GROUP2",
		}
		return nil, validateList(v, k, existGroups)
	}
}

func validateIntervalType() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		existIntervalTypes := []string{
			"DAILY",
			"WEEKLY",
			"MONTHLY",
		}
		return nil, validateList(v, k, existIntervalTypes)
	}
}

func validateAlgorithm() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		existAlgorithms := []string{
			"source",
			"roundrobin",
			"leastconn",
		}
		return nil, validateList(v, k, existAlgorithms)
	}
}

func validateTemplateFilter() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		existTemplateFilters := []string{
			"featured",
			"self",
			"selfexecutable",
			"sharedexecutable",
			"executable",
			"community",
		}
		return nil, validateList(v, k, existTemplateFilters)
	}
}

func validateProtocol(protocols []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		return nil, validateList(v, k, protocols)
	}
}

func validateNumber() schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		const (
			low  = 1
			high = 5
		)

		if v.(int) < low {
			errors = append(errors, fmt.Errorf("%d must be between %d and %d", v.(int), low, high))
		}

		if v.(int) > high {
			errors = append(errors, fmt.Errorf("%d must be between %d and %d", v.(int), low, high))
		}
		return
	}
}
