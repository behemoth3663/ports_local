PORTNAME=	terraform
PORTVERSION=	1.5.7
DISTVERSIONPREFIX=	v
PORTREVISION=	20240510
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER=	none@freebsd.org
COMMENT=	Provides a common configuration to launch infrastructure
WWW=		https://terraform.io/

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}${PKGNAMESUFFIX} \
		bin/${PORTNAME}-${PORTVERSION}

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+";"") | \
		select(test("^${PORTVERSION:C/[^.]+$/[0-9]+/:S/./\\\\./g}$$"; \
		"n"))][0]

#WITH_PIE?=	yes
#WITH_RELRO: -static-relro is not implemented in external linker called from /usr/local/go*/pkg/tool/freebsd_amd64/link

#https://stackoverflow.com/questions/64019336/go-compile-to-static-binary-with-pie
#go build  -ldflags '-linkmode external -s -w -extldflags "--static-pie"' -buildmode=pie -tags 'osusergo,netgo,static_build' -o /hello hello.go
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static${WITH_PIE:D-pie}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.VersionPrerelease='"

pre-extract: .PHONY .SILENT
	test -n '${WRKSRC_grpc_grpc_go_1}' && ${LN} -s ${WRKDIR}/grpc-go-cmd-protoc-gen-go-grpc-v1.1.0 ${WRKSRC_grpc_grpc_go_1} || :

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
