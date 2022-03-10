#

PORTNAME=	hclfmt
PORTVERSION=	0.1.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Format and prettify HCL files

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	fatih

GH_TAGNAME=	374cd0f

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#USES=go:modules but there is no go.mod
pre-patch: .PHONY
	cd ${WRKSRC} && \
	${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod init ${GO_PKGNAME} && \
	${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
