PORTNAME=	terragrunt-atlantis-config
PORTVERSION=	1.21.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Generate Atlantis config for Terragrunt projects
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	transcend-io

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.VERSION=${PORTVERSION}'"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

_OPTIONS_extract=	600:pre-post-extract
pre-post-extract: .PHONY
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=FiloSottile GHT_PROJ=age GHT_PATH=filippo.io/age )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=mozilla GHT_PROJ=sops GHT_PATH=go.mozilla.org/sops/v3 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

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

GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/Azure/go-autorest/(.*)|\1\3/\2|}

#-aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

#-GoogleCloudPlatform:opentelemetry-operations-go:v0.49.0:googlecloudplatform_opentelemetry_operations_go_metric/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric
#-GoogleCloudPlatform:opentelemetry-operations-go:v0.49.0:googlecloudplatform_opentelemetry_operations_go_resourcemapping/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping
#-GoogleCloudPlatform:opentelemetry-operations-go:v1.25.0:googlecloudplatform_opentelemetry_operations_go_gcp/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:opentelemetry-operations-go\:*}

#-google:go-genproto:49dd2c1f3d0b:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:49dd2c1f3d0b:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
GH_TUPLE:=	${GH_TUPLE:N*/google.golang.org/genproto/*}

#-googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2$|/gax-go|1}

GH_TUPLE:=	${GH_TUPLE:N*/cloud.google.com/go/*}

#-hashicorp:go-secure-stdlib:v0.1.2:hashicorp_go_secure_stdlib_mlock/vendor/github.com/hashicorp/go-secure-stdlib/mlock
#-hashicorp:go-secure-stdlib:v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib/strutil
#+hashicorp:go-secure-stdlib:mlock/v0.1.2:hashicorp_go_secure_stdlib_mlock
#+hashicorp:go-secure-stdlib:strutil/v0.1.2:hashicorp_go_secure_stdlib_strutil
#-hashicorp:go-secure-stdlib:v0.1.3:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
#+hashicorp:go-secure-stdlib:parseutil/v0.1.3:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/hashicorp/go-secure-stdlib/(mlock|strutil),\1\3/\2,}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/hashicorp/go-secure-stdlib)/(.*)|\1\3/\2|}

#-hashicorp:vault:v1.5.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault/api
#+hashicorp:vault:v1.5.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault
GH_TUPLE:=	${GH_TUPLE:S|/vault/api$|/vault|1}

.sinclude "${.CURDIR}/generated.deps.mk"
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
