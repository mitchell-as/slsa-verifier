package activestate

import (
	"context"

	"github.com/slsa-framework/slsa-verifier/v2/options"
	"github.com/slsa-framework/slsa-verifier/v2/register"
	"github.com/slsa-framework/slsa-verifier/v2/verifiers/utils"
)

const VerifierName = "ActiveState"

func init() {
	register.RegisterVerifier(VerifierName, ActiveStateVerifierNew())
}

type ActiveStateVerifier struct{}

func ActiveStateVerifierNew() *ActiveStateVerifier {
	return &ActiveStateVerifier{}
}

func (v *ActiveStateVerifier) IsAuthoritativeFor(builderIDName string) bool {
	return true
}

func (v *ActiveStateVerifier) VerifyArtifact(ctx context.Context,
	provenance []byte, artifactHash string,
	provenanceOpts *options.ProvenanceOpts,
	builderOpts *options.BuilderOpts,
) ([]byte, *utils.TrustedBuilderID, error) {
	return nil, nil, nil
}

func (v *ActiveStateVerifier) VerifyImage(ctx context.Context,
	provenance []byte, artifactImage string,
	provenanceOpts *options.ProvenanceOpts,
	builderOpts *options.BuilderOpts,
) ([]byte, *utils.TrustedBuilderID, error) {
	return nil, nil, nil
}

func (v *ActiveStateVerifier) VerifyNpmPackage(ctx context.Context,
	attestations []byte, tarballHash string,
	provenanceOpts *options.ProvenanceOpts,
	builderOpts *options.BuilderOpts,
) ([]byte, *utils.TrustedBuilderID, error) {
	return nil, nil, nil
}
