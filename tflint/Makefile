# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	tflint
PORTVERSION=	0.56.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	dutra@FreeBSD.org
COMMENT=	Terraform Linter
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export		M2T_OFFLINE

pre-deps: .PHONY
	if [ -d '${PATCHDIR}_' ]; then \
		${RM} -r ${PATCHDIR}; \
	else \
		${MV} ${PATCHDIR} ${PATCHDIR}_; \
	fi

post-deps-all: .PHONY
	test ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

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
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:*}
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:*}
GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:*}
GH_TUPLE:=	${GH_TUPLE:Nopentracing\:*}

#-cncf:xds:b4127c9b8d78:cncf_xds_go/vendor/github.com/cncf/xds/go
#+cncf:xds:b4127c9b8d78:cncf_xds_go/vendor/github.com/cncf/xds
GH_TUPLE:=	${GH_TUPLE:S|/xds/go|/xds|1}

#-googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-owenrumney:go-sarif:v2.3.3:owenrumney_go_sarif_v2/vendor/github.com/owenrumney/go-sarif/v2
#+owenrumney:go-sarif:v2.3.3:owenrumney_go_sarif_v2/vendor/github.com/owenrumney/go-sarif
GH_TUPLE:=	${GH_TUPLE:S|/go-sarif/v2|/go-sarif|1}

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
#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/0-9]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2\3_\4|}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/cloud.google.com/go/*}

GH_TUPLE+=	hashicorp:go-version:v1.7.0:hashicorp_go_version/vendor/github.com/hashicorp/go-version
GH_TUPLE+=	spf13:cast:v1.7.0:spf13_cast/vendor/github.com/spf13/cast

.sinclude "${.CURDIR}/generated.deps.mk"
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
