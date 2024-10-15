PORTNAME=	terraform-ls
PORTVERSION=	0.35.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform Language Server
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	test -f ${WRKSRC}/vendor/modules.txt || test ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	if [ -d '${PATCHDIR}-${PORTVERSION}' ]; then \
		${RM} -r ${PATCHDIR}; \
	else \
		${MV} ${PATCHDIR} ${PATCHDIR}-${PORTVERSION}; \
	fi

	${RM} -r ${PATCHDIR}-${PORTVERSION}
	${MKDIR} ${PATCHDIR}-${PORTVERSION}

	cd ${WRKSRC} && \
		${CP} -a go.mod go.mod.orig && \
		${CP} -a go.sum go.sum.orig && \
		${GO_CMD} mod tidy && \
		${SETENV} TZ=UTC ${DIFF} -udp go.mod.orig go.mod | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}-${PORTVERSION}/patch-go.mod && \
		${SETENV} TZ=UTC ${DIFF} -udp go.sum.orig go.sum | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}-${PORTVERSION}/patch-go.sum

	${FIND} ${PATCHDIR}-${PORTVERSION} -empty -delete

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}-${PORTVERSION}

post-deps-all: .PHONY
	test ! -d '${PATCHDIR}-${PORTVERSION}' || ${MV} ${PATCHDIR}-${PORTVERSION} ${PATCHDIR}

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-open-telemetry:opentelemetry-go-contrib:v0.50.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace
#-open-telemetry:opentelemetry-go-contrib:v0.50.0:open_telemetry_opentelemetry_go_contrib_1/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
#-open-telemetry:opentelemetry-go:v1.25.0:open_telemetry_opentelemetry_go/vendor/go.opentelemetry.io/otel
GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:opentelemetry-go*}

#-chigopher:pathlib:v1.0.0:chigopher_pathlib/vendor/github.com/chigopher/pathlib
#+chigopher:pathlib:v1.0.1:chigopher_pathlib/vendor/github.com/chigopher/pathlib
GH_TUPLE:=	${GH_TUPLE:C|^(chigopher:pathlib:)[^:]+|\1v1.0.1|1}

.include <bsd.port.mk>
