#

PORTNAME=	gogrep
PORTVERSION=	0.0.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Search for Go code using syntax trees

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mvdan
GH_TAGNAME=	624f953

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
