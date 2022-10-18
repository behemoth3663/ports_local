PORTNAME=	curlie
DISTVERSIONPREFIX=	v
DISTVERSION=	1.6.9
CATEGORIES=	ftp net www

MAINTAINER=	0mp@FreeBSD.org
COMMENT=	Frontend to cURL that adds the ease of use of HTTPie
WWW=		https://curlie.io/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	curl:ftp/curl

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	rs

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

#commit  = "0000000"
#version = "v0.0.0-LOCAL"
#date    = "0000-00-00T00:00:00Z"
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'main.commit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'main.date=${%FT%TZ:L:gmtime}'"

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
