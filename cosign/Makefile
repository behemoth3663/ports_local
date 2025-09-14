PORTNAME=	cosign
PORTVERSION=	2.6.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local security

MAINTAINER=	bofh@FreeBSD.org
COMMENT=	Signing OCI containers and other artifacts using Sigstore
WWW=		https://www.sigstore.dev/

LICENSE=	APACHE20

USES=		cpe go:modules
CPE_VENDOR=	sigstore

USE_GITHUB=	yes
GH_ACCOUNT=	sigstore

GO_TARGET=	./cmd/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}^{}

GO_BUILDFLAGS=	-ldflags="-s -w -buildid= \
	-X sigs.k8s.io/release-utils/version.gitVersion=${PORTVERSION} \
	-X sigs.k8s.io/release-utils/version.gitCommit=${GIT_COMMIT:sh:M*:[1]} \
	-X sigs.k8s.io/release-utils/version.gitTreeState=clean \
	-X sigs.k8s.io/release-utils/version.buildDate=${%FT%TZ:L:gmtime}"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY .SILENT
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-Azure:azure-sdk-for-go:v1.11.0:azure_azure_sdk_for_go_azidentity/vendor/github.com/Azure/azure-sdk-for-go/sdk/azidentity
#-Azure:azure-sdk-for-go:v1.11.2:azure_azure_sdk_for_go_internal/vendor/github.com/Azure/azure-sdk-for-go/sdk/internal
#-Azure:azure-sdk-for-go:v1.18.2:azure_azure_sdk_for_go_azcore/vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore
#-Azure:azure-sdk-for-go:v1.2.0:azure_azure_sdk_for_go_internal_1/vendor/github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal
#-Azure:azure-sdk-for-go:v1.4.0:azure_azure_sdk_for_go_azkeys/vendor/github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys
# Azure:azure-sdk-for-go:v68.0.0:azure_azure_sdk_for_go/vendor/github.com/Azure/azure-sdk-for-go
#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/Azure/azure-sdk-for-go/(.*)|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/Azure/azure-sdk-for-go/*}

#Azure:go-autorest:v0.11.29:azure_go_autorest_autorest/vendor/github.com/Azure/go-autorest/autorest
#Azure:go-autorest:v0.2.1:azure_go_autorest_logger/vendor/github.com/Azure/go-autorest/logger
#Azure:go-autorest:v0.3.0:azure_go_autorest_date/vendor/github.com/Azure/go-autorest/autorest/date
#Azure:go-autorest:v0.4.6:azure_go_autorest_cli/vendor/github.com/Azure/go-autorest/autorest/azure/cli
#Azure:go-autorest:v0.5.12:azure_go_autorest_auth/vendor/github.com/Azure/go-autorest/autorest/azure/auth
#Azure:go-autorest:v0.6.0:azure_go_autorest_tracing/vendor/github.com/Azure/go-autorest/tracing
#Azure:go-autorest:v0.9.23:azure_go_autorest_adal/vendor/github.com/Azure/go-autorest/autorest/adal
#Azure:go-autorest:v14.2.0:azure_go_autorest/vendor/github.com/Azure/go-autorest
#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/Azure/go-autorest/(.*)|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/Azure/go-autorest/*}

#-GoogleCloudPlatform:grpc-gcp-go:v1.5.2:googlecloudplatform_grpc_gcp_go_grpcgcp/vendor/github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp
#+GoogleCloudPlatform:grpc-gcp-go:grpcgcp/v1.5.2:googlecloudplatform_grpc_gcp_go_grpcgcp/vendor/github.com/GoogleCloudPlatform/grpc-gcp-go
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/GoogleCloudPlatform/grpc-gcp-go)/(grpcgcp)$|\1\3/\2|}

#-GoogleCloudPlatform:opentelemetry-operations-go:v0.49.0:googlecloudplatform_opentelemetry_operations_go_metric/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric
#-GoogleCloudPlatform:opentelemetry-operations-go:v0.49.0:googlecloudplatform_opentelemetry_operations_go_resourcemapping/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping
#-GoogleCloudPlatform:opentelemetry-operations-go:v1.25.0:googlecloudplatform_opentelemetry_operations_go_gcp/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:opentelemetry-operations-go\:*}

