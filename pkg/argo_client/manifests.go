package argo_client

import (
	"context"
	"fmt"
	"time"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/settings"
	argoappv1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	repoapiclient "github.com/argoproj/argo-cd/v2/reposerver/apiclient"
	"github.com/argoproj/argo-cd/v2/reposerver/repository"
	"github.com/argoproj/argo-cd/v2/util/git"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/zapier/kubechecks/telemetry"
)

func (a *ArgoClient) GetManifestsLocal(ctx context.Context, name, tempRepoDir, changedAppFilePath string, app argoappv1.Application) ([]string, error) {
	var err error

	ctx, span := tracer.Start(ctx, "GetManifestsLocal")
	defer span.End()

	log.Debug().Str("name", name).Msg("GetManifestsLocal")

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		getManifestsDuration.WithLabelValues(name).Observe(duration.Seconds())
	}()

	clusterCloser, clusterClient := a.GetClusterClient()
	defer clusterCloser.Close()

	settingsCloser, settingsClient := a.GetSettingsClient()
	defer settingsCloser.Close()

	log.Debug().
		Str("clusterName", app.Spec.Destination.Name).
		Str("clusterServer", app.Spec.Destination.Server).
		Msg("getting cluster")
	cluster, err := clusterClient.Get(ctx, &cluster.ClusterQuery{Name: app.Spec.Destination.Name, Server: app.Spec.Destination.Server})
	if err != nil {
		telemetry.SetError(span, err, "Argo Get Cluster")
		getManifestsFailed.WithLabelValues(name).Inc()
		return nil, errors.Wrap(err, "failed to get cluster")
	}

	argoSettings, err := settingsClient.Get(ctx, &settings.SettingsQuery{})
	if err != nil {
		telemetry.SetError(span, err, "Argo Get Settings")
		getManifestsFailed.WithLabelValues(name).Inc()
		return nil, errors.Wrap(err, "failed to get settings")
	}

	log.Debug().Str("name", name).Msg("generating diff for application...")
	res, err := a.generateManifests(ctx, fmt.Sprintf("%s/%s", tempRepoDir, changedAppFilePath), tempRepoDir, app, argoSettings, cluster)
	if err != nil {
		telemetry.SetError(span, err, "Generate Manifests")
		return nil, errors.Wrap(err, "failed to generate manifests")
	}

	if res.Manifests == nil {
		return nil, nil
	}
	getManifestsSuccess.WithLabelValues(name).Inc()
	return res.Manifests, nil
}

type repoRef struct {
	// revision is the git revision - can be any valid revision like a branch, tag, or commit SHA.
	revision string
	// commitSHA is the actual commit to which revision refers.
	commitSHA string
	// key is the name of the key which was used to reference this repo.
	key string
}

func (a *ArgoClient) generateManifests(
	ctx context.Context, appPath, tempRepoDir string, app argoappv1.Application, argoSettings *settings.Settings, cluster *argoappv1.Cluster,
) (*repoapiclient.ManifestResponse, error) {
	a.manifestsLock.Lock()
	defer a.manifestsLock.Unlock()

	source := app.Spec.GetSource()

	var projectSourceRepos []string
	var helmRepos []*argoappv1.Repository
	var helmCreds []*argoappv1.RepoCreds
	var enableGenerateManifests map[string]bool
	var helmOptions *argoappv1.HelmOptions
	var refSources map[string]*argoappv1.RefTarget

	q := repoapiclient.ManifestRequest{
		Repo:               &argoappv1.Repository{Repo: source.RepoURL},
		Revision:           source.TargetRevision,
		AppLabelKey:        argoSettings.AppLabelKey,
		AppName:            app.Name,
		Namespace:          app.Spec.Destination.Namespace,
		ApplicationSource:  &source,
		Repos:              helmRepos,
		KustomizeOptions:   argoSettings.KustomizeOptions,
		KubeVersion:        cluster.Info.ServerVersion,
		ApiVersions:        cluster.Info.APIVersions,
		HelmRepoCreds:      helmCreds,
		TrackingMethod:     argoSettings.TrackingMethod,
		EnabledSourceTypes: enableGenerateManifests,
		HelmOptions:        helmOptions,
		HasMultipleSources: app.Spec.HasMultipleSources(),
		RefSources:         refSources,
		ProjectSourceRepos: projectSourceRepos,
		ProjectName:        app.Spec.Project,
	}

	return repository.GenerateManifests(
		ctx,
		appPath,
		tempRepoDir,
		source.TargetRevision,
		&q,
		true,
		new(git.NoopCredsStore),
		resource.MustParse("0"),
		nil,
	)
}

func ConvertJsonToYamlManifests(jsonManifests []string) []string {
	var manifests []string
	for _, manifest := range jsonManifests {
		ret, err := yaml.JSONToYAML([]byte(manifest))
		if err != nil {
			log.Warn().Err(err).Msg("Failed to format manifest")
			continue
		}
		manifests = append(manifests, fmt.Sprintf("---\n%s", string(ret)))
	}
	return manifests
}
