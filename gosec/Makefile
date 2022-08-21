#

PORTNAME=	gosec
PORTVERSION=	2.13.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel security

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

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.Version=${PORTVERSION}' \
	-X 'main.GitTag=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'main.BuildDate=${%F %T:L:gmtime}'"

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
