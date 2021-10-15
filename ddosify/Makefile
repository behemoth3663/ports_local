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

GH_HASH=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
		-X main.GitVersion=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X main.GitCommit=${:!${GH_HASH}!:[1]} \
		-X main.BuildDate=${:!date +%Y-%m-%dT%H:%M:%SZ!}"

PLIST_FILES=	bin/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
