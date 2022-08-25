#

PORTNAME=	hclfmt
PORTVERSION=	2.13.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Format and prettify HCL files

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp
GH_PROJECT=	hcl

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

#const versionStr = "0.0.1-dev"
post-patch:
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*const[[:space:]]+versionStr[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1"${PORTVERSION}"|' \
	${WRKSRC}/${GO_TARGET}/main.go

.include <bsd.port.mk>
