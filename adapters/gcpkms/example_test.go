package gcpkms_test

import (
	"context"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	"github.com/soroauth/soroauth-go"
	"github.com/soroauth/soroauth-go/adapters/gcpkms"
	"github.com/stellar/go-stellar-sdk/xdr"
)

// ExampleNewSigner shows a complete end-to-end flow for using a Cloud KMS Ed25519
// key to authorize a Soroban entry.
//
// IAM Permissions required:
// To use this signer, the Google Cloud service account running the code must have:
// - roles/cloudkms.signerVerifier (to sign payloads)
// - roles/cloudkms.viewer (optional, if you need to fetch the public key dynamically)
//
// This example compiles but does not run by default so that tests stay offline.
func ExampleNewSigner() {
	// 1. Set up the Cloud KMS client using Application Default Credentials.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		fmt.Println("failed to create KMS client:", err)
		return
	}
	defer client.Close()

	// 2. Identify the key and the Stellar account it corresponds to.
	// The keyVersion string is the full resource name of the Ed25519 key version.
	// The address is the classic G... Stellar address matching the public key.
	keyVersion := "projects/my-project/locations/global/keyRings/my-ring/cryptoKeys/my-key/cryptoKeyVersions/1"
	address := "GABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890ABCDEFGHIJKLMNOP" // Replace with actual address

	// 3. Construct the GCP KMS signer.
	// We pass the KMS client directly; gcpkms.NewSigner validates the address format.
	signer, err := gcpkms.NewSigner(address, keyVersion, client)
	if err != nil {
		fmt.Println("failed to create signer:", err)
		return
	}

	// 4. Create an unsigned Soroban authorization entry.
	// In a real flow, this comes from a transaction builder or wallet.
	entry := xdr.SorobanAuthorizationEntry{
		Credentials: xdr.SorobanCredentials{
			Type: xdr.SorobanCredentialsTypeSorobanCredentialsSourceAccount,
		},
		RootInvocation: xdr.SorobanAuthorizedInvocation{
			Function: xdr.SorobanAuthorizedFunction{
				Type: xdr.SorobanAuthorizedFunctionTypeSorobanAuthorizedFunctionContractFn,
				ContractFn: &xdr.InvokeContractArgs{
					FunctionName: xdr.ScSymbol("transfer"),
				},
			},
		},
	}

	// 5. Authorize the entry.
	// AuthorizeEntry uses the signer to fetch the signature from Cloud KMS,
	// injects the signature into the entry's credentials, and assigns the expiration.
	passphrase := "Test SDF Network ; September 2015"
	validUntilLedger := uint32(100_000)

	signed, err := soroauth.AuthorizeEntry(ctx, entry, signer, validUntilLedger, passphrase, soroauth.ForAddress(address))
	if err != nil {
		fmt.Println("failed to authorize:", err)
		return
	}

	fmt.Printf("Successfully signed entry for %s\n", signed.Credentials.AddressV2.Nonce)
}
