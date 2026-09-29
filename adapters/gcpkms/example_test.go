package gcpkms_test

import (
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	"github.com/soroauth/soroauth-go/adapters/gcpkms"
)

func ExampleNewSigner() {
	// A nil client, to show what the constructor does with one rather than to
	// reach GCP: the example runs offline and makes no network call. In real
	// code this is a *kms.KeyManagementClient from kms.NewKeyManagementClient.
	//
	// Note this is a *typed* nil. NewSigner refuses it, which is the point —
	// accepting it would hand back a signer that panics later, inside Sign.
	var client *kms.KeyManagementClient
	signer, err := gcpkms.NewSigner("GAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAWHF", "keyVersion", client)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("signer address: %s\n", signer.Address())

	// Output: error: gcpkms: new signer: GAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAWHF: no signer for address
}
