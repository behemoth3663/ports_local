PORTNAME=	terraform-ls
PORTVERSION=	0.34.3
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

pre-deps: .PHONY
	if [ -d '${MASTERDIR}/files_' ]; then \
		${RM} -r ${MASTERDIR}/files; \
	else \
		${MV} ${MASTERDIR}/files ${MASTERDIR}/files_; \
	fi

pre-deps-all: .PHONY

post-deps-all: .PHONY
	test ! -d '${MASTERDIR}/files_' || ${MV} ${MASTERDIR}/files_ ${MASTERDIR}/files

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
