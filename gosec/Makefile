#

PORTNAME=	gosec
PORTVERSION=	2.9.1
DISTVERSIONPREFIX=	v
CATEGORIES=	devel security local

MAINTAINER=	none@freebsd.org
COMMENT=	Golang Security Checker

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	securego

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GH_SHORT_HASH=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} \
		| ${SED} -r 's|^(.{7}).*|\1|'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.Version=${PORTVERSION}' \
	-X 'main.GitTag=${:!${GH_SHORT_HASH}!}' \
	-X 'main.BuildDate=${:!date +%F_%T!:S/_/ /}'"

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
