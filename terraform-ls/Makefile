PORTNAME=	terraform-ls
PORTVERSION=	0.38.5
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
GO_BUILDFLAGS+=	-ldflags="-s -w"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

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

#-chigopher:pathlib:v1.0.0:chigopher_pathlib/vendor/github.com/chigopher/pathlib
#+chigopher:pathlib:v1.0.1:chigopher_pathlib/vendor/github.com/chigopher/pathlib
GH_TUPLE:=	${GH_TUPLE:C|^(chigopher:pathlib:)[^:]+|\1v1.0.1|1}

GH_TUPLE+=	hashicorp:go-version:v1.8.0:hashicorp_go_version/vendor/github.com/hashicorp/go-version

.include <bsd.port.mk>
