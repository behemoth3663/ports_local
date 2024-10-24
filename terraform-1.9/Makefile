PORTNAME=	terraform
PORTVERSION=	1.9.8
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER= 	dutra@FreeBSD.org
COMMENT=	Provides a common configuration to launch infrastructure
WWW=		https://terraform.io/

USES=		go:modules

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

post-patch: .PHONY
#	${CP} -a ${WRKSRC} ${MASTERDIR}
	cd ${WRKSRC} && \
		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} generate && \
		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} mod vendor
	${CAT} "${PATCHDIR}"/extra-* | ${PATCH} --forward --directory ${WRKSRC}
#	${REINPLACE_CMD} \
#		-e '/^[[:space:]]*go\.opencensus\.io[[:space:]]/d' \
#		-e '/^[[:space:]]*go\.opentelemetry\.io\//d' \
#	${WRKSRC}/go.mod ${WRKSRC}/go.sum
	${RM} -r ${WRKSRC}/vendor/go.opencensus.io ${WRKSRC}/vendor/go.opentelemetry.io \
		${WRKSRC}/telemetry.go \
		${WRKSRC}/internal/command/telemetry.go \
		${WRKSRC}/internal/promising/promise.go \
		${WRKSRC}/internal/promising/telemetry.go \
		${WRKSRC}/internal/rpcapi/telemetry.go \
		${WRKSRC}/internal/stacks/stackruntime/telemetry.go \
		${WRKSRC}/internal/stacks/stackruntime/internal/stackeval/telemetry.go \
		${WRKSRC}/vendor/github.com/hashicorp/aws-sdk-go-base/v2/logging/attributes.go

pre-extract: .PHONY .SILENT
	test -z '${WRKSRC_grpc_grpc_go_1}' || ${LN} -s ${WRKDIR}/grpc-go-cmd-protoc-gen-go-grpc-v1.1.0 ${WRKSRC_grpc_grpc_go_1}

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}${PKGNAMESUFFIX}

#post-patch: .PHONY
#	test -f ${WRKSRC}/vendor/modules.txt || test ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
#	cd ${WRKSRC} && \
#		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} generate && \
#		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} mod vendor

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

# XXX: Work around ports framework inability to regenerate GH_TUPLE for Go
# modules for modules with embedded slashes in version numbers.
GH_TUPLE:=	${GH_TUPLE:S|v1.1.0:grpc_grpc_go_1|cmd%2Fprotoc-gen-go-grpc%2Fv1.1.0:grpc_grpc_go_1/|1}

.include <bsd.port.mk>
