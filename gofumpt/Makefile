PORTNAME=	gofumpt
PORTVERSION=	0.7.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Stricter gofmt
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mvdan

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

#const fallbackVersion = "(devel)" // to match the default from runtime/debug
post-patch: .PHONY
	${REINPLACE_CMD} 's| = "(devel)"| = "${DISTVERSIONPREFIX}${PORTVERSION}"|' \
		${WRKSRC}/internal/version/version.go

pre-deps pre-deps-all post-deps-all: .PHONY

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
