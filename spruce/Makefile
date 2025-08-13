PORTNAME=	spruce
DISTVERSIONPREFIX=	v
DISTVERSION=	1.31.1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML & JSON merging tool
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	geofffranks

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

#post-extract: .PHONY
#	cd ${WRKSRC} && \
#		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -go=1.16 && \
#		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -go=1.17

pre-deps pre-deps-all post-deps-all: .PHONY

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
