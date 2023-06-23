PORTNAME=	spruce
DISTVERSIONPREFIX=	v
DISTVERSION=	1.30.2
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML & JSON merging tool
WWW=		https://github.com/${GH_ACCOUNT}/${PORTNAME}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	geofffranks

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

#post-extract: .PHONY
#	cd ${WRKSRC} && \
#		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -go=1.16 && \
#		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -go=1.17

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
