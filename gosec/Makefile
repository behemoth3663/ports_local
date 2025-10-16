PORTNAME=	gosec
PORTVERSION=	2.22.10
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel security

MAINTAINER=	none@freebsd.org
COMMENT=	Golang Security Checker
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	securego

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.GitTag=$$(${GIT_COMMIT})' \
	-X 'main.Version=${PORTVERSION}' \
	-X 'main.BuildDate=${%F %T:L:gmtime}'"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY .SILENT
	${MKDIR} ${WRKSRC}/vendor/cloud.google.com/go
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	${RM} ${WRKSRC}/vendor/cloud.google.com/go/longrunning/autogen/auxiliary_go123.go
	${RM} ${WRKSRC}/vendor/cloud.google.com/go/ai/generativelanguage/apiv1beta/auxiliary_go123.go
	${MKDIR} ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	test -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	test -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/auth
	test -z ${WRKSRC_googleapis_google_cloud_go_auth} || ${RLN} ${WRKSRC_googleapis_google_cloud_go_auth}/auth ${WRKSRC}/vendor/cloud.google.com/go
	${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/compute/metadata
	test -z ${WRKSRC_googleapis_google_cloud_go_compute_metadata} || ${RLN} ${WRKSRC_googleapis_google_cloud_go_compute_metadata}/compute/metadata ${WRKSRC}/vendor/cloud.google.com/go/compute

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=go-genai GHT_PATH=google.golang.org/genai )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:29210b9bc287:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#-google:go-genproto:e6fa225c2576:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:29210b9bc287:google_go_genproto_rpc
#+google:go-genproto:e6fa225c2576:google_go_genproto_api
GH_TUPLE:=	${GH_TUPLE:C|(:google_go_genproto).*/|\1_|}

#-googleapis:gax-go:v2.14.1:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.14.1:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#make -V GH_TUPLE | tr ' ' '\n' | i googleapis:google-cloud-go
#	googleapis:google-cloud-go:v0.112.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-	googleapis:google-cloud-go:v0.13.0:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+	googleapis:google-cloud-go:auth/v0.13.0:googleapis_google_cloud_go_auth
#-	googleapis:google-cloud-go:v0.6.0:googleapis_google_cloud_go_2/vendor/cloud.google.com/go/compute/metadata
#+	googleapis:google-cloud-go:compute/metadata/v0.6.0:googleapis_google_cloud_go_compute_metadata
GH_TUPLE:=	${GH_TUPLE:N*\:googleapis_google_cloud_go_[b-z]*}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(auth)$|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/0-9]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2\3_\4|}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/cloud.google.com/go/*}

#GH_TUPLE+=	open-telemetry:opentelemetry-go-contrib:instrumentation/google.golang.org/grpc/otelgrpc/v0.58.0:open_telemetry_opentelemetry_go_contrib
#GH_TUPLE+=	open-telemetry:opentelemetry-go-contrib:instrumentation/net/http/otelhttp/v0.58.0:open_telemetry_opentelemetry_go_contrib_net_http_otelhttp
#GH_TUPLE+=	open-telemetry:opentelemetry-go-instrumentation:sdk/v1.1.0:open_telemetry_opentelemetry_go_instrumentation_sdk
#GH_TUPLE+=	open-telemetry:opentelemetry-go:v1.34.0:open_telemetry_opentelemetry_go/vendor/go.opentelemetry.io/otel

.sinclude "${.CURDIR}/generated.deps.mk"

GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
