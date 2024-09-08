PORTNAME=	terraform
PORTVERSION=	1.8.5
DISTVERSIONPREFIX=	v
PORTREVISION=	20240902
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER= 	dutra@FreeBSD.org
COMMENT=	Provides a common configuration to launch infrastructure
WWW=		https://terraform.io/

USES=		go:modules,1.23

LICENSE_NAME=   Business Source License 1.1
LICENSE_FILE=	${WRKSRC}/LICENSE
LICENSE_PERMS=	dist-mirror pkg-mirror pkg-sell auto-accept

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
	test -z '${WRKSRC_grpc_grpc_go_1}' || ${LN} -s ${WRKDIR}/grpc-go-cmd-protoc-gen-go-grpc-v1.1.0 ${WRKSRC_grpc_grpc_go_1}

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}${PKGNAMESUFFIX}

pre-deps: .PHONY
	if [ -d '${MASTERDIR}/files_' ]; then \
		${RM} -r ${MASTERDIR}/files; \
	else \
		${MV} ${MASTERDIR}/files ${MASTERDIR}/files_; \
	fi
	${RM} -r ${MASTERDIR}/files-${PORTVERSION}
	${MKDIR} ${MASTERDIR}/files-${PORTVERSION}

	cd ${WRKSRC} && \
		${CP} -a go.mod go.mod.orig && \
		${CP} -a go.sum go.sum.orig && \
		${GO_CMD} get -u=patch && \
		${GO_CMD} mod tidy && \
		${SETENV} TZ=UTC ${DIFF} -udp go.mod.orig go.mod | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${MASTERDIR}/files-${PORTVERSION}/patch-go.mod && \
		${SETENV} TZ=UTC ${DIFF} -udp go.sum.orig go.sum | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${MASTERDIR}/files-${PORTVERSION}/patch-go.sum

pre-deps-all: .PHONY

post-deps-all: .PHONY
	test ! -d '${MASTERDIR}/files_' || ${MV} ${MASTERDIR}/files_ ${MASTERDIR}/files

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

# XXX: Work around ports framework inability to regenerate GH_TUPLE for Go
# modules for modules with embedded slashes in version numbers.
GH_TUPLE:=	${GH_TUPLE:S|v1.1.0:grpc_grpc_go_1|cmd%2Fprotoc-gen-go-grpc%2Fv1.1.0:grpc_grpc_go_1/|1}

.include <bsd.port.mk>
