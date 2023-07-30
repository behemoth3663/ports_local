PORTNAME=	hcltomd
PORTVERSION=	1.0.0
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Simple HashiCorp Configuration Language parser for your vars
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	idestis

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'main.commit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'main.date=${%FT%TZ:L:gmtime}'"

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
