#

PORTNAME=	ddosify
PORTVERSION=	0.4.1
DISTVERSIONPREFIX=	v
CATEGORIES=	benchmarks

MAINTAINER=	none@freebsd.org
COMMENT=	High-performance load testing tool

LICENSE=	AGPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	ddosify

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
		-X main.GitVersion=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X main.GitCommit=${:!${GH_HASH}!:[1]} \
		-X main.BuildDate=${%FT%TZ:L:gmtime}'"

PLIST_FILES=	bin/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>