# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	tflint
PORTVERSION=	0.60.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Terraform Linter
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export		M2T_OFFLINE

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} ${WRKSRC}/vendor/cloud.google.com/go/storage/grpc_metrics.go
	@${RM} ${WRKSRC}/vendor/github.com/envoyproxy/go-control-plane/envoy/config/trace/v3/opencensus.pb.go
	@${RM} ${WRKSRC}/vendor/github.com/envoyproxy/go-control-plane/envoy/config/trace/v3/opencensus.pb.validate.go
	@${RM} ${WRKSRC}/vendor/github.com/envoyproxy/go-control-plane/envoy/config/trace/v3/trace.pb.go
	@${RM} ${WRKSRC}/vendor/cloud.google.com/go/storage/trace.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/client/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/internal/stream/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/storage/internal/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/storage/gcp/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/storage/gcp/antispam/otel.go
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/auth
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/internal/trace
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/stats/opentelemetry
	@test ! -d ${WRKSRC}/vendor/cloud.google.com/go || ${RLN} ${WRKSRC_googleapis_google_cloud_go_auth}/auth ${WRKSRC}/vendor/cloud.google.com/go
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/compute/metadata
	@test ! -d ${WRKSRC}/vendor/cloud.google.com/go/compute || ${RLN} ${WRKSRC_googleapis_google_cloud_go_compute_metadata}/compute/metadata ${WRKSRC}/vendor/cloud.google.com/go/compute
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/api
	@test -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc
	@test -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=google GHT_PROJ=cel-spec GHT_PATH=cel.dev/expr )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-GoogleCloudPlatform:opentelemetry-operations-go:d637680bde0c:googlecloudplatform_opentelemetry_operations_go_gcp
#-census-instrumentation:opencensus-go:v0.24.0:census_instrumentation_opencensus_go/vendor/go.opencensus.io
#-open-telemetry:opentelemetry-go-contrib:v0.54.0:open_telemetry_opentelemetry_go_contrib_1/vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc
#-opentracing:opentracing-go:v1.2.0:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:opentelemetry-operations-go\:*}
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:*}
GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:*}
GH_TUPLE:=	${GH_TUPLE:Nopentracing\:*}
GH_TUPLE:=	${GH_TUPLE:Ngolang\:telemetry\:*}

#-GoogleCloudPlatform:grpc-gcp-go:v1.5.3:googlecloudplatform_grpc_gcp_go_grpcgcp/vendor/github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp
#+GoogleCloudPlatform:grpc-gcp-go:grpcgcp/v1.5.3:googlecloudplatform_grpc_gcp_go_grpcgcp/vendor/github.com/GoogleCloudPlatform/grpc-gcp-go
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/GoogleCloudPlatform/grpc-gcp-go)/(grpcgcp)$|\1\3/\2|}

#-cncf:xds:b4127c9b8d78:cncf_xds_go/vendor/github.com/cncf/xds/go
#+cncf:xds:b4127c9b8d78:cncf_xds_go/vendor/github.com/cncf/xds
GH_TUPLE:=	${GH_TUPLE:S|/xds/go|/xds|1}

#-googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-owenrumney:go-sarif:v2.3.3:owenrumney_go_sarif_v2/vendor/github.com/owenrumney/go-sarif/v2
#+owenrumney:go-sarif:v2.3.3:owenrumney_go_sarif_v2/vendor/github.com/owenrumney/go-sarif
GH_TUPLE:=	${GH_TUPLE:S|/go-sarif/v2|/go-sarif|1}

#-spiffe:go-spiffe:v2.5.0:spiffe_go_spiffe_v2/vendor/github.com/spiffe/go-spiffe/v2
#+spiffe:go-spiffe:v2.5.0:spiffe_go_spiffe_v2/vendor/github.com/spiffe/go-spiffe/v2
GH_TUPLE:=	${GH_TUPLE:S|/go-spiffe/v2|/go-spiffe|1}

#-envoyproxy:go-control-plane:v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane/envoy
#+envoyproxy:go-control-plane:envoy/v1.32.4:envoyproxy_go_control_plane/vendor/github.com/envoyproxy/go-control-plane
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/envoyproxy/go-control-plane)/(envoy)$|\1\3/\2|}

#-google:go-genproto:69823020774d:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:69823020774d:google_go_genproto_2
# google:go-genproto:e639e219e697:google_go_genproto/vendor/google.golang.org/genproto
#-google:go-genproto:e6fa225c2576:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:e6fa225c2576:google_go_genproto_1
GH_TUPLE:=	${GH_TUPLE:C|_[[:digit:]]/vendor/google.golang.org/genproto/googleapis/|_|}

#make -V GH_TUPLE | tr ' ' '\n' | i googleapis:google-cloud-go
#	googleapis:google-cloud-go:v0.112.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-	googleapis:google-cloud-go:v0.13.0:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+	googleapis:google-cloud-go:auth/v0.13.0:googleapis_google_cloud_go_auth
#-	googleapis:google-cloud-go:v0.6.0:googleapis_google_cloud_go_2/vendor/cloud.google.com/go/compute/metadata
#+	googleapis:google-cloud-go:compute/metadata/v0.6.0:googleapis_google_cloud_go_compute_metadata
GH_TUPLE:=	${GH_TUPLE:N*\:googleapis_google_cloud_go_[b-z]*}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(auth)$|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/cloud.google.com/go/*}

#-go-openapi:swag:v0.24.0:go_openapi_swag_cmdutils/vendor/github.com/go-openapi/swag/cmdutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_conv/vendor/github.com/go-openapi/swag/conv
#-go-openapi:swag:v0.24.0:go_openapi_swag_fileutils/vendor/github.com/go-openapi/swag/fileutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_jsonname/vendor/github.com/go-openapi/swag/jsonname
#-go-openapi:swag:v0.24.0:go_openapi_swag_jsonutils/vendor/github.com/go-openapi/swag/jsonutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_loading/vendor/github.com/go-openapi/swag/loading
#-go-openapi:swag:v0.24.0:go_openapi_swag_mangling/vendor/github.com/go-openapi/swag/mangling
#-go-openapi:swag:v0.24.0:go_openapi_swag_netutils/vendor/github.com/go-openapi/swag/netutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_stringutils/vendor/github.com/go-openapi/swag/stringutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_typeutils/vendor/github.com/go-openapi/swag/typeutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_yamlutils/vendor/github.com/go-openapi/swag/yamlutils
# go-openapi:swag:v0.24.1:go_openapi_swag/vendor/github.com/go-openapi/swag
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/go-openapi/swag/*}

.sinclude "${.CURDIR}/generated.deps.mk"
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