# asaskevich:govalidator:a9d515a09cc2:asaskevich_govalidator/vendor/github.com/asaskevich/govalidator
#-aws:aws-sdk-go-v2:v1.11.20:aws_aws_sdk_go_v2_presigned_url/vendor/github.com/aws/aws-sdk-go-v2/service/internal/presigned-url
#-aws:aws-sdk-go-v2:v1.11.5:aws_aws_sdk_go_v2_accept_encoding/vendor/github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding
#-aws:aws-sdk-go-v2:v1.16.14:aws_aws_sdk_go_v2_imds/vendor/github.com/aws/aws-sdk-go-v2/feature/ec2/imds
#-aws:aws-sdk-go-v2:v1.17.18:aws_aws_sdk_go_v2_s3shared/vendor/github.com/aws/aws-sdk-go-v2/service/internal/s3shared
#-aws:aws-sdk-go-v2:v1.17.22:aws_aws_sdk_go_v2_manager/vendor/github.com/aws/aws-sdk-go-v2/feature/s3/manager
#-aws:aws-sdk-go-v2:v1.17.34:aws_aws_sdk_go_v2_credentials/vendor/github.com/aws/aws-sdk-go-v2/credentials
#-aws:aws-sdk-go-v2:v1.23.0:aws_aws_sdk_go_v2_sso/vendor/github.com/aws/aws-sdk-go-v2/service/sso
#-aws:aws-sdk-go-v2:v1.27.0:aws_aws_sdk_go_v2_ssooidc/vendor/github.com/aws/aws-sdk-go-v2/service/ssooidc
#-aws:aws-sdk-go-v2:v1.27.36:aws_aws_sdk_go_v2_config/vendor/github.com/aws/aws-sdk-go-v2/config
#-aws:aws-sdk-go-v2:v1.3.18:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
#-aws:aws-sdk-go-v2:v1.3.18:aws_aws_sdk_go_v2_v4a/vendor/github.com/aws/aws-sdk-go-v2/internal/v4a
#-aws:aws-sdk-go-v2:v1.3.20:aws_aws_sdk_go_v2_checksum/vendor/github.com/aws/aws-sdk-go-v2/service/internal/checksum
# aws:aws-sdk-go-v2:v1.31.0:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.31.0:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
#-aws:aws-sdk-go-v2:v1.35.0:aws_aws_sdk_go_v2_dynamodb/vendor/github.com/aws/aws-sdk-go-v2/service/dynamodb
#-aws:aws-sdk-go-v2:v1.35.0:aws_aws_sdk_go_v2_sqs/vendor/github.com/aws/aws-sdk-go-v2/service/sqs
#-aws:aws-sdk-go-v2:v1.36.0:aws_aws_sdk_go_v2_iam/vendor/github.com/aws/aws-sdk-go-v2/service/iam
#-aws:aws-sdk-go-v2:v1.6.5:aws_aws_sdk_go_v2_eventstream/vendor/github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream
#-aws:aws-sdk-go-v2:v1.63.0:aws_aws_sdk_go_v2_s3/vendor/github.com/aws/aws-sdk-go-v2/service/s3
#-aws:aws-sdk-go-v2:v1.8.1:aws_aws_sdk_go_v2_ini/vendor/github.com/aws/aws-sdk-go-v2/internal/ini
#-aws:aws-sdk-go-v2:v1.9.19:aws_aws_sdk_go_v2_endpoint_discovery/vendor/github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery
#-aws:aws-sdk-go-v2:v2.6.18:aws_aws_sdk_go_v2_v2/vendor/github.com/aws/aws-sdk-go-v2/internal/endpoints/v2
# aws:aws-sdk-go:v1.44.122:aws_aws_sdk_go/vendor/github.com/aws/aws-sdk-go
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

#-decred:dcrd:v4.4.0:decred_dcrd_v4/vendor/github.com/decred/dcrd/dcrec/secp256k1/v4
#+                    decred:dcrd:v4.4.0:decred_dcrd_v4/vendor/github.com/decred/dcrd/dcrec/secp256k1/v4
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)(/vendor/github.com/decred/dcrd)/(dcrec/secp256k1)/.*|\1\4/\2\3|}

#-envoyproxy:go-control-plane:v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane/envoy
#+envoyproxy:go-control-plane:envoy/v1.32.4:envoyproxy_go_control_plane/vendor/github.com/envoyproxy/go-control-plane
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/envoyproxy/go-control-plane)/(envoy)$|\1\3/\2|}

# googleapis:google-cloud-go:v0.121.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v0.16.2:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#-googleapis:google-cloud-go:v0.6.7:googleapis_google_cloud_go_longrunning/vendor/cloud.google.com/go/longrunning
#-googleapis:google-cloud-go:v1.24.2:googleapis_google_cloud_go_monitoring/vendor/cloud.google.com/go/monitoring
#-googleapis:google-cloud-go:v1.5.2:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.55.0:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
#-googleapis:google-cloud-go:v1.82.0:googleapis_google_cloud_go_spanner/vendor/cloud.google.com/go/spanner
#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/cloud.google.com/go/(storage)|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:N*/cloud.google.com/go/*}

#-hashicorp:go-secure-stdlib:v0.1.2:hashicorp_go_secure_stdlib_mlock/vendor/github.com/hashicorp/go-secure-stdlib/mlock
#-hashicorp:go-secure-stdlib:v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib/strutil
#+hashicorp:go-secure-stdlib:mlock/v0.1.2:hashicorp_go_secure_stdlib_mlock
#+hashicorp:go-secure-stdlib:strutil/v0.1.2:hashicorp_go_secure_stdlib_strutil
#-hashicorp:go-secure-stdlib:v0.1.3:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
#+hashicorp:go-secure-stdlib:parseutil/v0.1.3:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/hashicorp/go-secure-stdlib/(mlock|strutil),\1\3/\2,}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/hashicorp/go-secure-stdlib)/(.*)|\1\3/\2|}

#-moby:moby:v0.7.0:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.7.0:docker_docker_credential_helpers/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/docker-credential-helpers)$|docker:docker-credential-helpers\1|}

#-withfig:autocomplete-tools:v1.2.1:withfig_autocomplete_tools_cobra/vendor/github.com/withfig/autocomplete-tools/integrations/cobra
#+withfig:autocomplete-tools:integrations/cobra/v1.2.1:withfig_autocomplete_tools_cobra/vendor/github.com/withfig/autocomplete-tools
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/withfig/autocomplete-tools)/(integrations/cobra)$|\1\3/\2|}

#-containerd:stargz-snapshotter:v0.16.3:containerd_stargz_snapshotter_estargz/vendor/github.com/containerd/stargz-snapshotter/estargz
#+containerd:stargz-snapshotter:v0.16.3:containerd_stargz_snapshotter_estargz/vendor/github.com/containerd/stargz-snapshotter
GH_TUPLE:=	${GH_TUPLE:S|/estargz$||1}

.include <bsd.port.mk>
