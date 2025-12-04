PORTNAME=	gh
PORTVERSION=	2.83.1
DISTVERSIONPREFIX=	v
PORTREVISION=	2
CATEGORIES=	local devel

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	GitHub's official command line tool
WWW=		https://cli.github.com

LICENSE=	MIT

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	cli
GH_PROJECT=	${GH_ACCOUNT}

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Version=${PORTVERSION} \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Date=${%FT%TZ:L:gmtime}"

_DOCS=		README.md docs/[!README]*.md
PORTDOCS=	${_DOCS:T}

OPTIONS_DEFINE=		DOCS MANPAGES
OPTIONS_DEFAULT=	MANPAGES
OPTIONS_SUB=		yes

M2T_OFFLINE=	true
.export	M2T_OFFLINE

do-build-MANPAGES-on:
	(cd ${WRKSRC} && \
		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} run ${GO_BUILDFLAGS} \
		./cmd/gen-docs --man-page --doc-path docs/man 2>/dev/null)

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	(cd ${WRKSRC} && \
		${INSTALL_DATA} ${_DOCS} \
		${STAGEDIR}${DOCSDIR})

post-install-MANPAGES-on:
	${INSTALL_MAN} ${WRKSRC}/docs/man/*.1 \
		${STAGEDIR}${PREFIX}/share/man/man1

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=google GHT_PROJ=cel-spec GHT_PATH=cel.dev/expr )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=imdario GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::(v[0-9.]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/}_${GHT_PROJ:tl:S/-/_/}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/client/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/internal/stream/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/storage/internal/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/storage/gcp/otel.go
	@${RM} ${WRKSRC}/vendor/github.com/transparency-dev/tessera/storage/gcp/antispam/otel.go
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/internal/trace
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/stats/opentelemetry
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/api
	@test -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc
	@test -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry_test.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentracing.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentracing_test.go
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/ansi
	@test -z '${WRKSRC_charmbracelet_x_ansi}' || ${RLN} ${WRKSRC_charmbracelet_x_ansi}/ansi ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/cellbuf
	@test -z '${WRKSRC_charmbracelet_x_cellbuf}' || ${RLN} ${WRKSRC_charmbracelet_x_cellbuf}/cellbuf ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/term
	@test -z '${WRKSRC_charmbracelet_x_term}' || ${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-GoogleCloudPlatform:opentelemetry-operations-go:d637680bde0c:googlecloudplatform_opentelemetry_operations_go_gcp
#-census-instrumentation:opencensus-go:v0.24.0:census_instrumentation_opencensus_go/vendor/go.opencensus.io
#-open-telemetry:opentelemetry-go-contrib:v0.54.0:open_telemetry_opentelemetry_go_contrib_1/vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc
#-opentracing:opentracing-go:v1.2.0:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:opentelemetry-operations-go\:*}
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:*}
#GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:*}
GH_TUPLE:=	${GH_TUPLE:Nopentracing\:*}
#GH_TUPLE:=	${GH_TUPLE:Ngolang\:telemetry\:*}

#-GoogleCloudPlatform:grpc-gcp-go:v1.5.2:googlecloudplatform_grpc_gcp_go_grpcgcp/vendor/github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp
#+GoogleCloudPlatform:grpc-gcp-go:grpcgcp/v1.5.2:googlecloudplatform_grpc_gcp_go_grpcgcp/vendor/github.com/GoogleCloudPlatform/grpc-gcp-go
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/GoogleCloudPlatform/grpc-gcp-go)/(grpcgcp)$|\1\3/\2|}

#-moby:moby:v0.7.0:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.7.0:docker_docker_credential_helpers/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/docker-credential-helpers)$|docker:docker-credential-helpers\1|}

#-google:go-genproto:c811ad7063a7:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:c811ad7063a7:google_go_genproto_api
#-google:go-genproto:94a12d6c2237:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:94a12d6c2237:google_go_genproto_rpc

#-google:go-genproto:513f23925822:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:513f23925822:google_go_genproto_api
#-google:go-genproto:513f23925822:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:513f23925822:google_go_genproto_rpc
#google:go-genproto:f936aa4a68b2:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:C,_./vendor/google\.golang\.org/genproto/googleapis/(api|rpc),_\1,}

#-googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|}

# googleapis:google-cloud-go:v0.110.7:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v1.1.1:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.30.1:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
#+googleapis:google-cloud-go:storage/v1.30.1:googleapis_google_cloud_go_storage

# googleapis:google-cloud-go:v0.121.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v0.16.2:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#-googleapis:google-cloud-go:v0.6.7:googleapis_google_cloud_go_longrunning/vendor/cloud.google.com/go/longrunning
#-googleapis:google-cloud-go:v1.24.2:googleapis_google_cloud_go_monitoring/vendor/cloud.google.com/go/monitoring
#-googleapis:google-cloud-go:v1.5.2:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.55.0:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
#-googleapis:google-cloud-go:v1.82.0:googleapis_google_cloud_go_spanner/vendor/cloud.google.com/go/spanner
#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/cloud.google.com/go/(storage)|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:N*/cloud.google.com/go/*}

# charmbracelet:x:v0.0.13:charmbracelet_x_cellbuf/vendor/github.com/charmbracelet/x/cellbuf
# charmbracelet:x:v0.2.1:charmbracelet_x_term/vendor/github.com/charmbracelet/x/term
#-charmbracelet:x:v0.8.0:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+charmbracelet:x:ansi/v0.8.0:charmbracelet_x_ansi
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(ansi\|cellbuf\|term)|\1\3/\2|}

#-charmbracelet:x:821143405392:charmbracelet_x_slice/vendor/github.com/charmbracelet/x/exp/slice
#-charmbracelet:x:821143405392:charmbracelet_x_strings/vendor/github.com/charmbracelet/x/exp/strings
#+charmbracelet:x:821143405392:charmbracelet_x/vendor/github.com/charmbracelet/x
GH_TUPLE:=	${GH_TUPLE:C|(.*:charmbracelet_x)_.*(/vendor/github.com/charmbracelet/x)/exp/.*|\1\2|:u}

#-containerd:stargz-snapshotter:v0.16.3:containerd_stargz_snapshotter_estargz/vendor/github.com/containerd/stargz-snapshotter/estargz
#+containerd:stargz-snapshotter:v0.16.3:containerd_stargz_snapshotter_estargz/vendor/github.com/containerd/stargz-snapshotter
GH_TUPLE:=	${GH_TUPLE:S|/estargz||1}

#-envoyproxy:go-control-plane:v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane/envoy
#+envoyproxy:go-control-plane:envoy/v1.32.4:envoyproxy_go_control_plane/vendor/github.com/envoyproxy/go-control-plane
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/envoyproxy/go-control-plane)/(envoy)$|\1\3/\2|}

#-cncf:xds:d92e9ce0af51:cncf_xds_go/vendor/github.com/cncf/xds/go
#+cncf:xds:d92e9ce0af51:cncf_xds_go/vendor/github.com/cncf/xds
GH_TUPLE:=	${GH_TUPLE:S|/xds/go|/xds|1}

#-spiffe:go-spiffe:v2.5.0:spiffe_go_spiffe_v2/vendor/github.com/spiffe/go-spiffe/v2
#+spiffe:go-spiffe:v2.5.0:spiffe_go_spiffe_v2/vendor/github.com/spiffe/go-spiffe
GH_TUPLE:=	${GH_TUPLE:S|/go-spiffe/v2|/go-spiffe|1}

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
GH_TUPLE:=	${GH_TUPLE:N*/go-openapi/swag/*}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
