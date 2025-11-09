PORTNAME=	oauth2-proxy
PORTVERSION=	7.13.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local www

MAINTAINER=	freebsd@rheinwolf.de
COMMENT=	OAuth 2.0 proxy server
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_RC_SUBR=	oauth2-proxy

USE_GITHUB=	yes

USER=		www
GROUP=		www

VAR_DIR?=	/var
LOG_DIR?=	${VAR_DIR}/log/${PORTNAME}
RUN_DIR?=	${VAR_DIR}/run/${PORTNAME}

SUB_LIST+=	GROUP=${GROUP} \
		LOG_DIR=${LOG_DIR} \
		PORTNAME=${PORTNAME} \
		PREFIX=${PREFIX} \
		RUN_DIR=${RUN_DIR} \
		USER=${USER}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'version.VERSION=${PORTVERSION}'"

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${LOG_DIR}
	${MKDIR} ${STAGEDIR}${RUN_DIR}
	${INSTALL_DATA} ${WRKSRC}/contrib/${PORTNAME}.cfg.example ${STAGEDIR}${PREFIX}/etc/${PORTNAME}.cfg.sample

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

pre-deps post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:6ceb2ff114de:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:6ceb2ff114de:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc$||}

#-googleapis:google-cloud-go:v0.2.5:googleapis_google_cloud_go/vendor/cloud.google.com/go/auth/oauth2adapt
#+googleapis:google-cloud-go:auth/oauth2adapt/v0.2.5:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v0.5.2:googleapis_google_cloud_go_1/vendor/cloud.google.com/go/compute/metadata
#+googleapis:google-cloud-go:compute/metadata/v0.5.2:googleapis_google_cloud_go_1/vendor/cloud.google.com/go
GH_TUPLE:=	${GH_TUPLE:C|^(googleapis:google-cloud-go:)(.*/go)/(.*)$|\1\3/\2|}

#-open-telemetry:opentelemetry-go-contrib:v0.55.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
#+open-telemetry:opentelemetry-go-contrib:instrumentation/net/http/otelhttp/v0.55.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib
GH_TUPLE:=	${GH_TUPLE:C|^(open-telemetry:opentelemetry-go-contrib:)(.*/contrib)/(.*)$|\1\3/\2|1}

#-open-telemetry:opentelemetry-go-instrumentation:v1.1.0:open_telemetry_opentelemetry_go_instrumentation/vendor/go.opentelemetry.io/auto/sdk
#+open-telemetry:opentelemetry-go-instrumentation:sdk/v1.1.0:open_telemetry_opentelemetry_go_instrumentation/vendor/go.opentelemetry.io/auto
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/go.opentelemetry.io/auto)/(sdk)|\1\3/\2|1}

.include <bsd.port.mk>
