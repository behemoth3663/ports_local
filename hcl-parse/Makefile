PORTNAME=	hcl-parse
PORTVERSION=	20240727
#DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Dump HCL native syntax AST by hclsyntax.Node
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters
GH_TAGNAME=	51f141e413f85cb439f0ed4fe4c04d7e47d883c4

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

#const versionStr = "0.0.1-dev"
post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
#	@${REINPLACE_CMD} -r \
#		-e 's|^([[:space:]]*const[[:space:]]+versionStr[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1"${PORTVERSION}"|' \
#	${WRKSRC}/${GO_TARGET}/main.go

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
