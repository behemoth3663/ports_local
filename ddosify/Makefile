PORTNAME=	ddosify
PORTVERSION=	1.0.6
DISTVERSIONPREFIX=	v
#PORTREVISION=	2
CATEGORIES=	local benchmarks

MAINTAINER=	none@freebsd.org
COMMENT=	High-performance load testing tool
WWW=		https://ddosify.com/

LICENSE=	AGPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	ddosify

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
		-X 'main.GitVersion=${DISTVERSIONPREFIX}${PORTVERSION}' \
		-X 'main.GitCommit=$$(${GIT_COMMIT})' \
		-X 'main.BuildDate=${%FT%TZ:L:gmtime}'"

PLIST_FILES=	bin/${PORTNAME}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
