#

PORTNAME=	sysbox
PORTVERSION=	0.14.0
DISTVERSIONPREFIX=	release-
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	sysadmin/scripting utilities, distributed as a single binary

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	skx

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.versionString=${PORTVERSION}'"

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>