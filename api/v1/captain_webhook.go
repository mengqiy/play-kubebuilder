package v1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// +kubebuilder:webhook:failurePolicy=fail,groups=crew.testproject.org,resources=captains,verbs=create;update,versions=v1,name=vcaptain.kb.io,path=/validate-crew-testproject-org-v1-captain,mutating=false

var _ webhook.Validator = &Captain{}

// ValidateCreate implements webhookutil.validator so a webhook will be registered for the type
func (c *Captain) ValidateCreate() error {
	//log.Info("validate create", "name", c.Name)

	if c.Spec.NextStop < 10 {
		return fmt.Errorf(".spec.nextStop must >= 10")
	}
	return nil
}

// ValidateUpdate implements webhookutil.validator so a webhook will be registered for the type
func (c *Captain) ValidateUpdate(old runtime.Object) error {
	//log.Info("validate update", "name", c.Name)

	if c.Spec.NextStop < 10 {
		return fmt.Errorf(".spec.nextStop must >= 10")
	}

	oldC, ok := old.(*Captain)
	if !ok {
		return fmt.Errorf("expect old object to be a %T instead of %T", oldC, old)
	}
	if oldC.Spec.NextStop < 10 {
		return fmt.Errorf("it is not allowed to delay.spec.nextStop for more than 1 hour")
	}
	return nil
}

// +kubebuilder:webhook:failurePolicy=fail,groups=crew.testproject.org,resources=captains,verbs=create;update,versions=v1,name=mcaptain.kb.io,path=/mutate-crew-testproject-org-v1-captain,mutating=true

var _ webhook.Defaulter = &Captain{}

// Default implements webhookutil.defaulter so a webhook will be registered for the type
func (c *Captain) Default() {
	//log.Info("default", "name", c.Name)

	if c.Annotations == nil {
		c.Annotations = map[string]string{}
	}
	c.Annotations["test-kb-v2"] = "hello"

	if c.Spec.NextStop < 12 {
		c.Spec.NextStop = 15
	}
}
