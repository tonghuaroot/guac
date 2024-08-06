package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"strings"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestVEXStatement is the resolver for the ingestVEXStatement field.
func (r *mutationResolver) IngestVEXStatement(ctx context.Context, subject model.PackageOrArtifactInput, vulnerability model.IDorVulnerabilityInput, vexStatement model.VexStatementInputSpec) (string, error) {
	funcName := "IngestVEXStatement"
	if err := validatePackageOrArtifactInput(&subject, funcName); err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}
	err := validateVexInput(vexStatement)
	if err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}
	if vulnerability.VulnerabilityInput != nil {
		err = validateNoVul(*vulnerability.VulnerabilityInput)
		if err != nil {
			return "", gqlerror.Errorf("%v ::  %s", funcName, err)
		}
		err = validateVulnerabilityIDInputSpec(*vulnerability.VulnerabilityInput)
		if err != nil {
			return "", gqlerror.Errorf("%v ::  %s", funcName, err)
		}
		// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
		return r.Backend.IngestVEXStatement(ctx, subject, model.IDorVulnerabilityInput{
			VulnerabilityTypeID: vulnerability.VulnerabilityTypeID,
			VulnerabilityNodeID: vulnerability.VulnerabilityNodeID,
			VulnerabilityInput:  &model.VulnerabilityInputSpec{Type: strings.ToLower(vulnerability.VulnerabilityInput.Type), VulnerabilityID: strings.ToLower(vulnerability.VulnerabilityInput.VulnerabilityID)},
		}, vexStatement)
	} else {
		return r.Backend.IngestVEXStatement(ctx, subject, vulnerability, vexStatement)
	}
}

// IngestVEXStatements is the resolver for the ingestVEXStatements field.
func (r *mutationResolver) IngestVEXStatements(ctx context.Context, subjects model.PackageOrArtifactInputs, vulnerabilities []*model.IDorVulnerabilityInput, vexStatements []*model.VexStatementInputSpec) ([]string, error) {
	funcName := "IngestVEXStatements"
	valuesDefined := 0
	if len(subjects.Packages) > 0 {
		if len(subjects.Packages) != len(vexStatements) {
			return []string{}, gqlerror.Errorf("uneven packages and vexStatements for ingestion")
		}
		valuesDefined = valuesDefined + 1
	}
	if len(subjects.Artifacts) > 0 {
		if len(subjects.Artifacts) != len(vexStatements) {
			return []string{}, gqlerror.Errorf("uneven artifact and vexStatements for ingestion")
		}
		valuesDefined = valuesDefined + 1
	}
	if len(vulnerabilities) != len(vexStatements) {
		return []string{}, gqlerror.Errorf("uneven vulnerabilities and vexStatements for ingestion")
	}
	if valuesDefined != 1 {
		return []string{}, gqlerror.Errorf("must specify at most packages or artifacts for %v", "IngestVEXStatements")
	}

	var lowercaseVulnList []*model.IDorVulnerabilityInput
	for _, v := range vulnerabilities {
		if v.VulnerabilityInput == nil {
			lowercaseVulnList = append(lowercaseVulnList, v)
			continue
		}
		err := validateNoVul(*v.VulnerabilityInput)
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		err = validateVulnerabilityIDInputSpec(*v.VulnerabilityInput)
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		lowercaseVulnInput := model.VulnerabilityInputSpec{
			Type:            strings.ToLower(v.VulnerabilityInput.Type),
			VulnerabilityID: strings.ToLower(v.VulnerabilityInput.VulnerabilityID),
		}

		lowercaseVulnList = append(lowercaseVulnList, &model.IDorVulnerabilityInput{
			VulnerabilityTypeID: v.VulnerabilityTypeID,
			VulnerabilityNodeID: v.VulnerabilityNodeID,
			VulnerabilityInput:  &lowercaseVulnInput,
		})
	}
	return r.Backend.IngestVEXStatements(ctx, subjects, lowercaseVulnList, vexStatements)
}

