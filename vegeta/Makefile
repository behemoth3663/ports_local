PORTNAME=	vegeta
PORTVERSION=	12.10.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local benchmarks

MAINTAINER=	none@freebsd.org
COMMENT=	HTTP load testing tool and library
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	tsenart

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'main.Commit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'main.Date=${%FT%TZ:L:gmtime}'"

post-post-extract: .PHONY .SILENT

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
