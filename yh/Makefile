#

PORTNAME=	yh
PORTVERSION=	0.2.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML syntax highlighter

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	andreazorzetto

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