// CertifyVEXStatement is the resolver for the CertifyVEXStatement field.
func (r *queryResolver) CertifyVEXStatement(ctx context.Context, certifyVEXStatementSpec model.CertifyVEXStatementSpec) ([]*model.CertifyVEXStatement, error) {
	if err := validatePackageOrArtifactQueryFilter(certifyVEXStatementSpec.Subject); err != nil {
		return nil, gqlerror.Errorf("CertifyVEXStatement :: %s", err)
	}

	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase

	if certifyVEXStatementSpec.Vulnerability != nil {
		var typeLowerCase *string = nil
		var vulnIDLowerCase *string = nil
		if certifyVEXStatementSpec.Vulnerability.Type != nil {
			lower := strings.ToLower(*certifyVEXStatementSpec.Vulnerability.Type)
			typeLowerCase = &lower
		}
		if certifyVEXStatementSpec.Vulnerability.VulnerabilityID != nil {
			lower := strings.ToLower(*certifyVEXStatementSpec.Vulnerability.VulnerabilityID)
			vulnIDLowerCase = &lower
		}

		lowercaseVulnFilter := &model.VulnerabilitySpec{
			ID:              certifyVEXStatementSpec.Vulnerability.ID,
			Type:            typeLowerCase,
			VulnerabilityID: vulnIDLowerCase,
			NoVuln:          certifyVEXStatementSpec.Vulnerability.NoVuln,
		}

		lowercaseCertifyVexFilter := &model.CertifyVEXStatementSpec{
			ID:               certifyVEXStatementSpec.ID,
			Subject:          certifyVEXStatementSpec.Subject,
			Vulnerability:    lowercaseVulnFilter,
			Status:           certifyVEXStatementSpec.Status,
			VexJustification: certifyVEXStatementSpec.VexJustification,
			Statement:        certifyVEXStatementSpec.Statement,
			StatusNotes:      certifyVEXStatementSpec.StatusNotes,
			KnownSince:       certifyVEXStatementSpec.KnownSince,
			Origin:           certifyVEXStatementSpec.Origin,
			Collector:        certifyVEXStatementSpec.Collector,
		}
		return r.Backend.CertifyVEXStatement(ctx, lowercaseCertifyVexFilter)
	} else {
		return r.Backend.CertifyVEXStatement(ctx, &certifyVEXStatementSpec)
	}
}

// CertifyVEXStatementList is the resolver for the CertifyVEXStatementList field.
func (r *queryResolver) CertifyVEXStatementList(ctx context.Context, certifyVEXStatementSpec model.CertifyVEXStatementSpec, after *string, first *int) (*model.VEXConnection, error) {
	if err := validatePackageOrArtifactQueryFilter(certifyVEXStatementSpec.Subject); err != nil {
		return nil, gqlerror.Errorf("CertifyVEXStatement :: %s", err)
	}

	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase

	if certifyVEXStatementSpec.Vulnerability != nil {
		var typeLowerCase *string = nil
		var vulnIDLowerCase *string = nil
		if certifyVEXStatementSpec.Vulnerability.Type != nil {
			lower := strings.ToLower(*certifyVEXStatementSpec.Vulnerability.Type)
			typeLowerCase = &lower
		}
		if certifyVEXStatementSpec.Vulnerability.VulnerabilityID != nil {
			lower := strings.ToLower(*certifyVEXStatementSpec.Vulnerability.VulnerabilityID)
			vulnIDLowerCase = &lower
		}

		lowercaseVulnFilter := &model.VulnerabilitySpec{
			ID:              certifyVEXStatementSpec.Vulnerability.ID,
			Type:            typeLowerCase,
			VulnerabilityID: vulnIDLowerCase,
			NoVuln:          certifyVEXStatementSpec.Vulnerability.NoVuln,
		}

		lowercaseCertifyVexFilter := model.CertifyVEXStatementSpec{
			ID:               certifyVEXStatementSpec.ID,
			Subject:          certifyVEXStatementSpec.Subject,
			Vulnerability:    lowercaseVulnFilter,
			Status:           certifyVEXStatementSpec.Status,
			VexJustification: certifyVEXStatementSpec.VexJustification,
			Statement:        certifyVEXStatementSpec.Statement,
			StatusNotes:      certifyVEXStatementSpec.StatusNotes,
			KnownSince:       certifyVEXStatementSpec.KnownSince,
			Origin:           certifyVEXStatementSpec.Origin,
			Collector:        certifyVEXStatementSpec.Collector,
		}
		return r.Backend.CertifyVEXStatementList(ctx, lowercaseCertifyVexFilter, after, first)
	} else {
		return r.Backend.CertifyVEXStatementList(ctx, certifyVEXStatementSpec, after, first)
	}
}
