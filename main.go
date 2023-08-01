package main

import (
	"context"
	"fmt"
	"log"

	crossplaneClient "github.com/crossplane/crossplane-runtime/pkg/client"
	"github.com/crossplane/crossplane/apis/stacks/v1alpha1"
	"k8s.io/client-go/rest"
)

func main() {
	// Create a Kubernetes REST client configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error creating in-cluster config: %v", err)
	}

	// Create a Crossplane client
	crossplaneClient, err := crossplaneClient.NewClient(config, 0)
	if err != nil {
		log.Fatalf("Error creating Crossplane client: %v", err)
	}

	// Create a MySQL resource claim object
	mysqlClaim := &v1alpha1.MySQLInstanceClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-mysql-claim",
			Namespace: "default",
		},
		Spec: v1alpha1.MySQLInstanceClaimSpec{
			ClassSelector: &v1alpha1.ClassSelector{
				Name: "standard-mysql-instance", // Replace with the name of the MySQL class you want to use
			},
		},
	}

	// Create or update the MySQL resource claim
	err = crossplaneClient.Create(context.Background(), mysqlClaim)
	if err != nil {
		log.Fatalf("Error creating MySQL resource claim: %v", err)
	}

	fmt.Println("MySQL resource claim created successfully.")
}

