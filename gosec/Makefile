PORTNAME=	gosec
PORTVERSION=	2.22.1
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

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

#post-extract: .PHONY
#	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
#	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/auth
#	@${RLN} ${WRKSRC_googleapis_google_cloud_go_auth}/auth ${WRKSRC}/vendor/cloud.google.com/go
#	@${MKDIR} ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
#	@${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
#	@${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
#	@${MKDIR} ${WRKSRC}/vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc
#	@${RLN} ${WRKSRC_open_telemetry_opentelemetry_go_contrib}/instrumentation/google.golang.org/grpc/otelgrpc ${WRKSRC}/vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc
#	@${MKDIR} ${WRKSRC}/vendor/go.opentelemetry.io/contrib/instrumentation/net/http
#	@${RLN} ${WRKSRC_open_telemetry_opentelemetry_go_contrib_1}/instrumentation/net/http/otelhttp ${WRKSRC}/vendor/go.opentelemetry.io/contrib/instrumentation/net/http

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:29210b9bc287:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#-google:go-genproto:e6fa225c2576:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:29210b9bc287:google_go_genproto_rpc
#+google:go-genproto:e6fa225c2576:google_go_genproto_api
#GH_TUPLE:=	${GH_TUPLE:C|(:google_go_genproto).*/|\1_|}

#-googleapis:gax-go:v2.14.1:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2 \
#+googleapis:gax-go:v2.14.1:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go \
#GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#make -V GH_TUPLE | tr ' ' '\n' | i googleapis:google-cloud-go
#	googleapis:google-cloud-go:v0.112.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-	googleapis:google-cloud-go:v0.13.0:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+	googleapis:google-cloud-go:auth/v0.13.0:googleapis_google_cloud_go_auth
#-	googleapis:google-cloud-go:v0.6.0:googleapis_google_cloud_go_2/vendor/cloud.google.com/go/compute/metadata
#+	googleapis:google-cloud-go:compute/metadata/v0.6.0:googleapis_google_cloud_go_compute_metadata
#GH_TUPLE:=	${GH_TUPLE:N*\:googleapis_google_cloud_go_[b-z]*}
#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(auth)$|\1\3/\2|}
#GH_TUPLE:=	${GH_TUPLE:N*/vendor/cloud.google.com/go/*}

#GH_TUPLE+=	open-telemetry:opentelemetry-go-contrib:instrumentation/google.golang.org/grpc/otelgrpc/v0.58.0:open_telemetry_opentelemetry_go_contrib
#GH_TUPLE+= 	open-telemetry:opentelemetry-go-contrib:instrumentation/net/http/otelhttp/v0.58.0:open_telemetry_opentelemetry_go_contrib_1 
#GH_TUPLE+=	open-telemetry:opentelemetry-go:v1.34.0:open_telemetry_opentelemetry_go/vendor/go.opentelemetry.io/otel

.include <bsd.port.mk>
