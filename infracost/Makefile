#

PORTNAME=	infracost
PORTVERSION=	0.10.7
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Cloud cost estimates for Terraform

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	infracost

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#.-include "${.CURDIR}/Makefile.deps"
#_OPTIONS_extract+=	600:pre-post-extract
#pre-post-extract: .PHONY
#	cd ${WRKSRC} && \
#		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod vendor -v && \
#		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v
#	${CP} ${.CURDIR}/files/go.mod ${.CURDIR}/files/go.sum ${WRKSRC}

pre-build: .PHONY
	cd ${WRKSRC} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod vendor -v && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v

.include <bsd.port.mk>
