#

PORTNAME=	gofumpt
PORTVERSION=	0.4.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Stricter gofmt
WWW=		https://github.com/${GH_ACCOUNT}/${PORTNAME}

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mvdan

PLIST_FILES=	bin/${PORTNAME}

#const fallbackVersion = "(devel)" // to match the default from runtime/debug
post-patch: .PHONY
	${REINPLACE_CMD} 's| = "(devel)"| = "${DISTVERSIONPREFIX}${PORTVERSION}"|' \
		${WRKSRC}/internal/version/version.go

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
