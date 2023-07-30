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

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'main.commit=$$(${GIT_COMMIT})' \
	-X 'main.date=${%FT%TZ:L:gmtime}'"

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
