PORTNAME=	cosign
PORTVERSION=	3.0.3
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	bofh@FreeBSD.org
COMMENT=	Signing OCI containers and other artifacts using Sigstore
WWW=		https://www.sigstore.dev/

LICENSE=	APACHE20

USES=		cpe go:modules
CPE_VENDOR=	sigstore

USE_GITHUB=	yes
GH_ACCOUNT=	sigstore

USE_GITLAB=	nodefault

GO_TARGET=	./cmd/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}^{}

CGO_ENABLED=	0
GO_BUILDFLAGS=	-ldflags="-s -w -buildid= \
	-X sigs.k8s.io/release-utils/version.gitVersion=${PORTVERSION} \
	-X sigs.k8s.io/release-utils/version.gitCommit=${GIT_COMMIT:sh:M*:[1]} \
	-X sigs.k8s.io/release-utils/version.gitTreeState=clean \
	-X sigs.k8s.io/release-utils/version.buildDate=${%FT%TZ:L:gmtime}"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY .SILENT
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/github.com/Azure/go-autorest/autorest/adal
	@${TEST} -z ${WRKSRC_azure_go_autorest_adal} || ${RLN} ${WRKSRC_azure_go_autorest_adal}/autorest/adal ${WRKSRC}/vendor/github.com/Azure/go-autorest/autorest/adal
	@${RM} -r ${WRKSRC}/vendor/github.com/Azure/go-autorest/logger
	@${TEST} -z ${WRKSRC_azure_go_autorest_logger} || ${RLN} ${WRKSRC_azure_go_autorest_logger}/logger ${WRKSRC}/vendor/github.com/Azure/go-autorest
	@${MKDIR} ${WRKSRC}/vendor/cuelang.org
	@${TEST} -z ${WRKSRC_cue_lang_cue} || ${RLN} ${WRKSRC_cue_lang_cue} ${WRKSRC}/vendor/cuelang.org/go
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@${TEST} ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v2 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v2 ${WRKSRC}/vendor/go.yaml.in/yaml/v2
	@${TEST} ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v3 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v3 ${WRKSRC}/vendor/go.yaml.in/yaml/v3
	@${MKDIR} ${WRKSRC}/vendor/gopkg.in
	@${TEST} ! -e ${WRKSRC}/vendor/github.com/go-ini/ini || ${RLN} ${WRKSRC}/vendor/github.com/go-ini/ini ${WRKSRC}/vendor/gopkg.in/ini.v1
	@${MKDIR} ${WRKSRC}/vendor/github.com/Azure/azure-sdk-for-go/sdk/security/keyvault
	@${TEST} -z ${WRKSRC_azure_azure_sdk_for_go_azcore} || ${RLN} ${WRKSRC_azure_azure_sdk_for_go_azcore}/sdk/azcore ${WRKSRC}/vendor/github.com/Azure/azure-sdk-for-go/sdk
	@${TEST} -z ${WRKSRC_azure_azure_sdk_for_go_internal} || ${RLN} ${WRKSRC_azure_azure_sdk_for_go_internal}/sdk/internal ${WRKSRC}/vendor/github.com/Azure/azure-sdk-for-go/sdk
	@${TEST} -z ${WRKSRC_azure_azure_sdk_for_go_azidentity} || ${RLN} ${WRKSRC_azure_azure_sdk_for_go_azidentity}/sdk/azidentity ${WRKSRC}/vendor/github.com/Azure/azure-sdk-for-go/sdk
	@${TEST} -z ${WRKSRC_azure_azure_sdk_for_go_internal_1} || ${RLN} ${WRKSRC_azure_azure_sdk_for_go_internal_1}/sdk/security/keyvault/internal ${WRKSRC}/vendor/github.com/Azure/azure-sdk-for-go/sdk/security/keyvault
	@${TEST} -z ${WRKSRC_azure_azure_sdk_for_go_azkeys} || ${RLN} ${WRKSRC_azure_azure_sdk_for_go_azkeys}/sdk/security/keyvault/azkeys ${WRKSRC}/vendor/github.com/Azure/azure-sdk-for-go/sdk/security/keyvault
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/api
	@test -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc
	@test -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry_test.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentracing.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentracing_test.go

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=json GHT_PATH=sigs.k8s.io/json )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=randfill GHT_PATH=sigs.k8s.io/randfill )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=release-utils GHT_PATH=sigs.k8s.io/release-utils )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=structured-merge-diff GHT_PATH=sigs.k8s.io/structured-merge-diff/v6 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=go-yaml GHT_PROJ=yaml GHT_PATH=go.yaml.in/yaml/v2 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=cue-lang GHT_PROJ=cue GHT_PATH=cuelang.org/go )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=cue-labs GHT_PROJ=oci GHT_PATH=cuelabs.dev/go/oci/ociregistry )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-AliyunContainerService:ack-ram-tool:v0.14.0:aliyuncontainerservice_ack_ram_tool_provider/vendor/github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider
#+AliyunContainerService:ack-ram-tool:pkg/credentials/provider/v0.14.0:aliyuncontainerservice_ack_ram_tool_provider/vendor/github.com/AliyunContainerService/ack-ram-tool
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/AliyunContainerService/ack-ram-tool)/(pkg/credentials/provider)$|\1\3/\2|}

