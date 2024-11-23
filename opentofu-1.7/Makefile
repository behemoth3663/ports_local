PORTNAME=	opentofu
DISTVERSIONPREFIX=	v
DISTVERSION=	1.7.6
#PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER=	bofh@FreeBSD.org
COMMENT=	Tool for building, changing, and versioning infrastructure
WWW=		https://opentofu.org/

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	opentofu

PLIST_FILES=	bin/tofu${PKGNAMESUFFIX} \
		bin/tofu-${PORTVERSION}

GO_TARGET=	./cmd/tofu

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.dev=no"

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+";"") | \
		select(test("^${PORTVERSION:C/[^.]+$/[0-9]+/:S/./\\\\./g}$$"; \
		"n"))][0]

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

pre-extract: .PHONY .SILENT
	test -n '${WRKSRC_grpc_grpc_go_1}' && ${LN} -s ${WRKDIR}/grpc-go-cmd-protoc-gen-go-grpc-v1.1.0 ${WRKSRC_grpc_grpc_go_1} || :

#post-patch: .PHONY .SILENT
#	${RM} ${WRKSRC}/cmd/tofu/telemetry.go ${WRKSRC}/internal/command/telemetry.go ${WRKSRC}/vendor/github.com/hashicorp/aws-sdk-go-base/v2/logging/attributes.go

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/tofu ${STAGEDIR}${PREFIX}/bin/tofu-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/tofu-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/tofu${PKGNAMESUFFIX}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

# XXX: Work around ports framework inability to regenerate GH_TUPLE for Go
# modules for modules with embedded slashes in version numbers.
GH_TUPLE:=	${GH_TUPLE:S|v1.1.0:grpc_grpc_go_1|cmd%2Fprotoc-gen-go-grpc%2Fv1.1.0:grpc_grpc_go_1/|1}

.include <bsd.port.mk>
