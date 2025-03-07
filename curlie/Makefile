PORTNAME=	curlie
DISTVERSIONPREFIX=	v
DISTVERSION=	1.8.2
CATEGORIES=	local ftp net www

MAINTAINER=	0mp@FreeBSD.org
COMMENT=	Frontend to cURL that adds the ease of use of HTTPie
WWW=		https://curlie.io/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	curl:ftp/curl
BUILD_DEPENDS=	git:devel/git

USES=		go:1.24,modules
USE_GITHUB=	yes
GH_ACCOUNT=	rs

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

#commit  = "0000000"
#version = "v0.0.0-LOCAL"
#date    = "0000-00-00T00:00:00Z"
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'main.commit=$$(${GIT_COMMIT})' \
	-X 'main.date=${%FT%TZ:L:gmtime}'"

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
