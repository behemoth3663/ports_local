PORTNAME=	terraform-ls
PORTVERSION=	0.39.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel
PKGNAMESUFFIX=	-local

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform Language Server
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

clean-unused: .PHONY clean patch
	( cd ${WRKSRC} && ${GO_CMD} mod tidy -v )
	${RM} -r ${WRKSRC}/vendor
	( ${MAKE} -C ${.CURDIR} gomod-vendor | \
		${GREP} --extended-regexp --invert-match '^(/|=+>|\*+[[:blank:]]+Error)' | \
		${SED} -r -e '/\r$$/d' -e '/^[[:blank:]]+"message":/d' -e '/\\}/d' -e '/^[[:blank:]]*$$/d' \
		> ${.CURDIR}/Makefile.deps )
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}
	${PATCH} --directory ${.CURDIR} < ${.CURDIR}/patch-Makefile.deps
	( ${MAKE} -C ${.CURDIR} makesum )

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	${TEST} ! -f ${PATCHDIR}/modules.txt || ${TEST} ! -d ${WRKSRC}/vendor || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-open-telemetry:opentelemetry-go-contrib:v0.50.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace
#-open-telemetry:opentelemetry-go-contrib:v0.50.0:open_telemetry_opentelemetry_go_contrib_1/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
#-open-telemetry:opentelemetry-go:v1.25.0:open_telemetry_opentelemetry_go/vendor/go.opentelemetry.io/otel
GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:opentelemetry-go*}

#unused:
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/cespare/xxhash/v2}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/felixge/httpsnoop}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/go-logr/logr}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/go-logr/stdr}

#-chigopher:pathlib:v1.0.0:chigopher_pathlib/vendor/github.com/chigopher/pathlib
#+chigopher:pathlib:v1.0.1:chigopher_pathlib/vendor/github.com/chigopher/pathlib
GH_TUPLE:=	${GH_TUPLE:C,^(chigopher:pathlib:)[^:]+,\1v1.0.1,1}

GH_TUPLE+=	hashicorp:go-version:v1.9.0:hashicorp_go_version/vendor/github.com/hashicorp/go-version

.include <bsd.port.mk>
