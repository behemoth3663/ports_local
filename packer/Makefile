PORTNAME=	packer
PORTVERSION=	1.15.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	brad@facefault.org
COMMENT=	Tool for creating machine images
WWW=		https://packer.io/

LICENSE=	MPL20

USES=		cpe go:modules
CPE_VENDOR=	hashicorp

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONFULL}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.GitCommit=${GIT_COMMIT:sh:M*:[1]}'"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${ECHO} ${PORTVERSION} >${WRKSRC}/version/VERSION
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/api/transport/http/internal/propagation
	@${RM} -r ${WRKSRC}/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
	@${RM} -r ${WRKSRC}/vendor/github.com/containerd/containerd/v2/pkg/tracing
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry.go
	@${RLN} ${WRKSRC}/vendor/github.com/mikelolasagasti/xz ${WRKSRC}/vendor/github.com/therootcompany/xz
	@${TEST} -z ${WRKSRC_hashicorp_go_secure_stdlib_parseutil} || ${RLN} ${WRKSRC_hashicorp_go_secure_stdlib_parseutil}/parseutil ${WRKSRC}/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
	@${MKDIR} ${WRKSRC}/vendor/github.com/hashicorp/go-getter
	@${TEST} -z ${WRKSRC_hashicorp_go_getter_v2} || ${RLN} ${WRKSRC_hashicorp_go_getter_v2} ${WRKSRC}/vendor/github.com/hashicorp/go-getter/v2
	@${MKDIR} ${WRKSRC}/vendor/github.com/hashicorp/go-getter/gcs
	@${TEST} -z ${WRKSRC_hashicorp_go_getter_v2} || ${RLN} ${WRKSRC_hashicorp_go_getter_v2}/gcs ${WRKSRC}/vendor/github.com/hashicorp/go-getter/gcs/v2
	@${MKDIR} ${WRKSRC}/vendor/github.com/hashicorp/go-getter/s3
	@${TEST} -z ${WRKSRC_hashicorp_go_getter_v2} || ${RLN} ${WRKSRC_hashicorp_go_getter_v2}/s3 ${WRKSRC}/vendor/github.com/hashicorp/go-getter/s3/v2
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc
	@${TEST} -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${MKDIR} ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${TEST} -z ${WRKSRC_charmbracelet_x_ansi} || ${RLN} ${WRKSRC_charmbracelet_x_ansi}/ansi ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${TEST} -z ${WRKSRC_charmbracelet_x_cellbuf} || ${RLN} ${WRKSRC_charmbracelet_x_cellbuf}/cellbuf ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${TEST} -z ${WRKSRC_charmbracelet_x_term} || ${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${MKDIR} ${WRKSRC}/vendor/github.com/moby/sys
	@${TEST} -z ${WRKSRC_moby_sys_userns} || ${RLN} ${WRKSRC_moby_sys_userns}/userns ${WRKSRC}/vendor/github.com/moby/sys
	@${TEST} -z ${WRKSRC_moby_sys_user} || ${RLN} ${WRKSRC_moby_sys_user}/user ${WRKSRC}/vendor/github.com/moby/sys
	@${TEST} -z ${WRKSRC_moby_sys_sequential} || ${RLN} ${WRKSRC_moby_sys_user}/sequential ${WRKSRC}/vendor/github.com/moby/sys
	@${TEST} -z ${WRKSRC_moby_sys_signal} || ${RLN} ${WRKSRC_moby_sys_signal}/signal ${WRKSRC}/vendor/github.com/moby/sys
	@${TEST} -z ${WRKSRC_moby_sys_mountinfo} || ${RLN} ${WRKSRC_moby_sys_mountinfo}/mountinfo ${WRKSRC}/vendor/github.com/moby/sys
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@${TEST} ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v3 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v3 ${WRKSRC}/vendor/go.yaml.in/yaml/v3
	@${RM} -r ${WRKSRC}/vendor/github.com/containerd/containerd/api
	@${TEST} -z ${WRKSRC_containerd_containerd_api} || ${RLN} ${WRKSRC_containerd_containerd_api}/api ${WRKSRC}/vendor/github.com/containerd/containerd

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=google GHT_PROJ=cel-spec GHT_PATH=cel.dev/expr )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=darccio GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"


#-GoogleCloudPlatform:opentelemetry-operations-go:v0.55.0:googlecloudplatform_opentelemetry_operations_go_metric/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric
#-GoogleCloudPlatform:opentelemetry-operations-go:v0.55.0:googlecloudplatform_opentelemetry_operations_go_resourcemapping/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping
#-GoogleCloudPlatform:opentelemetry-operations-go:v1.30.0:googlecloudplatform_opentelemetry_operations_go_gcp/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:opentelemetry-operations-go\:*}

#-aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

#-census-instrumentation:opencensus-proto:v0.4.1:census_instrumentation_opencensus_proto/vendor/github.com/census-instrumentation/opencensus-proto
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:*}

#-charmbracelet:x:v0.8.0:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+charmbracelet:x:ansi/v0.8.0:charmbracelet_x_ansi
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(ansi\|cellbuf\|term)|\1\3/\2|}

#-containerd:containerd:v1.10.0:containerd_containerd_api/vendor/github.com/containerd/containerd/api
#+containerd:containerd:api/v1.10.0:containerd_containerd_api/vendor/github.com/containerd/containerd
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/containerd/containerd/(api)|\1\3/\2|}