#-Azure:azure-sdk-for-go:v1.13.0:azure_azure_sdk_for_go_azcore/vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore
#+Azure:azure-sdk-for-go:sdk/azcore/v1.13.0:azure_azure_sdk_for_go_azcore
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^/]+).*/Azure/azure-sdk-for-go/(sdk/.*)|\1\3/\2|}

#-Azure:go-autorest:v0.9.20:azure_go_autorest_adal/vendor/github.com/Azure/go-autorest/autorest/adal
#+Azure:go-autorest:autorest/adal/v0.9.20:azure_go_autorest_adal
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/Azure/go-autorest/(autorest/adal|logger),\1\3/\2,}

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

# go-openapi:swag:v0.25.1:go_openapi_swag/vendor/github.com/go-openapi/swag
#-go-openapi:swag:v0.25.1:go_openapi_swag_cmdutils/vendor/github.com/go-openapi/swag/cmdutils
#-go-openapi:swag:v0.25.1:go_openapi_swag_conv/vendor/github.com/go-openapi/swag/conv
#-go-openapi:swag:v0.25.1:go_openapi_swag_fileutils/vendor/github.com/go-openapi/swag/fileutils
#-go-openapi:swag:v0.25.1:go_openapi_swag_jsonname/vendor/github.com/go-openapi/swag/jsonname
#-go-openapi:swag:v0.25.1:go_openapi_swag_jsonutils/vendor/github.com/go-openapi/swag/jsonutils
#-go-openapi:swag:v0.25.1:go_openapi_swag_loading/vendor/github.com/go-openapi/swag/loading
#-go-openapi:swag:v0.25.1:go_openapi_swag_mangling/vendor/github.com/go-openapi/swag/mangling
#-go-openapi:swag:v0.25.1:go_openapi_swag_netutils/vendor/github.com/go-openapi/swag/netutils
#-go-openapi:swag:v0.25.1:go_openapi_swag_stringutils/vendor/github.com/go-openapi/swag/stringutils
#-go-openapi:swag:v0.25.1:go_openapi_swag_typeutils/vendor/github.com/go-openapi/swag/typeutils
#-go-openapi:swag:v0.25.1:go_openapi_swag_yamlutils/vendor/github.com/go-openapi/swag/yamlutils
GH_TUPLE:=	${GH_TUPLE:N*/go-openapi/swag/*}

#-hashicorp:vault:v1.14.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault/api
#+hashicorp:vault:api/v1.14.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/vault)/(.*)|\1\3/\2|}

# sigstore:sigstore:181c5d3339b3:sigstore_sigstore/vendor/github.com/sigstore/sigstore
#-sigstore:sigstore:181c5d3339b3:sigstore_sigstore_gcp/vendor/github.com/sigstore/sigstore/pkg/signature/kms/gcp
#-sigstore:sigstore:v1.9.5:sigstore_sigstore_aws/vendor/github.com/sigstore/sigstore/pkg/signature/kms/aws
#-sigstore:sigstore:v1.9.5:sigstore_sigstore_azure/vendor/github.com/sigstore/sigstore/pkg/signature/kms/azure
#-sigstore:sigstore:v1.9.5:sigstore_sigstore_hashivault/vendor/github.com/sigstore/sigstore/pkg/signature/kms/hashivault
GH_TUPLE:=	${GH_TUPLE:N*/sigstore/sigstore/*}

#-google:go-genproto:513f23925822:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:513f23925822:google_go_genproto_api
#-google:go-genproto:513f23925822:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:513f23925822:google_go_genproto_rpc
# google:go-genproto:f936aa4a68b2:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:C,_./vendor/google\.golang\.org/genproto/googleapis/(api|rpc),_\1,}

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

#-awslabs:amazon-ecr-credential-helper:v0.10.1:awslabs_amazon_ecr_credential_helper_ecr_login/vendor/github.com/awslabs/amazon-ecr-credential-helper/ecr-login
#+awslabs:amazon-ecr-credential-helper:v0.10.1:awslabs_amazon_ecr_credential_helper_ecr_login/vendor/github.com/awslabs/amazon-ecr-credential-helper
GH_TUPLE:=	${GH_TUPLE:S|/ecr-login$||1}

#-AliyunContainerService:ack-ram-tool:v0.14.0:aliyuncontainerservice_ack_ram_tool_provider/vendor/github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider
#+AliyunContainerService:ack-ram-tool:v0.14.0:aliyuncontainerservice_ack_ram_tool_provider/vendor/github.com/AliyunContainerService/ack-ram-tool
GH_TUPLE:=	${GH_TUPLE:S|/pkg/credentials/provider$||1}

#-googleapis:gax-go:v2.15.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.15.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2$|/gax-go|1}

#-envoyproxy:go-control-plane:v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane/envoy
#+envoyproxy:go-control-plane:envoy/v1.32.4:envoyproxy_go_control_plane/vendor/github.com/envoyproxy/go-control-plane
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/envoyproxy/go-control-plane)/(envoy)$|\1\3/\2|}

#-cncf:xds:d92e9ce0af51:cncf_xds_go/vendor/github.com/cncf/xds/go
#+cncf:xds:d92e9ce0af51:cncf_xds_go/vendor/github.com/cncf/xds
GH_TUPLE:=	${GH_TUPLE:S|/xds/go|/xds|1}

# go-ini:ini:v1.67.0:go_ini_ini/vendor/github.com/go-ini/ini
#-go-ini:ini:v1.67.0:go_ini_ini_1/vendor/gopkg.in/ini.v1
GH_TUPLE:=	${GH_TUPLE:N*/vendor/gopkg.in/ini.v1}

.sinclude "${.CURDIR}/generated.deps.mk"

#-go-yaml:yaml:v2.4.2:go_yaml_yaml/vendor/go.yaml.in/yaml/v2
#+go-yaml:yaml:v2.4.0:go_yaml_yaml/vendor/gopkg.in/yaml.v2
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)[^:]+:([^/]+)/vendor/go.yaml.in/yaml/v2|\1v2.4.0:\2_2/vendor/gopkg.in/yaml.v2|1}

#-cue-lang:cue:v0.14.1:cue_lang_cue/vendor/cue-lang/cue
#+cue-lang:cue:v0.14.1:cue_lang_cue
GH_TUPLE:=	${GH_TUPLE:S|/vendor/cue-lang/cue$||1}

#-cue-labs:oci:49cab49c8e9d:cue_labs_oci/vendor/cuelabs.dev/go/oci/ociregistry
#+cue-labs:oci:49cab49c8e9d:cue_labs_oci/vendor/cuelabs.dev/go/oci
GH_TUPLE:=	${GH_TUPLE:S|/ociregistry$||1}

.include <bsd.port.mk>
