PORTNAME=	terraform
PORTVERSION=	1.6.5
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER=	none@freebsd.org
COMMENT=	Provides a common configuration to launch infrastructure
WWW=		https://terraform.io/

LICENSE=	BSL # not match 'Business Source License 1.1'
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}${PKGNAMESUFFIX} \
		bin/${PORTNAME}-${PORTVERSION}

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+";"") | \
		select(test("^${PORTVERSION:C/[^.]+$/[0-9]+/:S/./\\\\./g}$$"; \
		"n"))][0]

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.dev=no' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.VersionPrerelease='"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-extract: .PHONY .SILENT
	test -n '${WRKSRC_grpc_grpc_go_1}' && ${LN} -s ${WRKDIR}/grpc-go-cmd-protoc-gen-go-grpc-v1.1.0 ${WRKSRC_grpc_grpc_go_1} || :

post-patch: .PHONY .SILENT
	${RM} ${WRKSRC}/telemetry.go ${WRKSRC}/internal/command/telemetry.go

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}${PKGNAMESUFFIX}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

# XXX: Work around ports framework inability to regenerate GH_TUPLE for Go
# modules for modules with embedded slashes in version numbers.
GH_TUPLE:=	${GH_TUPLE:S|v1.1.0:grpc_grpc_go_1|cmd%2Fprotoc-gen-go-grpc%2Fv1.1.0:grpc_grpc_go_1/|1}

.include <bsd.port.mk>
