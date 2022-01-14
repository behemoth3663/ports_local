#

PORTNAME=	gofumpt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.2.1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Stricter gofmt

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mvdan

PLIST_FILES=	bin/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