#-googleapis:google-cloud-go:v1.1.3:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.35.1:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
GH_TUPLE:=	${GH_TUPLE:N*/cloud.google.com/go/*}

#-envoyproxy:go-control-plane:v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane/envoy
#+envoyproxy:go-control-plane:envoy/v1.32.4:envoyproxy_go_control_plane/vendor/github.com/envoyproxy/go-control-plane
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/envoyproxy/go-control-plane)/(envoy)$|\1\3/\2|}

#-hashicorp:consul:    v1.25.1:hashicorp_consul_api/vendor/github.com/hashicorp/consul/api
#+hashicorp:consul:api/v1.25.1:hashicorp_consul_api/vendor/github.com/hashicorp/consul
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/consul)/(.*)|\1\3/\2|}

#-hashicorp:go-secure-stdlib:        v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib/strutil
#+hashicorp:go-secure-stdlib:strutil/v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/go-secure-stdlib)/(strutil)|\1\3/\2|1}

#-hashicorp:go-secure-stdlib:          v0.1.8:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
#+hashicorp:go-secure-stdlib:parseutil/v0.1.8:hashicorp_go_secure_stdlib_parseutil
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/hashicorp/go-secure-stdlib/(parseutil)|\1\3/\2|1}

#-hashicorp:vault:    v1.14.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault/api
#+hashicorp:vault:api/v1.14.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/vault)/(.*)|\1\3/\2|}

#-hashicorp:go-getter:v2.2.2:hashicorp_go_getter_v2/vendor/github.com/hashicorp/go-getter/gcs/v2
#-hashicorp:go-getter:v2.2.2:hashicorp_go_getter_v2_1/vendor/github.com/hashicorp/go-getter/s3/v2
#-hashicorp:go-getter:v2.2.2:hashicorp_go_getter_v2_2/vendor/github.com/hashicorp/go-getter/v2
#+hashicorp:go-getter:v2.2.2:hashicorp_go_getter_v2
GH_TUPLE:=	${GH_TUPLE:C|(:hashicorp_go_getter_v2).*/go-getter/v2$|\1|}
GH_TUPLE:=	${GH_TUPLE:N*hashicorp_go_getter_v2?*}

#-googleapis:gax-go:v2.12.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.12.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-ugorji:go:v1.2.6:ugorji_go_codec/vendor/github.com/ugorji/go/codec
#+ugorji:go:v1.2.6:ugorji_go_codec/vendor/github.com/ugorji/go
GH_TUPLE:=	${GH_TUPLE:S|/go/codec|/go|1}

#-google:go-genproto:49dd2c1f3d0b:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:d783a09b4405:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:d783a09b4405:google_go_genproto_rpc
GH_TUPLE:=	${GH_TUPLE:C|(:google_go_genproto_).*/(rpc)|\1\2|}
GH_TUPLE:=	${GH_TUPLE:N*google_go_genproto_[0-9]*}

#-jcmturner:gokrb5:v8.4.4:jcmturner_gokrb5_v8/vendor/github.com/jcmturner/gokrb5/v8
#+jcmturner:gokrb5:v8.4.4:jcmturner_gokrb5_v8/vendor/github.com/jcmturner/gokrb5
GH_TUPLE:=	${GH_TUPLE:S|/gokrb5/v8|/gokrb5|1}

#-jcmturner:rpc:v2.0.3:jcmturner_rpc_v2/vendor/github.com/jcmturner/rpc/v2
#+jcmturner:rpc:v2.0.3:jcmturner_rpc_v2/vendor/github.com/jcmturner/rpc
GH_TUPLE:=	${GH_TUPLE:S|/jcmturner/rpc/v2|/jcmturner/rpc|1}

#-moby:moby:v0.7.0:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.7.0:docker_docker_credential_helpers/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/docker-credential-helpers)$|docker:docker-credential-helpers\1|}

#-moby:moby:v1.54.0:moby_moby_api/vendor/github.com/moby/moby/api
#+moby:moby:api/v1.54.0:moby_moby_api/vendor/github.com/moby/moby
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/moby/moby/(api)|\1\3/\2|}

#-moby:sys:v0.1.0:moby_sys_userns/vendor/github.com/moby/sys/userns
#+moby:sys:userns/v0.1.0:moby_sys_userns/vendor/github.com/moby/sys
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+).*/moby/sys/(userns|user|sequential|signal|mountinfo),\1\3/\2,}

#-cncf:xds:d92e9ce0af51:cncf_xds_go/vendor/github.com/cncf/xds/go
#+cncf:xds:d92e9ce0af51:cncf_xds_go/vendor/github.com/cncf/xds
GH_TUPLE:=	${GH_TUPLE:S|/xds/go|/xds|1}

#-containerd:containerd:v1.10.0:containerd_containerd_api/vendor/github.com/containerd/containerd/api
#+containerd:containerd:api/v1.10.0:containerd_containerd_api/vendor/github.com
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/containerd/containerd/(api)|\1\3/\2|}

GH_TUPLE+=	uber-go:atomic:v1.9.0:uber_go_atomic/vendor/go.uber.org/atomic
GH_TUPLE+=	uber-go:multierr:v1.8.0:uber_go_multierr/vendor/go.uber.org/multierr

.sinclude "${.CURDIR}/generated.deps.mk"
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
