PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.79.2
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Terragrunt is a thin wrapper for Terraform
WWW=		https://terragrunt.gruntwork.io/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/go-commons/version.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

_OPTIONS_extract=	600:pre-post-extract
pre-post-extract: .PHONY
#	github.com/gruntwork-io/terragrunt/tf/getproviders/mocks: no required module provides package github.com/gruntwork-io/terragrunt/tf/getproviders/mocks
	${RM} ${WRKSRC}/tf/getproviders/lock_test.go
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=darccio GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=FiloSottile GHT_PROJ=age GHT_PATH=filippo.io/age )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=FiloSottile GHT_PROJ=edwards25519 GHT_PATH=filippo.io/edwards25519 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=google GHT_PROJ=cel-spec GHT_PATH=cel.dev/expr )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-Azure:azure-sdk-for-go:v1.13.0:azure_azure_sdk_for_go_azcore/vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore
#+Azure:azure-sdk-for-go:sdk/azcore/v1.13.0:azure_azure_sdk_for_go_azcore/vendor/github.com/Azure/azure-sdk-for-go
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/Azure/azure-sdk-for-go)/(sdk/azcore)|\1\3/\2|1}

#-Azure:azure-sdk-for-go:v1.0.1                               :azure_azure_sdk_for_go_internal_1/vendor/github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal
#+Azure:azure-sdk-for-go:sdk/security/keyvault/internal/v1.0.1:azure_azure_sdk_for_go_internal_1
#-Azure:azure-sdk-for-go:v1.1.0                             :azure_azure_sdk_for_go_azkeys/vendor/github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys
#+Azure:azure-sdk-for-go:sdk/security/keyvault/azkeys/v1.1.0:azure_azure_sdk_for_go_azkeys
#-Azure:azure-sdk-for-go:v1.10.0              :azure_azure_sdk_for_go_internal/vendor/github.com/Azure/azure-sdk-for-go/sdk/internal
#+Azure:azure-sdk-for-go:sdk/internal/v1.10.0 :azure_azure_sdk_for_go_internal
#-Azure:azure-sdk-for-go:v1.7.0               :azure_azure_sdk_for_go_azidentity/vendor/github.com/Azure/azure-sdk-for-go/sdk/azidentity
#+Azure:azure-sdk-for-go:sdk/azidentity/v1.7.0:azure_azure_sdk_for_go_azidentity
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/Azure/azure-sdk-for-go/(.*)|\1\3/\2|}

#-aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

#-GoogleCloudPlatform:opentelemetry-operations-go:v0.49.0:googlecloudplatform_opentelemetry_operations_go_metric/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric
#-GoogleCloudPlatform:opentelemetry-operations-go:v0.49.0:googlecloudplatform_opentelemetry_operations_go_resourcemapping/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping
#-GoogleCloudPlatform:opentelemetry-operations-go:v1.25.0:googlecloudplatform_opentelemetry_operations_go_gcp/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:opentelemetry-operations-go\:*}

#-census-instrumentation:opencensus-proto:v0.4.1:census_instrumentation_opencensus_proto/vendor/github.com/census-instrumentation/opencensus-proto
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:opencensus-proto\:*}

#-envoyproxy:go-control-plane:v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane/envoy
#+envoyproxy:go-control-plane:envoy/v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+/vendor/github.com/envoyproxy/go-control-plane)/(envoy)|\1\3/\2|}

#-googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-spiffe:go-spiffe:v2.5.0:spiffe_go_spiffe_v2/vendor/github.com/spiffe/go-spiffe/v2
#+spiffe:go-spiffe:v2.5.0:spiffe_go_spiffe_v2/vendor/github.com/spiffe/go-spiffe
GH_TUPLE:=	${GH_TUPLE:S|/go-spiffe/v2|/go-spiffe|1}

#-google:go-genproto:e6fa225c2576:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:e6fa225c2576:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
GH_TUPLE:=	${GH_TUPLE:N*\:google_go_genproto_*}

# googleapis:google-cloud-go:v0.118.0:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v0.6.1:googleapis_google_cloud_go_longrunning/vendor/cloud.google.com/go/longrunning
#-googleapis:google-cloud-go:v1.2.1:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.20.0:googleapis_google_cloud_go_kms/vendor/cloud.google.com/go/kms
#-googleapis:google-cloud-go:v1.21.1:googleapis_google_cloud_go_monitoring/vendor/cloud.google.com/go/monitoring
#-googleapis:google-cloud-go:v1.47.0:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
#-googleapis:google-cloud-go:v0.13.0:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+googleapis:google-cloud-go:auth/v0.13.0:googleapis_google_cloud_go_auth
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(auth)|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:Ngoogleapis\:google-cloud-go\:*/cloud.google.com/go/*}

#-hashicorp:go-secure-stdlib:        v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib/strutil
#+hashicorp:go-secure-stdlib:strutil/v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/go-secure-stdlib)/(strutil)|\1\3/\2|1}

#-hashicorp:go-secure-stdlib:          v0.1.8:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
#+hashicorp:go-secure-stdlib:parseutil/v0.1.8:hashicorp_go_secure_stdlib_parseutil
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/hashicorp/go-secure-stdlib/(parseutil)|\1\3/\2|1}

#-hashicorp:vault:v1.4.2:hashicorp_vault/vendor/github.com/hashicorp/vault
#-hashicorp:vault:v1.14.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault/api
#+hashicorp:vault:api/v1.14.0:hashicorp_vault/vendor/github.com/hashicorp/vault
GH_TUPLE:=	${GH_TUPLE:Nhashicorp\:vault\:*/hashicorp/vault}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/vault)/(api)|\1\3/\2|1}

#-cncf:xds:b4127c9b8d78:cncf_xds_go/vendor/github.com/cncf/xds/go
#+cncf:xds:b4127c9b8d78:cncf_xds_go/vendor/github.com/cncf/xds
GH_TUPLE:=	${GH_TUPLE:S|/xds/go|/xds|1}

#-charmbracelet:x:     v0.1.4:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+charmbracelet:x:ansi/v0.1.4:charmbracelet_x_ansi
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(ansi\|cellbuf\|term)|\1\3/\2|}

.sinclude "${.CURDIR}/generated.deps.mk"
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
