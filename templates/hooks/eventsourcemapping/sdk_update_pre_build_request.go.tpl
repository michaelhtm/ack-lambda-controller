	if desired.ko.Spec.DestinationConfig != nil {
		if desired.ko.Spec.DestinationConfig.OnFailure != nil && desired.ko.Spec.DestinationConfig.OnFailure.Destination == nil {
			desired.ko.Spec.DestinationConfig.OnFailure = nil
		}
		if desired.ko.Spec.DestinationConfig.OnSuccess != nil && desired.ko.Spec.DestinationConfig.OnSuccess.Destination == nil {
			desired.ko.Spec.DestinationConfig.OnSuccess = nil
		}
		if desired.ko.Spec.DestinationConfig.OnFailure == nil && desired.ko.Spec.DestinationConfig.OnSuccess == nil {
			desired.ko.Spec.DestinationConfig = nil
		}
	}
