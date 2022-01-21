#

PORTNAME=	spruce
DISTVERSIONPREFIX=	v
DISTVERSION=	1.29.0
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML & JSON merging tool

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	geofffranks

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

post-extract: .PHONY
	cd ${WRKSRC} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -go=1.16 && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -go=1.17

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
