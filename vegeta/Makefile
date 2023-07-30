PORTNAME=	vegeta
PORTVERSION=	12.11.0
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

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'main.Commit=$$(${GIT_COMMIT})' \
	-X 'main.Date=${%FT%TZ:L:gmtime}'"

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
